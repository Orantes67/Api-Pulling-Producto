package main

import (

	"Practica/clienteServidor/Persona/server"
	
)

func main( ){
	go server.Run()

	select {}

}