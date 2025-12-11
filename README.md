# 📊 GoMonitor

> Um sistema de monitoramento de status de serviços em tempo real, desenvolvido com Golang e Goroutines.

![Badge em Desenvolvimento](http://img.shields.io/static/v1?label=STATUS&message=CONCLUIDO&color=GREEN&style=for-the-badge)
![Go Version](https://img.shields.io/badge/Go-1.25+-00ADD8?style=for-the-badge&logo=go)

---

## 💻 Sobre o Projeto

O **GoMonitor** é uma aplicação que permite cadastrar APIs, sites ou serviços de banco de dados para monitorar sua disponibilidade (Uptime).

O diferencial técnico deste projeto é o uso de **Goroutines** para realizar a checagem de múltiplos serviços de forma concorrente (paralela), garantindo que a aplicação não trave mesmo monitorando centenas de endpoints simultaneamente.

### ✨ Funcionalidades

- [x] **Cadastro de Serviços:** Adicione URLs para serem monitoradas.
- [x] **Monitoramento em Background:** Um *Scheduler* roda em uma thread separada testando os serviços automaticamente.
- [x] **Checagem Manual:** Possibilidade de forçar o teste de um serviço específico instantaneamente.
- [x] **Dashboard em Tempo Real:** Interface visual que atualiza os status via *Polling* sem recarregar a página.
- [x] **Soft Delete:** Remoção segura de serviços, mantendo histórico no banco.

---

## 🛠 Tecnologias Utilizadas

### Backend
- **[Go (Golang)](https://go.dev/):** Linguagem principal.
- **[Gin Gonic](https://github.com/gin-gonic/gin):** Framework Web de alta performance.
- **[GORM](https://gorm.io/):** ORM para manipulação do banco de dados.
- **SQLite:** Banco de dados relacional leve.
- **Concurrency:** Uso de `go func()` para tarefas assíncronas.

### Frontend
- **HTML5 & CSS3:** Estilização nativa com uso de CSS Variables (sem frameworks pesados).
- **JavaScript (Vanilla):** Consumo da API REST e manipulação do DOM.

---
