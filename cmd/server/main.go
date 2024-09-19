package main

import (
	"log"

	"go-todo/api"
	"go-todo/aws"
)

func main() {
	go api.StartGRPCServer()
	aws.ScheduleADRAnalysis()
	log.Println("Service is running...")
}
