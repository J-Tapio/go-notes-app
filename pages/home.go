package pages

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func Home() (string, g.Node) {
	return "Welcome!", html.Div(
		html.Img(html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/gopher-laptop_CveXuKP2B.webp?ik-sdk-version=javascript-1.4.3&updatedAt=1652719257522"),html.Class("mt-6 h-48 md:h-56 w-auto mx-auto md:mx-0")),
		html.H1(g.Text("Markdown archives")),
		html.P(g.Raw(`These pages have been created with <a target="_blank" rel="noopener noreferrer" href="https://go.dev">Go</a>.`)),
		html.Div(
			html.Class("flex items-center"),
			html.P(g.Raw(`Page is using <a target="_blank" rel="noopener noreferrer" href="https://tailwindcss.com">TailwindCSS</a> for styling with the help of excellent project by <a target="_blank" rel="noopener noreferrer" href="https://www.maragu.dk">Markus Wüstenberg</a>, <a target="_blank" rel="noopener noreferrer" href="https://www.gomponents.com">Gomponents</a>. <small></small>`)),
			html.Img(html.Class("h-24 w-auto mx-auto"), html.Alt("Gomponents gopher; Credit Markus Wüstenberg - Gomponents.com"), html.Src("https://www.gomponents.com/images/logo.png")),
		),
		html.H2(g.Text("What is this page about?")),
		html.P(g.Text("Here you will find selection of notes I have written while learning variety of topics related to Web development and programming in overall.")),
		html.P(g.Text("Main goal of this project was to implement pages which can render with certain theme my old notes, Markdown files, as html whenever I need to revive something.")),
		html.Div(
			html.Class("flex-col pt-10"),
			html.H3(g.Text("Why Golang?")),
			html.P(html.Class("pt-3"), g.Text("Originally I started learning coding and web development through Javascript. While I was learning Typescript later (which is still a bumpy road), I somehow started thinking if I should try to learn some statically-typed language.")),
			html.Div(
				html.Class("flex items-center"),
				html.P(html.Class("pt-3"), g.Text("For some odd reason I thought it would enhance my understanding of development with types and help with learning of Typescript?! Big brain time, indeed...not.")),
				html.Img(html.Class("h-24 w-auto mx-auto"), html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/golang-brainer_czxJV-r6U.png?ik-sdk-version=javascript-1.4.3&updatedAt=1652719256788")),
			),
			html.P(html.Class("pt-3"), g.Text(
				"Nevertheless, I was first considering C++ but somehow Golang and super-cool Gopher mascot caught my eyes which somehow intrigued me to learn more about it. There were also lot of good buzz around Rust language but well, mascot won over chainring. Difficult decisions."),
			),
			html.P(html.Class("pt-3"), g.Text(
				"Indeed, coming from Javascript language to statically-typed language has been somewhat intimidating in many ways but also rewarding experience."),
			),
			html.P(html.Class("pt-3"), g.Text(
				"Even though my mileage with Golang is short I already like it very much and intend to learn more. Also, hopefully Go best practices start to show in my code!"),
			),
			html.Div(
				html.Img(
					html.Class("mx-8 h-48 md:h-60 w-auto mx-auto md:mx-0"),
					html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/golang-go_peNjmp5pi.gif?ik-sdk-version=javascript-1.4.3&updatedAt=1652719257212"),
				),
			),
		),
	)
}