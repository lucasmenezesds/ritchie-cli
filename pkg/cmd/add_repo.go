package cmd

import (
	"fmt"

	"github.com/ZupIT/ritchie-cli/pkg/formula"

	"github.com/spf13/cobra"

	"github.com/ZupIT/ritchie-cli/pkg/prompt"
)

// addRepoCmd type for add repo command
type addRepoCmd struct {
	formula.AddLister
	prompt.InputText
	prompt.InputURL
	prompt.InputInt
	prompt.InputBool
}

// NewRepoAddCmd creates a new cmd instance
func NewAddRepoCmd(
	adl formula.AddLister,
	it prompt.InputText,
	iu prompt.InputURL,
	ii prompt.InputInt,
	ib prompt.InputBool) *cobra.Command {
	a := &addRepoCmd{
		adl,
		it,
		iu,
		ii,
		ib,
	}

	cmd := &cobra.Command{
		Use:     "repo",
		Short:   "Add a repository.",
		Example: "rit add repo ",
		RunE:    a.runFunc(),
	}

	return cmd
}

func (a addRepoCmd) runFunc() CommandRunnerFunc {
	return func(cmd *cobra.Command, args []string) error {
		rn, err := a.Text("Name of the repository: ", true)
		if err != nil {
			return err
		}

		repos, err := a.List()
		if err != nil {
			return err
		}
		for _, repo := range repos {
			if rn == repo.Name {
				fmt.Printf("Your repository %s is gonna be overwritten.\n", repo.Name)
				options := []string{"yes", "no"}
				choice, _ := a.Bool("Want to proceed?", options)
				if !choice {
					fmt.Println("Operation cancelled")
					return nil
				}
			}
		}

		ur, err := a.URL("URL of the tree [http(s)://host:port/tree.json]: ", "")
		if err != nil {
			return err
		}

		pr, err := a.Int("Priority [ps.: 0 is higher priority, the lower higher the priority] :")
		if err != nil {
			return err
		}

		r := formula.Repository{
			Priority: int(pr),
			Name:     rn,
			TreePath: ur,
		}

		if err = a.Add(r); err != nil {
			return err
		}

		return err
	}

}
