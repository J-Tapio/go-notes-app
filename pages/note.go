package pages

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func NotePage(mdBody string) g.Node {
	return html.Div(
		html.Class("prose md:prose-lg lg:prose-2xl mx-auto"),
		g.Raw(mdBody),
	)
}
