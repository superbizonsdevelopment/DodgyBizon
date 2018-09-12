package main

import (
  "log"
  //"./manager"
  "./constants"
	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/gxfont"
	"github.com/google/gxui/samples/flags"
  "github.com/knadh/go-get-youtube/youtube"
)

var option = &youtube.Option{
  Rename: true,
  Resume: true,
  Mp3:    true,
}

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	font, err := driver.CreateFont(gxfont.Default, 75)
	if err != nil {
		panic(err)
	}

	window := theme.CreateWindow(380, 100, constants.APP_NAME + constants.SPACE + constants.APP_VERSION)
	window.SetBackgroundBrush(gxui.CreateBrush(gxui.Gray50))

	label := theme.CreateLabel()
	label.SetFont(font)
	label.SetText("Hello world")

	window.AddChild(label)

	window.OnClose(driver.Terminate)
}

func main() {
  log.Println("Starting DodgyBizon 1.0.0")
  gl.StartDriver(appMain)
  /*
	video, err := youtube.Get("FTl0tl9BGdc")

  if err != nil {
    log.Fatalf("Error: ", err.Error())
  }

	video.Download(0, "video.mp3", option)
  */
}
