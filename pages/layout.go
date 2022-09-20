package pages

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

// Gomponents examples use dot import fot html.
// Use of 'html' is verbose but I didnt want to go against default
// behavior of Go (check can be disabled apparently)
// https://staticcheck.io/docs/configuration/options/
// https://github.com/golang/lint/issues/179

func Page(title, path string, body g.Node) g.Node {
	// HTML 5 boilerplate document
	return c.HTML5(c.HTML5Props{
		Title:    title,
		Language: "en",
		Head: []g.Node{
			html.Link(html.Rel("icon"), html.Type("image/x-icon"), html.Href("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/favicon-32x32_9t5v8KhSP.png?ik-sdk-version=javascript-1.4.3&updatedAt=1656174819184")),
			html.Link(html.Rel("stylesheet"), html.Href("/assets/css/base.min.css")),
			html.Link(html.Rel("stylesheet"), html.Href("/assets/css/components.min.css")),
			html.Link(html.Rel("stylesheet"), html.Href("/assets/css/typography.min.css")),
			html.Link(html.Rel("stylesheet"), html.Href("/assets/css/utilities.min.css")),
			html.Link(html.Rel("stylesheet"), html.Href("/assets/css/app.css")),
			html.Link(html.Rel("stylesheet"), html.Href("https://cdn.jsdelivr.net/simplemde/latest/simplemde.min.css")),
			html.Script(html.Src("https://cdn.jsdelivr.net/simplemde/latest/simplemde.min.js")),
			html.Script(html.Src("assets/editor.toolbar.js")),
		},
		Body: []g.Node{
			Navbar(),
			Container(
				body,
			),
			Footer(),
		},
	})
}

func Navbar() g.Node {
	return html.Nav(html.Class("navbar sticky top-0 px-2 py-4"),
		html.Div(html.Class("container flex max-w-5xl flex-wrap items-center mx-auto"),
			html.Div(html.Class("flex"),
				html.A(html.Href("/"),
					html.Img(html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/go-eyes-right_y1QEbTd8j.png?ik-sdk-version=javascript-1.4.3&updatedAt=1652719256767"), html.Alt("Gopher mascot"), html.Class("h-10 w-auto pl-5")),
				),
			),
			html.Ol(html.Class("flex"),
				html.Li(html.A(html.Class("active:text-white md:text-xl hover:text-gray-50 px-5"), html.Href("/"), g.Text("Home"))),
				html.Li(html.A(html.Class("active:text-white text-gray md:text-xl hover:text-gray-50 px-5"), html.Href("/notes"), g.Text("Notes"))),
				html.Li(html.A(html.Class("active:text-white text-gray md:text-xl hover:text-gray-50 px-5"), html.Href("/about"), g.Text("About"))),
				html.Li(html.A(html.Class("hidden sm:inline active:text-white text-gray md:text-xl hover:text-gray-50 px-5"), html.Href("/playground"), g.Text("Playground"))),
			),
		),
	)
}

func Container(children ...g.Node) g.Node {
	return html.Div(html.Class("max-w-5xl mx-auto prose md:prose-lg lg:prose-xl px-2 mt-10 sm:px-6 lg:px-8"), g.Group((children)))
}

type socialLink struct {
	link, img, alt, title string
}

var socialLinks = []socialLink{
	{
		alt:   "Github link",
		link:  "https://github.com/J-Tapio",
		img:   "https://ik.imagekit.io/htg3gsxgz/Markdown_archives/github_oskCjtxkj.svg?ik-sdk-version=javascript-1.4.3&updatedAt=1655457450302",
		title: "Main account",
	},
	{
		alt:   "Github link",
		link:  "https://github.com/JtapioT",
		img:   "https://ik.imagekit.io/htg3gsxgz/Markdown_archives/github_oskCjtxkj.svg?ik-sdk-version=javascript-1.4.3&updatedAt=1655457450302",
		title: "Account used while being in Strive school",
	},
	{
		alt:  "Twitter link",
		link: "https://twitter.com/JuhaTurpeinen",
		img:  "https://ik.imagekit.io/htg3gsxgz/Markdown_archives/twitter_yZMymzRM0.svg?ik-sdk-version=javascript-1.4.3&updatedAt=1652724047895",
	},
	{
		alt:  "Instagram link",
		link: "https://www.instagram.com/tapiolapio/",
		img:  "https://ik.imagekit.io/htg3gsxgz/Markdown_archives/instagram_sGHQ1hYua.svg?ik-sdk-version=javascript-1.4.3&updatedAt=1652724047449",
	},
	{
		alt:  "Researchgate link",
		link: "https://www.researchgate.net/profile/Juha_Tapio_Turpeinen",
		img:  "https://ik.imagekit.io/htg3gsxgz/Markdown_archives/researchgate_cLdis8c2D.svg?ik-sdk-version=javascript-1.4.3&updatedAt=1652724047744",
	},
	{
		alt:  "LinkedIn link",
		link: "https://fi.linkedin.com/in/juha-tapio-turpeinen-a89974137",
		img:  "https://ik.imagekit.io/htg3gsxgz/Markdown_archives/linkedin_01dXNmtjy.svg?ik-sdk-version=javascript-1.4.3&updatedAt=1652724047500",
	},
}

func contactLink(alt, link, title, img string) g.Node {
	return html.Li(
		html.Class("px-6"),
		html.A(
			html.Href(link),
			html.Rel("noopener noreferrer"),
			html.Target("_blank"),
			html.Alt(alt),
			html.TitleAttr(title),
			html.Img(
				html.Class("social h-8 w-auto"),
				html.Src(img),
			),
		),
	)
}

func Footer() g.Node {
	return html.Footer(html.Class("py-8 text-center"),
		html.Div(
			html.Class("flex justify-between items-end"),
			html.Img(html.Class("h-10 sm:h-14 md:h-20 w-auto"), html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/golang-gopher-left_Ep0uc1n6X.png?ik-sdk-version=javascript-1.4.3&updatedAt=1652719257184")),
			html.Div(
				html.Class("sm:text-lg md:text-xl"),
				html.P(g.Text("Juha-Tapio Turpeinen")),
				html.Ol(
					html.Class("prose-none flex mt-4"),
					g.Group(g.Map(len(socialLinks), func(i int) g.Node {
						return contactLink(socialLinks[i].alt, socialLinks[i].link, socialLinks[i].title, socialLinks[i].img)
					})),
				),
			),
			html.Img(html.Class("h-10 sm:h-14 md:h-20 w-auto"), html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/golang-gopher-right_Tcgh-recv.png?ik-sdk-version=javascript-1.4.3&updatedAt=1652719257280")),
		),
	)
}
