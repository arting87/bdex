package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"bdex.in/bdex/bdex-ex-backend/sync"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "data sync",
	Long:  `sync from contract table to database table`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sync")
	},
}

var syncServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "sync serve",
	Long:  `sync serve from contract table to database table`,
	Run: func(cmd *cobra.Command, args []string) {

		InitLogDbSet()

		ctx, cancel := context.WithCancel(context.Background())

		done := make(chan bool)
		go func() {
			defer close(done)
			sync.SyncServe(ctx)
		}()

		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-c:
			cancel()
			<-done
		case <-done:
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	syncCmd.AddCommand(syncServeCmd)
}
