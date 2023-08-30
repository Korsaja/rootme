package printer

import (
	"fmt"
	"io"

	"korsaj.io/rootme/pkg/models"

	. "github.com/xlab/treeprint"
)

type TreePrinter struct {
	stdout io.Writer
	stderr io.Writer
}

func NewTreePrinter(stdout, stderr io.Writer) *TreePrinter {
	return &TreePrinter{stdout: stdout, stderr: stderr}
}

func (p *TreePrinter) PrintText(text string) {
	_, _ = fmt.Fprintln(p.stdout, text)
}

func (p *TreePrinter) PrintError(text string) {
	_, _ = fmt.Fprint(p.stderr, text)
}

func (p *TreePrinter) PrintProfile(profile *models.Profile) {
	p.printHeader(profile.NickName, profile.Rank, profile.Position)
	p.printTaskSolved(profile.Solved)
}

func (p *TreePrinter) printHeader(username, rank string, position int) {
	const header = "UserName\tRank\t\tPosition\n%s\t%s\t\t%d\n"
	_, _ = fmt.Fprintf(p.stdout, header, username, rank, position)
}
func (p *TreePrinter) printTaskSolved(taskSolved map[string][]models.Task) {
	tree := New()
	solved := 0
	for rubric, tasks := range taskSolved {
		node := tree.AddBranch(rubric)
		for _, task := range tasks {
			node.AddNode(task.Title)
			solved++
		}
	}
	_, _ = fmt.Fprint(p.stdout, tree.String())
	_, _ = fmt.Fprintf(p.stdout, "solved: %d\n", solved)
}
