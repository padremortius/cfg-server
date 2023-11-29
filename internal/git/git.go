package git

import (
	"os"
)

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
