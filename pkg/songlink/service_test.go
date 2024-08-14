package songlink_test

import (
	"testing"

	"github.com/br0-space/musicbot/pkg/songlink"
	"github.com/smirzaei/parallel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type patternTest struct {
	in  string
	out bool
}

type getSonglinkEntryTest struct {
	in  string
	out *songlink.Entry
}

var patternTests = []patternTest{
	{in: "https://www.example.com/foo/bar", out: false},
	{in: "https://open.spotify.com/track/0Q5IOvNoREy7gzT0CWmayo?si=d2b1a4b4ae204358", out: true},
	{in: "https://open.spotify.com/album/2Gbv0Wjtwn9zQYMvWtTHnK", out: true},
	{in: "https://music.apple.com/de/album/by-1899-the-age-of-outlaws-and-gunslingers-was-at-an-end/1472283462?i=1472283463", out: true},
	{in: "https://music.apple.com/de/album/hi/1140071785?i=1140071869&l=en", out: true},
	{in: "https://music.apple.com/de/album/the-music-of-red-dead-redemption-2-original-score/1472283462", out: true},
	{in: "https://music.apple.com/de/album/life-thrills/1140071785?l=en", out: true},
	{in: "https://www.youtube.com/watch?v=Q3enDjbwXWc", out: true},
	{in: "https://www.youtube.com/playlist?list=OLAK5uy_n3vkQMwLHzd3vClPzPEU9Oiy7COOwA89I", out: true},
	{in: "https://metrikmusic.bandcamp.com/album/life-thrills", out: true},
	{in: "https://listen.tidal.com/album/112998896", out: true},
	{in: "https://listen.tidal.com/track/291949037", out: true},
	{in: "https://tidal.com/browse/album/112998896?u", out: true},
	{in: "https://tidal.com/browse/track/291949037?u", out: true},
}

