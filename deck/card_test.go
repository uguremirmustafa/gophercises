package deck

import (
	"fmt"
	"math/rand"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Eight, Suit: Spade})
	fmt.Println(Card{Rank: Five, Suit: Club})
	fmt.Println(Card{Rank: Jack, Suit: Diamond})
	fmt.Println(Card{Suit: Joker})

	// Output:
	// Ace of Hearts
	// Eight of Spades
	// Five of Clubs
	// Jack of Diamonds
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cards in a new deck!")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	want := Card{Suit: Spade, Rank: Ace}
	if cards[0] != want {
		t.Errorf("want Ace of Spades; got %s", cards[0])
	}
}

func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	want := Card{Suit: Spade, Rank: Ace}
	if cards[0] != want {
		t.Errorf("want Ace of Spades; got %s", cards[0])
	}
}

func TestJokers(t *testing.T) {
	jokerCount := 3
	cards := New(Jokers(jokerCount))
	count := 0
	for _, card := range cards {
		if card.Suit == Joker {
			count++
		}
	}

	if count != jokerCount {
		t.Errorf("want %d Jokers; got %d Jokers", jokerCount, count)

	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// First call to shuffleRand.Perm(52) should be:
	// [40 35 ...]
	shuffleRand = rand.New(rand.NewSource(0)) // updating package level variable
	original := New()
	first := original[40]
	second := original[35]
	shuffled := New(Shuffle)

	if shuffled[0] != first {
		t.Errorf("expected the first card to be %s, got %s", first, shuffled[0])
	}
	if shuffled[1] != second {
		t.Errorf("expected the second card to be %s, got %s", second, shuffled[1])
	}
}

func TestFilter(t *testing.T) {
	filterTwoThree := func(c Card) bool {
		return c.Rank == Two || c.Rank == Three
	}
	cards := New(Filter(filterTwoThree))
	for _, card := range cards {
		if card.Rank == Two || card.Rank == Three {
			t.Error("Expected all two and three cards filtered out!")
		}
	}
}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))

	want := 13 * 4 * 3
	if len(cards) != want {
		t.Errorf("wanted %d;got %d", want, len(cards))
	}

}
