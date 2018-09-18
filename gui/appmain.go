package gui

import (
  "../listener"
  "../constants"
  "github.com/andlabs/ui"
)

func SetUpUI() {
  mainWindow := ui.NewWindow(constants.APP_NAME + constants.SPACE + constants.APP_VERSION, constants.APP_HEIGHT, constants.APP_WIDTH, true)

  mainWindow.OnClosing(func(*ui.Window) bool {
    mainWindow.Destroy()
    ui.Quit()
    return false
  })

  ui.OnShouldQuit(func() bool {
		mainWindow.Destroy()
		return true
	})

  hbox := ui.NewHorizontalBox()
  hbox.SetPadded(true)
  mainWindow.SetChild(hbox)

  vbox := ui.NewVerticalBox()
  vbox.SetPadded(true)
  hbox.Append(vbox, false)

  linkInfoLabel := ui.NewLabel("test")

  linkField := ui.NewEntry()

  downloadButton := ui.NewButton("Download")
  downloadButton.OnClicked(func(b *ui.Button) {
    listener.OnDownloadButtonClicked(b, linkField.Text())
  })

  vbox.Append(linkInfoLabel, false)
  vbox.Append(linkField, false)
  vbox.Append(downloadButton, false)

  mainWindow.Show()
}
