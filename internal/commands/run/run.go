package run

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	pakelib "github.com/pake-go/pake-lib"
	"github.com/pake-go/pake-lib/config"
	"github.com/pake-go/pakeos/internal/utils/pathutil"
)

type run struct {
	args []string
}

func New(args []string) pakelib.Command {
	return &run{
		args: args,
	}
}

func (r *run) Execute(cfg *config.Config, logger *log.Logger) error {
	logger.Printf("Running  %s\n", strings.Join(r.args, " "))

	var cmd *exec.Cmd
	if len(r.args) != 1 {
		cmd = exec.Command("sh", "-c", strings.Join(r.args, " "))
	} else {
		executablePath := r.args[0]
		extension := filepath.Ext(executablePath)
		switch extension {
		case ".class":
			basename := filepath.Base(executablePath)
			directory := filepath.Dir(executablePath)
			filename := strings.TrimSuffix(basename, extension)
			command := fmt.Sprintf("java -cp %s %s", directory, filename)
			cmd = exec.Command("sh", "-c", command)
		case ".jar":
			cmd = exec.Command("javar", "-jar", executablePath)
		case ".java":
			basename := filepath.Base(executablePath)
			directory := filepath.Dir(executablePath)
			filename := strings.TrimSuffix(basename, extension)
			cmdFmt := "javac -d %[1]s %[3]s && java -cp %[1]s %[2]s"
			command := fmt.Sprintf(cmdFmt, directory, filename, executablePath)
			cmd = exec.Command("sh", "-c", command)
		case ".py":
			cmd = exec.Command("python", executablePath)
		case ".rb":
			cmd = exec.Command("ruby", executablePath)
		default:
			if exists := pathutil.Exists(executablePath); exists {
				cmd = exec.Command(executablePath)
			} else {
				cmd = exec.Command("sh", "-c", executablePath)
			}
		}
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type RunValidator struct {
}

func (rv *RunValidator) CanHandle(line string) bool {
	return strings.HasPrefix(line, "run ")
}

func (rv *RunValidator) ValidateArgs(args []string) bool {
	return len(args) > 0
}
