package listener

import (
  "log"
  "../downloader"
  "github.com/andlabs/ui"
)

func OnDownloadButtonClicked(b *ui.Button, linkToYoutube string) {
  log.Println("Download button clicked!")
  err := downloader.DownloadVideoFromYoutube(linkToYoutube)
  if err != nil {
    log.Fatalf(err.Error())
  } else {
      log.Println("Video Downloaded!")
  }
}
