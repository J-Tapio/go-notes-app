package pages

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

func About() (string, g.Node) {
	return "About", html.Div(
		html.Div(
			html.Class("flex items-center"),
			html.Img(html.Class("h-60 w-60 object-cover rounded-full"), html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/picture-compressed_OSOBwCmtC.webp?ik-sdk-version=javascript-1.4.3&updatedAt=1652719500043")),
			html.Div(
				html.Class("flex-column pl-5"),
				html.H1(g.Text("Whoami?")),
				html.H4(html.Class(""), g.Text("Juha-Tapio Turpeinen")),
				html.H6(html.Class(""), g.Text("Physiotherapist OMT")),
				html.H6(html.Class(""), g.Text("MSc. Sport Sciences")),
				html.H6(html.Class(""), g.Text("Fullstack Web developer")),
			),
		),
		html.Div(
			html.Class("flex items-center"),
			html.P(g.Text("Currently living in Berlin, Germany with my girlfriend and our little daughter. Originally from Finland.")),
			html.Img(html.Class("h-32 md:h-96 object-cover w-auto mx-auto"), html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/suomikartta_soOAbgHNG.webp?ik-sdk-version=javascript-1.4.3&updatedAt=1652719258402")),
		),
		html.Div(
			html.Class("flex items-center"),
			html.Img(html.Class("h-32 md:h-96 object-cover w-auto mx-auto"), html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/therapist_u9uHNSxRu.webp?ik-sdk-version=javascript-1.4.3&updatedAt=1652719258616")),
			html.P(html.Class("sm: pl-8 md: pl-3"), g.Raw(`I have worked as a Physiotherapist for almost a decade and also have MSc. in High Performance Sports. Besides health and fitness, I have been very interested in technology but never thought that I could learn to code. Well, eventually came writing of my publication for <a target="_blank" rel="noopener noreferrer" href="https://onlinelibrary.wiley.com/doi/10.1111/sms.13733">Scandinavian Journal of Medicine &  Science in Sports</a>.<small class="text-xs">(Publication link)</small>`)),
		),
		html.Div(
			html.Class("flex items-center"),
			html.P(html.Class("sm: pr-8 md: pr-3"), g.Raw(`I had to learn use of <a target="_blank" rel="noopener noreferrer" href="https://www.r-project.org">R-language</a> for statistical analyses. Learning R to only accomplish statistical analyses was a difficult task. I spend many, many hours on Stack Overflow pages and eventually got it working. Success followed by closing of hundred browser tabs is gratifying. Somewhat familiar process still sometimes with difficult development issues!`)),
			html.Img(html.Class("h-32 md:h-96 object-cover w-auto mx-auto"), html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/web-design_BdNBH_yUE.webp?ik-sdk-version=javascript-1.4.3&updatedAt=1652719258578")),
		),
		html.Div(
			html.Class("flex items-center"),
			html.Img(html.Class("h-32 md:h-96 object-cover w-auto mx-auto"), html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/shredder_RtL3xIjmL.webp?ik-sdk-version=javascript-1.4.3&updatedAt=1652719257995")),
			html.P(html.Class("sm: pl-8 md: pl-3"), g.Raw(`Later in 2019 a good friend of mine, <a target="_blank" rel="noopener noreferrer" href="https://www.discogs.com/artist/3531308-Juuso-Soinio">shredder of guitars and keyboard</a>, suggested to start learning Javascript. I was interested for 1-2 months but then somehow curiosity ended to programming in overall. Then almost a year later I started learning again and kept going. Credits where it's due as I think it has been a blessing to have a mentor with decade of experience shredding keyboard and vast knowledge of the field.`)),
		),
		html.Div(
			html.Class("flex items-center justify-center"),
			html.Img(html.Class("h-32 md:h-96 w-auto object-cover"), html.Src("https://ik.imagekit.io/htg3gsxgz/Markdown_archives/iron-mike_6Jgof9EOE.webp?ik-sdk-version=javascript-1.4.3&updatedAt=1652719257859")),
			html.P(g.Raw(`Success breeds confidence.<br>Confidence breeds success.<br><small>- Iron Mike</small>`)),
		),
	)
}