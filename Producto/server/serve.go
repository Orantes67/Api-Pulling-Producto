package server

import (
	"Practica/clienteServidor/Producto/handerl"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/addProduct", handerl.CreateProductHandler)
	http.HandleFunc("/deleteProduct", handerl.DeleteProductHandler)
	http.HandleFunc("/isNewProductAdded", handerl.IsNewProductAddedHandler)
	http.HandleFunc("/countProductsInDiscount", handerl.CountProductsInDiscountHandler)

	log.Println("Servidor principal corriendo en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
