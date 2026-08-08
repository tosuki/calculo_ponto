package overlay

import (
	"log"
	"time"

	"github.com/4mti/ponto/src/core"
	"github.com/4mti/ponto/src/render"
	"github.com/4mti/ponto/src/server"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "overlay",
		Short: "Start overlay",
		Run: func(cmd *cobra.Command, args []string) {
			cfg := core.NewConfig()
			timer := core.NewTimer(core.TimerModeRegressive, time.Now(), time.Hour*6)

			go server.StartServer(timer, cfg)

			if err := render.RunApp(timer, cfg); err != nil {
				log.Fatalln(err.Error())
			}
		},
	}

	return cmd
}
