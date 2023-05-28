package songlink

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type apiResponse struct {
	EntityUniqueId string                       `json:"entityUniqueId"`
	PageURL        string                       `json:"pageUrl"`
	Entities       map[string]apiResponseEntity `json:"entitiesByUniqueId"`
	Links          apiResponseLinks             `json:"linksByPlatform"`
}

type apiResponseEntity struct {
	Type   string `json:"type"`
	Title  string `json:"title"`
	Artist string `json:"artistName"`
}

type apiResponseLinks struct {
	Spotify    apiResponseLink
	AppleMusic apiResponseLink
	Youtube    apiResponseLink
	SoundCloud apiResponseLink
	Bandcamp   apiResponseLink
}

type apiResponseLink struct {
	URL            string `json:"url"`
	EntityUniqueId string `json:"entityUniqueId"`
}

// Example URLs:
//
// Spotify album: https://api.song.link/v1-alpha.1/links?platform=spotify&type=album&id=2Gbv0Wjtwn9zQYMvWtTHnK
// Spotify track: https://api.song.link/v1-alpha.1/links?platform=spotify&type=song&id=0Q5IOvNoREy7gzT0CWmayo
// Apple Music album: https://api.song.link/v1-alpha.1/links?platform=appleMusic&type=album&id=1472283462
// Apple Music track: https://api.song.link/v1-alpha.1/links?platform=appleMusic&type=song&id=1472283463

func songlinkResponse(queryURL string) (*apiResponse, error) {
	url := fmt.Sprintf(
		"https://api.song.link/v1-alpha.1/links?url=%s&userCountry=DE&songIfSingle=true",
		queryURL,
	)

	request, _ := http.NewRequest("GET", url, nil)
	response, _ := http.DefaultClient.Do(request)

	responseBody := &apiResponse{}
	if err := json.NewDecoder(response.Body).Decode(responseBody); err != nil {
		return nil, err
	}

	return responseBody, nil
}
