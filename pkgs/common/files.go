package common

import (
	"errors"
	"fmt"
	"os"

	goyaml "github.com/goccy/go-yaml"
)

func ReadFile(aFileName string) ([]byte, error) {
	return os.ReadFile(aFileName)
}

func InitDir(localPath string) error {
	_, err := os.Stat(localPath)
	if err != nil {
		if err := os.Mkdir(localPath, os.ModeDir+os.ModeAppend+os.ModePerm); err != nil {
			return err
		}
	} else {
		if err := os.RemoveAll(localPath); err != nil {
			return err
		}
		_, err = os.Stat(localPath)
		if err != nil {
			if err := os.Mkdir(localPath, os.ModeDir+os.ModeAppend+os.ModePerm); err != nil {
				return err
			}
		}
	}
	return nil
}

func DirExists(pathToDir string) (bool, error) {
	_, err := os.Stat(pathToDir)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

func FileExists(fName string) bool {
	res := true
	if _, err := os.Stat(fName); errors.Is(err, os.ErrNotExist) {
		res = false
	}
	return res
}

func GetDataFromFile(fName string) (res map[string]any, err error) {
	var rawData []byte
	if FileExists(fName) {
		if rawData, err = ReadFile(fName); err != nil {
			return res, fmt.Errorf("error reading file %v with error message: %v", fName, err)
		}
		if err = goyaml.Unmarshal(rawData, &res); err != nil {
			svcErr := fmt.Errorf("error unmarshalling file %v with error message: %v", fName, err)
			return res, svcErr
		}
	}
	return res, nil
}
