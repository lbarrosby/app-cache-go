package main

import (
    "fmt"
    "log"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
    "meu-projeto/app-globo-go/internal/middleware"
)

var rdb *redis.Client

// Definição das métricas Prometheus
var (
    httpRequests = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Número total de requisições HTTP",
        },
        []string{"method", "endpoint"},
    )

    requestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "Duração das requisições HTTP",
            Buckets: prometheus.DefBuckets,
        },
        []string{"method", "endpoint"},
    )
)

func init() {
    // Registrar as métricas no Prometheus
    prometheus.MustRegister(httpRequests)
    prometheus.MustRegister(requestDuration)
}

func main() {
    // Inicializa o cliente Redis
    rdb = redis.NewClient(&redis.Options{
        Addr: "redis:6379", // Assumindo que o Redis está rodando localmente
        DB:   0,            // Usando o banco de dados padrão 0
    })

    // Inicializa o middleware de cache
    middleware.InitCache(rdb)

    // Cria o router Gin
    router := gin.Default()

    // Adiciona o endpoint de métricas Prometheus
    router.GET("/metrics", gin.WrapH(promhttp.Handler()))

    // Rota que retorna um texto fixo, cacheado por 10 segundos
    router.GET("/texto_fixo", middleware.CacheMiddleware(10*time.Second), func(c *gin.Context) {
        start := time.Now()
        c.String(200, "Texto fixo Golang!")
        duration := time.Since(start).Seconds()

        // Incrementa o contador de requisições e registra a duração
        httpRequests.WithLabelValues("GET", "/texto_fixo").Inc()
        requestDuration.WithLabelValues("GET", "/texto_fixo").Observe(duration)
    })

    // Rota que retorna o horário atual, cacheado por 1 minuto
    router.GET("/hora", middleware.CacheMiddleware(1*time.Minute), func(c *gin.Context) {
        start := time.Now()
        c.String(200, fmt.Sprintf("Horário atual: %s Golang", time.Now().Format(time.RFC1123)))
        duration := time.Since(start).Seconds()

        // Incrementa o contador de requisições e registra a duração
        httpRequests.WithLabelValues("GET", "/hora").Inc()
        requestDuration.WithLabelValues("GET", "/hora").Observe(duration)
    })

    // Inicia o servidor na porta 8080
    err := router.Run(":8085")
    if err != nil {
        log.Fatal("Erro ao rodar o servidor:", err)
    }
}

