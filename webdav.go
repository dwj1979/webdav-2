package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/webdav"
)

func main() {
	// default directory to serve from is current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	port := flag.String("port", "8080", "the port to serve on")
	dir := flag.String("dir", cwd, "the path to the directory to serve")
	flag.Parse()

	// print help if user gives no options
	if len(os.Args) == 1 {
		flag.Usage()
		os.Exit(1)
	}
	// set up the webdav server
	server := &http.Server{
		Addr: "localhost:" + *port,
		Handler: &webdav.Handler{
			FileSystem: webdav.Dir(*dir),
			LockSystem: webdav.NewMemLS(),
		},
	}
	// stop the server on the KILL and INTERRUPT signals
	stopper := make(chan os.Signal, 2)
	signal.Notify(stopper, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stopper
		// kills all current connections
		err := server.Close()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}()
	fmt.Printf("serving WebDAV at http://localhost:%s for %s\n", *port, *dir)
	log.Fatal(server.ListenAndServe())
}
