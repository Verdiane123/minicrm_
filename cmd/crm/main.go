package main

import (
	"flag"
	"fmt"
	"log"
	"minicrm_Verdiane/internal/app"
	"minicrm_Verdiane/internal/storage"
	"os"
	"strings"
)

func main() {
	// Flags
	name := flag.String("name", "", "Nom du client")
	email := flag.String("email", "", "Email du client")
	useMemory := flag.Bool("memory", false, "Utiliser stockage en mémoire (par défaut JSON)")
	flag.Parse()

	// Choix du backend de stockage
	var store storage.Storer
	var err error

	if *useMemory {
		store = storage.NewMemoryStorage()
	} else {
		store, err = storage.NewJSONStorage("clients.json")
		if err != nil {
			log.Fatal("Impossible d'initialiser le stockage JSON :", err)
		}
	}

	// Mode "ajout rapide via flags"
	if strings.TrimSpace(*name) != "" || strings.TrimSpace(*email) != "" {
		n := strings.TrimSpace(*name)
		e := strings.TrimSpace(*email)
		if n == "" || e == "" {
			fmt.Println("Erreur : Merci de renseigner -name et -email pour une saisie unique :)")
			os.Exit(1)
		}
		client := &storage.Client{Name: n, Email: e}
		if err := store.Add(client); err != nil {
			fmt.Println("Erreur ajout:", err)
			os.Exit(1)
		}
		fmt.Printf("Client '%s' ajouté avec succès ! ID %d via flags\n", n, client.ID)
		return
	}

	// Sinon → on lance le menu interactif
	app.Run(store)
}
