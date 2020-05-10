package sorting

import (
	"fmt"
	"sort"
	"testing"
)

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func TestByArtiest(t *testing.T) {
	source := make([]*Track, len(tracks), cap(tracks))
	copy(source, tracks)
	printTracks(source)
	sort.Sort(byArtist(source))
	printTracks(source)
}

func TestByArtiestReverse(t *testing.T) {
	source := make([]*Track, len(tracks), cap(tracks))
	copy(source, tracks)
	printTracks(source)
	sort.Sort(sort.Reverse(byArtist(source)))
	printTracks(source)
}

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func TestByYear(t *testing.T) {
	source := make([]*Track, len(tracks), cap(tracks))
	copy(source, tracks)
	printTracks(source)
	sort.Sort(byYear(source))
	printTracks(source)
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func TestByCustomSort(t *testing.T) {
	source := make([]*Track, len(tracks), cap(tracks))
	copy(source, tracks)
	printTracks(source)
	sort.Sort(customSort{source, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(source)
}

func TestSort(t *testing.T) {
	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values))
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
}
