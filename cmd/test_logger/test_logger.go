package main

import (
	"github.com/br0-space/musicbot/container"
	"github.com/spf13/pflag"
)

func main() {
	pflag.Parse()

	logger := container.ProvideLogger()

	logger.Debug("debug")
	logger.Info("info")
	logger.Warning("warning")
	logger.Error("error")
}
