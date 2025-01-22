// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 187.

// Sorting sorts a music playlist into a variety of orders.
package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

// !+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Delilah", "From the Roots Up", 2013, length("3m38s")},
	{"Go", "Delilah", "From the Roots Up", 2014, length("3m38s")},
	{"Come", "Delilah", "Come Back", 2014, length("4m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go", "Moby", "Sun", 1991, length("2m33s")},
	{"Come", "Moby", "Moon", 1995, length("5m27s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Go Ahead", "Alicia Keys", "Who you are", 2007, length("4m36s")},
	{"Ready 2 Go", "Alicia Keys", "Who you are", 2007, length("5m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Come", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Ready 2 Go", "Martin Solveig", "Smile Face", 1988, length("3m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

//!-main

// !+printTracks
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

//!-printTracks

//!-lengthcode

func main() {

	sortFlag := flag.String("sort", "Artist", "eg: -sort=\"Artist,Year,Title\", Default: Artist") // bind sortFLog to -sort
	flag.Parse()                                                                                  // parse

	sortField := strings.Split(*sortFlag, ",")
	fmt.Println("Parsed sort order=", sortField)

	sort.SliceStable(tracks, func(i, j int) bool {

		for _, t := range sortField {
			t = strings.ToLower(t)
			switch t {
			case "artist":
				if tracks[i].Artist != tracks[j].Artist {
					return tracks[i].Artist < tracks[j].Artist
				}
			case "year":
				if tracks[i].Year != tracks[j].Year {
					return tracks[i].Year < tracks[j].Year
				}
			case "title":
				if tracks[i].Title != tracks[j].Title {
					return tracks[i].Title < tracks[j].Title
				}
			case "album":
				if tracks[i].Album != tracks[j].Album {
					return tracks[i].Album < tracks[j].Album
				}
			case "length":
				if tracks[i].Length != tracks[j].Length {
					return tracks[i].Length < tracks[j].Length
				}

			default:
				fmt.Println("Unsupported sort:", t)
				os.Exit(1)
			}

		}
		return false
	})
	fmt.Println("Sorted sort order=", sortField)
	printTracks(tracks)

}
