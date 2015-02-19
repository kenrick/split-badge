package main

import (
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/ajstarks/svgo"
)

var height = 20

func main() {
	port := flag.String("port", "8080", "Server Port")
	flag.Parse()

	log.Println("Starting server on port", *port)
	http.HandleFunc("/", badgeHandler)
	http.ListenAndServe(":"+*port, nil)
}

func badgeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	points := r.URL.Query()["points"]
	if len(points) > 0 {
		log.Println("Creating a badge worth", points[0])
		createBadge(w, points[0]+"pts")
	}
}

func createBadge(w io.Writer, points string) {
	wordWidth := 7 * len(points)
	width := 45 + wordWidth

	s := svg.New(w)
	s.Start(width, height)
	s.Rect(0, 0, width, height, "fill: #000000; stroke: #000000")
	s.Path("M0,0 L28.8975184,0 L42,0 L42,10.1277146 L42,20 L0,20 L0,0 Z", "fill: #4990E2; stroke: #4990E2")
	s.Text(9, 13, "split", "font-size: 11; font-weight: bold; fill: #FFFFFF; font-family:'Open Sans'")
	s.Text(47, 13, points, "font-size: 11; font-weight: bold; fill: #FFFFFF; font-family:'Open Sans'")
	s.End()
}
