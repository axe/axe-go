package axe

import (
	"encoding/xml"
	"io"
	"regexp"
	"strings"
)

type XmlGenericAssetLoader struct{}

type XmlNode struct {
	Token    xml.Token
	Children []*XmlNode
}

var _ AssetFormat = &XmlGenericAssetLoader{}
var xmlGenericAssetLoaderRegex, _ = regexp.Compile(`\.xml$`)

func (loader *XmlGenericAssetLoader) Handles(ref AssetRef) bool {
	return xmlGenericAssetLoaderRegex.MatchString(ref.URI)
}

func (loader *XmlGenericAssetLoader) Types() []AssetType {
	return []AssetType{AssetTypeXml}
}

func (loader *XmlGenericAssetLoader) Load(asset *Asset) error {
	decoder := xml.NewDecoder(asset.SourceReader)

	root := make([]*XmlNode, 0)
	stack := make([]*XmlNode, 0, 16)

	asset.LoadStatus.Start()

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				asset.LoadStatus.Fail(err)
				return err
			}
		}
		if _, isEnd := token.(xml.EndElement); isEnd {
			stack = stack[:len(stack)-1]
			continue
		}
		node := &XmlNode{
			Token: xml.CopyToken(token),
		}
		if len(stack) == 0 {
			root = append(root, node)
		} else {
			last := stack[len(stack)-1]
			last.Children = append(last.Children, node)
		}
		if _, isStart := token.(xml.StartElement); isStart {
			node.Children = make([]*XmlNode, 0)
			stack = append(stack, node)
		}
	}

	asset.LoadStatus.Success()
	asset.Data = root

	return nil
}

func (loader *XmlGenericAssetLoader) Unload(asset *Asset) error {
	asset.LoadStatus.Reset()
	asset.Data = nil
	return nil
}

func (loader *XmlGenericAssetLoader) Activate(asset *Asset) error {
	asset.ActivateStatus.Success()
	return nil
}

func (loader *XmlGenericAssetLoader) Deactivate(asset *Asset) error {
	asset.ActivateStatus.Reset()
	return nil
}

type XmlAssetLoader[T any] struct {
	EmptyValue        T
	Suffix            string
	SuffixInsensitive bool
	Regex             *regexp.Regexp
	CustomTypes       []AssetType
}

var _ AssetFormat = &XmlAssetLoader[any]{}

func (loader *XmlAssetLoader[T]) Handles(ref AssetRef) bool {
	if loader.Suffix != "" {
		suffix := loader.Suffix
		uri := ref.URI
		if loader.SuffixInsensitive {
			suffix = strings.ToLower(suffix)
			uri = strings.ToLower(uri)
		}
		return strings.HasSuffix(uri, suffix)
	}
	if loader.Regex != nil {
		return loader.Regex.MatchString(ref.URI)
	}
	return false
}

func (loader *XmlAssetLoader[T]) Types() []AssetType {
	return loader.CustomTypes
}

func (loader *XmlAssetLoader[T]) Load(asset *Asset) error {
	copy := loader.EmptyValue
	decoder := xml.NewDecoder(asset.SourceReader)
	err := decoder.Decode(&copy)

	if err != nil {
		asset.LoadStatus.Fail(err)
	} else {
		asset.LoadStatus.Success()
		asset.Data = copy
	}
	return err
}

func (loader *XmlAssetLoader[T]) Unload(asset *Asset) error {
	asset.LoadStatus.Reset()
	asset.Data = nil
	return nil
}

func (loader *XmlAssetLoader[T]) Activate(asset *Asset) error {
	asset.ActivateStatus.Success()
	return nil
}

func (loader *XmlAssetLoader[T]) Deactivate(asset *Asset) error {
	asset.ActivateStatus.Reset()
	return nil
}
