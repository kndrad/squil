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

// deleteDogCmd represents the delete command.
var deleteDogCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes a dog from a shelter",
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

		logger.Info("Deleting a dog from a database",
			slog.String("name", DogName),
		)

		queries := shelter.New(conn)

		if err := queries.DeleteDog(ctx, DogName); err != nil {
			logger.Error("Deleting dog from db", "err", err.Error())

			return fmt.Errorf("delete dog from db: %w", err)
		}

		logger.Info("Deletied a dog in a database",
			slog.String("name", DogName),
		)
		logger.Info("Program completed successfully.")

		return nil
	},
}

func init() {
	dogsCmd.AddCommand(deleteDogCmd)

	deleteDogCmd.Flags().StringVar(&DogName, "name", "", "dog's name to be deleted")
	cobra.MarkFlagRequired(deleteDogCmd.Flags(), "name")
}
