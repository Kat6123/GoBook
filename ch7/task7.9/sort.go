package main

import (
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func byTrack(reversed bool) {
	custom := sort.Interface(customSort{
		tracks,
		func(x, y *Track) bool { return x.Title < y.Title }})

	if reversed {
		custom = sort.Reverse(custom)
	}

	sort.Sort(custom)
}

func byArtist(reversed bool) {
	custom := sort.Interface(customSort{
		tracks,
		func(x, y *Track) bool { return x.Artist < y.Artist }})
	if reversed {
		custom = sort.Reverse(custom)
	}

	sort.Sort(custom)
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
