package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"minicrm_Verdiane/internal/storage"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Veuillez spécifier une commande: add ou list")
	}

	command := os.Args[1]

	store, err := storage.NewGORMStore("clients.db")
	if err != nil {
		log.Fatal("Erreur connexion DB:", err)
	}

	switch command {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		name := addCmd.String("name", "", "Nom du client")
		email := addCmd.String("email", "", "Email du client")
		addCmd.Parse(os.Args[2:])

		if *name == "" || *email == "" {
			log.Fatal("Merci de fournir --name et --email")
		}

		client := &storage.Client{Name: *name, Email: *email}
		if err := store.Add(client); err != nil {
			log.Fatal("Erreur ajout:", err)
		}
		fmt.Println("Client ajouté:", client)

	case "list":
		all, _ := store.GetAll()
		fmt.Println("Clients en DB:", all)

	default:
		log.Fatal("Commande inconnue:", command)
	}
}
