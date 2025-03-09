package web

import (
	"fmt"
	"go-password-manager/internal/crypto"
	"go-password-manager/internal/storage"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// TemplateCache stocke les templates compilés
var TemplateCache *template.Template

// InitTemplates charge les templates au démarrage
func InitTemplates() {
	var err error
	TemplateCache, err = template.ParseFiles(
		"web/templates/base.html",
		"web/templates/home.html",
		"web/templates/list.html",
		"web/templates/add.html",
	)
	if err != nil {
		log.Fatalf("❌ Erreur de chargement des templates : %v", err)
	}
	fmt.Println("📂 Templates chargés avec succès !")
}

// HomeHandler affiche la page d'accueil
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("<h1 style='color:red;'>Test réussi !</h1><p>Si tu vois ce message, le serveur fonctionne.</p>"))
}

// ListPasswords affiche la liste des mots de passe
func ListPasswords(w http.ResponseWriter, r *http.Request) {
	db := storage.InitDB("super-securise-passphrase")
	rows, err := db.Query("SELECT site, username FROM passwords")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des mots de passe", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var passwords []struct {
		Site     string
		Username string
	}
	for rows.Next() {
		var site, username string
		if err := rows.Scan(&site, &username); err == nil {
			passwords = append(passwords, struct {
				Site     string
				Username string
			}{site, username})
		}
	}

	renderTemplate(w, "list.html", passwords)
}

// AddPasswordHandler gère l'ajout d'un mot de passe
func AddPasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		site := r.FormValue("site")
		username := r.FormValue("username")
		password := r.FormValue("password")

		if site == "" || username == "" || password == "" {
			http.Error(w, "Tous les champs sont obligatoires", http.StatusBadRequest)
			return
		}

		// Récupérer la clé de chiffrement
		encryptionKey := os.Getenv("ENCRYPTION_KEY")
		if len(encryptionKey) != 32 {
			log.Fatal("❌ ERREUR : ENCRYPTION_KEY doit être une clé de 32 caractères dans .env")
		}

		// Chiffrer le mot de passe
		encryptedPassword, err := crypto.Encrypt(password)
		if err != nil {
			http.Error(w, "Erreur de chiffrement du mot de passe", http.StatusInternalServerError)
			return
		}

		// Stocker dans la base de données
		db := storage.InitDB("super-securise-passphrase")
		_, err = db.Exec("INSERT INTO passwords (site, username, password) VALUES (?, ?, ?)", site, username, encryptedPassword)
		if err != nil {
			http.Error(w, "Erreur lors de l'enregistrement", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Mot de passe enregistré pour %s !", site)
		return
	}
	renderTemplate(w, "add.html", nil)
}

// renderTemplate rend les fichiers HTML
func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	// Assurer que la réponse est bien du HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := TemplateCache.ExecuteTemplate(w, templateName, data)
	if err != nil {
		log.Printf("❌ Erreur de rendu (%s) : %v", templateName, err)
		http.Error(w, "Erreur de rendu : "+err.Error(), http.StatusInternalServerError)
	}
}

// StartServer démarre le serveur web
func StartServer() {
	InitTemplates() // Charger les templates au démarrage

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/passwords", ListPasswords).Methods("GET")
	r.HandleFunc("/add-password", AddPasswordHandler).Methods("GET", "POST")

	// Servir les fichiers statiques (HTMX, Tailwind)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Println("🚀 Serveur lancé sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
