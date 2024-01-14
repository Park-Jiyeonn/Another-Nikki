package main

import (
	"Another-Nikki/job/service"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	s := service.New()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer-group-id",
		Topic:    "hello",
		MaxBytes: 1e6, // 10MB
	})

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for {
			sign := <-sigchan
			switch sign {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
				fmt.Printf("Received signal %v. Closing reader...\n", sign)
				if err := r.Close(); err != nil {
					fmt.Println(err)
				}
			case syscall.SIGHUP:
			default:
				return
			}
		}
	}()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
			log.Println(m.Value)
			break
		}
		s.MessageHandle(m.Value)
	}
}
