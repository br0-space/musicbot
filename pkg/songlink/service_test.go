package songlink_test

import (
	"testing"

	"github.com/br0-space/musicbot/pkg/songlink"
	"github.com/smirzaei/parallel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type songlinkTest struct {
	in  string
	out *songlink.Entry
}

var tests = []songlinkTest{
	{
		in: "https://open.spotify.com/track/0Q5IOvNoREy7gzT0CWmayo?si=d2b1a4b4ae204358",
		out: &songlink.Entry{
			Type:   songlink.Song,
			Title:  "By 1899, The Age Of Outlaws And Gunslingers Was At An End",
			Artist: "Jeff Silverman, Luke O’Malley, Woody Jackson",
			Links: []songlink.EntryLink{
				{songlink.PlatformSonglink, "https://song.link/s/0Q5IOvNoREy7gzT0CWmayo"},
				{songlink.PlatformSpotify, "https://open.spotify.com/track/0Q5IOvNoREy7gzT0CWmayo"},
				{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1472283462?i=1472283463&mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
				{songlink.PlatformYoutube, "https://www.youtube.com/watch?v=atmy8uAI8K0"},
			},
		},
	},
	{
		in: "https://open.spotify.com/album/2Gbv0Wjtwn9zQYMvWtTHnK",
		out: &songlink.Entry{
			Type:   songlink.Album,
			Title:  "The Music of Red Dead Redemption 2 (Original Score)",
			Artist: "Various Artists",
			Links: []songlink.EntryLink{
				{songlink.PlatformSonglink, "https://album.link/s/2Gbv0Wjtwn9zQYMvWtTHnK"},
				{songlink.PlatformSpotify, "https://open.spotify.com/album/2Gbv0Wjtwn9zQYMvWtTHnK"},
				{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1472283462?mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
				{songlink.PlatformYoutube, "https://www.youtube.com/playlist?list=OLAK5uy_myT6DLJmO1jsviiIR4li7oyaHXWpyIVWo"},
			},
		},
	},
	{
		in: "https://music.apple.com/de/album/by-1899-the-age-of-outlaws-and-gunslingers-was-at-an-end/1472283462?i=1472283463",
		out: &songlink.Entry{
			Type:   songlink.Song,
			Title:  "By 1899, The Age of Outlaws and Gunslingers Was At an End",
			Artist: "Jeff Silverman, Luke O'Malley & Woody Jackson",
			Links: []songlink.EntryLink{
				{songlink.PlatformSonglink, "https://song.link/de/i/1472283463"},
				{songlink.PlatformSpotify, "https://open.spotify.com/track/0Q5IOvNoREy7gzT0CWmayo"},
				{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1472283462?i=1472283463&mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
				{songlink.PlatformYoutube, "https://www.youtube.com/watch?v=atmy8uAI8K0"},
			},
		},
	},
	{
		in: "https://music.apple.com/de/album/hi/1140071785?i=1140071869&l=en",
		out: &songlink.Entry{
			Type:   songlink.Song,
			Title:  "Hi!",
			Artist: "Metrik",
			Links: []songlink.EntryLink{
				{songlink.PlatformSonglink, "https://song.link/de/i/1140071869"},
				{songlink.PlatformSpotify, "https://open.spotify.com/track/6pRgr64gnVjL2tHj2zXpfY"},
				{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1686060156?i=1686060159&mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
				{songlink.PlatformYoutube, "https://www.youtube.com/watch?v=Q3enDjbwXWc"},
			},
		},
	},
	{
		in: "https://music.apple.com/de/album/the-music-of-red-dead-redemption-2-original-score/1472283462",
		out: &songlink.Entry{
			Type:   songlink.Album,
			Title:  "The Music of Red Dead Redemption 2 (Original Score)",
			Artist: "Verschiedene Interpreten",
			Links: []songlink.EntryLink{
				{songlink.PlatformSonglink, "https://album.link/de/i/1472283462"},
				{songlink.PlatformSpotify, "https://open.spotify.com/album/2Gbv0Wjtwn9zQYMvWtTHnK"},
				{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1472283462?mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
				{songlink.PlatformYoutube, "https://www.youtube.com/playlist?list=PLasqSXwCIDXcSbf1CrOrSDirCh8bywoCy"},
			},
		},
	},
	{
		in: "https://music.apple.com/de/album/life-thrills/1140071785?l=en",
		out: &songlink.Entry{
			Type:   songlink.Album,
			Title:  "LIFE/THRILLS",
			Artist: "Metrik",
			Links: []songlink.EntryLink{
				{songlink.PlatformSonglink, "https://album.link/de/i/1140071785"},
				{songlink.PlatformSpotify, "https://open.spotify.com/album/4mYRuvOJ4uWHm5G94pTQw9"},
				{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1686060156?mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
				{songlink.PlatformYoutube, "https://www.youtube.com/playlist?list=OLAK5uy_n3vkQMwLHzd3vClPzPEU9Oiy7COOwA89I"},
				{songlink.PlatformBandcamp, "https://metrikmusic.bandcamp.com/album/life-thrills"},
			},
		},
	},
	{
		in: "https://www.youtube.com/watch?v=Q3enDjbwXWc",
		out: &songlink.Entry{
			Type:   songlink.Song,
			Title:  "Hi!",
			Artist: "Metrik - Topic",
			Links: []songlink.EntryLink{
				{songlink.PlatformSonglink, "https://song.link/y/Q3enDjbwXWc"},
				{songlink.PlatformSpotify, "https://open.spotify.com/track/6pRgr64gnVjL2tHj2zXpfY"},
				{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1686060156?i=1686060159&mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
				{songlink.PlatformYoutube, "https://www.youtube.com/watch?v=Q3enDjbwXWc"},
			},
		},
	},
	{
		in: "https://www.youtube.com/playlist?list=OLAK5uy_n3vkQMwLHzd3vClPzPEU9Oiy7COOwA89I",
		out: &songlink.Entry{
			Type:   songlink.Album,
			Title:  "Album - LIFE/THRILLS",
			Artist: "Metrik",
			Links: []songlink.EntryLink{
				{songlink.PlatformSonglink, "https://album.link/y/OLAK5uy_n3vkQMwLHzd3vClPzPEU9Oiy7COOwA89I"},
				{songlink.PlatformSpotify, "https://open.spotify.com/album/4mYRuvOJ4uWHm5G94pTQw9"},
				{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1686060156?mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
				{songlink.PlatformYoutube, "https://www.youtube.com/playlist?list=OLAK5uy_n3vkQMwLHzd3vClPzPEU9Oiy7COOwA89I"},
				{songlink.PlatformBandcamp, "https://metrikmusic.bandcamp.com/album/life-thrills"},
			},
		},
	},
	{
		in: "https://metrikmusic.bandcamp.com/album/life-thrills",
		out: &songlink.Entry{
			Type:   songlink.Album,
			Title:  "LIFE/THRILLS",
			Artist: "Metrik",
			Links: []songlink.EntryLink{
				{songlink.PlatformSonglink, "https://album.link/b/3803701310"},
				{songlink.PlatformSpotify, "https://open.spotify.com/album/4mYRuvOJ4uWHm5G94pTQw9"},
				{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1686060156?mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
				{songlink.PlatformYoutube, "https://www.youtube.com/playlist?list=OLAK5uy_n3vkQMwLHzd3vClPzPEU9Oiy7COOwA89I"},
				{songlink.PlatformBandcamp, "https://metrikmusic.bandcamp.com/album/life-thrills"},
			},
		},
	},
}

func TestGetSonglinkEntry(t *testing.T) {
	t.Parallel()

	parallel.ForEach(tests, func(tt songlinkTest) {
		entry, err := songlink.MakeService().EntryForURL(tt.in)
		require.NoError(t, err)
		assert.NotNil(t, entry)
		assert.Equal(t, tt.out, entry)
	})
}
