package asset

import (
	"errors"
	"io"
	"io/fs"
	"net/http"
	"os"
	"regexp"
)

var relativeRegex = regexp.MustCompile(`[^/]+$`)
var previousRegex = regexp.MustCompile(`[^/]+/\.\./`)

type LocalAssetSource struct{}

var _ Source = LocalAssetSource{}
var localAssetSourceRegex = regexp.MustCompile(`^(/|\./|[a-zA-Z]:)`)

func (local LocalAssetSource) Handles(ref Ref) bool {
	return localAssetSourceRegex.MatchString(ref.URI)
}
func (local LocalAssetSource) Read(ref Ref) (io.Reader, error) {
	return os.Open(ref.URI)
}
func (local LocalAssetSource) Relative(uri string, relative string) string {
	uri = relativeRegex.ReplaceAllString(uri, relative)
	uri = previousRegex.ReplaceAllString(uri, "")
	return uri
}

type WebAssetSource struct{}

var _ Source = WebAssetSource{}
var webAssetSourceRegex = regexp.MustCompile("^https?:")

func (local WebAssetSource) Handles(ref Ref) bool {
	return webAssetSourceRegex.MatchString(ref.URI)
}
func (local WebAssetSource) Read(ref Ref) (io.Reader, error) {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(ref.URI)
	return resp.Body, err
}
func (local WebAssetSource) Relative(uri string, relative string) string {
	uri = relativeRegex.ReplaceAllString(uri, relative)
	uri = previousRegex.ReplaceAllString(uri, "")
	return uri
}

type EmbedAssetSource struct{}

var embedFiles fs.FS

func SetEmbedAssetRoot(root fs.FS) {
	embedFiles = root
}

var _ Source = EmbedAssetSource{}
var embedAssetSourceRegex = regexp.MustCompile("^embed:")

var ErrNoEmbeddedFileSystem = errors.New("no embedded files set via axe.SetEmbedAssetRoot")

func (local EmbedAssetSource) Handles(ref Ref) bool {
	return embedAssetSourceRegex.MatchString(ref.URI)
}
func (local EmbedAssetSource) Read(ref Ref) (io.Reader, error) {
	if embedFiles == nil {
		return nil, ErrNoEmbeddedFileSystem
	}
	return embedFiles.Open(ref.URI)
}
func (local EmbedAssetSource) Relative(uri string, relative string) string {
	uri = relativeRegex.ReplaceAllString(uri, relative)
	uri = previousRegex.ReplaceAllString(uri, "")
	return uri
}

type FileSystemAssetSource struct {
	Files          fs.FS
	HandlesPattern *regexp.Regexp
	CustomRelative func(uri string, relative string) string
}

var _ Source = FileSystemAssetSource{}

func (src FileSystemAssetSource) Handles(ref Ref) bool {
	return src.HandlesPattern.MatchString(ref.URI)
}
func (src FileSystemAssetSource) Read(ref Ref) (io.Reader, error) {
	return src.Files.Open(ref.URI)
}
func (src FileSystemAssetSource) Relative(uri string, relative string) string {
	if src.CustomRelative != nil {
		return src.CustomRelative(uri, relative)
	}
	uri = relativeRegex.ReplaceAllString(uri, relative)
	uri = previousRegex.ReplaceAllString(uri, "")
	return uri
}
