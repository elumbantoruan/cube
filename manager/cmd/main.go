package main

import (
	"cube/manager"
	"fmt"
	"os"
	"strconv"
)

func main() {

	mhost := os.Getenv("CUBE_MANAGER_HOST")
	mport, _ := strconv.Atoi(os.Getenv("CUBE_MANAGER_PORT"))

	whost1 := os.Getenv("CUBE_WORKER1_HOST")
	whost2 := os.Getenv("CUBE_WORKER2_HOST")
	whost3 := os.Getenv("CUBE_WORKER3_HOST")

	wport, _ := strconv.Atoi(os.Getenv("CUBE_WORKER_PORT"))

	fmt.Println("Starting Cube manager")

	workers := []string{
		fmt.Sprintf("%s:%d", whost1, wport),
		fmt.Sprintf("%s:%d", whost2, wport),
		fmt.Sprintf("%s:%d", whost3, wport),
	}
	m := manager.New(workers, "epvm")
	mapi := manager.Api{Address: mhost, Port: mport, Manager: m}

	go m.ProcessTasks()
	go m.UpdateTasks()
	go m.DoHealthChecks()

	mapi.Start()
}
