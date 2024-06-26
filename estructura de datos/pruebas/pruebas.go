package pruebas

import "fmt"

type persona struct {
	nombre string
	edad   int
}

func main() {
	estructuras()
}

func (p *persona) diHola() {
	fmt.Println("Hola soy ", p.nombre)
}

func punteros() {
	var x int = 10

	var y *int = &x

	editar(y)

	fmt.Println(x)
	fmt.Println(*y)
}

func editar(x *int) {
	*x = 20
}

func estructuras() {

	p := persona{nombre: "Juan", edad: 20}

	fmt.Println(p)

	p.nombre = "Pedro"

	fmt.Println(p)

	p2 := persona{nombre: "", edad: 20}

	p2.nombre = "Maria"

	p.diHola()
	p2.diHola()

	fmt.Println(p)

}

func mapas() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	colors["black"] = "#000000"

	delete(colors, "white")

	fmt.Println(colors)

	for key, value := range colors {
		fmt.Printf("key %s value %s \n", key, value)
	}

}

func rebanadas() {

	diasSemanas := []string{"Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado", "Domingo"}

	diasEntreSemana := diasSemanas[:4]

	diasEntreSemana = append(diasEntreSemana, diasSemanas[4])

	fmt.Println(diasEntreSemana)

	diaRebanada := append(diasEntreSemana[:2], diasEntreSemana[4:]...)

	fmt.Println(diaRebanada)

	fmt.Println(diasSemanas)

	fmt.Println(cap(diasSemanas))
	fmt.Println(len(diasSemanas))
}

func matrices() {

	var matriz [5]int

	var matriz2 = [...]int{1, 2, 3, 4, 5}

	matriz[0] = 1
	matriz[1] = 2
	matriz[2] = 3
	matriz[3] = 4
	matriz[4] = 5

	for i := 0; i < len(matriz); i++ {
		fmt.Println(matriz[i])
	}

	for i := 0; i < len(matriz2); i++ {
		fmt.Println(matriz2[i])
	}

	for i, value := range matriz2 {
		fmt.Printf("indice %d valor %d \n", i, value)
	}

	var matriz3 [2][2]int

	matriz3[0][0] = 1
	matriz3[0][1] = 2
	matriz3[1][0] = 3
	matriz3[1][1] = 4

	for i := 0; i < len(matriz3); i++ {
		for j := 0; j < len(matriz3[i]); j++ {
			fmt.Println(matriz3[i][j])
		}
	}

}
