package cmd

import (
	"fmt"
	"log"
	"minicrm_Verdiane/internal/storage"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var store storage.Storer

var rootCmd = &cobra.Command{
	Use:   "minicrm",
	Short: "MiniCRM - Gestion simple de clients",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initConfig()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Erreur lecture config:", err)
	}

	storageType := viper.GetString("storage.type")

	switch storageType {
	case "gorm":
		s, err := storage.NewGORMStore(viper.GetString("storage.db_file"))
		if err != nil {
			log.Fatal(err)
		}
		store = s
	case "json":
		s, err := storage.NewJSONStorage(viper.GetString("storage.json_file"))
		if err != nil {
			log.Fatal(err)
		}
		store = s
	case "memory":
		store = storage.NewMemoryStorage()
	default:
		log.Fatal("Type de stockage invalide :", storageType)
	}
	fmt.Println("Stockage actif :", storageType)
}
