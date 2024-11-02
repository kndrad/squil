/*
Copyright © 2024 Konrad Nowara

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/kndrad/squil/cmd/internal/config"
	"github.com/kndrad/squil/internal/shelter"
	"github.com/spf13/cobra"
)

var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
}

// pingdbCmd represents the pingdb command.
var pingdbCmd = &cobra.Command{
	Use:   "ping",
	Short: "Pings a database",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Info("Command pingdb called")

		cfg, err := config.LoadShelterConfig(config.DefaultEnvFilePath)
		if err != nil {
			logger.Error("Error", "err", err.Error())

			return fmt.Errorf("loading shelter db config: %w", err)
		}

		logger.Info("Establishing connection to DB.")

		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()

		ping := true
		pool, err := shelter.DatabasePool(ctx, *cfg, ping)
		if err != nil {
			logger.Error("Error", "err", err.Error())

			return fmt.Errorf("database pool: %w", err)
		}
		defer pool.Close()

		logger.Info("DB Connection established.")

		logger.Info("Pinging DB.")
		if err := shelter.Ping(ctx, pool); err != nil {
			logger.Error("Error", "err", err.Error())

			return fmt.Errorf("ping: %w", err)
		}

		logger.Info("Ping ok.")
		logger.Info("Program compeleted successfully.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(pingdbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingdbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingdbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
