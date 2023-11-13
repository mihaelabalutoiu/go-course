package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Card interface {
	fmt.Stringer
	Name() string
}

type CarteDeJoc struct {
	Color  string
	Number string
}

func NewBook(c, n string) *CarteDeJoc {
	return &CarteDeJoc{
		Color:  c,
		Number: n,
	}
}

func (c *CarteDeJoc) String() string {
	return fmt.Sprintf("%s de tipul %s", c.Number, c.Color)
}

func (c *CarteDeJoc) Name() string {
	return c.String()
}

type MagicBook struct {
	BookName string
	Power    int
}

func NewMagicBook(n string, p int) *MagicBook {
	return &MagicBook{
		BookName: n,
		Power:    p,
	}
}

func (cp *MagicBook) String() string {
	return fmt.Sprintf("Caught %s with power %d", cp.BookName, cp.Power)
}

func (cp *MagicBook) Name() string {
	return cp.String()
}

type Pachet[T Card] struct {
	books []T
}

func NewPachetDeCarti() *Pachet[*CarteDeJoc] {
	colors := []string{"RedHeart", "Trefla", "Diamond", "BlackHeart"}
	numbers := []string{"2", "3", "4", "5", "6", "7", "J", "Q", "K"}

	pachet := &Pachet[*CarteDeJoc]{}
	for _, c := range colors {
		for _, n := range numbers {
			pachet.AddBook(NewBook(c, n))
		}
	}
	return pachet
}

func NewPacketDeMagicBooks() *Pachet[*MagicBook] {
	magicBooks := []string{"salazar", "harry", "pika"}
	powers := []int{100, 200, 300}

	pachet := &Pachet[*MagicBook]{}
	for _, mb := range magicBooks {
		for _, p := range powers {
			pachet.AddBook(NewMagicBook(mb, p))
		}
	}
	return pachet
}

func (p *Pachet[T]) AddBook(carte T) {
	p.books = append(p.books, carte)
}

func (p *Pachet[T]) BookRandom() T {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))

	return p.books[r.Intn(len(p.books))]
}

func main() {
	pachet := NewPachetDeCarti()
	pachetDeMagicBooks := NewPacketDeMagicBooks()

	book := pachet.BookRandom()
	fmt.Println(book)

	magicBook := pachetDeMagicBooks.BookRandom()
	fmt.Println(magicBook)
}
