package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/MicroOps-cn/fuck/log"
	"github.com/MicroOps-cn/fuck/signals"
	"github.com/go-kit/log/level"
	"github.com/gogo/protobuf/jsonpb"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/MicroOps-cn/fuck-web/config"
	"github.com/MicroOps-cn/fuck-web/pkg/server"
)

var (
	cfgFile       string
	configDisplay bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gateway",
	Short: "The fuck-web gateway server.",
	Long:  `The fuck-web gateway server.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		initConfig()
		if len(config.Get().Security.Secret) == 0 {
			return fmt.Errorf("`security.secret` cannot be empty")
		}
		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		logger := log.GetContextLogger(cmd.Context())
		ch := signals.SetupSignalHandler(logger)
		ctx, cancelFunc := context.WithCancel(cmd.Context())
		go func() {
			<-ch.Channel()
			cancelFunc()
		}()
		ctx = context.WithValue(ctx, "command", cmd.Use)
		s, err := server.New(ctx, logger, ch)
		if err != nil {
			return err
		}
		return s.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	server.AddCommand(rootCmd)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./fuck-web.yaml", "config file")
	rootCmd.PersistentFlags().BoolVar(&configDisplay, "config.display", false, "display config")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile == "" {
		cfgFile = "./fuck-web.yaml"
	}

	logger := log.NewTraceLogger()
	if err := config.ReloadConfigFromFile(logger, cfgFile); err != nil {
		level.Error(logger).Log("msg", "failed to load config", "err", err)
		os.Exit(1)
	}
	if configDisplay {
		var buf bytes.Buffer
		err := (&jsonpb.Marshaler{OrigName: true}).Marshal(&buf, config.Get())
		if err != nil {
			level.Error(logger).Log("msg", "failed to marshaller config", "err", err)
			os.Exit(1)
		}
		var tmpObj map[string]interface{}
		err = json.NewDecoder(&buf).Decode(&tmpObj)
		if err != nil {
			level.Error(logger).Log("msg", "failed to marshaller config", "err", err)
			os.Exit(1)
		}
		encoder := yaml.NewEncoder(os.Stdout)
		encoder.SetIndent(2)
		if err = encoder.Encode(tmpObj); err != nil {
			level.Error(logger).Log("msg", "failed to encode config", "err", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
}
