package container

import (
	"flag"
	logger "github.com/br0-space/bot-logger"
	matcher "github.com/br0-space/bot-matcher"
	telegramclient "github.com/br0-space/bot-telegramclient"
	"github.com/br0-space/musicbot/interfaces"
	"github.com/br0-space/musicbot/pkg/config"
	"github.com/br0-space/musicbot/pkg/matchers/musiclinks"
	"github.com/br0-space/musicbot/pkg/songlink"
	"sync"
)

var configInstance *interfaces.ConfigStruct
var configLock = &sync.Mutex{}
var matcherRegistryInstance *matcher.Registry
var matcherRegistryLock = &sync.Mutex{}

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
	} else {
		return telegramclient.NewClient(
			ProvideConfig().Telegram,
		)
	}
}

func ProvideSonglinkService() interfaces.SonglinkServiceInterface {
	return songlink.MakeService()
}
