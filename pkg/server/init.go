package server

import (
	"context"
	"fmt"

	logs "github.com/MicroOps-cn/fuck/log"
	"github.com/MicroOps-cn/fuck/signals"
	w "github.com/MicroOps-cn/fuck/wrapper"
	"github.com/go-kit/log/level"
	"github.com/spf13/cobra"

	"github.com/MicroOps-cn/fuck-web/pkg/endpoint"
	"github.com/MicroOps-cn/fuck-web/pkg/service"
)

var (
	adminUsername string
	skipInitData  bool
)

// migrateCmd represents the migrate command
var initDataCmd = &cobra.Command{
	Use:   "init",
	Short: "Data initialization tool",
	Long:  `The data initialization tool will create a table with missing columns and indexes. And create the required user and application data.`,

	Run: func(cmd *cobra.Command, args []string) {
		logger := logs.GetDefaultLogger()
		ctx := context.WithValue(cmd.Context(), "command", cmd.Use)
		InitData(ctx, signals.SetupSignalHandler(logger))
	},
}

func InitData(ctx context.Context, _ *signals.Handler) {
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
	if err := svc.InitData(ctx, adminUsername); err != nil {
		logger := logs.GetContextLogger(ctx)
		level.Error(logs.WithPrint(w.NewStringer(func() string {
			return fmt.Sprintf("%+v", err)
		}))(logger)).Log("msg", "failed to http request", "err", err)
		panic(err)
	}
}

func init() {
	initDataCmd.PersistentFlags().StringVar(&adminUsername, "admin", "admin", "admin username.")
	initDataCmd.PersistentFlags().BoolVar(&skipInitData, "skip-init-data", false, "Don't execute data init.")
}

func AddInitCommand(rootCmd *cobra.Command) {
	rootCmd.AddCommand(initDataCmd)
	initDataCmd.PreRun = rootCmd.PreRun
	initDataCmd.PreRunE = rootCmd.PreRunE
}
