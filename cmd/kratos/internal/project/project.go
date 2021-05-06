package project

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// CmdNew represents the new command.
var CmdNew = &cobra.Command{
	Use:   "new",
	Short: "Create a service template",
	Long:  "Create a service project using the repository template. Example: kratos new helloworld",
	Run:   run,
}

var repoURL string

func init() {
	if repoURL = os.Getenv("KRATOS_LAYOUT_REPO"); repoURL == "" {
		repoURL = "https://github.com/go-kratos/kratos-layout.git"
	}
	CmdNew.Flags().StringVarP(&repoURL, "-repo-url", "r", repoURL, "layout repo")
}

func run(cmd *cobra.Command, args []string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: project name is required.\033[m Example: kratos new helloworld\n")
		return
	}
	p := &Project{Name: args[0]}
	if err := p.New(ctx, wd, repoURL); err != nil {
		fmt.Fprintf(os.Stderr, "\033[31mERROR: %s\033[m\n", err)
		return
	}
}
