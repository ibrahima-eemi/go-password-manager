package storage

import (
	"sync"
	"time"
)

// lastAuthTime enregistre le dernier moment où l'utilisateur s'est authentifié
var lastAuthTime time.Time
var authLock sync.Mutex

// SetAuthenticated enregistre le temps d'authentification
func SetAuthenticated() {
	authLock.Lock()
	defer authLock.Unlock()
	lastAuthTime = time.Now()
}

// IsAuthenticated vérifie si l'authentification est toujours valide
func IsAuthenticated() bool {
	authLock.Lock()
	defer authLock.Unlock()

	// Timeout après 5 minutes d'inactivité
	expiration := lastAuthTime.Add(5 * time.Minute)
	return time.Now().Before(expiration)
}
