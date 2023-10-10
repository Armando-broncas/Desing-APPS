//codigo preliminar//

package main

import (
    "fmt"
    "net/http"

    "github.com/Armandobroncas/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "¡Bienvenido al servidor Go con httprouter!")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    name := ps.ByName("name")
    fmt.Fprintf(w, "Hola, %s!", name)
}

func main() {
    // Crear un enrutador httprouter
    router := httprouter.New()

    // Definir rutas y manejadores
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)

    // Iniciar el servidor HTTP en el puerto 8080
    fmt.Println("Servidor escuchando en el puerto 8080...")
    http.ListenAndServe(":8080", router)
}
