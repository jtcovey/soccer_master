package main

import (
	"bytes"
	"testing"
)

var team1name = "Bengal Tigers"
var team2name = "Chicago Bears"
var team3name = "Dallas Cowboys"
var team4name = "New York Yankees"
var team5name = "Cleveland Browns"
var team6name = "Staten Island Stevedores"

// matchDay is built over the course of the tests
var matchDay = make(map[string]int)

func TestGetMatchResults(t *testing.T) {
	match, err := getMatchResult(team1name+" 79, "+team2name+" 6", matchRegex)
	if err != nil {
		t.Fatalf("Failed to match with valid match string: %s", err)
	}
	if match.team1 != team1name || match.team1score != 79 || match.team2 != team2name || match.team2score != 6 {
		t.Fatalf("Failed to get correct match results")
	}
	_, err = getMatchResult("nil", matchRegex)
	if err == nil {
		t.Fatalf("Failed to return error on bad string")
	}
	_, err = getMatchResult(team6name+" 26, ", matchRegex)
	if err == nil {
		t.Fatalf("Failed to return error on incomplete string")
	}
}

func TestUpdateMatchDay(t *testing.T) {
	match := matchResult{team1: team1name, team1score: 5, team2: team2name, team2score: 79}

	updateMatchDay(match, matchDay)
	if len(matchDay) != 2 || matchDay[team1name] != 0 || matchDay[team2name] != 3 {
		t.Fatalf("Fail to update match 1")
	}

	match2 := matchResult{team1: team3name, team1score: 42, team2: team4name, team2score: 42}
	updateMatchDay(match2, matchDay)
	if len(matchDay) != 4 || matchDay[team1name] != 0 || matchDay[team2name] != 3 ||
		matchDay[team3name] != 1 || matchDay[team4name] != 1 {
		t.Fatalf("Fail to update match 2")
	}
}

func TestIsMatchDayOver(t *testing.T) {
	match3 := matchResult{team1: team5name, team1score: -1, team2: team6name, team2score: 113}

	if isMatchDayOver(match3, matchDay) {
		t.Fatalf("Failed to declare match day not over")
	}
	updateMatchDay(match3, matchDay)

	match4 := matchResult{team1: team6name, team1score: 21, team2: team1name, team2score: 6}

	if !isMatchDayOver(match4, matchDay) {
		t.Fatalf("Failed to call match day over")
	}
}

func TestAnnounceMatchDay(t *testing.T) {
	expectedOutput := "Matchday 1\nChicago Bears, 3 pts\nStaten Island Stevedores, 3 pts\nDallas Cowboys, 1 pts\n\n"
	var output bytes.Buffer
	announceMatchDay(matchDay, 1, &output)
	if output.String() != expectedOutput {
		t.Fatalf("Failed to make correct announcement. Got\n%sbut expected\n%s", output.String(), expectedOutput)
	}
}
