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
	"time"

	"github.com/kndrad/squil/cmd/internal/config"
	"github.com/kndrad/squil/cmd/internal/logging"
	"github.com/kndrad/squil/internal/shelter"
	"github.com/spf13/cobra"
)

// readDogCmd represents the read command.
var readDogCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a dog from a database",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.DefaultLogger()

		logger.Info("Connecting to database")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		cfg, err := config.LoadShelterConfig(config.DefaultEnvFilePath)
		if err != nil {
			logger.Error("Error", "err", err.Error())

			return fmt.Errorf("getting db config: %w", err)
		}
		pool, err := shelter.DatabasePool(ctx, *cfg, false)
		if err != nil {
			logger.Error("Error", "err", err.Error())

			return fmt.Errorf("getting db config: %w", err)
		}
		defer pool.Close()

		conn, err := shelter.Connection(ctx, pool)
		if err != nil {
			logger.Error("Error", "err", err.Error())

			return fmt.Errorf("getting db config: %w", err)
		}
		defer conn.Close(ctx)

		logger.Info("Database connection established")

		logger.Info("Reading a dog from a database",
			slog.String(DogName, "name"),
		)

		queries := shelter.New(conn)
		dog, err := queries.ReadDog(ctx, DogName)
		if err != nil {
			logger.Error("Reading dog from db", "err", err.Error())

			return fmt.Errorf("read dog from db: %w", err)
		}
		logger.Info("Read dog from a database",
			slog.Int64("id", dog.ID),
			slog.String("name", dog.Name),
			slog.String("breed", dog.Breed),
		)
		logger.Info("Program completed successfully.")

		return nil
	},
}

func init() {
	dogsCmd.AddCommand(readDogCmd)

	readDogCmd.Flags().StringVar(&DogName, "name", "", "dog's name")
	cobra.MarkFlagRequired(readDogCmd.Flags(), "name")
}
