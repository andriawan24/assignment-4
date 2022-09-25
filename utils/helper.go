package utils

import "os"

func GetFile(filename string) ([]byte, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return file, err
	}
	return file, nil
}

func SaveFile(outputPath string, data []byte) error {
	err := os.WriteFile("raw/stats.json", data, 0644)
	if err != nil {
		return err
	}

	return nil
}
