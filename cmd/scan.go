package cmd

import (
	"context"

	"github.com/navidrome/navidrome/log"
	"github.com/spf13/cobra"
)

var fullRescan bool
var folder string

func init() {
	scanCmd.Flags().StringVar(&folder, "folder", "", "folder to scan")
	scanCmd.Flags().BoolVarP(&fullRescan, "full", "f", false, "check all subfolders, ignoring timestamps")
	rootCmd.AddCommand(scanCmd)
}

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan music folder",
	Long:  "Scan music folder for updates",
	Run: func(cmd *cobra.Command, args []string) {
		runScanner()
	},
}

func runScanner() {
	if folder == "" {
		log.Fatal("folder must not be empty")
	}

	scanner := GetScanner()

	err := scanner.RescanFolder(context.Background(), folder)
	if err != nil {
		log.Error(err)
	}
}
