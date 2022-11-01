package axe

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

type AssetSystem struct {
	FormatMap map[AssetType]AssetFormat
	Formats   []AssetFormat
	Sources   []AssetSource
	Assets    map[string]*Asset
}

var _ GameSystem = &AssetSystem{}

type AssetType string

const (
	AssetTypeUnknown        AssetType = ""
	AssetTypeTexture        AssetType = "texture"
	AssetTypeAudio          AssetType = "audio"
	AssetTypeVertexShader   AssetType = "vertex-shader"
	AssetTypeFragmentShader AssetType = "fragment-shader"
	AssetTypeXml            AssetType = "xml"
	AssetTypeJson           AssetType = "json"
)

type AssetStatus struct {
	Progress float32
	Done     bool
	Error    error
}

func (status *AssetStatus) Reset() {
	status.Progress = 0
	status.Done = false
	status.Error = nil
}
func (status *AssetStatus) Fail(err error) {
	status.Done = true
	status.Error = err
}
func (status *AssetStatus) Success() {
	status.Done = true
	status.Progress = 1
}
func (status *AssetStatus) IsSuccess() bool {
	return status.Done && status.Error == nil
}

type AssetFormat interface {
	// Does this format handle refs like this one.
	Handles(ref AssetRef) bool
	// Native types this format handles.
	Types() []AssetType
	// Load the asset from the source. This could be done in any thread.
	Load(asset *Asset) error
	// Unload the asset, we don't need it anymore. This could be done in any thread.
	Unload(asset *Asset) error
	// The asset is loaded by we need to use it now. This is done in the main thread.
	Activate(asset *Asset) error
	// The asset is not needed right now, but don't unload it yet. This is done in the main thread.
	Deactivate(asset *Asset) error
}

type AssetSource interface {
	Handles(ref AssetRef) bool
	Read(ref AssetRef) (io.Reader, error)
}

type AssetRef struct {
	Name string
	URI  string
	Type AssetType
}

type Asset struct {
	Ref            AssetRef
	Source         AssetSource
	SourceReader   io.Reader
	Format         AssetFormat
	LoadStatus     AssetStatus
	Data           any
	ActivateStatus AssetStatus
}

func (a Asset) IsValid() bool {
	return a.Source != nil && a.Format != nil
}

func (a *Asset) Load() error {
	if !a.IsValid() {
		return fmt.Errorf("Asset %s (%s) must have a source and loader to be loaded.", a.Ref.Name, a.Ref.URI)
	}
	if err := a.LoadReader(); err != nil {
		return err
	}
	if err := a.LoadData(); err != nil {
		return nil
	}
	return nil
}

func (a *Asset) LoadReader() error {
	if a.SourceReader != nil {
		return nil
	}
	if a.Source == nil {
		return fmt.Errorf("Asset %s (%s) must have a source to be loaded.", a.Ref.Name, a.Ref.URI)
	}
	reader, err := a.Source.Read(a.Ref)
	if err != nil {
		return err
	}
	a.SourceReader = reader
	return nil
}

func (a *Asset) LoadData() error {
	if a.Data != nil {
		return nil
	}
	if a.Format == nil || a.SourceReader == nil {
		return fmt.Errorf("Asset %s (%s) must have a loader and source to be loaded.", a.Ref.Name, a.Ref.URI)
	}
	return a.Format.Load(a)
}

func (a *Asset) Activate() error {
	if a.ActivateStatus.IsSuccess() {
		return nil
	}
	err := a.Load()
	if err != nil {
		return err
	}
	return a.Format.Activate(a)
}

func NewAssetSystem() AssetSystem {
	return AssetSystem{
		FormatMap: make(map[AssetType]AssetFormat),
		Formats:   make([]AssetFormat, 0, 16),
		Sources:   make([]AssetSource, 0, 16),
		Assets:    make(map[string]*Asset, 128),
	}
}

func (assets *AssetSystem) AddLoader(loader AssetFormat) {
	assets.Formats = append(assets.Formats, loader)
	types := loader.Types()
	if types != nil && len(types) > 0 {
		for _, t := range types {
			assets.FormatMap[t] = loader
		}
	}
}

func (assets *AssetSystem) AddSource(source AssetSource) {
	assets.Sources = append(assets.Sources, source)
}

func (assets *AssetSystem) Add(ref AssetRef) *Asset {
	if loaded, ok := assets.Assets[ref.URI]; ok {
		return loaded
	}
	asset := &Asset{
		Ref: ref,
	}
	for _, source := range assets.Sources {
		if source.Handles(ref) {
			asset.Source = source
			break
		}
	}
	for _, format := range assets.Formats {
		if format.Handles(ref) {
			asset.Format = format
			break
		}
	}
	if asset.Format == nil && ref.Type != AssetTypeUnknown {
		if typeLoader, ok := assets.FormatMap[ref.Type]; ok {
			asset.Format = typeLoader
		}
	}
	assets.Assets[ref.URI] = asset
	return asset
}

func (assets *AssetSystem) AddMany(refs []AssetRef) []*Asset {
	many := make([]*Asset, len(refs))
	for i, ref := range refs {
		many[i] = assets.Add(ref)
	}
	return many
}

func (assets *AssetSystem) Init(game *Game) error {
	assets.AddSource(&LocalAssetSource{})
	assets.AddSource(&WebAssetSource{})
	assets.AddLoader(&XmlGenericAssetLoader{})
	assets.AddLoader(&JsonGenericAssetLoader{})
	return nil
}

func (assets *AssetSystem) Update(game *Game) {

}

func (assets *AssetSystem) Destroy() {
	for _, asset := range assets.Assets {
		if asset.ActivateStatus.IsSuccess() {
			asset.Format.Deactivate(asset) // ignore error for now
		}
		if asset.LoadStatus.IsSuccess() {
			asset.Format.Unload(asset) // ignore error for now
		}
		if asset.SourceReader != nil {
			if closer, ok := asset.SourceReader.(io.Closer); ok {
				closer.Close()
			}
		}
	}
}

type LocalAssetSource struct{}

var _ AssetSource = &LocalAssetSource{}
var localAssetSourceRegex, _ = regexp.Compile("^(/|\\./|[a-zA-Z]:)")

func (local *LocalAssetSource) Handles(ref AssetRef) bool {
	return localAssetSourceRegex.MatchString(ref.URI)
}
func (local *LocalAssetSource) Read(ref AssetRef) (io.Reader, error) {
	return os.Open(ref.URI)
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
