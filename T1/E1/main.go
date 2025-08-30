package main

import (
	"fmt"
	"net/http"
)

func main() {

	// Se define el contenido HTML.
	htmlContent := `
			<!DOCTYPE html>
			<html>
			<head><title>¡Hola, Mundo!</title></head>
			<body>
				<h1>¡Servidor funcionando!</h1>
				<p>Ejercicio 1.</p>
			</body>
			</html>
		`

	// Se registra un manejador (handler) para la ruta raíz "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if esRutaInvalida(w, r, "/", http.MethodGet) {
			return
		}

		// Se establece la cabecera Content-Type.
		// Es decir, el tipo de contenido que se recibirá y la codificación.
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Se escribe en HTML la respuesta.
		// La biblioteca lo hace automáticamente.
		fmt.Fprint(w, htmlContent)
	})

	// Se registra un manejador (handler) para la ruta "/about"
	// En lugar de definir la función de forma anónima, le damos un nombre.
	// (Hecho diferente solo para probar).
	http.HandleFunc("/about", serveAbout)

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

func esRutaInvalida(w http.ResponseWriter, r *http.Request, path string, method string) bool {

	if r.URL.Path != path || r.Method != method {
		http.NotFound(w, r)
		return true
	}

	return false
}

// serveForm: Maneja GET /about para mostrar la información del servidor.
func serveAbout(w http.ResponseWriter, r *http.Request) {

	if esRutaInvalida(w, r, "/about", http.MethodGet) {
		return
	}

	information := "Server running in Debian 13 (Trixie)."

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, information)
}
