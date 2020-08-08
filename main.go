package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	scryfall "github.com/BlueMonday/go-scryfall"
	yaml "gopkg.in/yaml.v2"
)

// ConfigCard defines the Card object in config
type ConfigCard struct {
	Name     string `yaml:"name,omitempty"`
	Quantity uint   `yaml:"quantity,omitempty"`
	// Foil     bool   `yaml:"foil,omitempty"`
}

// Config is the config for mtg-trade-helper
type Config struct {
	Wants []ConfigCard `yaml:"wants"`
}

func getConfigFileName() (string, error) {
	if len(os.Args) < 1 {
		return "", fmt.Errorf("No arguments passed. Expected 1 argument: config file")
	}
	args := os.Args[1:]
	if len(args) != 1 {
		return "", fmt.Errorf("Log file not passed, expected 1 arg, got: %v", args)
	}

	return args[0], nil
}

func getConfig(configFileName string) (config Config, retErr error) {
	f, err := os.Open(configFileName)
	if err != nil {
		retErr = fmt.Errorf("problem opening config file: %v", err)
		return
	}
	defer f.Close()
	configBytes, err := ioutil.ReadAll(f)
	if err != nil {
		retErr = fmt.Errorf("unable to read contents of the config file: %v", err)
		return
	}

	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		retErr = fmt.Errorf("unable to interpret config file: %v", err)
		return
	}

	// Now that we have the config properly, we populate defaults

	for i := range config.Wants {
		// If the quantity number is 0, that means it (hopefully) wasn't
		// populated.
		if config.Wants[i].Quantity == 0 {
			config.Wants[i].Quantity = 1
		}
	}

	return
}

func getCards(cards []ConfigCard) ([]scryfall.Card, error) {
	var scryfallCards []scryfall.Card

	ctx := context.Background()
	client, err := scryfall.NewClient()
	if err != nil {
		return nil, fmt.Errorf("unable to establish scryfall client: %v", err)
	}

	for _, c := range cards {
		card, err := client.GetCardByName(ctx, c.Name, true, scryfall.GetCardByNameOptions{})
		if err != nil {
			log.Fatal(err)
		}

		for i := uint(0); i < c.Quantity; i++ {
			scryfallCards = append(scryfallCards, card)
		}
	}

	return scryfallCards, nil
}

func main() {

}
