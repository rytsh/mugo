package exec

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/cli/safeexec"
	"github.com/kballard/go-shellquote"
	"github.com/rakunlabs/logi/logadapter"
	"github.com/rytsh/liz/shutdown"

	"github.com/rytsh/mugo/fstore/errors"
)

type Exec struct {
	trust bool
	log   logadapter.Adapter
}

func New(trust bool, log logadapter.Adapter) *Exec {
	if log == nil {
		log = logadapter.Noop{}
	}

	return &Exec{
		trust: trust,
		log:   log,
	}
}

func (e Exec) Exec(cli string) (map[string]interface{}, error) {
	if !e.trust {
		return nil, errors.ErrTrustRequired
	}

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
		e.log.Error("failed to run exec")
		fmt.Println(string(cmdErrorResult))
		return nil, err
	}

	return map[string]interface{}{
		"stdout": cmdOutputResult,
		"stderr": cmdErrorResult,
		"status": cmd.ProcessState.ExitCode(),
	}, nil
}
