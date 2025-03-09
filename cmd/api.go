package cmd

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-password-manager/internal/crypto"
	"go-password-manager/internal/storage"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// Credentials représente une requête envoyée par l'extension
type Credentials struct {
	Site string `json:"site"`
	Auth string `json:"auth"`
}

// PasswordResponse représente la réponse envoyée à l'extension
type PasswordResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *sql.DB

// handleGetPassword renvoie le mot de passe stocké pour un site donné
func handleGetPassword(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Requête invalide", http.StatusBadRequest)
		return
	}

	// Vérifier l'authentification via un token
	expectedToken := os.Getenv("API_TOKEN")
	if creds.Auth != expectedToken {
		http.Error(w, "Authentification refusée", http.StatusUnauthorized)
		return
	}

	// Rechercher le mot de passe dans la base de données
	row := db.QueryRow("SELECT username, password FROM passwords WHERE site = ?", creds.Site)
	var username, encryptedPassword string
	err = row.Scan(&username, &encryptedPassword)
	if err != nil {
		http.Error(w, "Site non trouvé", http.StatusNotFound)
		return
	}

	// Déchiffrer le mot de passe
	decryptedPassword, err := crypto.Decrypt(encryptedPassword, "super-secret-key")
	if err != nil {
		http.Error(w, "Erreur de déchiffrement", http.StatusInternalServerError)
		return
	}

	// Réponse JSON sécurisée
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PasswordResponse{Username: username, Password: decryptedPassword})
}

// StartServer lance le serveur API
func StartServer() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Avertissement : Impossible de charger le fichier .env")
	}

	db = storage.InitDB("super-securise-passphrase")

	// Création du routeur
	r := mux.NewRouter()
	r.HandleFunc("/get-password", handleGetPassword).Methods("POST")

	// Démarrage du serveur API
	fmt.Println("Serveur API démarré sur http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Démarre le serveur API",
	Run: func(cmd *cobra.Command, args []string) {
		StartServer()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
