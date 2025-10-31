package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	// 1. Se connecter à NATS
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	// 2. Activer JetStream
	js, err := nc.JetStream()
	if err != nil {
		panic(err)
	}

	// 3. Publier un événement toutes les 3 secondes
	for i := 1; i <= 5; i++ {
		user := fmt.Sprintf("user%d@example.com", i)
		msg := []byte(fmt.Sprintf(`{"action":"registered","email":"%s"}`, user))

		// Publie sur le sujet "events.user.registered"
		js.Publish("events.user.registered", msg)
		fmt.Printf("Envoyé : %s\n", msg)

		time.Sleep(3 * time.Second)
	}

	fmt.Println("Publisher terminé.")
}