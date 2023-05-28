package interfaces

import telegramclient "github.com/br0-space/bot-telegramclient"

type ConfigStruct struct {
	Verbose  bool `mapstructure:"verbose"`
	Quiet    bool `mapstructure:"quiet"`
	Server   ServerConfigStruct
	Telegram telegramclient.ConfigStruct
}

type ServerConfigStruct struct {
	ListenAddr string
}

type MatcherConfigStruct struct {
	Enabled bool
}
