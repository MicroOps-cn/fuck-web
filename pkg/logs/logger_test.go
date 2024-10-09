package logs

import (
	"bytes"
	"testing"

	log "github.com/MicroOps-cn/fuck/log"
	"github.com/go-kit/log/level"
	"github.com/stretchr/testify/require"
)

func TestRegisterLogger(t *testing.T) {
	log.RegisterLogFormat(FormatFuck, NewFuckLogger)
	buf := bytes.NewBuffer(nil)
	l := log.New(log.WithWriter(buf), log.WithConfig(log.MustNewConfig("info", string(FormatFuck))))
	level.Error(l).Log("msg", "test message", WrapKeyName("Name"), "Test")
	const matchExpr = `(?m)^[-.\d:TZ]+ \[error] [-\w]+ \S+ - test message - \n\[Name]:\s+Test`
	require.Regexp(t, matchExpr, buf.String())
}
