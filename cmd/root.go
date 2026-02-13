package cmd

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/somnifobia/internal/app"
)

var (
	outputDir string
)

var rootCmd = &cobra.Command{
	Use:   "medown [url...]",
	Short: "Download media files from URLs",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error{
		for _, u := range args {
			if err := app.Download(u, outputDir); err != nil {
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
	rootCmd.PersistentFlags().StringVarP(&outputDir, "output", "o", ".", "Output Directory")
}
