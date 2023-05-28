package songlink

import (
	"github.com/br0-space/musicbot/interfaces"
	"regexp"
)

func Pattern() *regexp.Regexp {
	return regexp.MustCompile(`(https?://open.spotify.com/(album|track)/.+?|https?://music.apple.com/[a-z]{2}/album/.+?|https?://(www.)?youtube.com/(watch\?v=|playlist\?list=).+?|https?://youtu.be/.+?|https?://.+?.bandcamp.com/(album|track)/.+?)(\s|$)`)
}

type Service struct{}

func MakeService() interfaces.SonglinkServiceInterface {
	return Service{}
}

func (s Service) EntryForUrl(url string) (interfaces.SonglinkEntryInterface, error) {
	response, err := songlinkResponse(url)
	if err != nil {
		return nil, err
	}

	entry := Entry{
		Type:   EntryType(response.Entities[response.EntityUniqueId].Type),
		Title:  response.Entities[response.EntityUniqueId].Title,
		Artist: response.Entities[response.EntityUniqueId].Artist,
		Links:  make([]EntryLink, 0),
	}

	// Now we add all links
	entry.Links = append(entry.Links, EntryLink{
		Platform: PlatformSonglink,
		URL:      response.PageURL,
	})
	if response.Links.Spotify.URL != "" {
		entry.Links = append(entry.Links, EntryLink{
			Platform: PlatformSpotify,
			URL:      response.Links.Spotify.URL,
		})
	}
	if response.Links.AppleMusic.URL != "" {
		entry.Links = append(entry.Links, EntryLink{
			Platform: PlatformAppleMusic,
			URL:      response.Links.AppleMusic.URL,
		})
	}
	if response.Links.Youtube.URL != "" {
		entry.Links = append(entry.Links, EntryLink{
			Platform: PlatformYoutube,
			URL:      response.Links.Youtube.URL,
		})
	}
	if response.Links.Bandcamp.URL != "" {
		entry.Links = append(entry.Links, EntryLink{
			Platform: PlatformBandcamp,
			URL:      response.Links.Bandcamp.URL,
		})
	}

	return &entry, nil
}
