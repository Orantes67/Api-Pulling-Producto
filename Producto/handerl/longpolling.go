package handerl

import (
	"Practica/clienteServidor/Producto/app"
	"encoding/json"
	"net/http"
	"time"
)

// Handler para obtener productos en descuento con long polling
func CountProductsInDiscountHandler(w http.ResponseWriter, r *http.Request) {
	products := getDiscountedProducts()

	// Si hay productos con descuento, los devuelve de inmediato
	if len(products) > 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
		return
	}

	// Si no hay productos, espera hasta 30 segundos
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			products = getDiscountedProducts()
			if len(products) > 0 {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(products)
				return
			}
		case <-timeout:
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
}

// Obtener productos en descuento desde el árbol binario
func getDiscountedProducts() []app.Producto {
	var productos []app.Producto
	lista := Arbol.ObtenerTodos()
	for _, p := range lista {
		if p.Descuento {
			productos = append(productos, p)
		}
	}
	return productos
}
