package pages

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func PlayGround() (string, g.Node) {
	return "Markdown editor", html.Div(
		html.H1(g.Text("Markdown Editor")),
		html.Textarea(),
		g.Raw(`<script>
		var mdNotes = new SimpleMDE();
		mdNotes.value("# Markdown editor\n## [Simple markdown editor](https://github.com/sparksuite/simplemde-markdown-editor)\n\n\n> Future feature could be authentication followed by possibility for me to online edit my old documents or write new one online.\nAt the moment though, I like writing notes with VsCode and use Extended Preview extension.\n\nIf you are not familiar with [Markdown](https://www.markdownguide.org), I highly suggest you to try it out 🙂\nCheck the preview of Markdown by clicking the eye-icon from the toolbar.");
		</script>`),
	)
}
