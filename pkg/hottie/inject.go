package hottie

import (
	"strings"
)

var (
	TARGETED_TAGS  = []string{"</body>", "</head>"}
	WEBSOCKET_CODE = `
<script type="text/javascript">
  var source = new EventSource("/_/sse");
  source.onmessage = function (event) {
    console.log(event.data);
  };
</script>`
)

func injectWebsocketCode(originalHTML []byte) []byte {
	strHTML := string(originalHTML)
	for _, tag := range TARGETED_TAGS {
		tagIndex := strings.Index(strHTML, tag)
		if tagIndex > -1 {
			newSection := WEBSOCKET_CODE + tag
			if tag == "</body>" {
				newSection = tag + WEBSOCKET_CODE
			}
			strHTML = strings.Replace(strHTML, tag, newSection, 1)
			break
		}
	}

	return []byte(strHTML)
}
