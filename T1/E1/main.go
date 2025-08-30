package main

import (
	"fmt"
	"net/http"
)

const (
	port        = ":8080"  // Puerto en el que escucharemos.
	rootPath    = "/"      // Ruta raíz
	aboutPath   = "/about" // Ruta /about
	contentType = "text/html; charset=utf-8"
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
	http.HandleFunc(rootPath, func(w http.ResponseWriter, r *http.Request) {

		if esRutaInvalida(w, r, rootPath, http.MethodGet) {
			return
		}

		// Se establece la cabecera Content-Type.
		// Es decir, el tipo de contenido que se recibirá y la codificación.
		w.Header().Set("Content-Type", contentType)

		// Se escribe en HTML la respuesta.
		// La biblioteca lo hace automáticamente.
		fmt.Fprint(w, htmlContent)
	})

	// Se registra un manejador (handler) para la ruta "/about"
	// En lugar de definir la función de forma anónima, le damos un nombre.
	// (Hecho diferente solo para probar).
	http.HandleFunc(aboutPath, serveAbout)

	// ----------------------------------------------------------------------------------------- //
	// Poner a correr el servidor.
	// ----------------------------------------------------------------------------------------- //

	// Mensaje que se mostrará al poner a correr el servidor.
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)

	// Se inicia el servidor HTTP.
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Error al iniciar el servidor: %s\n", err)
	}
}

func esRutaInvalida(w http.ResponseWriter, r *http.Request, path string, method string) bool {

	if r.URL.Path != path || r.Method != method {
		http.Error(w, "404 página no encontrada", http.StatusNotFound)
		return true
	}

	return false
}

// serveForm: Maneja GET /about para mostrar la información del servidor.
func serveAbout(w http.ResponseWriter, r *http.Request) {

	if esRutaInvalida(w, r, aboutPath, http.MethodGet) {
		return
	}

	information := "Server running in Debian 13 (Trixie)."

	w.Header().Set("Content-Type", contentType)
	fmt.Fprint(w, information)
}
