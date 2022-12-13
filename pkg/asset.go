package axe

import (
	"fmt"
	"io"
)

type AssetSystem struct {
	FormatMap     map[AssetType]AssetFormat
	Formats       []AssetFormat
	Sources       []AssetSource
	DefaultSource AssetSource
	Assets        map[string]*Asset
	NamedAssets   map[string]*Asset
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
	AssetTypeModel          AssetType = "model"
	AssetTypeMaterials      AssetType = "materials"
	AssetTypeMaterial       AssetType = "material"
)

type AssetStatus struct {
	Progress float32
	Started  bool
	Done     bool
	Error    error
}

func (status *AssetStatus) Reset() {
	status.Progress = 0
	status.Done = false
	status.Error = nil
	status.Started = false
}
func (status *AssetStatus) Start() {
	status.Reset()
	status.Started = true
}
func (status *AssetStatus) Fail(err error) {
	status.Done = true
	status.Started = true
	status.Error = err
}
func (status *AssetStatus) Success() {
	status.Done = true
	status.Started = true
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
	Relative(uri string, relative string) string
}

type AssetRef struct {
	Name string
	URI  string
	Type AssetType
}

func (ref AssetRef) UniqueName() string {
	if ref.Name != "" {
		return ref.Name
	}
	return ref.URI
}
func (ref AssetRef) String() string {
	if ref.Name != "" && ref.Name != ref.URI {
		return fmt.Sprintf("%s (%s)", ref.Name, ref.URI)
	} else {
		return ref.URI
	}
}

type Asset struct {
	Ref            AssetRef
	Source         AssetSource
	SourceReader   io.Reader
	Format         AssetFormat
	LoadStatus     AssetStatus
	Data           any
	ActivateStatus AssetStatus
	Dependent      []AssetRef
	Next           []AssetRef
}

func (a Asset) IsValid() bool {
	return a.Source != nil && a.Format != nil
}

func (a *Asset) Load() error {
	if !a.IsValid() {
		return fmt.Errorf("Asset %s must have a source and format to be loaded", a.Ref.String())
	}
	if err := a.LoadReader(); err != nil {
		return err
	}
	if err := a.LoadData(); err != nil {
		return nil
	}
	return nil
}

func (a *Asset) AddNext(relative string, dependent bool) {
	rel := a.Relative(relative)
	a.Next = append(a.Next, rel)
	if dependent {
		a.Dependent = append(a.Dependent, rel)
	}
}

func (a *Asset) Relative(relative string) AssetRef {
	return AssetRef{
		URI:  a.Source.Relative(a.Ref.URI, relative),
		Name: relative,
	}
}

func (a *Asset) LoadReader() error {
	if a.SourceReader != nil {
		return nil
	}
	if a.Source == nil {
		return fmt.Errorf("Asset %s must have a source to be loaded", a.Ref.String())
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
		return fmt.Errorf("Asset %s must have a format and source to be loaded", a.Ref.String())
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

func (a *Asset) Unload() error {
	if a.ActivateStatus.IsSuccess() {
		err := a.Format.Deactivate(a)
		if err != nil {
			return err
		}
	}
	if a.LoadStatus.IsSuccess() {
		err := a.Format.Unload(a)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewAssetSystem() AssetSystem {
	return AssetSystem{
		FormatMap:   make(map[AssetType]AssetFormat),
		Formats:     make([]AssetFormat, 0, 16),
		Sources:     make([]AssetSource, 0, 16),
		Assets:      make(map[string]*Asset, 128),
		NamedAssets: make(map[string]*Asset),
	}
}

func (assets *AssetSystem) AddFormat(format AssetFormat) {
	assets.Formats = append(assets.Formats, format)
	types := format.Types()
	if len(types) > 0 {
		for _, t := range types {
			assets.FormatMap[t] = format
		}
	}
}

func (assets *AssetSystem) AddSource(source AssetSource) {
	assets.Sources = append(assets.Sources, source)
}

func (assets *AssetSystem) Add(ref AssetRef) *Asset {
	if ref.Name != "" && ref.URI == "" {
		return assets.NamedAssets[ref.Name]
	}
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
	if asset.Source == nil {
		asset.Source = assets.DefaultSource
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
	if ref.Name != "" && assets.NamedAssets[ref.Name] == nil {
		assets.NamedAssets[ref.Name] = asset
	}
	return asset
}

func (assets *AssetSystem) AddMany(refs []AssetRef) []*Asset {
	many := make([]*Asset, len(refs))
	if len(refs) > 0 {
		for i, ref := range refs {
			many[i] = assets.Add(ref)
		}
	}
	return many
}

func (assets *AssetSystem) AddManyMap(refs []AssetRef) map[string]*Asset {
	many := make(map[string]*Asset)
	if len(refs) > 0 {
		for _, ref := range refs {
			asset := assets.Add(ref)
			many[asset.Ref.URI] = asset
		}
	}
	return many
}

func (assets *AssetSystem) Init(game *Game) error {
	assets.DefaultSource = &LocalAssetSource{}
	assets.AddSource(LocalAssetSource{})
	assets.AddSource(WebAssetSource{})
	assets.AddSource(EmbedAssetSource{})

	assets.AddFormat(&XmlGenericAssetFormat{})
	assets.AddFormat(&JsonGenericAssetFormat{})
	assets.AddFormat(&ObjFormat{})
	assets.AddFormat(&MtlFormat{})

	if len(game.Settings.Assets) > 0 {
		assets.AddMany(game.Settings.Assets)
	}

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

func (assets *AssetSystem) Get(uri string) *Asset {
	return assets.Assets[uri]
}

func (assets *AssetSystem) GetNamed(name string) *Asset {
	return assets.NamedAssets[name]
}

func (assets *AssetSystem) GetEither(nameOrURI string) *Asset {
	asset := assets.Assets[nameOrURI]
	if asset != nil {
		return asset
	}
	return assets.NamedAssets[nameOrURI]
}

func (assets *AssetSystem) GetRef(ref AssetRef) *Asset {
	if ref.URI != "" {
		return assets.Assets[ref.URI]
	}
	if ref.Name != "" {
		return assets.NamedAssets[ref.Name]
	}
	return nil
}
