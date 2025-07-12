package shortcuts

import (
	"slices"

	"ok.build/cli/arg"
)

var (
	shortcuts = map[string]string{
		"b": "build",
		"t": "test",
		"q": "query",
		"r": "run",
	}

	defaultTargetCommands = []string{"aquery", "build", "coverage", "cquery", "test", "query"}
)

// HandleShorcuts finds the first non-flag command and tries to expand it
func HandleShortcuts(args []string) []string {
	command, idx := arg.GetCommandAndIndex(args)
	if expanded, ok := shortcuts[command]; ok {
		args[idx] = expanded
	}

	newCommand := args[idx]

	if len(arg.GetTargets(args)) > 0 || !slices.Contains(defaultTargetCommands, newCommand) {
		return args
	}

	return append(args, "//...")
}
