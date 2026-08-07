package src

import (
	"log"
	"time"

	"github.com/4mti/ponto/src/core"
	"github.com/4mti/ponto/src/render"
	"github.com/4mti/ponto/src/server"
)

func Start() {
	config := core.NewConfig()
	timer := core.NewTimer(core.TimerModeRegressive, time.Now(), time.Hour*6)

	go server.StartServer(timer, config)
	if err := render.RunApp(timer, config); err != nil {
		log.Fatalln(err.Error())
	}
}
