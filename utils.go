package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	jsonextract "github.com/blainemoser/JsonExtract"
)

// BaseDir gets the base directory
func BaseDir(remove []string, projectRoot string) (string, error) {
	o, err := exec.Command("pwd").CombinedOutput()
	if err != nil {
		return "", err
	}
	pwd := strings.Trim(string(o), "\n")
	if remove != nil {
		for _, rem := range remove {
			if strings.Contains(pwd, fmt.Sprintf("%s/%s", projectRoot, rem)) {
				return strings.Replace(pwd, fmt.Sprintf("/%s", rem), "", 1), nil
			}
		}
	}
	return pwd, nil
}

// GetFileContent parses a file and retrieves its content
func GetFileContent(name string) ([]byte, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, os.ModeDevice)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	b := make([]byte, stat.Size())
	_, err = file.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// FileConfigs returns a jsonextract from a JSON file
func FileConfigs(path string) (jsonextract.JSONExtract, error) {
	b, err := GetFileContent(path)
	if err != nil {
		return jsonextract.JSONExtract{}, err
	}
	js := string(b)
	conf := jsonextract.JSONExtract{
		RawJSON: js,
	}
	return conf, nil
}

// Float64Interface checks if interface is a float64 and returns it if so
func Float64Interface(input interface{}) float64 {
	if input == nil {
		return 0
	}
	if result, ok := input.(float64); ok {
		return result
	}
	return 0
}

// Int64Interface checks if interface is an int64 and returns it if so
func Int64Interface(input interface{}) int64 {
	if input == nil {
		return 0
	}
	if result, ok := input.(int64); ok {
		return result
	}
	return 0
}

// StringInterface checks that an interface is a string and returns it if so
func StringInterface(input interface{}) string {
	if input == nil {
		return ""
	}
	if result, ok := input.(string); ok {
		return result
	}
	return ""
}
