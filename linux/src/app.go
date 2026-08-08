package src

import (
	"log"

	"github.com/4mti/ponto/src/cli"
)

func Start() {
	// config := core.NewConfig()
	// timer := core.NewTimer(core.TimerModeRegressive, time.Now(), time.Hour*6)

	// go server.StartServer(timer, config)
	// if err := render.RunApp(timer, config); err != nil {
	// 	log.Fatalln(err.Error())
	// }
	if err := cli.RunCli(); err != nil {
		log.Fatalln(err.Error())
	}
}
