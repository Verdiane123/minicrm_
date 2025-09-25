package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprimer un client par son ID",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")

		if err := store.Delete(id); err != nil {
			fmt.Println("Erreur:", err)
			return
		}
		fmt.Printf("Client avec ID %d supprimé.\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Int("id", 0, "ID du client à supprimer")
	deleteCmd.MarkFlagRequired("id")
}
