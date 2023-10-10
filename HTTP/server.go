//codigo preliminar//
package main

import (
    "fmt"
    "net/http"

)

func main() {
    // Crear un enrutador httprouter
    router := httprouter.New()

    // Definir una ruta y su manejador
    router.GET("/saludo/:nombre", SaludoHandler)

    // Iniciar el servidor HTTP en el puerto 8080
    fmt.Println("Servidor escuchando en el puerto 8080...")
    http.ListenAndServe(":8080", router)
}

func SaludoHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    nombre := ps.ByName("nombre")
    fmt.Fprintf(w, "¡Hola, %s!", nombre)
}
