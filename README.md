# Application Golang com Redis e Prometheus

Este projeto é uma aplicação Go que utiliza o Redis para caching e Prometheus para monitoramento de métricas.

## Tecnologias

- **Go (Golang)**: Linguagem de programação usada para desenvolver a aplicação.
- **Gin**: Framework web para Go.
- **go-redis**: Cliente Redis para Go.
- **Prometheus Client**: Biblioteca para coletar e expor métricas para o Prometheus.
- **Redis**: Banco de dados em memória utilizado para caching.
- **Docker**: Containerização da aplicação e do Redis.

## Funcionalidades

1. **Cache com Redis**:
   - A aplicação utiliza o Redis para cache de respostas, minimizando o tempo de resposta e sobrecarga no servidor.
   - As rotas `/texto_fixo` e `/hora` possuem cache configurado com TTL específico.

2. **Métricas para Prometheus**:
   - A aplicação expõe métricas sobre o número total de requisições HTTP recebidas, a duração das requisições, e outras métricas de performance.
   - As métricas são acessíveis através da rota `/metrics`.

3. **Middleware**:
   - O middleware `CacheMiddleware` lida com o cache das respostas usando o Redis, proporcionando cache de alta performance.

## Instalação

### Pré-requisitos

1. **Go**: Certifique-se de ter o Go instalado na sua máquina. Caso não tenha, instale a partir de [golang.org](https://golang.org/dl/).
2. **Docker**: A aplicação utiliza o Docker para rodar o Redis em um container. Se ainda não tem o Docker, siga a documentação em [docker.com](https://www.docker.com/get-started) para instalar.
3. **Docker Compose** (opcional): Caso queira rodar o Redis e a aplicação com Docker Compose.
