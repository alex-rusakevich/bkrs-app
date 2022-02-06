package bkrs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

func ReadFileIntoString(filePath string) string {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

type Config struct {
	Features struct {
		Settings                       bool `json:"settings"`
		TopMenu                        bool `json:"topMenu"`
		NoAds                          bool `json:"noAds"`
		FocusOnInput                   bool `json:"focusOnInput"`
		ChineseCharMouseWheelClick     bool `json:"chineseCharMouseWheelClick"`
		HandwritingButtonAlwaysVisible bool `json:"handwritingButtonAlwaysVisible"`
		SearchSelectedOnShortcut       bool `json:"searchSelectedOnShortcut"`
		AdaptiveScrollbarColor         bool `json:"adaptiveScrollbarColor"`
		DisableBrowserAutocomplete     bool `json:"disableBrowserAutocomplete"`
		StrokeOrderSearch              bool `json:"strokeOrderSearch"`
	} `json:"features"`

	Shortcuts struct {
		SearchSelectedOnShortcut string `json:"searchSelectedOnShortcut"`
		StrokeOrderSearch        string `json:"strokeOrderSearch"`
	} `json:"shortcuts"`
}

var globalConfig Config

func LoadConfig() {
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(b, &globalConfig)
}

func SaveConfig() {
	data, err := json.MarshalIndent(globalConfig, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("config.json", data, 0644)
}

func IsFeatureEnabled(featureName string) (bool, error) {
	r := reflect.ValueOf(globalConfig)
	f := reflect.Indirect(r).FieldByName("Features").FieldByName(strings.Title(featureName))
	return bool(f.Bool()), nil
}

func GetKeysByFeatureName(featureName string) (string, error) {
	r := reflect.ValueOf(globalConfig)
	f := reflect.Indirect(r).FieldByName("Shortcuts").FieldByName(strings.Title(featureName))
	return string(f.String()), nil
}

func SetFeatureState(featureName string, newState bool) {
	r := reflect.ValueOf(&globalConfig)
	reflect.Indirect(r).FieldByName("Features").FieldByName(strings.Title(featureName)).SetBool(newState)
}

func SetKeys(featureName string, newKeys string) {
	r := reflect.ValueOf(&globalConfig)
	reflect.Indirect(r).FieldByName("Shortcuts").FieldByName(strings.Title(featureName)).SetString(newKeys)
}
