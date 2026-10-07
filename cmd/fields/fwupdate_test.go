package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLatestUnifiVersionNoDebianMatch(t *testing.T) {
	assert := assert.New(t)

	fwVersion, err := version.NewVersion("7.3.83+atag-7.3.83-19645")
	require.NoError(t, err)

	respData := firmwareUpdateApiResponse{
		Embedded: firmwareUpdateApiResponseEmbedded{
			Firmware: []firmwareUpdateApiResponseEmbeddedFirmware{
				{Channel: releaseChannel, Platform: "windows", Product: unifiControllerProduct, Version: fwVersion},
				{Channel: releaseChannel, Platform: "document", Product: unifiControllerProduct, Version: fwVersion},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		resp, err := json.Marshal(respData)
		assert.NoError(err)
		_, err = rw.Write(resp)
		assert.NoError(err)
	}))
	defer server.Close()

	firmwareUpdateApi = server.URL
	defer func() { firmwareUpdateApi = "https://fw-update.ubnt.com/api/firmware-latest" }()

	gotVersion, gotDownload, err := latestUnifiVersion()
	assert.NoError(err)
	assert.Nil(gotVersion)
	assert.Nil(gotDownload)
}

func TestLatestUnifiVersionURLParseError(t *testing.T) {
	firmwareUpdateApi = "http://example.com/%zz"
	defer func() { firmwareUpdateApi = "https://fw-update.ubnt.com/api/firmware-latest" }()

	_, _, err := latestUnifiVersion()
	assert.Error(t, err)
}

func TestLatestUnifiVersionRequestError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}))
	server.Close() // closed before use so the request is refused.

	firmwareUpdateApi = server.URL
	defer func() { firmwareUpdateApi = "https://fw-update.ubnt.com/api/firmware-latest" }()

	_, _, err := latestUnifiVersion()
	assert.Error(t, err)
}

func TestLatestUnifiVersionDecodeError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, _ *http.Request) {
		_, _ = rw.Write([]byte("not json"))
	}))
	defer server.Close()

	firmwareUpdateApi = server.URL
	defer func() { firmwareUpdateApi = "https://fw-update.ubnt.com/api/firmware-latest" }()

	_, _, err := latestUnifiVersion()
	assert.Error(t, err)
}

func TestFirmwareLinkUnmarshalJSON(t *testing.T) {
	var l firmwareUpdateApiResponseEmbeddedFirmwareDataLink

	require.NoError(t, l.UnmarshalJSON([]byte(`{"href": "https://example.com/x.deb"}`)))
	require.NotNil(t, l.Href)
	assert.Equal(t, "https://example.com/x.deb", l.Href.String())
}

func TestFirmwareLinkUnmarshalJSONMissingHref(t *testing.T) {
	var l firmwareUpdateApiResponseEmbeddedFirmwareDataLink
	require.NoError(t, l.UnmarshalJSON([]byte(`{}`)))
	assert.Nil(t, l.Href)
}

func TestFirmwareLinkUnmarshalJSONInvalidHref(t *testing.T) {
	var l firmwareUpdateApiResponseEmbeddedFirmwareDataLink
	err := l.UnmarshalJSON([]byte(`{"href": "http://example.com/%zz"}`))
	assert.Error(t, err)
}

func TestFirmwareLinkUnmarshalJSONInvalidJSON(t *testing.T) {
	var l firmwareUpdateApiResponseEmbeddedFirmwareDataLink
	err := l.UnmarshalJSON([]byte(`not json`))
	assert.Error(t, err)
}

func TestFirmwareLinkMarshalJSON(t *testing.T) {
	l := firmwareUpdateApiResponseEmbeddedFirmwareDataLink{}
	data, err := l.MarshalJSON()
	require.NoError(t, err)
	assert.JSONEq(t, `{"href": ""}`, string(data))
}
