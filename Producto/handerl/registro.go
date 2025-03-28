package handerl

import (
	"Practica/clienteServidor/Producto/app"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	Arbol       = &app.ArbolBinario{}
	Subscribers = make([]chan bool, 0)
	Mu          sync.Mutex
)

// Agregar un nuevo producto
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var p app.Producto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Error en el JSON", http.StatusBadRequest)
		return
	}

	Mu.Lock()
	Arbol.Insertar(p)
	Mu.Unlock()

	notifySubscribers()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

// Eliminar un producto por código
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	codigo := r.URL.Query().Get("codigo")
	if codigo == "" {
		http.Error(w, "Código inválido", http.StatusBadRequest)
		return
	}

	Mu.Lock()
	Arbol.Eliminar(codigo)
	Mu.Unlock()

	notifySubscribers()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Producto con código %s eliminado", codigo)
}

// Notificar cambios a los suscriptores
func notifySubscribers() {
	Mu.Lock()
	defer Mu.Unlock()
	for _, ch := range Subscribers {
		ch <- true
	}
	Subscribers = nil
}
