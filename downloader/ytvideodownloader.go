package downloader

import (
  "errors"
  "../manager"
  "../constants"
  "github.com/knadh/go-get-youtube/youtube"
)

var invalidLinkErr = errors.New("Invalid link error!")

func DownloadVideoFromYoutube(link string) error {

  if !manager.IsLinkFromYoutube(link) {
    return invalidLinkErr
  }

  videoID, err := manager.GetVideoIDFromLink(link)

  if err != nil {
    return err
  }

  video, err := youtube.Get(videoID)

  if err != nil {
    return err
  }

  option := &youtube.Option{
    Rename: true,
    Resume: true,
    Mp3: true,
  }
  video.Download(0, constants.APP_DIR + "/video.mp3", option)
  return nil
}
