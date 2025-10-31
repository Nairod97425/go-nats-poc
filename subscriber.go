package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func main() {
	// 1. Connexion
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	// 2. JetStream
	js, err := nc.JetStream()
	if err != nil {
		panic(err)
	}

	// 2.1. Créer le stream s'il n'existe pas déjà
	// _, err = js.AddStream(&nats.StreamConfig{
	// 	Name:     "USERS",
	// 	Subjects: []string{"events.user.*"},
	// })
	// if err != nil && err != nats.ErrStreamNameAlreadyInUse {
	// 	panic(err)
	// }

	// 3. S'abonner au sujet
	sub, err := js.Subscribe("events.user.registered", func(m *nats.Msg) {
		fmt.Printf("Reçu : %s\n", string(m.Data))
		// Ici : on pourrait envoyer un WhatsApp !
		m.Ack() // Confirme qu'on a traité
	})
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	fmt.Println("Subscriber en écoute... (Ctrl+C pour arrêter)")
	select {} // Bloque indéfiniment
}
