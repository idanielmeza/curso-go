package main

import "fmt"

type Tarea struct {
	nombre     string
	desc       string
	compeltado bool
}

type ListaTareas struct {
	usuario string
	tareas  []Tarea
}

func (lt *ListaTareas) agregarTarea(t Tarea) {
	lt.tareas = append(lt.tareas, t)
}

func (lt *ListaTareas) completarTarea(index int) {
	lt.tareas[index].compeltado = true
}

func (lt *ListaTareas) eliminarTarea(index int) {
	lt.tareas = append(lt.tareas[:index], lt.tareas[index+1:]...)
}

func (lt *ListaTareas) editarTarea(index int, t Tarea) {
	lt.tareas[index] = t
}

func (lt *ListaTareas) mostrarTareas() {
	for i, t := range lt.tareas {

		estaCompletado := "No"

		if t.compeltado {
			estaCompletado = "Si"
		}

		fmt.Println("Tarea ", i+1, " Nombre: ", t.nombre, " Descripcion: ", t.desc, " Completado: ", estaCompletado)
	}
}

func main() {

	lista := ListaTareas{}

	fmt.Println("Bienvenido a la lista de tareas, cual es tu nombre?")

	fmt.Scanln(&lista.usuario)

	for {
		menu()

		var opcion int
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var nombre string
			var desc string

			fmt.Println("Nombre de la tarea: ")
			fmt.Scanln(&nombre)

			fmt.Println("Descripcion de la tarea: ")
			fmt.Scanln(&desc)

			tarea := Tarea{nombre: nombre, desc: desc}

			lista.agregarTarea(tarea)

		case 2:
			lista.mostrarTareas()

			var index int
			fmt.Println("Cual tarea quieres completar? (ingresa el numero de la tarea)")
			fmt.Scanln(&index)

			lista.completarTarea(index - 1)

		case 3:
			lista.mostrarTareas()

			var index int
			fmt.Println("Cual tarea quieres eliminar? (ingresa el numero de la tarea)")
			fmt.Scanln(&index)

			lista.eliminarTarea(index - 1)

		case 4:
			lista.mostrarTareas()

			var index int
			fmt.Println("Cual tarea quieres editar? (ingresa el numero de la tarea)")
			fmt.Scanln(&index)

			var nombre string
			var desc string

			fmt.Println("Nombre de la tarea: ")
			fmt.Scanln(&nombre)

			fmt.Println("Descripcion de la tarea: ")
			fmt.Scanln(&desc)

			tarea := Tarea{nombre: nombre, desc: desc}

			lista.editarTarea(index-1, tarea)

		case 5:
			fmt.Println("Adios ", lista.usuario)
			return
		}

	}

}

func menu() {
	fmt.Println("==========================")
	fmt.Println("Menu de opciones")
	fmt.Println("1. Agregar tarea")
	fmt.Println("2. Completar tarea")
	fmt.Println("3. Eliminar tarea")
	fmt.Println("4. Editar tarea")
	fmt.Println("5. Salir")
	fmt.Println("==========================")
}
