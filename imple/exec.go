package imple

import (
	"bytes"
	"fmt"
	"github.com/amine-khemissi/lets/errors"
	"io"
	"os"
	exec2 "os/exec"
	"strings"
)

func execAction(cmd string, args ...string) error {

	args = []string{"-c", fmt.Sprintf("%s %s", cmd, strings.TrimSuffix(strings.Join(args, " "), " "))}
	command := exec2.Command("bash", args...)
	var errOut bytes.Buffer
	w := io.MultiWriter(&errOut)
	command.Stdout = os.Stdout
	command.Stderr = w
	if err := command.Run(); err != nil {
		return errors.New("failed to execAction", errOut.String())
	}
	return nil
}
