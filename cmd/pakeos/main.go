package main

import (
	"os"

	"github.com/PGo-Projects/output"
	"github.com/pake-go/pake-lib/executor"
	"github.com/pake-go/pake-lib/parser"
	"github.com/pake-go/pakeos/internal/commands"
	"github.com/pake-go/pakeos/internal/comments"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pakeos <file>",
	Short: "pakeos is a tool to help bootstrap OS/dotfiles",
	Args:  cobra.ExactArgs(1),
	Run:   runFile,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		output.Error(err)
		os.Exit(1)
	}
}

func runFile(cmd *cobra.Command, args []string) {
	filename := args[0]

	pparser := parser.New(commands.Candidates, &comments.Validator{})
	commands, err := pparser.ParseFile(filename)
	if err != nil {
		output.Error(err)
		os.Exit(1)
	}
	executor.Run(commands)
}
