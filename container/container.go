package container

import (
	"flag"
	"sync"

	logger "github.com/br0-space/bot-logger"
	matcher "github.com/br0-space/bot-matcher"
	telegramclient "github.com/br0-space/bot-telegramclient"
	"github.com/br0-space/musicbot/interfaces"
	"github.com/br0-space/musicbot/pkg/config"
	"github.com/br0-space/musicbot/pkg/matchers/musiclinks"
	"github.com/br0-space/musicbot/pkg/matchers/ping"
	"github.com/br0-space/musicbot/pkg/songlink"
)

var (
	configInstance          *interfaces.ConfigStruct
	configLock              = &sync.Mutex{}
	matcherRegistryInstance *matcher.Registry
	matcherRegistryLock     = &sync.Mutex{}
)

func runsAsTest() bool {
	return flag.Lookup("test.v") != nil
}

func ProvideLogger() logger.Interface {
	return logger.New()
}

func ProvideConfig() *interfaces.ConfigStruct {
	configLock.Lock()
	defer configLock.Unlock()

	if configInstance == nil {
		if runsAsTest() {
			configInstance = config.NewTestConfig()
		} else {
			configInstance = config.NewConfig()
		}
	}

	return configInstance
}

func ProvideMatchersRegistry() *matcher.Registry {
	matcherRegistryLock.Lock()
	defer matcherRegistryLock.Unlock()

	if matcherRegistryInstance == nil {
		matcherRegistryInstance = matcher.NewRegistry(
			ProvideLogger(),
			ProvideTelegramClient(),
		)
		matcherRegistryInstance.Register(musiclinks.MakeMatcher(ProvideSonglinkService()))
		matcherRegistryInstance.Register(ping.MakeMatcher())
	}

	return matcherRegistryInstance
}

func ProvideTelegramWebhookHandler() telegramclient.WebhookHandlerInterface {
	matchersRegistry := ProvideMatchersRegistry()

	return telegramclient.NewHandler(
		&ProvideConfig().Telegram,
		func(messageIn telegramclient.WebhookMessageStruct) {
			matchersRegistry.Process(messageIn)
		},
	)
}

func ProvideTelegramClient() telegramclient.ClientInterface {
	if runsAsTest() {
		return telegramclient.NewMockClient()
	}

	return telegramclient.NewClient(
		ProvideConfig().Telegram,
	)
}

func ProvideSonglinkService() interfaces.SonglinkServiceInterface {
	return songlink.MakeService()
}
