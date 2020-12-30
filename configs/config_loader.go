package configs

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

// Configuration struct for client usage
type Configuration struct {
	BotConfig struct {
		BroadbandProvider          string `yaml:"broadbandProvider"`
		BroadbandProviderTwitterID string `yaml:"broadbandProviderTwitterID"`
		ExpectedDownloadSpeed      int    `yaml:"expectedDownloadSpeed"`
		ExpectedUploadSpeed        int    `yaml:"expectedUploadSpeed"`
		MeasureFrequency           int    `yaml:"measureFrequency"`
	} `yaml:"botConfig"`

	TwitterAPIConfig struct {
		ConsumerKey       string `yaml:"consumerKey"`
		ConsumerSecret    string `yaml:"consumerSecret"`
		AccessToken       string `yaml:"accessToken"`
		AccessTokenSecret string `yaml:"accessTokenSecret"`
	} `yaml:"twitterAPIConfig"`
}

// ConfigLoader interface for loading configuration values
type ConfigLoader interface {
	load() (Configuration, error)
}

const configFileName = "configs/config.yml"

func (c *Configuration) load() (*Configuration, error) {
	file, err := os.Open(configFileName)
	if err != nil {
		return c, err
	}

	defer file.Close()
	err = yaml.NewDecoder(file).Decode(&c)
	return c, err
}
