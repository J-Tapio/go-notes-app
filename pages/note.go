package pages

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func NotePage(mdBody string) g.Node {
	return html.Div(
		g.Raw(mdBody),
	)
}
