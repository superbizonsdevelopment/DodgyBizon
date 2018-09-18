package downloader

import (
  //"../manager"
  "github.com/knadh/go-get-youtube/youtube"
)

var option = &youtube.Option{
  Rename: true,
  Resume: true,
  Mp3:    true,
}

func DownloadVideoFromYoutube(link string) error {
  /*videoID, err := manager.GetVideoIDFromLink(link)

  if err != nil {
    return err
  }
*/
  video, err := youtube.Get("FTl0tl9BGdc")

  if err != nil {
    return err
  }

  option := &youtube.Option{
    Rename: true,
    Resume: true,
    Mp3: true,
  }
  video.Download(0, "video.mp3", option)
  return nil
}
