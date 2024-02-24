package songlink

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type apiResponse struct {
	EntityUniqueID string                       `json:"entityUniqueId"`
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
	EntityUniqueID string `json:"entityUniqueId"`
}

func newAPIResponse() *apiResponse {
	apiResponse := makeAPIResponse()

	return &apiResponse
}

func makeAPIResponse() apiResponse {
	return apiResponse{
		EntityUniqueID: "",
		PageURL:        "",
		Entities:       map[string]apiResponseEntity{},
		Links:          makeAPIResponseLinks(),
	}
}

func makeAPIResponseLinks() apiResponseLinks {
	return apiResponseLinks{
		Spotify:    makeAPIResponseLink(),
		AppleMusic: makeAPIResponseLink(),
		Youtube:    makeAPIResponseLink(),
		SoundCloud: makeAPIResponseLink(),
		Bandcamp:   makeAPIResponseLink(),
	}
}

func makeAPIResponseLink() apiResponseLink {
	return apiResponseLink{
		URL:            "",
		EntityUniqueID: "",
	}
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

	request, _ := http.NewRequest(http.MethodGet, url, nil)
	response, _ := http.DefaultClient.Do(request) //nolint:bodyclose

	responseBody := newAPIResponse()
	if err := json.NewDecoder(response.Body).Decode(responseBody); err != nil { //nolint:musttag
		return nil, err
	}

	return responseBody, nil
}
