package cmd

import (
	"os"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/somnifobia/media-downloader/internal/app"
)

var (
	outputDir string
)

func defaultOutputDir() string {
    home, err := os.UserHomeDir()
    if err != nil {
        return "."
    }
    return filepath.Join(home, "Videos")
}

var rootCmd = &cobra.Command{
	Use:   "medown [url...]",
	Short: "Download media files from URLs",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error{
		dir := outputDir
		if dir == "" {
			dir = defaultOutputDir()
		}

		for _, u := range args {
			if err := app.Download(u, dir); err != nil {
				fmt.Fprintf(os.Stderr, "Download Error %s: %v\n", u, err)
			} else {
				fmt.Printf("Download completed successfully\n")
			}
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&outputDir,
		"output",
		"o",
		"",
		"Output Directory (default: $HOME/Videos)")
}
