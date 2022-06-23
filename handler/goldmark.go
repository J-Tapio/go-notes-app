package handler

import(
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/aziis98/goldmark-latex"
	"github.com/yuin/goldmark-highlighting"
)

// Styles
// https://github.com/alecthomas/chroma/tree/master/styles


func initGoldMark() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(extension.GFM, highlighting.NewHighlighting(
			highlighting.WithStyle("fruity"),
		)),
		goldmark.WithExtensions(latex.NewLatex()),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
		),
	)
}
