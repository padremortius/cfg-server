package common

import (
	"io"
	"net/http"

	"github.com/BurntSushi/toml"
	gojson "github.com/goccy/go-json"
	goyaml "github.com/goccy/go-yaml"
)

// StructToJSONBytes is ...
func StructToJSONBytes(v any) ([]byte, error) {
	res, err := gojson.Marshal(v)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func StructToYamlBytes(v any) ([]byte, error) {
	res, err := goyaml.Marshal(v)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func StructToTomlBytes(v any) ([]byte, error) {
	res, err := toml.Marshal(v)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetFileByURL(URL string) ([]byte, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil && err == nil {
			err = closeErr // Propagate the close error if no other error occurred
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func GetPubKey(URL string) (string, error) {
	rawBytes, err := GetFileByURL(URL)
	if err != nil {
		return "", err
	}
	var answer map[string]string
	if err = gojson.Unmarshal(rawBytes, &answer); err != nil {
		return "", err
	}
	return answer["value"], nil
}
