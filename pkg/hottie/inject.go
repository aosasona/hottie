package hottie

import (
	"strings"
)

var (
	TARGETED_TAGS  = []string{"</body>", "</head>"}
	WEBSOCKET_CODE = `
<!-- This code was injected by Hottie (https://github.com/aosaona/hottie) -->
<script type="text/javascript">
	if (typeof(EventSource) !== "undefined") {
		function reloadStylesheets() {
			let qs = '?reload=' + new Date().getTime();
			var sheets = document.getElementsByTagName('link');
			for (var i = 0; i < sheets.length; i++) {
				var elem = sheets[i];
				var rel = elem.rel;
				if (elem.href && typeof rel != 'string' || rel.length == 0 || rel.toLowerCase() == 'stylesheet') {
					var url = elem.href.replace(/(&|\?)_cacheOverride=\d+/, '');
					elem.href = url + (url.indexOf('?') >= 0 ? '&' : '?') + '_cacheOverride=' + (new Date().valueOf());
				}
			}
		}

  	var source = new EventSource("/_/sse");
  	source.onmessage = function (event) {
  		if (event.data == "_full") {
				window.location.reload();
  		} else if (event.data == "_style") {
				reloadStylesheets();
  		}
  	}
  } else {
  console.log("Your browser doesn't support SSE unforrunately :()");
  }
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
