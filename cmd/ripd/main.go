package main

import (
	"fmt"

	"github.com/coreswitch/zebra/pkg/server/rip"
)

func main() {
	server := rip.NewServer()
	fmt.Println("ripd", server)
}
