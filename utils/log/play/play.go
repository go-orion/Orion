package main

import (
	"context"
	"fmt"

	"github.com/go-orion/Orion/utils/errors"
	"github.com/go-orion/Orion/utils/errors/notifier"
	"github.com/go-orion/Orion/utils/log"
	"github.com/go-orion/Orion/utils/log/loggers"
)

func main() {
	ctx := context.Background()
	ctx = loggers.AddToLogContext(ctx, "hello", "world")
	logger := log.GetLogger()
	//logger := log.NewLogger(stdlog.NewLogger())
	logger.SetLevel(loggers.InfoLevel)
	logger.Error(ctx, "error")
	logger.Warn(ctx, "warning")
	logger.Info(ctx, "info")
	logger.Debug(ctx, "debug")
	logger.SetLevel(loggers.DebugLevel)
	logger.Debug(ctx, "debug2")
	log.Debug(ctx, "debug3")

	e := errors.New("hello world")
	notifier.Notify(e, ctx)

	e2 := fmt.Errorf("generic error")
	notifier.Notify(e2, ctx)
}
