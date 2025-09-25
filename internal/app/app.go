package app

import (
	"bufio"
	"fmt"
	"os"
	"minicrm_Verdiane/internal/storage"
	"strconv"
	"strings"
)

func Run(store storage.Storer) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Mini CRM v3!")

	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Add a client")
		fmt.Println("2. List clients")
		fmt.Println("3. Update a client")
		fmt.Println("4. Delete a client")
		fmt.Println("5. Exit")
		fmt.Print("Your choice: ")

		choice := readUserChoice(reader)

		switch choice {
		case 1:
			handleAddClient(reader, store)
		case 2:
			handleListClients(store)
		case 3:
			handleUpdateClient(reader, store)
		case 4:
			handleDeleteClient(reader, store)
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again")

		}
	}
}



func handleAddClient(reader *bufio.Reader, storer storage.Storer) {
	fmt.Print("Enter client name: ")
	name := readLine(reader)

	fmt.Print("Enter client email: ")
	email := readLine(reader)

	client := &storage.Client{
		Name:  name,
		Email: email,
	}
	err := storer.Add(client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Client '%s' added with ID %d.\n", client.Name, client.ID)
}

func handleListClients(store storage.Storer) {
	clients, err := store.GetAll()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	if len(clients) == 0 {
		fmt.Println(" No clients to display.")
		return
	}

	fmt.Println("\n--- Client List ---")
	for _, client := range clients {
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", client.ID, client.Name, client.Email)
	}
}

func handleUpdateClient(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Enter the ID of the client to update: ")
	id := readInteger(reader)
	if id == -1 {
		return
	}

	// Vérification de l'existence du client
	existingClient, err := store.GetByID(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Updating '%s'. Leave blank to keep current value.\n", existingClient.Name)

	fmt.Printf("New name (%s): ", existingClient.Name)
	newName := readLine(reader)

	fmt.Printf("New email (%s): ", existingClient.Email)
	newEmail := readLine(reader)

	err = store.Update(id, newName, newEmail)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Client updated successfully.")
}

func handleDeleteClient(reader *bufio.Reader, store storage.Storer) {
	fmt.Print("Enter the ID of the client to delete: ")
	id := readInteger(reader)
	if id == -1 {
		return
	}

	err := store.Delete(id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Client with ID %d has been deleted.\n", id)
}

func readLine(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func readUserChoice(reader *bufio.Reader) int {
	choice, err := strconv.Atoi(readLine(reader))
	if err != nil {
		return -1 
	}
	return choice
}

func readInteger(reader *bufio.Reader) int {
	id, err := strconv.Atoi(readLine(reader))
	if err != nil {
		fmt.Println("Error: Invalid ID. Please enter a number.")
		return -1
	}
	return id
}
