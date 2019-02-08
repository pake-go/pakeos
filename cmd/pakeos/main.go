package main

import (
	"log"
	"os"

	"github.com/PGo-Projects/output"
	"github.com/pake-go/pake-lib/executor"
	"github.com/pake-go/pake-lib/parser"
	"github.com/pake-go/pakeos/internal/commands"
	"github.com/pake-go/pakeos/internal/comments"
	"github.com/spf13/cobra"
)

// RootCmd defines the CLI app.
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

// RunFile parses and runs the file.
func runFile(cmd *cobra.Command, args []string) {
	filename := args[0]

	logger := log.New(os.Stderr, "", log.LstdFlags)
	pparser := parser.New(commands.Candidates, &comments.Validator{})
	commands, err := pparser.ParseFile(filename, logger)
	if err != nil {
		output.Error(err)
		os.Exit(1)
	}
	executor.Run(commands, logger)
}
