package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Configuration struct {
	RepositoryFile   string
	SubscriptionFile string
}

func GetConfig() Configuration {

	configuration := Configuration{}

	if _, err := os.Stat("podcasthub.config"); err == nil {

		file, _ := os.Open("podcasthub.config")
		decoder := json.NewDecoder(file)
		decoder.Decode(&configuration)
	} else {
		configuration.RepositoryFile = "PodcastHub.bolt"
		configuration.SubscriptionFile = "subscriptions"

		json, _ := json.Marshal(configuration)
		ioutil.WriteFile("podcasthub.config", json, 0644)
	}

	return configuration
}
