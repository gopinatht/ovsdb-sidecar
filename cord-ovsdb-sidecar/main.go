package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	var ovsDBSocket string
	flag.StringVar(&ovsDBSocket, "ovs-socket", "/run/openvswitch/db.sock", "The path to the OVSDB unix domain socket")

	// Wait for unix socket created by OVSDB to exist.
	// This implies the OVSDB container has been created
	_, err := os.Stat(ovsDBSocket)
	for os.IsNotExist(err) {
		log.Println("Waiting for the OVSDB server unix socket to appear")
		time.Sleep(5 * time.Second) // Sleep 5 seconds before checking for domain socket again
		_, err = os.Stat(ovsDBSocket)
	}

	// Start the OVSDB server with the passive TCP port specification
	cmd := exec.Command("ovs-appctl", "-t", "ovsdb-server", "ovsdb-server/add-remote", "ptcp:6641")
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...\n")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v\n", err)

	// listening OS shutdown singal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
