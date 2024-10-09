package cmd

import (
	"context"
	"fmt"
	"io"
	"net/http/httptest"
	"os"

	logs "github.com/MicroOps-cn/fuck/log"
	"github.com/MicroOps-cn/fuck/signals"
	"github.com/spf13/cobra"

	"github.com/MicroOps-cn/fuck-web/pkg/endpoint"
	"github.com/MicroOps-cn/fuck-web/pkg/global"
	"github.com/MicroOps-cn/fuck-web/pkg/transport"
)

var outputFile string

// migrateCmd represents the migrate command
var rootCmd = &cobra.Command{
	Use:   "openapi",
	Short: "OpenAPI generator",
	Long:  `OpenAPI generator`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logs.GetDefaultLogger()
		ch := signals.SetupSignalHandler(logger)
		ctx, cancelFunc := context.WithCancel(cmd.Context())
		go func() {
			<-ch.Channel()
			cancelFunc()
		}()
		ctx = context.WithValue(ctx, global.HTTPWebPrefixKey, "/")

		handler := transport.NewHTTPHandler(ctx, logger, endpoint.Set{}, "/apidocs.json")

		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/apidocs.json", nil)
		handler.ServeHTTP(w, req)
		if w.Body != nil {
			if len(outputFile) == 0 {
				_, _ = io.Copy(os.Stdout, w.Body)
			} else if err := os.WriteFile(outputFile, w.Body.Bytes(), 0o600); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFile, "output", "o", "", "Output openAPI to the specified file")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
