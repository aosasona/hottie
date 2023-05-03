package hottie

import (
	"fmt"
	"strings"
)

var (
	TARGETED_TAGS  = []string{"</body>", "</head>"}
	WEBSOCKET_CODE = `<script type="text/javascript">
/* This code is injected by Hottie to enable hot reloading. */
console.log("Hottie is watching for changes...");

const ws = new WebSocket("%s");
ws.onmessage = function(e) {
  if (e.data == "reload") {
    window.location.reload();
  }
}
</script>`
)

func injectWebsocketCode(originalHTML []byte, websocketAddr string) []byte {
	strHTML := string(originalHTML)
	for _, tag := range TARGETED_TAGS {
		tagIndex := strings.Index(strHTML, tag)
		if tagIndex > -1 {
			newSection := WEBSOCKET_CODE + tag
			if tag == "</body>" {
				newSection = tag + WEBSOCKET_CODE
			}
			newSection = fmt.Sprintf(newSection, websocketAddr)
			strHTML = strings.Replace(strHTML, tag, newSection, 1)
			break
		}
	}

	return []byte(strHTML)
}
