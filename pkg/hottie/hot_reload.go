package hottie

import (
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func (h *hottie) watchForFileChanges() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		h.log.Fatal(err)
	}
	defer watcher.Close()

	dirContent, err := os.ReadDir(h.dir)
	if err != nil {
		h.log.Fatal(err)
	}

	for _, file := range dirContent {
		if file.IsDir() {
			err = watcher.Add(h.dir + "/" + file.Name())
			if err != nil {
				// chug along if we can't watch a directory
				h.log.Errorf("failed to watch directory: ", err)
				continue
			}
		}
	}

	err = watcher.Add(h.dir)
	if err != nil {
		h.log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) ||
					event.Has(fsnotify.Remove) {
					if strings.Contains(event.Name, ".null-ls") {
						continue
					}
					evt := full_reload
					if strings.Contains(event.Name, ".css") {
						evt = css_reload
					}
					h.log.Infof("change detected -> %s", event.Name)
					notifChan <- evt
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				h.log.Errorf("error: %v", err)
			}
		}
	}()

	<-make(chan struct{})
}
