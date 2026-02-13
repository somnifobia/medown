package ytdl

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kkdai/youtube/v2"
	"github.com/schollz/progressbar/v3"
)

func Download(videoURL, outputDir string) error {
	client := youtube.Client{}

	video, err := client.GetVideo(videoURL)
	if err != nil {
		return fmt.Errorf("failed getting video info: %w", err)
	}

	// Prefer muxed formats (video + audio)
	muxed := video.Formats.Type("video/mp4").
		Select(func(f youtube.Format) bool {
			return f.AudioQuality != ""
		})

	var format *youtube.Format

	if len(muxed) > 0 {
		muxed.Sort()
		f := muxed[len(muxed)-1]
		format = &f
	} else {
		// Fallback: any format with audio channels
		formats := video.Formats.WithAudioChannels()
		if len(formats) == 0 {
			return fmt.Errorf("no video+audio formats found")
		}
		formats.Sort()
		f := formats[len(formats)-1]
		format = &f
	}

	stream, _, err := client.GetStream(video, format)
	if err != nil {
		return fmt.Errorf("failed getting stream: %w", err)
	}
	defer stream.Close()

	title := sanitizeFilename(video.Title)
	filename := filepath.Join(outputDir, title+".mp4")

	if err := os.MkdirAll(filepath.Dir(filename), 0o755); err != nil {
		return fmt.Errorf("failed creating output directory: %w", err)
	}

	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed creating output file: %w", err)
	}
	defer out.Close()

	// Progress bar
	size := format.ContentLength
	if size <= 0 {
		// unknown size, progressbar will still work but without total
		size = -1
	}

	bar := progressbar.DefaultBytes(
		size,
		fmt.Sprintf("Downloading %s", title),
	)

	writer := io.MultiWriter(out, bar)

	if _, err := io.Copy(writer, stream); err != nil {
		return fmt.Errorf("failed writing to output file: %w", err)
	}

	fmt.Printf("\nYouTube '%s' -> %s (itag %d, %s)\n",
		video.Title, filename, format.ItagNo, format.QualityLabel)

	return nil
}

var invalidChars = regexp.MustCompile(`[<>:"/\\|?*]+`)

func sanitizeFilename(name string) string {
	name = strings.TrimSpace(name)
	name = invalidChars.ReplaceAllString(name, "_")
	if len(name) == 0 {
		return "video"
	}
	if len(name) > 128 {
		name = name[:128]
	}
	return name
}
