package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func serve(file, listen string) error {
	h := Handler{target: file}

	http.Handle("/", &h)
	log.Println("starting server on", listen)
	log.Fatal(http.ListenAndServe(listen, nil))
	return nil
}

func run() error {
	flag.Parse()

	listen := flag.String("http", ":2345", "listen port")
	if len(flag.Args()) > 1 {
		return errors.New("too many args")
	} else if len(flag.Args()) == 0 {
		return errors.New("no files specified")
	}
	file := flag.Args()[0]

	return serve(file, *listen)
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
