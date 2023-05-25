package main

import (
	"cube/worker"
	"log"
	"os"
	"strconv"
)

func main() {
	whost := os.Getenv("CUBE_WORKER_HOST")
	wport, _ := strconv.Atoi(os.Getenv("CUBE_WORKER_PORT"))
	wName := os.Getenv("CUBE_WORKER_NAME")
	wrk := worker.New(wName, "memory")
	log.Printf("worker name: %s host: %s, port: %d\n", wName, whost, wport)

	wrkAPI := worker.Api{Address: whost, Port: wport, Worker: wrk}

	go wrk.RunTasks()
	go wrk.UpdateTasks()
	wrkAPI.Start()
}
