package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// configurar linea de codigo en la que se produjo el error
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.SetPrefix("Main: ")

	log.Println("Inicio del programa")
	// log.Fatalln("Error fatal")

	file, errr := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if errr != nil {
		log.Fatalln("Error al abrir el archivo")
	}

	log.SetOutput(file)

	log.Print("Hola mundo")

	log.Println("Fin del programa")

	defer file.Close()

}

func divide(a, b int) int {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	validarZero(b)

	return a / b

}

func validarZero(b int) {
	if b == 0 {
		panic("No se puede dividir por cero")
	}
}

func usoDefer() {

	file, err := os.Create("test.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hello World"))

	if err != nil {
		fmt.Println(err)
		return
	}

}
