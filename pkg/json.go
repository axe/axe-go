package axe

import (
	"encoding/json"
	"io"
	"regexp"
)

type JsonGenericAssetLoader struct{}
type JsonValueKind int

const (
	JsonValueKindUnknown JsonValueKind = iota
	JsonValueKindNull
	JsonValueKindNumber
	JsonValueKindString
	JsonValueKindBoolean
	JsonValueKindObject
	JsonValueKindArray
)

type JsonArray []*JsonValue
type JsonObject map[string]*JsonValue

type JsonValue struct {
	Kind   JsonValueKind
	Value  any
	Parent *JsonValue
}

var _ AssetFormat = &JsonGenericAssetLoader{}
var jsonGenericAssetLoaderRegex, _ = regexp.Compile("\\.json$")

func (loader *JsonGenericAssetLoader) Handles(ref AssetRef) bool {
	return xmlGenericAssetLoaderRegex.MatchString(ref.URI)
}
func (loader *JsonGenericAssetLoader) Types() []AssetType {
	return []AssetType{AssetTypeJson}
}
func (loader *JsonGenericAssetLoader) Load(asset *Asset) error {
	decoder := json.NewDecoder(asset.SourceReader)

	var setValue func(out *JsonValue) error
	setValue = func(out *JsonValue) error {
		for {
			token, err := decoder.Token()
			if err != nil {
				return err
			}
			switch tt := token.(type) {
			case json.Number:
				out.Kind = JsonValueKindNumber
				out.Value, err = tt.Float64()
				return err
			case string:
				out.Kind = JsonValueKindString
				out.Value = tt
				return nil
			case float64:
				out.Kind = JsonValueKindNumber
				out.Value = tt
				return nil
			case nil:
				out.Kind = JsonValueKindNull
				out.Value = nil
				return nil
			case bool:
				out.Kind = JsonValueKindBoolean
				out.Value = tt
				return nil
			case json.Delim:
				switch tt {
				case '[':
					out.Kind = JsonValueKindArray
					jsonArray := JsonArray{}

					for {
						value := JsonValue{}
						err := setValue(&value)
						if err != nil {
							return err
						}
						if value.Kind == JsonValueKindUnknown {
							out.Value = jsonArray
							return nil
						}
						jsonArray = append(jsonArray, &value)
					}
				case '{':
					out.Kind = JsonValueKindObject
					jsonObject := JsonObject{}

					for {
						property, err := decoder.Token()
						if err != nil {
							return err
						}
						if delim, isEnd := property.(json.Delim); isEnd && delim == '}' {
							out.Value = jsonObject
							return nil
						}
						value := JsonValue{}
						err = setValue(&value)
						if err != nil {
							return err
						}
						jsonObject[property.(string)] = &value
					}

				default:
					return nil
				}
			default:
				return nil
			}
		}
	}

	asset.LoadStatus.Start()

	root := JsonValue{}
	err := setValue(&root)

	if err == io.EOF {
		err = nil
	}

	if err != nil {
		asset.LoadStatus.Fail(err)
	} else {
		asset.LoadStatus.Success()
		asset.Data = root
	}

	return err
}
func (loader *JsonGenericAssetLoader) Unload(asset *Asset) error {
	asset.LoadStatus.Reset()
	return nil
}
func (loader *JsonGenericAssetLoader) Activate(asset *Asset) error {
	asset.ActivateStatus.Success()
	return nil
}
func (loader *JsonGenericAssetLoader) Deactivate(asset *Asset) error {
	asset.ActivateStatus.Reset()
	return nil
}
