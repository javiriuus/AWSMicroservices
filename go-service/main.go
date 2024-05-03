package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// Estructura para los datos del usuario
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Obtiene la cadena de conexión desde las variables de entorno
	dbConnStr := os.Getenv("DB_CONN_STR")

	// Conecta a la base de datos PostgreSQL
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Define un manejador HTTP para la ruta "/users"
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// Consulta a la base de datos para obtener los usuarios
		rows, err := db.Query("SELECT id, name, email FROM user")
		if err != nil {
			log.Printf("Error al consultar la base de datos: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Procesa los resultados de la consulta
		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Name, &user.Email)
			if err != nil {
				log.Printf("Error al escanear la fila: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		// Convierte la lista de usuarios a JSON
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			log.Printf("Error al codificar JSON: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	// Escucha en el puerto 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Servidor escuchando en el puerto %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
