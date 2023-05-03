package hottie

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type hottie struct {
	addr            string
	dir             string
	port            int
	enableHotReload bool
	log             *log.Logger
}

type HottieOpts struct {
	Address         string
	Dir             string
	Port            int
	EnableHotReload bool
}

var notifChan = make(chan bool)

func New() *hottie {
	logger := log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: false,
		Prefix:          "🔥",
	})
	return &hottie{
		log: logger,
	}
}

func (h *hottie) SetOpts(opts HottieOpts) *hottie {
	if opts.Address == "" {
		opts.Address = "localhost"
	}
	h.addr = opts.Address
	h.dir = opts.Dir
	h.port = opts.Port
	h.enableHotReload = opts.EnableHotReload
	return h
}

func (h *hottie) Run() error {
	router := router.New()

	router.SaveMatchedRoutePath = true
	router.GET("/{path:*}", h.serveFile)

	h.log.Info(
		fmt.Sprintf(
			"Starting Hottie on port %d - serving %s\n\nVisit http://%s:%d in your browser\n",
			h.port,
			h.dir,
			h.addr,
			h.port,
		),
	)

	go h.watchForFileChanges()

	return fasthttp.ListenAndServe(fmt.Sprintf(":%d", h.port), router.Handler)
}
