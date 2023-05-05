package hottie

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/log"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type event int

const (
	_ event = iota
	full_reload
	css_reload
)

type hottie struct {
	addr            string
	dir             string
	port            int
	enableHotReload bool
	openBrowser     bool
	log             *log.Logger
}

type HottieOpts struct {
	Address         string
	Dir             string
	Port            int
	EnableHotReload bool
	OpenBrowser     bool
}

var notifChan = make(chan event, 16)

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
		opts.Address = "127.0.0.1"
	}
	h.addr = opts.Address
	h.dir = opts.Dir
	h.port = opts.Port
	h.enableHotReload = opts.EnableHotReload
	h.openBrowser = opts.OpenBrowser
	return h
}

func (h *hottie) openInBrowser() {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}

	args = append(args, fmt.Sprintf("http://%s:%d", h.addr, h.port))
	if err := exec.Command(cmd, args...).Start(); err != nil {
		h.log.Error(err)
	}
}

func (h *hottie) Run() error {
	router := router.New()

	router.SaveMatchedRoutePath = true
	router.GET("/_/sse", h.handleSSE)
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

	if h.openBrowser {
		h.openInBrowser()
	}
	go h.watchForFileChanges()
	defer close(notifChan)

	return fasthttp.ListenAndServe(fmt.Sprintf(":%d", h.port), router.Handler)
}
