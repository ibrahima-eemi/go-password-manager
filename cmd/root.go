package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	_ "github.com/barry/go-password-manager/cmd/add"
	_ "github.com/barry/go-password-manager/cmd/list"
	_ "github.com/barry/go-password-manager/cmd/get"
	_ "github.com/barry/go-password-manager/cmd/api"
	_ "github.com/barry/go-password-manager/cmd/setmaster"
	_ "github.com/barry/go-password-manager/cmd/tui"
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
