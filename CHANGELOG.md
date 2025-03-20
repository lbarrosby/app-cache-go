# Changelog

Todas as mudanças importantes para este projeto serão documentadas neste arquivo.

## [v1.0.0] - 2025-03-20

### Adicionado
- Implementação da aplicação em Go utilizando o framework **Gin**.
- Integração com o **Redis** para caching de respostas das rotas.
- Criação de métricas com **Prometheus** para monitoramento de desempenho (total de requisições, duração das requisições, etc.).
- Middleware de cache para otimizar respostas de endpoints.
- Endpoints principais:
  - `/texto_fixo`: Retorna um texto fixo com cache de 10 segundos.
  - `/hora`: Retorna o horário atual com cache de 1 minuto.
  - `/metrics`: Exposição de métricas no formato Prometheus.
- Configuração do **Docker Compose** para rodar a aplicação junto com o Redis em containers separados.
- Suporte a cache Redis utilizando **go-redis**.
- Configuração de **Dockerfile** para containerizar a aplicação Go.
- Arquivo `docker-compose.yml` para facilitar o setup dos containers.

### Melhorias
- Uso de **Prometheus** para coletar métricas de performance da aplicação.
- Aplicação de cache eficiente nas respostas com Redis, garantindo alta performance.
- Uso de **Gin** para criação de rotas e middleware de forma rápida e eficiente.

### Correções
- Nenhuma correção de bugs nesta versão, pois é a versão inicial do projeto.

## [v0.1.0] - 2025-03-10

### Inicial
- Criação do repositório.
- Configuração inicial do Go.
