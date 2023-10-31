package asset

import "io"

type Assets struct {
	FormatMap     map[Type]Format
	Formats       []Format
	Sources       []Source
	DefaultSource Source
	Assets        map[string]*Asset
	NamedAssets   map[string]*Asset
}

func NewAssets() Assets {
	return Assets{
		FormatMap:   make(map[Type]Format),
		Formats:     make([]Format, 0, 16),
		Sources:     make([]Source, 0, 16),
		Assets:      make(map[string]*Asset, 128),
		NamedAssets: make(map[string]*Asset),
	}
}

func (assets *Assets) AddFormat(format Format) {
	assets.Formats = append(assets.Formats, format)
	types := format.Types()
	if len(types) > 0 {
		for _, t := range types {
			assets.FormatMap[t] = format
		}
	}
}

func (assets *Assets) AddSource(source Source) {
	assets.Sources = append(assets.Sources, source)
}

func (assets *Assets) Add(ref Ref) *Asset {
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
	if asset.Format == nil && ref.Type != TypeUnknown {
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

func (assets *Assets) AddMany(refs []Ref) []*Asset {
	many := make([]*Asset, len(refs))
	if len(refs) > 0 {
		for i, ref := range refs {
			many[i] = assets.Add(ref)
		}
	}
	return many
}

func (assets *Assets) AddManyMap(refs []Ref) map[string]*Asset {
	many := make(map[string]*Asset)
	if len(refs) > 0 {
		for _, ref := range refs {
			asset := assets.Add(ref)
			many[asset.Ref.URI] = asset
		}
	}
	return many
}

func (assets *Assets) AddDefaults() {
	assets.DefaultSource = &LocalAssetSource{}
	assets.AddSource(LocalAssetSource{})
	assets.AddSource(WebAssetSource{})
	assets.AddSource(EmbedAssetSource{})

	assets.AddFormat(&XmlGenericAssetFormat{})
	assets.AddFormat(&JsonGenericAssetFormat{})
}

func (assets *Assets) Destroy() {
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

func (assets *Assets) Get(uri string) *Asset {
	return assets.Assets[uri]
}

func (assets *Assets) GetNamed(name string) *Asset {
	return assets.NamedAssets[name]
}

func (assets *Assets) GetEither(nameOrURI string) *Asset {
	asset := assets.Assets[nameOrURI]
	if asset != nil {
		return asset
	}
	return assets.NamedAssets[nameOrURI]
}

func (assets *Assets) GetRef(ref Ref) *Asset {
	if ref.URI != "" {
		return assets.Assets[ref.URI]
	}
	if ref.Name != "" {
		return assets.NamedAssets[ref.Name]
	}
	return nil
}
