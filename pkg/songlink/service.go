package songlink

import (
	"regexp"

	"github.com/br0-space/musicbot/interfaces"
)

func Pattern() *regexp.Regexp {
	return regexp.MustCompile(`(https?://open\.spotify\.com/(album|track)/.+?|https?://music\.apple\.com/[a-z]{2}/album/.+?|https?://(www\.)?youtube\.com/(watch\?v=|playlist\?list=).+?|https?://youtu\.be/.+?|https?://.+?\.bandcamp\.com/(album|track)/.+?|https?://(listen\.)?tidal\.com/(browse/)?(album|track)/.+?)(\s|$)`)
}

type Service struct{}

func MakeService() interfaces.SonglinkServiceInterface {
	return Service{}
}

func (s Service) EntryForURL(url string) (interfaces.SonglinkEntryInterface, error) {
	response, err := songlinkResponse(url)
	if err != nil {
		return nil, err
	}

	entry := Entry{
		Type:   EntryType(response.Entities[response.EntityUniqueID].Type),
		Title:  response.Entities[response.EntityUniqueID].Title,
		Artist: response.Entities[response.EntityUniqueID].Artist,
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

	if response.Links.Tidal.URL != "" {
		entry.Links = append(entry.Links, EntryLink{
			Platform: PlatformTidal,
			URL:      response.Links.Tidal.URL,
		})
	}

	return &entry, nil
}
