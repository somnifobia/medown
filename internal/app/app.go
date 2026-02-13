package app

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/somnifobia/medown/internal/ytdl"
	"github.com/somnifobia/medown/internal/twitterdl"
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
		case strings.Contains(host, "twitter.com") || strings.Contains(host, "x.com"):
			return twitterdl.Download(rawURL, outputDir)
		default:
			return fmt.Errorf("unsupported URL: %s", host)
	}
}
