package utils

import (
	"fmt"
	"testing"

	jsonextract "github.com/blainemoser/JsonExtract"
)

func TestBaseDir(t *testing.T) {
	dir, err := BaseDir([]string{}, "goutils")
	if err != nil {
		t.Fatal(err)
	}
	if len(dir) < 1 {
		t.Errorf("expected directory, got %s", dir)
	}
}

func TestGetFileContents(t *testing.T) {
	js, err := testFile()
	if err != nil {
		t.Fatal(err)
	}
	if testInterface, err := js.Extract("test/file_content"); err == nil {
		if testStr, ok := testInterface.(string); ok {
			if testStr != "This is the result" {
				t.Errorf("expected test string in file '%s', got '%s'", "This is the result", testStr)
			}
		} else {
			t.Errorf("expected a test string, got %v", testInterface)
		}
	} else {
		t.Fatal(err)
	}
}

func TestInterfaceFunctions(t *testing.T) {
	js, err := testFile()
	if err != nil {
		t.Fatal(err)
	}
	stringInterface, _ := js.Extract("interfaces/string")
	if StringInterface(stringInterface) != "test string" {
		t.Errorf("expected string to be %s, got %s", "test string", StringInterface(stringInterface))
	}
	floatInterface, _ := js.Extract("interfaces/float64")
	if Float64Interface(floatInterface) != 1.12345 {
		t.Errorf("expected float to be %f, got %f", 1.12345, Float64Interface(floatInterface))
	}
	intInterface := createInt64Interface()
	if Int64Interface(intInterface) != 5 {
		t.Errorf("expected int to be %d, got %d", 5, Int64Interface(intInterface))
	}
}

func createInt64Interface() interface{} {
	convertThis := int64(5)
	return convertThis
}

func testFile() (jsonextract.JSONExtract, error) {
	path, err := BaseDir([]string{}, "goutils")
	if err != nil {
		return jsonextract.JSONExtract{}, err
	}
	content, err := GetFileContent(fmt.Sprintf("%s/%s", path, "test_file.json"))
	if err != nil {
		return jsonextract.JSONExtract{}, err
	}
	return jsonextract.JSONExtract{RawJSON: string(content)}, nil
}
