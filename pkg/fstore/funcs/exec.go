package funcs

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/cli/safeexec"
	"github.com/kballard/go-shellquote"
	"github.com/rytsh/liz/shutdown"
	"github.com/worldline-go/logz"

	"github.com/rytsh/mugo/pkg/fstore/registry"
)

func init() {
	registry.CallReg.AddFunction("exec", new(Exec).init, "trust", "log")
}

type Exec struct {
	trust bool
	log   logz.Adapter
}

func (e Exec) init(trust bool, log logz.Adapter) any {
	e.trust = trust
	e.log = log
	if e.log == nil {
		e.log = logz.AdapterNoop{}
	}

	return e.Exec
}

func (e Exec) Exec(cli string) (map[string]interface{}, error) {
	if !e.trust {
		return nil, registry.ErrTrustRequired
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
