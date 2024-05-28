package flyweight

import (
	"testing"
)

// 測試卡牌工廠的獲取卡牌功能
func TestCardFactory(t *testing.T) {
	factory := NewCardFactory()

	card1 := factory.GetCard("Hearts", "A")
	card2 := factory.GetCard("Hearts", "A")

	if card1 != card2 {
		t.Errorf("Expected card1 and card2 to be the same instance")
	}

	card3 := factory.GetCard("Diamonds", "A")

	if card1 == card3 {
		t.Errorf("Expected card1 and card3 to be different instances")
	}
}

// 測試牌局的初始化功能
func TestNewGame(t *testing.T) {
	factory := NewCardFactory()
	game := NewGame(factory)

	if len(game.Deck) != 52 {
		t.Errorf("Expected deck size to be 52, but got %d", len(game.Deck))
	}
}

// 測試牌局的洗牌功能
func TestShuffle(t *testing.T) {
	factory := NewCardFactory()
	game := NewGame(factory)
	initialDeck := make([]*Card, len(game.Deck))
	copy(initialDeck, game.Deck)

	game.Shuffle()
	shuffledDeck := game.Deck

	sameOrder := true
	for i, card := range initialDeck {
		if card != shuffledDeck[i] {
			sameOrder = false
			break
		}
	}

	if sameOrder {
		t.Errorf("Expected deck to be shuffled, but order is the same")
	}
}
