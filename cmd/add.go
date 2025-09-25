package cmd

import (
	"fmt"
	"minicrm_Verdiane/internal/storage"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un client",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")

		client := &storage.Client{Name: name, Email: email}
		if err := store.Add(client); err != nil {
			fmt.Println("Erreur:", err)
			return
		}
		fmt.Printf("Client '%s' ajouté avec ID %d\n", client.Name, client.ID)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().String("name", "", "Nom du client")
	addCmd.Flags().String("email", "", "Email du client")
	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("email")
}
