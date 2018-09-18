package manager

import (
  "errors"
  "strings"
)

func GetVideoIDFromLink(link string) (string, error) {
  var err = errors.New("Invalid link!")

  parts := strings.Split(link, "v=")

  if len(parts) != 2 {
    return "", err
  }

  return parts[1], nil
}
