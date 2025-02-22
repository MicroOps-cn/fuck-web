/*
 Copyright © 2022 MicroOps-cn.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package server

import (
	"context"

	logs "github.com/MicroOps-cn/fuck/log"
	"github.com/MicroOps-cn/fuck/signals"
	"github.com/spf13/cobra"

	"github.com/MicroOps-cn/fuck-web/pkg/endpoint"
	"github.com/MicroOps-cn/fuck-web/pkg/service"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Automatic migration tool",
	Long:  `Automatic migration only creates tables, lacks columns and indexes, and does not change the type of existing columns or delete unused columns to protect data.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logs.GetDefaultLogger()
		ctx := context.WithValue(cmd.Context(), "command", cmd.Use)
		Migrate(ctx, signals.SetupSignalHandler(logger))
	},
}

func Migrate(ctx context.Context, _ *signals.Handler) {
	svc, err := service.New(ctx)
	if err != nil {
		panic(err)
	}
	if err := svc.AutoMigrate(ctx); err != nil {
		panic(err)
	}
	if err := svc.RegisterPermission(ctx, endpoint.GetPermissions()); err != nil {
		panic(err)
	}
}

func AddMigrateCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.PreRun = rootCmd.PreRun
	migrateCmd.PreRunE = rootCmd.PreRunE
}
