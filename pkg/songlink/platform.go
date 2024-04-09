package songlink

type Platform string

const (
	PlatformSonglink   Platform = "songlink"
	PlatformSpotify    Platform = "spotify"
	PlatformAppleMusic Platform = "appleMusic"
	PlatformYoutube    Platform = "youtube"
	PlatformBandcamp   Platform = "bandcamp"
	PlatformTidal      Platform = "tidal"
)

func (p Platform) Natural() string {
	switch p {
	case PlatformSonglink:
		return "Songlink"
	case PlatformSpotify:
		return "Spotify"
	case PlatformAppleMusic:
		return "Apple Music"
	case PlatformYoutube:
		return "YouTube"
	case PlatformBandcamp:
		return "Bandcamp"
	case PlatformTidal:
		return "Tidal"
	default:
		return "Unknown"
	}
}
