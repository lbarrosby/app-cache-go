package middleware

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gin-gonic/gin"
)

// Cliente Redis global
var rdb *redis.Client

// InitCache inicializa o cliente Redis para ser utilizado globalmente
func InitCache(redisClient *redis.Client) {
	rdb = redisClient
}

// CacheMiddleware cria um middleware de cache usando o Redis com expiração configurável.
func CacheMiddleware(expiration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cria uma chave única para o cache com base na URL da requisição
		cacheKey := fmt.Sprintf("cache:%s", c.Request.URL.Path)

		// Verifica se existe cache no Redis
		cachedData, err := rdb.Get(context.Background(), cacheKey).Result()
		if err == nil {
			// Se o cache existir, retorna os dados diretamente do Redis
			c.String(200, cachedData)
			c.Abort() // Impede a execução do próximo handler
			return
		}

		// Se não houver cache, cria um ResponseWriter personalizado para capturar a resposta
		writer := &responseWriter{ResponseWriter: c.Writer}
		c.Writer = writer

		// Chama o próximo handler da rota
		c.Next()

		// Após a execução da rota, captura os dados da resposta
		dataToCache := writer.Body

		// Armazena a resposta no Redis com o tempo de expiração configurado
		err = rdb.Set(context.Background(), cacheKey, dataToCache, expiration).Err()
		if err != nil {
			log.Println("Erro ao armazenar cache no Redis:", err)
		}
	}
}

// responseWriter é uma estrutura customizada para interceptar a resposta.
type responseWriter struct {
	gin.ResponseWriter
	Body string
}

// Write captura os dados da resposta e os armazena.
func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.Body += string(b) // Concatena o conteúdo da resposta
	return rw.ResponseWriter.Write(b)
}

