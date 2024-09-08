package config

import (
	"bytes"
	"testing"
	"testing/fstest"
)

type configFile struct {
	DbString string `yaml:"dbString"`
}

func TestConfigFile(t *testing.T) {
	var data = `
dbString: mongodb://mongodb:27017
`

	var fileName = "file.yml"

	var fileSystem = fstest.MapFS{
		fileName: {Data: []byte(data)},
	}

	t.Run("should return error on nil data", func(t *testing.T) {
		t.Skip()
		_, err := getConfigFile[configFile](nil)
		if err == nil {
			t.Errorf("err should not be nil. err: %v", err)
		}
	})

	t.Run("should read from a reader", func(t *testing.T) {
		r := bytes.NewBufferString(data)
		_, err := getConfigFile[configFile](r)
		if err != nil {
			t.Errorf("%v", err)
		}
	})

	t.Run("should read data", func(t *testing.T) {
		r := bytes.NewBufferString(data)
		c, err := getConfigFile[configFile](r)
		if err != nil {
			t.Errorf("%v", err)
		}
		if c.DbString == "" {
			t.Errorf("DB_STRING is empty")
		}
	})

	t.Run("should read from a file", func(t *testing.T) {
		_, err := ReadConfigFile[configFile](fileSystem, fileName)
		if err != nil {
			t.Errorf("%v", err)
		}
	})
}
