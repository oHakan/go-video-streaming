package helpers

import "os"

func CreateNewStaticDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func IsDirectoryExists(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func GetCurrentPath() string {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	return dir
}
