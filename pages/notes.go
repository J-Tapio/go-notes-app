package pages

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"

	"go-notes-app/db"
)

func noteLink(title string, tags []string) g.Node {
	noteLink := "/notes/" + title

	return html.Div(
		html.Class("mb-8 note-container pb-2 pl-2 rounded"),
		html.A(
			html.Class("note-link"),
			html.Href(noteLink),
			html.P(
				html.Class("note-title"),
				g.Text(title),
			),
			html.Div(
				html.Class("flex pt-5"),
				g.Group(g.Map(len(tags), func(i int) g.Node {
					return html.Div(
						html.Class("rounded-full tag mr-5 py-1 px-3 text-black text-sm"),
						g.Text(tags[i]),
					)
				})),
			),
		))
}

func Notes() (string, g.Node) {
	//TODO: Reconsider if available notes should be referred from elsewhere than db
	//db.GetAvailableNotes()

	if len(db.AvailableNotes) == 0 {
		// return html with error message
		return "Notes", html.Div(
			html.H1(g.Text("Unfortunately something went wrong, notes not available.")),
		)
	} else {
		return "Notes", html.Div(
			html.H1(g.Text("Notes")),
			html.Div(html.Class("flex-column"),
				g.Group(g.Map(len(db.AvailableNotes), func(i int) g.Node {
					return noteLink(db.AvailableNotes[i].Title, db.AvailableNotes[i].Tags)
				})),
			),
		)
	}
}
