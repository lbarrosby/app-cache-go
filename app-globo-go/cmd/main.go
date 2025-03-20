package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"meu-projeto/app-globo-go/internal/middleware"
)

var rdb *redis.Client

func main() {
	// Inicializa o cliente Redis
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis:6379", // Assumindo que o Redis está rodando localmente
		DB:   0,                // Usando o banco de dados padrão 0
	})

	// Inicializa o middleware de cache
	middleware.InitCache(rdb)

	// Cria o router Gin
	router := gin.Default()

	// Rota que retorna um texto fixo, cacheado por 10 segundos
	router.GET("/texto_fixo", middleware.CacheMiddleware(10*time.Second), func(c *gin.Context) {
		c.String(200, "Texto fixo Golang!")
	})

	// Rota que retorna o horário atual, cacheado por 1 minuto
	router.GET("/hora", middleware.CacheMiddleware(1*time.Minute), func(c *gin.Context) {
		c.String(200, fmt.Sprintf("Horário atual: %s Golang", time.Now().Format(time.RFC1123)))
	})

	// Inicia o servidor na porta 8080
	err := router.Run(":8085")
	if err != nil {
		log.Fatal("Erro ao rodar o servidor:", err)
	}
}

