package interfaces

type SonglinkEntryInterface interface {
	ToMarkdown() string
}

type SonglinkServiceInterface interface {
	EntryForUrl(url string) (SonglinkEntryInterface, error)
}
