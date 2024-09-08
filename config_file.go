package config

import (
	"bytes"
	"errors"
	"io"
	"io/fs"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

// Read a ConfigFile by name from an [fs.FS].
func ReadConfigFile[T any](fsys fs.FS, name string) (T, error) {
	buffer, err := getBuffer(fsys, name)
	if err != nil {
		var t T
		return t, err
	}

	conf, err := getConfigFile[T](buffer)
	if err != nil {
		return conf, err
	}

	conf = GetConfigMap(conf)

	return conf, nil
}

func getBuffer(fsys fs.FS, name string) (*bytes.Buffer, error) {
	b, err := fs.ReadFile(fsys, name)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(b)
	if buf == nil {
		return nil, ErrNilReader
	}

	return buf, nil
}

var ErrNilReader = errors.New("nil Reader")

func getConfigFile[T any](r io.Reader) (T, error) {
	var configFile T

	decoder := yaml.NewDecoder(r)
	err := decoder.Decode(&configFile)

	return configFile, err
}

func LoadEnv(filenames ...string) error {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return nil
}

func GetEnvString(name string, defaultValues ...string) string {
	value := os.Getenv(name)
	if value != "" {
		return value
	}

	for _, newValue := range defaultValues {
		value = newValue
		if value != "" {
			break
		}
	}

	return value
}

func GetEnvInt[U ~int](name string, defaultValues ...*U) U {
	if i, ok := getOsInt(name); ok {
		return U(i)
	}

	for _, value := range defaultValues {
		if value != nil {
			return U(*value)
		}
	}

	return 0
}

func GetEnvInt64[U ~int64](name string, defaultValues ...*U) U {
	if i, ok := getOsInt(name); ok {
		return U(i)
	}

	for _, value := range defaultValues {
		if value != nil {
			return U(*value)
		}
	}

	return 0
}

func Nullable[U any](value U) *U {
	return &value
}

func getOsInt(name string) (int, bool) {
	var i int

	s := os.Getenv(name)
	if s == "" {
		return 0, false
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}

	return i, true
}
