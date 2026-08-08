package serve

import (
	"fmt"

	"github.com/4mti/ponto/src/core"
	"github.com/4mti/ponto/src/render"
	"github.com/4mti/ponto/src/server"
	"github.com/spf13/cobra"
)

func NewCommand(timer *core.Timer, cfg *core.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "Inicia a overlay com o servidor HTTP pra interações via GUI/CLI",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Inicializando servidor HTTP")
			go server.StartServer(timer, cfg)
			fmt.Println("Inicializando overlay")
			render.RunApp(timer, cfg)
		},
	}

	cmd.Flags().Int("port", 8080, "porta que será usada pra rodar a API")

	return cmd
}
