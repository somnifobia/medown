package ytdl

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func Download(videoURL, outputDir string) error {
	client := youtube.Client{}

	video, err := client.GetVideo(videoURL)
	if err != nil {
		return fmt.Errorf("Failed getting video info: %w", err)
	}

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
        formats := video.Formats.WithAudioChannels()
        if len(formats) == 0 {
            return fmt.Errorf("nenhum formato com áudio+vídeo encontrado")
        }
        formats.Sort()
        f := formats[len(formats)-1]
        format = &f
    }

	formats := video.Formats.WithAudioChannels()
	if len(formats) == 0 {
		return fmt.Errorf("no audio formats available")
	}

	stream, _, err := client.GetStream(video, format)
	if err != nil {
		return fmt.Errorf("Failed getting stream: %w", err)
	}
	defer stream.Close()

	title := sanitizeFilename(video.Title)
	filename := filepath.Join(outputDir, title+".mp4")

	if err := os.MkdirAll(filepath.Dir(filename), 0o755); err != nil {
		return fmt.Errorf("Failed creating output directory: %w", err)
	}

	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Failed creating output file: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, stream); err != nil {
		return fmt.Errorf("Failed writing to output file: %w", err)
	}

	fmt.Printf("Youtube '%s' -> %s\n", video.Title, filename)
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
