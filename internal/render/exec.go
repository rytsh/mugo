package render

import (
	"io"
	"os/exec"

	"github.com/cli/safeexec"
	"github.com/kballard/go-shellquote"
	"github.com/rytsh/mugo/internal/shutdown"
)

func Exec(cli string) (map[string]interface{}, error) {
	commands, err := shellquote.Split(cli)
	if err != nil {
		return nil, err
	}

	bin, err := safeexec.LookPath(commands[0])
	if err != nil {
		return nil, err
	}

	args := []string{}
	if len(commands) > 1 {
		args = commands[1:]
	}

	cmd := exec.Command(bin, args...)

	cmdOutput, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	cmdError, err := cmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	// record the PID of the process
	v := shutdown.Global.AddAnonymous(cmd.Process.Kill)
	defer shutdown.Global.Delete(v)

	cmdOutputResult, err := io.ReadAll(cmdOutput)
	if err != nil {
		return nil, err
	}

	cmdErrorResult, err := io.ReadAll(cmdError)
	if err != nil {
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"stdout": cmdOutputResult,
		"stderr": cmdErrorResult,
		"status": cmd.ProcessState.ExitCode(),
	}, nil
}
