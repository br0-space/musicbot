package songlink

import (
	"fmt"
	"strings"

	telegramclient "github.com/br0-space/bot-telegramclient"
)

const minLinksPerEntry = 3

type Entry struct {
	Type   EntryType
	Title  string
	Artist string
	Links  []EntryLink
}

type EntryLink struct {
	Platform Platform
	URL      string
}

func (e Entry) ToMarkdown() string {
	if len(e.Links) < minLinksPerEntry {
		return ""
	}

	text := fmt.Sprintf(
		"*%s*\n*%s* · %s\n\n",
		telegramclient.EscapeMarkdown(e.Title),
		telegramclient.EscapeMarkdown(e.Artist),
		e.Type.Natural(),
	)

	var textSb35 strings.Builder
	for i := range e.Links {
		if e.Links[i].Platform == PlatformSonglink {
			continue
		}

		textSb35.WriteString(fmt.Sprintf(
			"🎧 [%s](%s)\n\n",
			e.Links[i].Platform.Natural(),
			e.Links[i].URL,
		))
	}
	text += textSb35.String()

	text += fmt.Sprintf(
		"🔗 [%s](%s)",
		PlatformSonglink.Natural(),
		e.Links[0].URL,
	)

	return text
}
