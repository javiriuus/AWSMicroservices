package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Usuario representa un usuario de la tabla "user" en la base de datos PostgreSQL
type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

func main() {
	// Configura la conexión a PostgreSQL
	dsn := "postgres://username:password@hostname:5432/database_name?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Verifica la conexión a la base de datos
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error al verificar conexión a la base de datos: %v", err)
	}

	// Crea un enrutador Gin
	router := gin.Default()

	// Ruta para obtener la lista de usuarios
	router.GET("/usuarios", func(c *gin.Context) {
		usuarios, err := obtenerUsuarios(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
			return
		}
		c.JSON(http.StatusOK, usuarios)
	})

	// Inicia el servidor HTTP
	log.Println("Servidor en ejecución en el puerto 8080")
	router.Run(":8080")
}

// obtenerUsuarios obtiene la lista de usuarios de la base de datos PostgreSQL
func obtenerUsuarios(db *sql.DB) ([]Usuario, error) {
	rows, err := db.Query("SELECT id, nombre, email FROM user")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []Usuario
	for rows.Next() {
		var usuario Usuario
		if err := rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Email); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}

	// Verifica si hay errores en la iteración
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return usuarios, nil
}
