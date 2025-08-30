package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Se define un directorio que contiene los archivos estáticos.
	fileDir := "./E3/files"

	// Se crea un manejador (handler) de servidor de archivos.
	// http.Dir convierte la ruta del directorio en un sistema de archivos HTTP.
	// http.FileServer crea un manejador que sirve archivos
	// ¡Automáticamente sirve index.html para directorios!
	fileServer := http.FileServer(http.Dir(fileDir))

	// Usamos http.Handle porque fileServer es un http.Handler.
	http.Handle("/", fileServer)

	// ----------------------------------------------------------------------------------------- //
	// Poner a correr el servidor.
	// ----------------------------------------------------------------------------------------- //

	// Definimos el puerto en el que estaremos escuchando.
	port := ":8080"

	// Mensaje que se mostrará al poner a correr el servidor.
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	// Se inicia el servidor HTTP.
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}
