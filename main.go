package main

import (
	"fmt"

	log "github.com/jeanphorn/log4go"

	"github.com/jeavenwong/study/routers"
)

func main() {
	fmt.Println("start program...")
	log.LoadConfiguration("./log.json")
	//log.LOGGER("DEBUG").Debug("test print log, %s", "hello, world!")
	log.Debug("test print log, %s", "hello, world!")
	routers.Init()
	routers.Run()
}
