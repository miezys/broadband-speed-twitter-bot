package configs

import (
	"log"
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

const configFileName = "configs/config.yml"

// Load config
func (c *Configuration) Load() *Configuration {
	log.Printf("Loading configuration from file: %s\n", configFileName)
	file, err := os.Open(configFileName)
	if err != nil {
		log.Fatalf("Could not load configuration: %d\n", err)
		os.Exit(1)
	}

	defer file.Close()
	err = yaml.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatalf("Could not load the configuration: %d\n", err)
		os.Exit(1)
	}
	log.Printf("Loaded config values: %+v\n", c)
	return c
}
