package main

import (
	"fmt"
	"net/http"
)

const (
	port        = ":8080"         // Puerte que se escucha.
	rootPath    = "/"             // Ruta raíz
	contactPath = "/contacto"     // Ruta contacto
	contactGet  = "/contacto-get" // Ruta contacto GET
	contentType = "text/html; charset=utf-8"
)

func main() {

	// Se define el contenido HTML.
	htmlContent := `
			<!DOCTYPE html>
			<html>
			<head><title>¡Hola, Mundo!</title></head>
			<body><h1>¡Servidor funcionando!</h1></body>
			<p>Ejercicio 2. Pruebe ir a "/contacto" o a "/contacto-get"</p>
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

	// No puede haber dos handle asociados a una misma ruta.
	http.HandleFunc(contactPath, manageForm)

	// Variación para probar.
	http.HandleFunc(contactGet, manageFormGet)

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

func manageForm(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		serveForm(w, r)
	} else if r.Method == http.MethodPost {
		handleLogin(w, r)
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// serveForm: Maneja GET /login para mostrar el formulation
// Ver "Tipos de Inputs en HTML" --> https://www.w3schools.com/html/html_form_input_types.asp
func serveForm(w http.ResponseWriter, r *http.Request) {

	if esRutaInvalida(w, r, contactPath, http.MethodGet) {
		return
	}

	// Agregando "required" dentro del "input" obliga a que el campo no esté vacío.
	// type="email" obliga a que lo que se coloque en el campo tenga formato de email.
	loginForm := `
		<!DOCTYPE html>
		<html>
		<head><title>Login</title></head>
		<body>
			<h2>Login</h2>
			<form action="/contacto" method="POST">
			<label>Nombre:</label><input type="text" name="name" required> <br>
			<label>Email:</label><input type="email" name="email" required> <br>
			<label>Mensaje:</label><input type="text" name="message" required> <br>
			<button type="submit">Login</button></form>
		</body>
		</html>
		`

	w.Header().Set("Content-Type", contentType)
	fmt.Fprint(w, loginForm)
}

// handleLogin: Maneja POST /login para procesar datos
// También sirve si es GET.
func handleLogin(w http.ResponseWriter, r *http.Request) {

	// Se parsean los datos del formulario (crucial)
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al parsear", http.StatusBadRequest)
		return
	}

	// Obtención de los valores de los campos.
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")
	if name == "" || email == "" || message == "" {
		http.Error(w, "Los campos no pueden estar vacíos.", http.StatusBadRequest)
		return
	}

	// Se genera y se envía la respuesta HTML.
	w.Header().Set("Content-Type", contentType)

	salida := `
		<!DOCTYPE html>
		<html>
		<head><title>Bienvenido</title></head>
		<body><h1>¡Hola, %s!</h1>
		<p>Recibimos tus datos. <br>
		Email: %s. <br> 
		Mensaje: "%s".</p>
		<a href="/">Volver</a></body></html>
	`

	fmt.Fprintf(w, salida, name, email, message)
}

// Variación Get.
func manageFormGet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.RawQuery != "" {
		handleLogin(w, r)
	} else {
		serveFormGet(w, r)
	}
}

func serveFormGet(w http.ResponseWriter, r *http.Request) {

	if esRutaInvalida(w, r, contactGet, http.MethodGet) {
		return
	}

	// Agregando "required" dentro del "input" obliga a que el campo no esté vacío.
	// type="email" obliga a que lo que se coloque en el campo tenga formato de email.
	loginForm := `
		<!DOCTYPE html>
		<html>
		<head><title>Login</title></head>
		<body>
			<h2>Login</h2>
			<form action="/contacto-get" method="GET">
			<label>Nombre:</label><input type="text" name="name" required> <br>
			<label>Email:</label><input type="email" name="email" required> <br>
			<label>Mensaje:</label><input type="text" name="message" required> <br>
			<button type="submit">Login</button></form>
		</body>
		</html>
		`

	w.Header().Set("Content-Type", contentType)
	fmt.Fprint(w, loginForm)
}
