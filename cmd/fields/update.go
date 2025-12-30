package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hashicorp/go-version"
)

var firmwareUpdateApi = "https://fw-update.ubnt.com/api/firmware-latest"

const (
	debianPlatform         = "debian"
	releaseChannel         = "release"
	unifiControllerProduct = "unifi-controller"
)

type firmwareUpdateApiResponse struct {
	Embedded firmwareUpdateApiResponseEmbedded `json:"_embedded"`
}

type firmwareUpdateApiResponseEmbedded struct {
	Firmware []firmwareUpdateApiResponseEmbeddedFirmware `json:"firmware"`
}

type firmwareUpdateApiResponseEmbeddedFirmware struct {
	Channel  string                                         `json:"channel"`
	Created  string                                         `json:"created"`
	Id       string                                         `json:"id"`
	Platform string                                         `json:"platform"`
	Product  string                                         `json:"product"`
	Version  *version.Version                               `json:"version"`
	Links    firmwareUpdateApiResponseEmbeddedFirmwareLinks `json:"_links"`
}

type firmwareUpdateApiResponseEmbeddedFirmwareDataLink struct {
	Href *url.URL `json:"href"`
}

func (l firmwareUpdateApiResponseEmbeddedFirmwareDataLink) MarshalJSON() ([]byte, error) {
	var href string
	if l.Href != nil {
		href = l.Href.String()
	}

	aux := struct {
		Href string `json:"href"`
	}{
		Href: href,
	}

	return json.Marshal(aux)
}

func (l *firmwareUpdateApiResponseEmbeddedFirmwareDataLink) UnmarshalJSON(j []byte) error {
	var m map[string]any

	err := json.Unmarshal(j, &m)
	if err != nil {
		return err
	}

	if href := m["href"]; href != nil {
		shref, ok := href.(string)
		if !ok {
			return fmt.Errorf("expected href to be a string, got %T", href)
		}
		url, err := url.Parse(shref)
		if err != nil {
			return err
		}

		l.Href = url
	}

	return nil
}

type firmwareUpdateApiResponseEmbeddedFirmwareLinks struct {
	Data firmwareUpdateApiResponseEmbeddedFirmwareDataLink `json:"data"`
}

func firmwareUpdateApiFilter(operator, key string, value any) string {
	return fmt.Sprintf("%s~~%s~~%v", operator, key, value)
}

func latestUnifiVersion() (*version.Version, *url.URL, error) {
	url, err := url.Parse(firmwareUpdateApi)
	if err != nil {
		return nil, nil, err
	}

	query := url.Query()
	query.Add("filter", firmwareUpdateApiFilter("eq", "channel", releaseChannel))
	query.Add("filter", firmwareUpdateApiFilter("eq", "product", unifiControllerProduct))
	query.Add("filter", firmwareUpdateApiFilter("eq", "version_major", "9"))
	url.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var respData firmwareUpdateApiResponse
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, nil, err
	}

	for _, firmware := range respData.Embedded.Firmware {
		if firmware.Platform != debianPlatform {
			continue
		}

		return firmware.Version.Core(), firmware.Links.Data.Href, nil
	}

	return nil, nil, nil
}