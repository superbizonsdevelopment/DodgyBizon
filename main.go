package main

import (
	"log"
	"./gui"
	"./util"
  "./constants"
  "github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func main() {
  log.Printf("Starting %s %s", constants.APP_NAME, constants.APP_VERSION)
	log.Println("Making app dir.")
	util.CreateDirIfNotExist(constants.APP_DIR)
	log.Println("Making Application GUI.")
	ui.Main(gui.SetUpUI)
}
