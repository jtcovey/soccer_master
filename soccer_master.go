package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type matchResult struct {
	team1      string
	team2      string
	team1score int
	team2score int
}

type TeamScoring struct {
	TeamName string
	Score    int
}

type MatchDay []TeamScoring

func (m MatchDay) Len() int      { return len(m) }
func (m MatchDay) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m MatchDay) Less(i, j int) bool {
	if m[i].Score == m[j].Score {
		return m[i].TeamName < m[j].TeamName
	}
	return m[i].Score > m[j].Score
}

var matchRegex = regexp.MustCompile(`(\D+\s)+(\d+)$`)

func main() {
	var scanner *bufio.Scanner
	if len(os.Args) < 2 {
		// use stdin
		scanner = bufio.NewScanner(os.Stdin)
	} else {
		// use file passed as arg
		filepath := os.Args[1]

		file, err := os.Open(filepath)
		if err != nil {
			exit(err)
		}
		defer file.Close()

		scanner = bufio.NewScanner(file)
	}

	season := make(map[string]int)

	matchDay := make(map[string]int)
	matchDayCount := 1

	for scanner.Scan() {
		line := scanner.Text()
		match, err := getMatchResult(line, matchRegex)

		if err == nil {
			// If day is over, announce winners and start new day
			if isMatchDayOver(match, matchDay) {
				updateSeason(matchDay, season)
				announceMatchDay(season, matchDayCount, os.Stdout)
				matchDay = make(map[string]int)
				matchDayCount++
			}

			updateMatchDay(match, matchDay)
		}
	}

	// Announce final results
	updateSeason(matchDay, season)
	announceMatchDay(season, matchDayCount, os.Stdout)
}

// Parse a match string into team names and their scores based on supplied regex
func getMatchResult(matchString string, regex *regexp.Regexp) (matchResult, error) {
	match := matchResult{}
	teams := strings.Split(matchString, ",")

	if len(teams) != 2 {
		return match, errors.New("invalid match string")
	}

	result1 := regex.FindStringSubmatch(teams[0])
	if len(result1) != 3 {
		return match, errors.New("invalid match string")
	}
	match.team1 = strings.TrimSpace(result1[1])
	match.team1score, _ = strconv.Atoi(strings.TrimSpace(result1[2]))

	result2 := regex.FindStringSubmatch(teams[1])
	if len(result2) != 3 {
		return match, errors.New("invalid match string")
	}
	match.team2 = strings.TrimSpace(result2[1])
	match.team2score, _ = strconv.Atoi(strings.TrimSpace(result2[2]))

	return match, nil
}

// Updates the matchDay map with the points. 3 for win, 0 for loss, 1 for tie
func updateMatchDay(match matchResult, matchDay map[string]int) {
	if match.team1score == match.team2score {
		matchDay[match.team1] = 1
		matchDay[match.team2] = 1
	} else if match.team1score > match.team2score {
		matchDay[match.team1] = 3
		matchDay[match.team2] = 0
	} else {
		matchDay[match.team1] = 0
		matchDay[match.team2] = 3
	}
}

// Update the season's scores with the results of the days match
func updateSeason(day map[string]int, season map[string]int) {
	for team := range day {
		season[team] = season[team] + day[team]
	}
}

// The match day is over if a team in the current match has already played this day
func isMatchDayOver(match matchResult, matchDay map[string]int) bool {
	if _, ok := matchDay[match.team1]; ok {
		return true
	}
	if _, ok := matchDay[match.team2]; ok {
		return true
	}
	return false
}

func announceMatchDay(season map[string]int, matchDayCount int, w io.Writer) {
	sortedSeason := make(MatchDay, len(season))
	i := 0
	for t, s := range season {
		sortedSeason[i] = TeamScoring{t, s}
		i++
	}
	sort.Sort(sortedSeason)

	fmt.Fprintf(w, "Matchday %d\n", matchDayCount)

	for i := 0; i < 3; i++ {
		if len(sortedSeason) > i {
			fmt.Fprintf(w, sortedSeason[i].TeamName+", %d pts\n", sortedSeason[i].Score)
		}
	}
	fmt.Fprintln(w)
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
