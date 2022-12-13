package axe

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

var _ AssetSource = LocalAssetSource{}
var localAssetSourceRegex = regexp.MustCompile(`^(/|\./|[a-zA-Z]:)`)

func (local LocalAssetSource) Handles(ref AssetRef) bool {
	return localAssetSourceRegex.MatchString(ref.URI)
}
func (local LocalAssetSource) Read(ref AssetRef) (io.Reader, error) {
	return os.Open(ref.URI)
}
func (local LocalAssetSource) Relative(uri string, relative string) string {
	uri = relativeRegex.ReplaceAllString(uri, relative)
	uri = previousRegex.ReplaceAllString(uri, "")
	return uri
}

type WebAssetSource struct{}

var _ AssetSource = WebAssetSource{}
var webAssetSourceRegex = regexp.MustCompile("^https?:")

func (local WebAssetSource) Handles(ref AssetRef) bool {
	return webAssetSourceRegex.MatchString(ref.URI)
}
func (local WebAssetSource) Read(ref AssetRef) (io.Reader, error) {
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

var _ AssetSource = EmbedAssetSource{}
var embedAssetSourceRegex = regexp.MustCompile("^embed:")

var ErrNoEmbeddedFileSystem = errors.New("no embedded files set via axe.SetEmbedAssetRoot")

func (local EmbedAssetSource) Handles(ref AssetRef) bool {
	return embedAssetSourceRegex.MatchString(ref.URI)
}
func (local EmbedAssetSource) Read(ref AssetRef) (io.Reader, error) {
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

var _ AssetSource = FileSystemAssetSource{}

func (src FileSystemAssetSource) Handles(ref AssetRef) bool {
	return src.HandlesPattern.MatchString(ref.URI)
}
func (src FileSystemAssetSource) Read(ref AssetRef) (io.Reader, error) {
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
