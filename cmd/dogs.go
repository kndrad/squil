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
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/kndrad/squil/cmd/internal/config"
	"github.com/kndrad/squil/cmd/internal/logging"
	"github.com/kndrad/squil/internal/shelter"
	"github.com/spf13/cobra"
)

// dogsCmd represents the dogs command.
var dogsCmd = &cobra.Command{
	Use:   "dogs",
	Short: "Lists dogs from the database",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logging.DefaultLogger()

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

			return fmt.Errorf("db pool: %w", err)
		}
		defer pool.Close()

		logger.Info("Database connection established")

		conn, err := pgx.Connect(ctx, pool.Config().ConnString())
		if err != nil {
			logger.Error("Error", "err", err.Error())

			return fmt.Errorf("connecting db: %w", err)
		}
		defer conn.Close(ctx)

		if err := shelter.Ping(ctx, pool); err != nil {
			logger.Error("Error", "err", err.Error())

			return fmt.Errorf("pinging db: %w", err)
		}

		// Using sqlc queries to get all dogs from a database
		queries := shelter.New(conn)
		dogs, err := queries.AllDogs(ctx)
		if err != nil {
			logger.Error("Error", "err", err.Error())

			return fmt.Errorf("fetiching all dogs: %w", err)
		}
		logger.Info("Fetched dogs from a database", "len_dogs", len(dogs))
		for _, dog := range dogs {
			fmt.Println(dog)
		}
		logger.Info("Program completed successfully.")

		return nil
	},
}

func init() {
	rootCmd.AddCommand(dogsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dogsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dogsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
