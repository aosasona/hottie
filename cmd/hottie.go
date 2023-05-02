package main

import (
	"flag"
	"fmt"

	"github.com/aosasona/hottie/pkg/hottie"
	log "github.com/charmbracelet/log"
)

func main() {
	d := flag.String("d", ".", "Directory to serve")
	p := flag.Int("p", 3000, "Port to serve the directory on")

	flag.Parse()

	// dereference once to get flag value
	port := *p
	dir := *d
	h := hottie.New().SetDir(dir).SetPort(port)

	log.Info(
		fmt.Sprintf(
			"Starting Hottie on port %d - serving %s\n\nVisit http://127.0.0.1:%d in your browser\n",
			port,
			dir,
			port,
		),
	)

	if err := h.Run(); err != nil {
		log.Fatal(err)
	}
}
