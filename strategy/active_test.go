package strategy

import "testing"

import "github.com/vitocchi/match-match/table/card"

import "reflect"

func TestNewPickedMemoryList(t *testing.T) {
	cm := card.CardMap{
		card.NewCard(card.Diamond, 1): 1,
		card.NewCard(card.Spade, 1):   2,
		card.NewCard(card.Diamond, 2): 5,
		card.NewCard(card.Spade, 2):   4,
		card.NewCard(card.Club, 3):    2,
		card.NewCard(card.Heart, 3):   4,
		card.NewCard(card.Diamond, 3): 3,
		card.NewCard(card.Heart, 3):   1,
	}

	list := newPickedMemoryList(cm)
	expected := pickedMemoryList{
		pickedMemory{
			number:           1,
			lastPickedCard:   card.NewCard(card.Diamond, 1),
			lastPickedAt:     1,
			secondPickedCard: card.NewCard(card.Spade, 1),
			secondPickedAt:   2,
		},
		pickedMemory{
			number:           2,
			lastPickedCard:   card.NewCard(card.Spade, 2),
			lastPickedAt:     4,
			secondPickedCard: card.NewCard(card.Diamond, 2),
			secondPickedAt:   5,
		},
		pickedMemory{
			number:           3,
			lastPickedCard:   card.NewCard(card.Heart, 3),
			lastPickedAt:     1,
			secondPickedCard: card.NewCard(card.Club, 3),
			secondPickedAt:   2,
		},
	}
	if reflect.DeepEqual(list, expected) {
		t.Fatalf("expected %+v acutual %+v", expected, list)
	}
}

func TestGetMostLikelyMemory(t *testing.T) {
	list := pickedMemoryList{
		pickedMemory{
			number:           1,
			lastPickedCard:   card.NewCard(card.Diamond, 1),
			lastPickedAt:     6,
			secondPickedCard: card.NewCard(card.Spade, 1),
			secondPickedAt:   1,
		},
		pickedMemory{
			number:           2,
			lastPickedCard:   card.NewCard(card.Spade, 2),
			lastPickedAt:     3,
			secondPickedCard: card.NewCard(card.Diamond, 2),
			secondPickedAt:   5,
		},
		pickedMemory{
			number:           3,
			lastPickedCard:   card.NewCard(card.Heart, 3),
			lastPickedAt:     4,
			secondPickedCard: card.NewCard(card.Club, 3),
			secondPickedAt:   2,
		},
	}
	expected := pickedMemory{
		number:           2,
		lastPickedCard:   card.NewCard(card.Spade, 2),
		lastPickedAt:     3,
		secondPickedCard: card.NewCard(card.Diamond, 2),
		secondPickedAt:   5,
	}

	if list.getMostLikelyMemory() != expected {
		t.Fatalf("expected %+v acutual %+v", expected, list.getMostLikelyMemory())
	}
}
