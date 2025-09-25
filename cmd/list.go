package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister tous les clients",
	Run: func(cmd *cobra.Command, args []string) {
		clients, err := store.GetAll()
		if err != nil {
			fmt.Println("Erreur:", err)
			return
		}

		if len(clients) == 0 {
			fmt.Println("Aucun client trouvé")
			return
		}

		fmt.Println("Liste des clients :")
		for _, c := range clients {
			fmt.Printf("ID: %d | Nom: %s | Email: %s\n", c.ID, c.Name, c.Email)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
