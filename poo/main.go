package main

import (
	"poo/animal"
	"poo/book"
)

func main() {

	var myBook = book.NewBook("La verdad sobre el caso Harry Quebert", "Joël Dicker", 670)

	myBook.PrintInfo()

	miPerro := animal.NewDog("Firulais")

	miGato := animal.NewCat("Garfield")

	miGato.Speak()
	miPerro.Speak()

}
