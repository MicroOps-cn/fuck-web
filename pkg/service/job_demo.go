package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-kit/log/level"

	"github.com/MicroOps-cn/fuck/errors"
	logs "github.com/MicroOps-cn/fuck/log"
)

type JobDemoParams struct {
	Timeout int `json:"timeout"`
}

func init() {
	RegisterCronJobs(JobEntry{
		Name: "demo",
		Handler: func(ctx context.Context, params string) (err error) {
			logger := logs.GetContextLogger(ctx)
			var demoParams JobDemoParams
			err = json.Unmarshal([]byte(params), &demoParams)
			if err != nil {
				return err
			}
			for i := 0; i < demoParams.Timeout; i++ {
				select {
				case <-ctx.Done():
					return errors.NewError(500, "user cancel", "500")
				default:
					time.Sleep(time.Second)
				}
				fmt.Println("runing")
				level.Debug(logger).Log("msg", "running...", "params", params)
			}
			return nil
		},
	})
}
