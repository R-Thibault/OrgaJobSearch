package config

// © Rossa Thibault 2024. Tous droits réservés.
// Ce code est la propriété de Rossa Thibault et ne peut être utilisé,
// distribué ou modifié sans autorisation explicite.
import (
	"log"

	"github.com/spf13/viper"
)

func SetupConfig() {
	// Looking for ".env" file and try read it
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	viper.AutomaticEnv() // Automatically read environment variables
}

func GetConfig(key string) string {
	return viper.GetString(key)
}
