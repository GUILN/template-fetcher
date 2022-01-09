package helpers

import "os"

func CheckPathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateFile(fContent []byte, fPath string) error {
	f, err := os.Create(fPath)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(fContent)
	if err != nil {
		return err
	}

	return nil
}
