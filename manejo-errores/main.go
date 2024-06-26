package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func saveConctactsToFile(contacts []Contact) error {

	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(contacts)

	if err != nil {
		return err
	}

	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {

	file, err := os.Open("contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&contacts)

	if err != nil {
		return err
	}

	return nil
}

func main() {

	var contacts []Contact

	err := loadContactsFromFile(&contacts)

	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\033[H\033[2J")

		mostrarMenu()

		var option int

		_, err = fmt.Scanln(&option)

		if err != nil {
			fmt.Println(err)

			continue
		}

		switch option {
		case 1:
			fmt.Println("Contactos")
			for _, contact := range contacts {
				fmt.Println("Nombre: ", contact.Name)
				fmt.Println("Telefono: ", contact.Phone)
				fmt.Println("Email: ", contact.Email)
				fmt.Println("===================================")
			}
		case 2:
			fmt.Println("Agregar contacto")
			var contact Contact

			fmt.Print("Nombre: ")
			contact.Name, _ = reader.ReadString('\n')

			fmt.Print("Telefono: ")
			contact.Phone, _ = reader.ReadString('\n')

			fmt.Print("Email: ")
			contact.Email, _ = reader.ReadString('\n')

			contacts = append(contacts, contact)

			err = saveConctactsToFile(contacts)

			if err != nil {
				fmt.Println(err)
			}

		case 3:
			fmt.Println("Saliendo del programa")
			return
		default:
			fmt.Println("Opcion no valida")

		}

	}

}

func mostrarMenu() {
	fmt.Println("===================================")
	fmt.Println("Menu de opciones")
	fmt.Println("1. Mostrar contactos")
	fmt.Println("2. Agregar contacto")
	fmt.Println("3. Salir")
	fmt.Println("===================================")
}
