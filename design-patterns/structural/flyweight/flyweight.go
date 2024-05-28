package flyweight

import (
	"fmt"
	"math/rand"
	"time"
)

// Card 表示一張撲克牌
type Card struct {
	Suit  string
	Value string
}

func (c *Card) Display() string {
	return fmt.Sprintf("%s of %s", c.Value, c.Suit)
}

// CardFactory 管理卡牌對象的享元工廠
type CardFactory struct {
	cards map[string]*Card
}

func NewCardFactory() *CardFactory {
	return &CardFactory{
		cards: make(map[string]*Card),
	}
}

func (f *CardFactory) GetCard(suit, value string) *Card {
	key := fmt.Sprintf("%s-%s", suit, value)
	if card, exists := f.cards[key]; exists {
		return card
	}
	card := &Card{Suit: suit, Value: value}
	f.cards[key] = card
	return card
}

type Game struct {
	Deck    []*Card
	Factory *CardFactory
}

func NewGame(factory *CardFactory) *Game {
	game := &Game{
		Deck:    []*Card{},
		Factory: factory,
	}
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	for _, suit := range suits {
		for _, value := range values {
			card := factory.GetCard(suit, value)
			game.Deck = append(game.Deck, card)
		}
	}
	return game
}

func (g *Game) Shuffle() {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(g.Deck), func(i, j int) {
		g.Deck[i], g.Deck[j] = g.Deck[j], g.Deck[i]
	})
}

func (g *Game) DisplayDeck() {
	for _, card := range g.Deck {
		fmt.Println(card.Display())
	}
}
