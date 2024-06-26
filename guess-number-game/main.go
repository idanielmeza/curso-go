package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {

	jugar()

}

func jugar() {
	fmt.Println("Empezando a jugar")
	const limiteIntentos = 5
	numeroAleatorio := rand.Intn(100)
	intentos := 0
	var numeroIngresado int

	for intentos < limiteIntentos {
		fmt.Println("Ingrese un numero")
		fmt.Scan(&numeroIngresado)

		if numeroIngresado == numeroAleatorio {
			fmt.Println("Ganaste")
			break
		} else {
			intentos++
			fmt.Println("Te quedan ", limiteIntentos-intentos, " intentos")
		}

	}

	if intentos == limiteIntentos {
		fmt.Println("Perdiste")
		fmt.Println("El numero era ", numeroAleatorio)
	}

	// volver a jugar
	var respuesta string
	fmt.Println("Desea volver a jugar? s/n")
	fmt.Scan(&respuesta)

	if respuesta == "s" {
		jugar()
	} else {
		fmt.Println("Gracias por jugar")
	}

}

func retonar2Valores() (int, int) {

	return 1, 2
}

func hello(texto string) string {
	fmt.Println(texto)
	return texto + " desde la funcion"
}

func bucleFor() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

}

func estructurasControl() {
	if t := time.Now(); t.Hour() < 12 {
		fmt.Println("Es de mañana")
	} else if t.Hour() < 18 {
		fmt.Println("Es de tarde")
	} else {
		fmt.Println("Es de noche")
	}

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println(os)
	case "darwin":
		fmt.Println(os)
	case "linux":
		fmt.Println(os)
	default:
		fmt.Println("No se reconoce el sistema operativo")
	}
}
