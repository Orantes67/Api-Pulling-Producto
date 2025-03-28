package main

import (

	"Practica/clienteServidor/Producto/server"
	
)

func main( ){
	go server.Run()

	select {}

}