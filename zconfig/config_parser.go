package zconfig

import (
	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
	"strconv"
	"strings"
)

var configReader *ini.File
var NotInitError = errors.New("config not init")

func Load(configFile string) error {
	conf, err := ini.Load(configFile)
	if err != nil {
		configReader = nil
		return errors.Wrap(err, "Load config error")
	}
	configReader = conf
	return nil
}

func GetString(section, key string) string {
	if configReader == nil {
		return ""
	}
	return configReader.Section(section).Key(key).String()
}

func GetInt32(section, key string) int32 {
	if configReader == nil {
		return 0
	}
	res, _ := configReader.Section(section).Key(key).Int()
	return int32(res)
}

func GetInt64(section, key string) int64 {
	if configReader == nil {
		return 0
	}
	res, _ := configReader.Section(section).Key(key).Int64()
	return res
}

func GetFloat64(section, key string) float64 {
	if configReader == nil {
		return 0
	}
	res, _ := configReader.Section(section).Key(key).Float64()
	return res
}

func GetBool(section, key string) bool {
	if configReader == nil {
		return false
	}
	res, _ := configReader.Section(section).Key(key).Bool()
	return res
}

func GetStringList(section, key, sep string) []string {
	result := GetString(section, key)
	return strings.Split(result, sep)
}

func GetInt64List(section, key, sep string) []int64 {
	var result []int64
	fields := GetStringList(section, key, sep)
	for _, item := range fields {
		res, err := strconv.ParseInt(item, 0, 64)
		if err != nil {
			continue
		}
		result = append(result, res)
	}
	return result
}

func GetFloat64List(section, key, sep string) []float64 {
	var result []float64
	fields := GetStringList(section, key, sep)
	for _, item := range fields {
		res, err := strconv.ParseFloat(item, 64)
		if err != nil {
			continue
		}
		result = append(result, res)
	}
	return result
}

func GetConfig() *ini.File {
	return configReader
}

func Unmarshal(v interface{}) error {
	if configReader == nil {
		return NotInitError
	}
	err := configReader.MapTo(v)
	return err
}

func UnmarshalWithSection(section string, v interface{}) error {
	if configReader == nil {
		return NotInitError
	}
	return configReader.Section(section).MapTo(v)
}
