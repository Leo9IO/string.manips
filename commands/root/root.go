package root

import (
	"sync"

	"github.com/spf13/cobra"
)

var characterCmd = &cobra.Command{
	Use:   "character",
	Short: "character performs character lookups and conversions",
}

var errorCount struct {
	sync.Mutex
	value int
}

// AddCommand is the hook used by our sub-commands to register themselves
// into our CLI dispatch system.  Per-module init() hooks should use this.
func AddCommand(cmds ...*cobra.Command) {
	characterCmd.AddCommand(cmds...)
}

// Errored bumps an error count, used to determine if the main program should
// exit false or not.
func Errored() {
	errorCount.Lock()
	errorCount.value++
	errorCount.Unlock()
}

// GetErrorCount is intended for use by main(), to determine how to exit.
func GetErrorCount() int {
	errorCount.Lock()
	defer errorCount.Unlock()
	return errorCount.value
}

// Start is the entry-point used by main(), after all the sub-modules have
// registered their availability via AddCommand calls in their init functions.
func Start() {
	characterCmd.Execute()
}
