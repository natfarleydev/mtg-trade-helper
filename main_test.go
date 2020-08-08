package main

import (
	"os"
	"testing"
)

const configName = "config.yaml"

func TestGetConfigFileName(t *testing.T) {
	os.Args = []string{"mtg-trade-helper", configName}

	name, err := getConfigFileName()
	if err != nil {
		t.Errorf("got an error I didn't expect: %v", err)
	}

	if name != configName {
		t.Errorf("unexpected filename: %v", name)
	}

}

func TestGetConfig(t *testing.T) {
	config, err := getConfig(configName)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(config.Wants) == 0 {
		t.Error("config.wants is empty?")
	}
}

func TestGetCards(t *testing.T) {
	config, err := getConfig(configName)
	if err != nil {
		t.Errorf("unable to get test config: %v", err)
	}

	cards, err := getCards(config.Wants)
	if len(cards) != 15 {
		t.Errorf("unexpected number of cards: %v", len(cards))
		var cardNames []string
		for _, c := range cards {
			cardNames = append(cardNames, c.Name)
		}
		t.Errorf("got cards: %v", cardNames)
	}
}

// func TestGetCardsGetsFoils(t *testing.T) {
// 	cards, err := getCards([]ConfigCard{
// 		{Name: "Primal Empathy", Foil: true, Quantity: 1},
// 	})
// 	if len(cards) == 0 {
// 		t.Errorf("failed to get card: %v", err)
// 	}

// 	if !cards[0].Foil {
// 		t.Errorf("failed to get foil version of card")
// 	}
// }

// func TestGetCardsDoesntGetFoilsWhenNotWanted(t *testing.T) {
// 	cards, err := getCards([]ConfigCard{
// 		{Name: "Primal Empathy", Foil: false, Quantity: 1},
// 	})
// 	if len(cards) == 0 {
// 		t.Errorf("failed to get card: %v", err)
// 	}

// 	if cards[0].Foil {
// 		t.Errorf("got foil but didn't want foil")
// 	}
// }
