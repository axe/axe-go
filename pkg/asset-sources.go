package axe

import (
	"io"
	"net/http"
	"os"
	"regexp"
)

var relativeRegex, _ = regexp.Compile(`[^/]+$`)
var previousRegex, _ = regexp.Compile(`[^/]+/\.\./`)

type LocalAssetSource struct{}

var _ AssetSource = &LocalAssetSource{}
var localAssetSourceRegex, _ = regexp.Compile(`^(/|\./|[a-zA-Z]:)`)

func (local *LocalAssetSource) Handles(ref AssetRef) bool {
	return localAssetSourceRegex.MatchString(ref.URI)
}
func (local *LocalAssetSource) Read(ref AssetRef) (io.Reader, error) {
	return os.Open(ref.URI)
}
func (local *LocalAssetSource) Relative(uri string, relative string) string {
	uri = relativeRegex.ReplaceAllString(uri, relative)
	uri = previousRegex.ReplaceAllString(uri, "")
	return uri
}

type WebAssetSource struct{}

var _ AssetSource = &WebAssetSource{}
var webAssetSourceRegex, _ = regexp.Compile("^https?:")

func (local *WebAssetSource) Handles(ref AssetRef) bool {
	return webAssetSourceRegex.MatchString(ref.URI)
}
func (local *WebAssetSource) Read(ref AssetRef) (io.Reader, error) {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(ref.URI)
	return resp.Body, err
}
func (local *WebAssetSource) Relative(uri string, relative string) string {
	uri = relativeRegex.ReplaceAllString(uri, relative)
	uri = previousRegex.ReplaceAllString(uri, "")
	return uri
}
