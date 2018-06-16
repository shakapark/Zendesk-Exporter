package config

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

//Config Represent the Yaml config of Zendesk-Exporter
type Config struct {
	URL      string `yaml:"url"`
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
	Token    string `yaml:"token"`

	XXX map[string]interface{} `yaml:",inline"`
}

func checkOverflow(m map[string]interface{}, ctx string) error {
	if len(m) > 0 {
		var keys []string
		for k := range m {
			keys = append(keys, k)
		}
		return fmt.Errorf("unknown fields in %s: %s", ctx, strings.Join(keys, ", "))
	}
	return nil
}

//UnmarshalYAML Decoding yaml config file
func (s *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type plain Config
	if err := unmarshal((*plain)(s)); err != nil {
		return err
	}
	if err := checkOverflow(s.XXX, "config"); err != nil {
		return err
	}
	return nil
}

//SafeConfig Represent a config locked
type SafeConfig struct {
	sync.RWMutex
	C *Config
}

//ReloadConfig Reload Config from new yaml file
func (sc *SafeConfig) ReloadConfig(confFile string) (err error) {
	var c = &Config{}

	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		return fmt.Errorf("Error reading config file: %s", err)
	}

	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		return fmt.Errorf("Error parsing config file: %s", err)
	}

	sc.Lock()
	sc.C = c
	sc.Unlock()

	return nil
}
