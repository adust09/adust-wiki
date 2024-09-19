package main

import (
	"log"
	"project-root/api"
	"project-root/aws"
)

func main() {
	go api.StartGRPCServer()
	aws.ScheduleADRAnalysis()
	log.Println("Service is runnning...")
}