var getSonglinkEntryTests = []getSonglinkEntryTest{
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
				{songlink.PlatformTidal, "https://listen.tidal.com/track/112998897"},
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
				{songlink.PlatformTidal, "https://listen.tidal.com/album/112998896"},
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
				{songlink.PlatformTidal, "https://listen.tidal.com/track/112998897"},
			},
		},
	},
	//{
	//	in: "https://music.apple.com/de/album/hi/1140071785?i=1140071869&l=en",
	//	out: &songlink.Entry{
	//		Type:   songlink.Song,
	//		Title:  "Hi!",
	//		Artist: "Metrik",
	//		Links: []songlink.EntryLink{
	//			{songlink.PlatformSonglink, "https://song.link/de/i/1140071869"},
	//			{songlink.PlatformSpotify, "https://open.spotify.com/track/6pRgr64gnVjL2tHj2zXpfY"},
	//			{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1686060156?i=1686060159&mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
	//			{songlink.PlatformYoutube, "https://www.youtube.com/watch?v=Q3enDjbwXWc"},
	//			{songlink.PlatformTidal, "https://listen.tidal.com/track/291949037"},
	//		},
	//	},
	//},
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
				{songlink.PlatformTidal, "https://listen.tidal.com/album/112998896"},
			},
		},
	},
	//{
	//	in: "https://music.apple.com/de/album/life-thrills/1140071785?l=en",
	//	out: &songlink.Entry{
	//		Type:   songlink.Album,
	//		Title:  "LIFE/THRILLS",
	//		Artist: "Metrik",
	//		Links: []songlink.EntryLink{
	//			{songlink.PlatformSonglink, "https://album.link/de/i/1140071785"},
	//			{songlink.PlatformSpotify, "https://open.spotify.com/album/4mYRuvOJ4uWHm5G94pTQw9"},
	//			{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1686060156?mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
	//			{songlink.PlatformYoutube, "https://www.youtube.com/playlist?list=OLAK5uy_n3vkQMwLHzd3vClPzPEU9Oiy7COOwA89I"},
	//			{songlink.PlatformBandcamp, "https://metrikmusic.bandcamp.com/album/life-thrills"},
	//			{songlink.PlatformTidal, "https://listen.tidal.com/album/291949036"},
	//		},
	//	},
	//},
	//{
	//	in: "https://www.youtube.com/watch?v=Q3enDjbwXWc",
	//	out: &songlink.Entry{
	//		Type:   songlink.Song,
	//		Title:  "Hi!",
	//		Artist: "Metrik - Topic",
	//		Links: []songlink.EntryLink{
	//			{songlink.PlatformSonglink, "https://song.link/y/Q3enDjbwXWc"},
	//			{songlink.PlatformSpotify, "https://open.spotify.com/track/6pRgr64gnVjL2tHj2zXpfY"},
	//			{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1686060156?i=1686060159&mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
	//			{songlink.PlatformYoutube, "https://www.youtube.com/watch?v=Q3enDjbwXWc"},
	//			{songlink.PlatformTidal, "https://listen.tidal.com/track/291949037"},
	//		},
	//	},
	//},
	//{
	//	in: "https://www.youtube.com/playlist?list=OLAK5uy_n3vkQMwLHzd3vClPzPEU9Oiy7COOwA89I",
	//	out: &songlink.Entry{
	//		Type:   songlink.Album,
	//		Title:  "Album - LIFE/THRILLS",
	//		Artist: "Metrik",
	//		Links: []songlink.EntryLink{
	//			{songlink.PlatformSonglink, "https://album.link/y/OLAK5uy_n3vkQMwLHzd3vClPzPEU9Oiy7COOwA89I"},
	//			{songlink.PlatformSpotify, "https://open.spotify.com/album/4mYRuvOJ4uWHm5G94pTQw9"},
	//			{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1686060156?mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
	//			{songlink.PlatformYoutube, "https://www.youtube.com/playlist?list=OLAK5uy_n3vkQMwLHzd3vClPzPEU9Oiy7COOwA89I"},
	//			{songlink.PlatformBandcamp, "https://metrikmusic.bandcamp.com/album/life-thrills"},
	//			{songlink.PlatformTidal, "https://listen.tidal.com/album/291949036"},
	//		},
	//	},
	//},
	//{
	//	in: "https://metrikmusic.bandcamp.com/album/life-thrills",
	//	out: &songlink.Entry{
	//		Type:   songlink.Album,
	//		Title:  "LIFE/THRILLS",
	//		Artist: "Metrik",
	//		Links: []songlink.EntryLink{
	//			{songlink.PlatformSonglink, "https://album.link/b/3803701310"},
	//			{songlink.PlatformSpotify, "https://open.spotify.com/album/4mYRuvOJ4uWHm5G94pTQw9"},
	//			{songlink.PlatformAppleMusic, "https://geo.music.apple.com/de/album/_/1686060156?mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m"},
	//			{songlink.PlatformYoutube, "https://www.youtube.com/playlist?list=OLAK5uy_n3vkQMwLHzd3vClPzPEU9Oiy7COOwA89I"},
	//			{songlink.PlatformBandcamp, "https://metrikmusic.bandcamp.com/album/life-thrills"},
	//			{songlink.PlatformTidal, "https://listen.tidal.com/album/291949036"},
	//		},
	//	},
	//},
}

func TestPattern(t *testing.T) {
	t.Parallel()

	parallel.ForEach(patternTests, func(tt patternTest) {
		assert.Equal(t, tt.out, songlink.Pattern().MatchString(tt.in), tt.in)
	})
}

func TestGetSonglinkEntry(t *testing.T) {
	t.Parallel()

	parallel.ForEach(getSonglinkEntryTests, func(tt getSonglinkEntryTest) {
		entry, err := songlink.MakeService().EntryForURL(tt.in)
		require.NoError(t, err)
		assert.NotNil(t, entry)
		assert.Equal(t, tt.out, entry)
	})
}
