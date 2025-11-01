# PoC Event-Driven avec Go + NATS

## Objectif
Simuler un flux d'événements :  
`user.registered` → notification (console, extensible à WhatsApp)

## Stack
- Go
- NATS JetStream (via Docker)
- Event-driven architecture

## Lancer
```bash
docker run -d --name nats -p 4222:4222 nats:latest -js