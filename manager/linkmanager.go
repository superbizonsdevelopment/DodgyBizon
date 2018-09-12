package manager

import (
  "strings"
)

func GetVideoIDFromLink(link string) string {
  parts := strings.Split(link, "v=")
  return parts[1]
}
