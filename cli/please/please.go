package please

import (
	"flag"
	"os"
	"strings"

	"ok.build/cli/claude"
)

var (
	flags = flag.NewFlagSet("ask", flag.ContinueOnError)
)

var (
	usage = `
usage: ok ` + flags.Name() + `

Asks ok to perform a task.
`
)

func HandleAsk(args []string) (int, error) {

	claudePrompt := strings.Join(args, " ")

	claude.Run(os.Stdin, []string{claudePrompt}, true)

	return 0, nil
}
