package pages

import (
	"net/http"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

func ErrorPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errorPage().Render(w)
	}
}

func errorPage() g.Node {
	// HTML 5 boilerplate document
	return c.HTML5(c.HTML5Props{
		Title:		"404 - Not Found",
		Language: "en",
		Head: []g.Node{
			html.Link(html.Rel("stylesheet"), html.Href("https://unpkg.com/tailwindcss@2.2.19/dist/base.min.css")),
			html.Link(html.Rel("stylesheet"), html.Href("https://unpkg.com/tailwindcss@2.2.19/dist/components.min.css")),
			html.Link(html.Rel("stylesheet"), html.Href("https://unpkg.com/@tailwindcss/typography@0.4.0/dist/typography.min.css")),
			html.Link(html.Rel("stylesheet"), html.Href("https://unpkg.com/tailwindcss@2.2.19/dist/utilities.min.css")),
			html.Link(html.Rel("stylesheet"), html.Href("/assets/css/app.css")),
		},
		Body: []g.Node{
			Navbar(),
			errorBody(),
		},
	})
}

func errorBody() g.Node {
	return html.Div(
		html.Class("max-w-5xl mx-auto prose md:prose-lg lg:prose-xl px-2 mt-10 sm:px-6 lg:px-8"),
		html.Div(
			html.Class("flex-col"),
			html.H1(html.Class("text-center"), g.Text("Oops, page not found")),
			html.Img(html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/gopherize-me_7Azjt_g4T.png?ik-sdk-version=javascript-1.4.3&updatedAt=1652731068953")),
		),
	)
}