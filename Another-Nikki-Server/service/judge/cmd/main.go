package main

import (
	"Another-Nikki/service/judge/internal/service"
	"Another-Nikki/util/log"
)

func main() {
	log.Init()
	service.New()
}
