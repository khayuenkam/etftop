package main

import "github.com/khayuenkam/etftop/etftop/etftop"

import "log"

func main() {
	et, err := etftop.NewEtfTop()
	if err != nil {
		log.Fatal(err)
	}
	et.Run()
}
