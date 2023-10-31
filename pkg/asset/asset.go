package asset

import (
	"fmt"
	"io"
)

type Type string

const (
	TypeUnknown        Type = ""
	TypeTexture        Type = "texture"
	TypeAudio          Type = "audio"
	TypeVertexShader   Type = "vertex-shader"
	TypeFragmentShader Type = "fragment-shader"
	TypeXml            Type = "xml"
	TypeJson           Type = "json"
	TypeModel          Type = "model"
	TypeMaterials      Type = "materials"
	TypeMaterial       Type = "material"
	TypeFontBitmap     Type = "font-bitmap"
)

type Format interface {
	// Does this format handle refs like this one.
	Handles(ref Ref) bool
	// Native types this format handles.
	Types() []Type
	// Load the asset from the source. This could be done in any thread.
	Load(asset *Asset) error
	// Unload the asset, we don't need it anymore. This could be done in any thread.
	Unload(asset *Asset) error
	// The asset is loaded by we need to use it now. This is done in the main thread.
	Activate(asset *Asset) error
	// The asset is not needed right now, but don't unload it yet. This is done in the main thread.
	Deactivate(asset *Asset) error
}

type Source interface {
	Handles(ref Ref) bool
	Read(ref Ref) (io.Reader, error)
	Relative(uri string, relative string) string
}

type Ref struct {
	Name    string
	URI     string
	Type    Type
	Options any
}

func (ref Ref) UniqueName() string {
	if ref.Name != "" {
		return ref.Name
	}
	return ref.URI
}
func (ref Ref) String() string {
	if ref.Name != "" && ref.Name != ref.URI {
		return fmt.Sprintf("%s (%s)", ref.Name, ref.URI)
	} else {
		return ref.URI
	}
}

type Asset struct {
	Ref            Ref
	Source         Source
	SourceReader   io.Reader
	Format         Format
	LoadStatus     Status
	Data           any
	ActivateStatus Status
	Dependent      []Ref
	Next           []Ref
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

func (a *Asset) Relative(relative string) Ref {
	return Ref{
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
