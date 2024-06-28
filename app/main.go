package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"go.elastic.co/apm/module/apmgin"
	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/pq"
)

var db *sql.DB

func main() {
	// Configuração do banco de dados
	pgURL, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	db, err = apmsql.Open("postgres", pgURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criação da tabela
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT)`)
	if err != nil {
		log.Fatal(err)
	}

	// Configuração do Gin com APM e logging
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(apmgin.Middleware(r))

	// Servir arquivos estáticos
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	// Rota para a página inicial
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Rotas da API
	r.POST("/api/users", createUser)
	r.GET("/api/users", getUsers)

	// Iniciar o servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

func createUser(c *gin.Context) {
	var user struct {
		Name string `json:"name"`
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO users (name) VALUES ($1)", user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func getUsers(c *gin.Context) {
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	for rows.Next() {
		var user struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}
