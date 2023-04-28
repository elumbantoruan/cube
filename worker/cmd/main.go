package main

import (
	"cube/task"
	"cube/worker"
	"log"
	"os"
	"strconv"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

func main() {
	whost := os.Getenv("CUBE_WORKER_HOST")
	wport, _ := strconv.Atoi(os.Getenv("CUBE_WORKER_PORT"))
	wrk := worker.Worker{
		Queue: *queue.New(),
		Db:    make(map[uuid.UUID]*task.Task),
	}
	log.Printf("host: %s, port: %d\n", whost, wport)

	wrkAPI := worker.Api{Address: whost, Port: wport, Worker: &wrk}

	go wrk.RunTasks()
	go wrk.UpdateTasks()
	wrkAPI.Start()
}
