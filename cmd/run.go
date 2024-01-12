package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pdstuber/gameboy-emulator/internal/emulator"
	"github.com/spf13/cobra"
)

var (
	debug bool
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run pathToBootRom pathToBootRom pathToRom",
	Short: "Run the emulator with the provided rom(s).",
	Args:  cobra.RangeArgs(0, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var (
			bootRomFilePath string
			romFilePath     string
		)

		if len(args) >= 1 {
			bootRomFilePath = args[0]
		}

		if len(args) >= 2 {
			romFilePath = args[1]
		}
		config := emulator.NewConfig(debug, bootRomFilePath, romFilePath)

		emulator, err := emulator.New(config)
		if err != nil {
			log.Panic(err)
		}
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		go func() {
			if err := emulator.Start(ctx); err != nil {
				log.Fatalf("starting emulator: %v\n", err)
			}
		}()

		<-ctx.Done()
		emulator.Stop()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Toggle debug mode.")
}
