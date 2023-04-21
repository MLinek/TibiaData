package internal

import (
	"encoding/json"
	"os"
)

func ExportJsonToFile(v any, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = f.Write(bytes)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}

func ImportJsonFromFile(v any, path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	return nil
}
