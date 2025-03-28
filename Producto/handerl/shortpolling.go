package handerl

import (
	"encoding/json"
	"net/http"
	"sync"
)

var (
	newProductAdded = false
	spMu            sync.Mutex
)

// Handler para verificar si hay un nuevo producto
func IsNewProductAddedHandler(w http.ResponseWriter, r *http.Request) {
	spMu.Lock()
	defer spMu.Unlock()

	resp := map[string]bool{"newProduct": newProductAdded}
	newProductAdded = false // Reseteamos el estado después de consultarlo

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Función para marcar que hay un nuevo producto
func MarkNewProductAdded() {
	spMu.Lock()
	newProductAdded = true
	spMu.Unlock()
}
