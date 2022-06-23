package handler

import (
	"bytes"
	"flag"
	"log"
	"net/http"

	"go-notes-app/db"
	"go-notes-app/pages"

	"github.com/gorilla/mux"
	g "github.com/maragudk/gomponents"
	"github.com/yuin/goldmark"
)

var gMark goldmark.Markdown = initGoldMark()
var AppRouter *mux.Router

func InitRouter() {
	var dir string
	flag.Parse()
	flag.StringVar(&dir, "assets", ".", "the directory to serve files from. Defaults to the current dir")
	
	router := mux.NewRouter()
	router.HandleFunc("/", createHandler(pages.Home())).Methods("GET")
	//router.NotFoundHandler = pages.ErrorPageHandler()
	router.HandleFunc("/about", createHandler(pages.About())).Methods("GET")
	router.HandleFunc("/playground", createHandler(pages.PlayGround())).Methods("GET")
	router.HandleFunc("/notes", createHandler(pages.Notes())).Methods("GET")
	router.HandleFunc("/notes/{title}", createNoteHandler("Notes")).Methods("GET")
	router.PathPrefix("/assets").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))
	router.NotFoundHandler = pages.ErrorPageHandler()
	AppRouter = router
}


func createHandler(title string, body g.Node) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pages.Page(title, r.URL.Path, body).Render(w)
	}
}


func createNoteHandler(title string) http.HandlerFunc {
	// Success - Render note document
	// Fail - Render fail page.
	return func(w http.ResponseWriter, r *http.Request) {
		var document db.NoteFile
		documentTitle := mux.Vars(r)["title"]
		log.Println(documentTitle)

		// Check from database the requested note
		documentFound, document := db.GetDocument(documentTitle)

		if(!documentFound) {
			pages.Page("Document not found!", r.URL.Path, pages.NotePage("<h1 class=text-center>Unfortunately the document was not found</h1>")).Render(w)
		} else {
			// Parse and return html as writable byte buffer
			var documentBuffer bytes.Buffer

			if err := gMark.Convert(document.File, &documentBuffer); err != nil {
				log.Println(err)
			}
			pages.Page(title, r.URL.Path, pages.NotePage(documentBuffer.String())).Render(w)
		}
	}
}

