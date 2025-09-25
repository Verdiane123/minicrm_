package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Mettre à jour un client existant",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")

		if err := store.Update(id, name, email); err != nil {
			fmt.Println("Erreur:", err)
			return
		}
		fmt.Printf("Client avec ID %d mis à jour.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().Int("id", 0, "ID du client à mettre à jour")
	updateCmd.Flags().String("name", "", "Nouveau nom")
	updateCmd.Flags().String("email", "", "Nouvel email")
	updateCmd.MarkFlagRequired("id")
}
