# go-projects

API REST per gestione città con database PostgreSQL.

## Descrizione

Servizio Go che espone API CRUD per entità City. Implementa routing avanzato con Gorilla mux, middleware per content-type e audit logging, e persistenza su PostgreSQL tramite GORM.

## Endpoint

- `POST /city` - Crea una nuova città
- `GET /city` - Elenca tutte le città
- `GET /city/{id}` - Ottiene una città specifica
- `PATCH /city/{id}` - Aggiorna una città
- `DELETE /city/{id}` - Elimina una città

## Tecnologie

- **Linguaggio**: Go 1.15
- **Routing**: Gorilla Mux
- **ORM**: GORM
- **Database**: PostgreSQL

## Configurazione

Configurare la connessione al database tramite variabili d'ambiente nel Connection Factory.

## Build & Run

```bash
go build -o go-projects .
./go-projects
```

Il server ascolta su porta 8085.