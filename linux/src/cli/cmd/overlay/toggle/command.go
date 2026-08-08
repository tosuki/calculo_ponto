package pause

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "toggle",
		Short: "Pausa o relógio.",
		Run: func(cmd *cobra.Command, args []string) {
			client := &http.Client{}

			host, hostErr := cmd.Flags().GetString("host")

			if hostErr != nil {
				log.Fatalln(hostErr.Error())
			}

			port, portErr := cmd.Flags().GetInt("port")

			if portErr != nil {
				log.Fatalln(portErr.Error())
			}

			resp, err := client.Get(fmt.Sprintf("http://%s:%d/timer/toggle", host, port))

			if err != nil {
				log.Fatalln(err.Error())
			}

			if resp.StatusCode != 200 {
				fmt.Printf("Não foi possivel realizar a operação devido a status code inesperado: %s\n", resp.StatusCode)
				return
			}

			fmt.Println("Feito!")
		},
	}

	cmd.Flags().String("host", "localhost", "Host que a API se encontra rodando.")
	cmd.Flags().Int("port", 8080, "Número da porta que a API se encontra rodando.")

	return cmd
}
