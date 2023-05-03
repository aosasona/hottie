package main

import (
	"flag"

	"github.com/aosasona/hottie/pkg/hottie"
	log "github.com/charmbracelet/log"
)

func main() {
	addr := flag.String("addr", "localhost", "Address to serve the directory on")
	dir := flag.String("dir", ".", "Directory to serve")
	port := flag.Int("port", 3000, "Port to serve the directory on")
	reload := flag.Bool("reload", true, "Enable or disable hot reloading")

	flag.Parse()

	h := hottie.New().SetOpts(hottie.HottieOpts{
		Address:         *addr,
		Dir:             *dir,
		Port:            *port,
		EnableHotReload: *reload,
	})

	if err := h.Run(); err != nil {
		log.Fatal(err)
	}
}
