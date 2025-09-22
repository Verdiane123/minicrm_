package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)


type Client struct {
	ID    int
	Name  string
	Email string
}


var clients = make(map[int]Client)
var nextID = 1

func main() {
	// Flags
	name := flag.String("name", "", "Nom du client")
	email := flag.String("email", "", "Email du client")
	flag.Parse()


	if strings.TrimSpace(*name) != "" || strings.TrimSpace(*email) != "" {
		n := strings.TrimSpace(*name)
		e := strings.TrimSpace(*email)
		if n == "" || e == "" {
			fmt.Println("Erreur : Merci de renseigner -name et -email pour une saisie unique :)")
			os.Exit(1)
		}
		id := addClient(n, e)
		fmt.Printf("Client '%s' ajouté avec succès ! ID %d via flags\n", n, id)
		return
	}

	// Mode menu interactif
	reader := bufio.NewReader(os.Stdin)
	for {
		printMenu()
		fmt.Print("Votre action : ")
		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erreur de lecture :", err)
			continue
		}
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addClientInteractive(reader)
		case "2":
			listClients()
		case "3":
			deleteClientInteractive(reader)
		case "4":
			updateClientInteractive(reader)
		case "5", "q", "Q", "":
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide, réessayez.")
		}
	}
}

func printMenu() {
	fmt.Println("\n--- Verdiane CRM ---")
	fmt.Println("1. Nouveau client")
	fmt.Println("2. Lister les clients")
	fmt.Println("3. Supprimer un client")
	fmt.Println("4. Mettre à jour un client")
	fmt.Println("5. Quitter")
}

func addClient(name, email string) int {
	id := nextID
	clients[id] = Client{ID: id, Name: name, Email: email}
	nextID++
	return id
}

func addClientInteractive(r *bufio.Reader) {
	fmt.Print("Entrez le nom du client : ")
	name, _ := r.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Entrez son email : ")
	email, _ := r.ReadString('\n')
	email = strings.TrimSpace(email)

	if name == "" || email == "" {
		fmt.Println("Nom et email requis. Client non ajouté.")
		return
	}
	id := addClient(name, email)
	fmt.Printf("Client '%s' ajouté avec succès ID %d.\n", name, id)
}

func listClients() {
	if len(clients) == 0 {
		fmt.Println("Vous n'avez pas de client enregistré.")
		return
	}
	fmt.Println("\n--- Liste des Clients ---")
	// tri avec l'ID 
	ids := make([]int, 0, len(clients))
	for id := range clients {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	for _, id := range ids {
		c := clients[id]
		fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Name, c.Email)
	}
}

func deleteClientInteractive(r *bufio.Reader) {
	fmt.Print("Entrez l'ID du client à supprimer : ")
	s, _ := r.ReadString('\n')
	s = strings.TrimSpace(s)
	id, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ID invalide.")
		return
	}

	// comma-ok idiom : vérification du client dans la map
	if _, ok := clients[id]; !ok {
		fmt.Printf("Client ID %d introuvable.\n", id)
		return
	}
	delete(clients, id)
	fmt.Printf("Client ID %d supprimé.\n", id)
}

func updateClientInteractive(r *bufio.Reader) {
	fmt.Print("Entrez l'ID du client à mettre à jour : ")
	s, _ := r.ReadString('\n')
	s = strings.TrimSpace(s)
	id, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ID invalide.")
		return
	}

	client, ok := clients[id]
	if !ok {
		fmt.Printf("Client ID %d introuvable.\n", id)
		return
	}

	fmt.Printf("Nom actuel: %s\n", client.Name)
	fmt.Print("Nouveau nom (laisser vide si vous changez d'avis) : ")
	name, _ := r.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = client.Name
	}

	fmt.Printf("Email actuel: %s\n", client.Email)
	fmt.Print("Nouvel email (laisser vide si vous changez d'avis) : ")
	email, _ := r.ReadString('\n')
	email = strings.TrimSpace(email)
	if email == "" {
		email = client.Email
	}

	clients[id] = Client{ID: id, Name: name, Email: email}
	fmt.Printf("Client ID %d mis à jour avec succès.\n", id)
}
