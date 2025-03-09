# 🔐 Go Password Manager  

**Un gestionnaire de mots de passe ultra sécurisé en Go avec chiffrement AES et stockage local chiffré.**  

## 🌟 **Fonctionnalités**  

✅ **Stockage sécurisé** : Chiffrement AES-GCM pour protéger les mots de passe.  
✅ **Mot de passe maître** : Accès restreint avec une authentification préalable.  
✅ **Interface CLI complète** : Ajout, récupération et suppression de mots de passe via un terminal.  
✅ **Serveur API local** : Communication avec une extension navigateur pour l'auto-remplissage.  
✅ **Timeout d’inactivité** : Sécurisation automatique après une période d’inactivité.  
✅ **Mode interactif (TUI)** : Interface terminal ergonomique pour une navigation simplifiée.  
✅ **Clés d’API sécurisées** : Fichiers `.env` générés automatiquement pour stocker les clés sensibles.  
✅ **Sauvegarde et restauration** : Export des mots de passe chiffrés pour les réimporter en toute sécurité.  

## 📂 **Architecture du projet**  

```txt
go-password-manager/
│── cmd/ # Commandes CLI
│   ├── root.go # Commande principale
│   ├── add.go # Ajouter un mot de passe
│   ├── list.go # Lister les mots de passe
│   ├── get.go # Récupérer un mot de passe
│   ├── setmaster.go # Définir le mot de passe maître
│   ├── api.go # Serveur API pour l'extension navigateur
│   ├── tui.go # Mode interactif en ligne de commande (TUI)
│
│── internal/
│   ├── crypto/ # Gestion du chiffrement
│   │   ├── encryption.go # Chiffrement AES-GCM des mots de passe
│   │   ├── hashing.go # Hashage du mot de passe maître
│   ├── storage/ # Gestion de la base de données SQLite chiffrée
│   │   ├── database.go # Connexion et opérations sur la base
│   │   ├── env.go # Gestion des fichiers .env pour les clés d’API
│   ├── extension/ # Extension navigateur (Chrome/Firefox)
│       ├── manifest.json # Fichier de configuration de l’extension
│       ├── popup.html # Interface utilisateur
│       ├── background.js # Communication avec le serveur Go
│
│── main.go # Point d’entrée du programme
│── go.mod # Dépendances et gestion du module Go
│── go.sum # Fichier des versions de dépendances
│── .env # Clés API et clé de chiffrement (non suivi par Git)
│── passwords.db # Base de données SQLite chiffrée
│── README.md # Documentation du projet
```

## 🛠️ **Installation et Configuration**  

### **1️⃣ Cloner le projet**

```sh
git clone git@github.com:ibrahima-eemi/go-password-manager.git
cd go-password-manager
```

### **2️⃣ Installer les dépendances**

```sh
go mod tidy
```

### **3️⃣ Générer le fichier .env avec la clé de chiffrement et l’API token**

```sh
go run main.go generate-env
```

### **4️⃣ Lancer le gestionnaire de mots de passe**

```sh
go run main.go menu
```

## 💻 **Utilisation des commandes CLI**  

### 🔹 **Ajouter un mot de passe**

```sh
go run main.go add
```

👉 Exemple : Ajouter example.com avec l’utilisateur user123 et un mot de passe généré.

### 🔹 **Lister les mots de passe**

```sh
go run main.go list
```

👉 Affiche tous les sites enregistrés.

### 🔹 **Récupérer un mot de passe en clair**

```sh
go run main.go get
```

👉 Demande le site et affiche l’identifiant et le mot de passe déchiffré.

### 🔹 **Démarrer l’API locale**

```sh
go run main.go api
```

👉 Lance un serveur sur http://localhost:8080 pour l’extension navigateur.

## 🔒 **Sécurité et Bonnes Pratiques**

✅ Chiffrement AES-GCM pour garantir la sécurité des mots de passe.  
✅ Mot de passe maître stocké en hash (bcrypt) pour éviter toute récupération brute.  
✅ Base de données SQLite chiffrée avec SQLCipher.  
✅ Clé de chiffrement stockée uniquement dans .env et jamais hardcodée.  
✅ Timeout automatique en cas d’inactivité pour éviter les accès non autorisés.

## 🤝 **Contribuer au projet**

1. Forker le projet  
2. Créer une branche pour une nouvelle feature  
3. Soumettre une Pull Request (PR)  

## 📄 **Licence**

Ce projet est sous licence MIT.
