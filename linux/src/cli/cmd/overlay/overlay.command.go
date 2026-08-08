package overlay

import (
	"time"

	"github.com/4mti/ponto/src/cli/cmd/overlay/serve"
	toggle "github.com/4mti/ponto/src/cli/cmd/overlay/toggle"
	"github.com/4mti/ponto/src/core"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	timer := core.NewTimer(core.TimerModeRegressive, time.Now(), time.Hour*6)
	cfg := core.NewConfig()

	cmd := &cobra.Command{
		Use:   "overlay",
		Short: "Start overlay",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(serve.NewCommand(timer, cfg))
	cmd.AddCommand(toggle.NewCommand())

	return cmd
}
