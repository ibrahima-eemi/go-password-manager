package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd is the main command
var rootCmd = &cobra.Command{
	Use:   "password-manager",
	Short: "Gestionnaire de mots de passe sécurisé",
	Long:  "Un gestionnaire de mots de passe ultra sécurisé en Go avec chiffrement AES et stockage chiffré.",
}

// Execute runs the CLI application
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Erreur :", err)
	}
}
