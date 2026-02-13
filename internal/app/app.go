package app

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/somnifobia/media-downloader/internal/ytdl"
)

func Download(rawURL, outputDir string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	host := strings.ToLower(u.Host)

	switch {
		case strings.Contains(host, "youtube") || strings.Contains(host, "youtu.be"):
			return ytdl.Download(rawURL, outputDir)
		default:
			return fmt.Errorf("unsupported URL: %s", host)
	}
}
