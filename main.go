package main

import (
	"fmt"
	"go-password-manager/cmd"
	"go-password-manager/internal/storage"
	"go-password-manager/web"
	"os"
)

func main() {
	fmt.Println("Secure Password Manager")

	// Générer .env si nécessaire
	storage.GenerateEnvFile()

	// Charger les variables d'environnement
	storage.LoadEnv()

	// Vérifier si l'utilisateur veut démarrer l'interface web
	if len(os.Args) > 1 && os.Args[1] == "web" {
		web.StartServer()
		return
	}

	// Exécuter les commandes CLI
	cmd.Execute()
}
