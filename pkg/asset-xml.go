package axe

import (
	"encoding/xml"
	"io"
	"regexp"
	"strings"
)

type XmlGenericAssetFormat struct{}

type XmlNode struct {
	Token    xml.Token
	Children []*XmlNode
}

var _ AssetFormat = &XmlGenericAssetFormat{}
var xmlGenericAssetLoaderRegex, _ = regexp.Compile(`\.xml$`)

func (format *XmlGenericAssetFormat) Handles(ref AssetRef) bool {
	return xmlGenericAssetLoaderRegex.MatchString(ref.URI)
}

func (format *XmlGenericAssetFormat) Types() []AssetType {
	return []AssetType{AssetTypeXml}
}

func (format *XmlGenericAssetFormat) Load(asset *Asset) error {
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

func (format *XmlGenericAssetFormat) Unload(asset *Asset) error {
	asset.LoadStatus.Reset()
	asset.Data = nil
	return nil
}

func (format *XmlGenericAssetFormat) Activate(asset *Asset) error {
	asset.ActivateStatus.Success()
	return nil
}

func (format *XmlGenericAssetFormat) Deactivate(asset *Asset) error {
	asset.ActivateStatus.Reset()
	return nil
}

type XmlAssetFormat[T any] struct {
	EmptyValue        T
	Suffix            string
	SuffixInsensitive bool
	Regex             *regexp.Regexp
	CustomTypes       []AssetType
}

var _ AssetFormat = &XmlAssetFormat[any]{}

func (format *XmlAssetFormat[T]) Handles(ref AssetRef) bool {
	if format.Suffix != "" {
		suffix := format.Suffix
		uri := ref.URI
		if format.SuffixInsensitive {
			suffix = strings.ToLower(suffix)
			uri = strings.ToLower(uri)
		}
		return strings.HasSuffix(uri, suffix)
	}
	if format.Regex != nil {
		return format.Regex.MatchString(ref.URI)
	}
	return false
}

func (format *XmlAssetFormat[T]) Types() []AssetType {
	return format.CustomTypes
}

func (format *XmlAssetFormat[T]) Load(asset *Asset) error {
	copy := format.EmptyValue
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

func (format *XmlAssetFormat[T]) Unload(asset *Asset) error {
	asset.LoadStatus.Reset()
	asset.Data = nil
	return nil
}

func (format *XmlAssetFormat[T]) Activate(asset *Asset) error {
	asset.ActivateStatus.Success()
	return nil
}

func (format *XmlAssetFormat[T]) Deactivate(asset *Asset) error {
	asset.ActivateStatus.Reset()
	return nil
}
