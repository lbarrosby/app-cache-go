package handlers

import (
    "github.com/gin-gonic/gin"
    "app-globo-go/internal/cache"
    "time"
)

func RetornarTexto(c *gin.Context) {
    cacheKey := "/texto"
    cachedValue, err := cache.GetCache(cacheKey)
    if err != nil {
        c.JSON(500, gin.H{"error": "Erro ao acessar o cache"})
        return
    }

    if cachedValue != "" {
        c.JSON(200, gin.H{"cache": cachedValue})
        return
    }

    texto := "Texto fixo!"
    cache.SetCache(cacheKey, texto, 10*time.Second)
    c.JSON(200, gin.H{"message": texto})
}

func RetornarHorario(c *gin.Context) {
    cacheKey := "/horario"
    cachedValue, err := cache.GetCache(cacheKey)
    if err != nil {
        c.JSON(500, gin.H{"error": "Erro ao acessar o cache"})
        return
    }

    if cachedValue != "" {
        c.JSON(200, gin.H{"horario": cachedValue})
        return
    }

    currentTime := time.Now().Format("2006-01-02 15:04:05")
    cache.SetCache(cacheKey, currentTime, 30*time.Second)
    c.JSON(200, gin.H{"horario": currentTime})
}
