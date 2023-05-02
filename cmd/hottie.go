package main

import (
	"flag"

	"github.com/aosasona/hottie/pkg/hottie"
	log "github.com/charmbracelet/log"
)

func main() {
	dir := flag.String("d", ".", "Directory to serve")
	port := flag.Int("p", 3000, "Port to serve the directory on")
	hotReload := flag.Bool("h", false, "Enable or disable hot reloading")

	flag.Parse()

	h := hottie.New().SetDir(*dir).SetPort(*port).SetEnableHotReload(*hotReload)

	if err := h.Run(); err != nil {
		log.Fatal(err)
	}
}
