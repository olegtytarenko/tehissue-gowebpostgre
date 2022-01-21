package service

import (
	"os"
	"reflect"
	"testing"
)

type testCaseJsonConfig struct {
	fileContent   []byte
	isError       bool
	configCompare *Config
}

var casesJsonConfigurations = []testCaseJsonConfig{
	{
		fileContent: []byte("{\"http\": {}, \"database\": {}}"),
		isError:     false,
		configCompare: &Config{
			Http:     &ConfigHttp{},
			Database: &ConfigDatabase{},
		},
	},
	{
		fileContent: []byte("{\"http\": {\"endpoint\": \"api\"}, \"database\": {}}"),
		isError:     false,
		configCompare: &Config{
			Http: &ConfigHttp{
				Endpoint: "api",
			},
			Database: &ConfigDatabase{},
		},
	},
	{
		fileContent:   []byte("{}"),
		isError:       false,
		configCompare: &Config{},
	},
}

func TestGetConfiguration(t *testing.T) {
	for _, caseConfig := range casesJsonConfigurations {
		// generate tmp file
		err := os.WriteFile("/tmp/config_import_test", caseConfig.fileContent, 0777)
		if err != nil {
			t.Error(err)
			break
		}

		conf, errConf := GetConfiguration("/tmp/config_import_test")

		if errConf != nil {
			if caseConfig.isError {
				continue
			}

			t.Error(errConf)
			break
		}

		if reflect.DeepEqual(caseConfig.configCompare, conf) {
			continue
		}

		t.Log("Configuration is not equal")
		t.Fail()
	}
}
