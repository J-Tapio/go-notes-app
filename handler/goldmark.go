package handler

import (
	"github.com/aziis98/goldmark-latex"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// Styles
// https://github.com/alecthomas/chroma/tree/master/styles

func initGoldMark() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.Table, latex.NewLatex(), highlighting.NewHighlighting(
			highlighting.WithStyle("fruity"),
		)),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		),
	)
}
