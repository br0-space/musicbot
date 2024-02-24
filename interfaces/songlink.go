package interfaces

type SonglinkEntryInterface interface {
	ToMarkdown() string
}

type SonglinkServiceInterface interface {
	EntryForURL(url string) (SonglinkEntryInterface, error)
}
