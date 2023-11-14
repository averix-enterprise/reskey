package backend

import (
	"encoding/json"
	"os"
)

func LoadOrCreateFile(path string, value interface{}, defaultValue interface{}) {
	if _, err := os.Stat(path); err == nil {
		data, err := os.ReadFile(path)
		checkFileError(err, "Could not read file: "+path)
		err = json.Unmarshal(data, value)
		checkFileError(err, "Could not parse json")
	} else if os.IsNotExist(err) {
		data, err := json.Marshal(defaultValue)
		checkFileError(err, "Could not create defaultValue json data")
		file, err := os.Create(path)
		checkFileError(err, "Could not create default file")
		stringData := string(data)
		_, err = file.WriteString(stringData)
		checkFileError(err, "Could not write file")
		err = file.Close()
		checkFileError(err, "Could not close file")
	}
}

func WriteToFile(path string, value interface{}) {
	if _, err := os.Stat(path); err == nil {
		file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
		checkFileError(err, "Cannot open file: "+path)
		data, err := json.Marshal(value)
		checkFileError(err, "Cannot convert object to json: "+path)
		_, err = file.WriteString(string(data))
		checkFileError(err, "Cannot write string to file: "+path)
		err = file.Close()
		checkFileError(err, "Cannot close file: "+path)
	}
}

func checkFileError(err interface{}, message string) {
	if err != nil {
		panic(message)
	}
}
