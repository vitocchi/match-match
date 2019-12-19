package card

import (
	"reflect"
	"testing"
)

func TestDrop(t *testing.T) {
	cm := CardMap{
		NewCard(Spade, 1): 1,
		NewCard(Spade, 2): 2,
		NewCard(Spade, 3): 3,
	}
	cm.Drop(NewCard(Spade, 1))
	expected := CardMap{
		NewCard(Spade, 2): 2,
		NewCard(Spade, 3): 3,
	}
	if !reflect.DeepEqual(cm, expected) {
		t.Fatalf("expected %v, actual %v", expected, cm)
	}
}

func TestFlip(t *testing.T) {
	cm := CardMap{
		NewCard(Spade, 1): 1,
		NewCard(Spade, 2): 2,
		NewCard(Spade, 3): 3,
	}
	cm.Flip(NewCard(Spade, 1), 4)
	expected := CardMap{
		NewCard(Spade, 1): 4,
		NewCard(Spade, 2): 2,
		NewCard(Spade, 3): 3,
	}
	if !reflect.DeepEqual(cm, expected) {
		t.Fatalf("expected %v, actual %v", expected, cm)
	}
}

func TestGetMostFormalyFlipped(t *testing.T) {
	cm := CardMap{
		NewCard(Spade, 1): 3,
		NewCard(Spade, 2): 2,
		NewCard(Spade, 3): 1,
		NewCard(Heart, 3): 4,
	}
	card := cm.GetMostFormalyFlipped()
	expected := NewCard(Spade, 3)
	if card != expected {
		t.Fatalf("expected %v, actual %v", expected, card)
	}
	cm = CardMap{
		NewCard(Spade, 1): 0,
		NewCard(Spade, 2): 0,
		NewCard(Spade, 3): 0,
		NewCard(Spade, 4): 1,
	}
	card = cm.GetMostFormalyFlipped()
	if !(card == NewCard(Spade, 1) || card == NewCard(Spade, 2) || card == NewCard(Spade, 3)) {
		t.Fatalf("expected 1, 2, 3 actual %v", card)
	}
}

func TestGetMostRecentlyFlippedInNumber(t *testing.T) {
	cm := CardMap{
		NewCard(Spade, 1):   3,
		NewCard(Spade, 2):   2,
		NewCard(Spade, 3):   1,
		NewCard(Heart, 3):   4,
		NewCard(Diamond, 3): 5,
		NewCard(Club, 3):    6,
		NewCard(Spade, 4):   0,
		NewCard(Heart, 4):   0,
		NewCard(Diamond, 4): 0,
		NewCard(Club, 4):    0,
	}
	card := cm.GetMostRecentlyFlippedInNumber(1)
	expected := NewCard(Spade, 1)
	if card != expected {
		t.Fatalf("expected %v, actual %v", expected, card)
	}
	card = cm.GetMostRecentlyFlippedInNumber(3)
	expected = NewCard(Club, 3)
	if card != expected {
		t.Fatalf("expected %v, actual %v", expected, card)
	}
	card = cm.GetMostRecentlyFlippedInNumber(4)
	if card.Number() != 4 {
		t.Fatalf("expected %v, actual %v", expected, card)
	}
}
