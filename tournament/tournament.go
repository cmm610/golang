package tournament

import (
	"errors"
	"strings"
	"sort"
	"io"
	"fmt"
)
const (
	headerLine = "Team                           | MP |  W |  D |  L |  P\n"
	lineFmt    = "%-31s| %2d | %2d | %2d | %2d | %2d\n"
)

type TeamStats struct {
	w, l, d, mp, p	int
	name 			string
}

func (t TeamStats) String() string {
	return fmt.Sprintf(lineFmt, t.name, t.mp, t.w, t.d, t.l, t.p)
}
type Teams map[string]*TeamStats
type TeamsStats []TeamStats
type By func(t1, t2 *TeamStats) bool
func (by By) Sort(teamStats []TeamStats) {
	ps := &teamSorter{
		teams: teamStats,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ps)
}
type teamSorter struct {
	teams []TeamStats
	by func(t1, t2 *TeamStats) bool
}

// Len is part of sort.Interface.
func (s *teamSorter) Len() int {
	return len(s.teams)
}

// Swap is part of sort.Interface.
func (s *teamSorter) Swap(i, j int) {
	s.teams[i], s.teams[j] = s.teams[j], s.teams[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *teamSorter) Less(i, j int) bool {
	return s.by(&s.teams[i], &s.teams[j])
}


var teams = []string{"Allegoric Alaskians", "Blithering Badgers", "Devastating Donkeys", "Courageous Californians"}
var outcomes = []string{"win", "lose", "draw"}

func checkValid(s string, ss []string) error {
	i := sort.SearchStrings(ss, s)
	if i == -1 {
		return errors.New("string not valid")
	}
	return nil
}

func buildStats(outcome string, tms Teams, t1 string, t2 string) error {
	tms[t1].name = t1
	tms[t2].name = t2
	switch outcome {
	case "win" :
		tms[t1].w ++
		tms[t2].l ++
		tms[t1].p += 3
		tms[t1].mp ++
		tms[t2].mp ++
	case "loss" :
		tms[t1].l ++
		tms[t2].w ++
		tms[t2].p += 3
		tms[t1].mp ++
		tms[t2].mp ++
	case "draw" :
		tms[t1].d ++
		tms[t2].d ++
		tms[t1].p ++
		tms[t2].p ++
		tms[t1].mp ++
		tms[t2].mp ++
	//
	//default :
	//	fmt.Println(outcome)
	//	return errors.New("outcome not valid")
	}
	return nil
}

func Tally(r io.Reader, w io.Writer) error {
	bs := make([]byte, 400)
	_, err := r.Read(bs)
	if err != nil {
		errors.New("error reading from file")
	}
	tms := make(Teams)
	for i:=0; i<len(teams); i++ {
		tms[teams[i]] = &TeamStats{mp: 0, w: 0, d: 0, l: 0, p: 0}
	}

	sa := strings.Split(strings.TrimSpace(string(bs)), "\n")
	teamSlice := []TeamStats{}
	for _, oneLine := range sa {
		if oneLine == "" {
			continue
		} else {
			statsArr := strings.Split(oneLine, ";")
			if len(statsArr) != 3 {
				return errors.New("stats not fully filled in")
			}
			t1, t2, outcome := statsArr[0], statsArr[1], statsArr[2]

			err = buildStats(outcome, tms, t1, t2)
			if err != nil {
				return err
			}
		}
	}
	for _, value := range tms {
		teamSlice = append(teamSlice, *value)
	}

	wins := func(t1, t2 *TeamStats) bool {
		return t1.w > t2.w
	}

	By(wins).Sort(teamSlice)
	printOut := headerLine
	for _, team := range teamSlice {
		printOut += team.String()
	}
	_, err = w.Write([]byte(printOut))
	return err
}