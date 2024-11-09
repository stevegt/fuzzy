package fuzzy

import (
	"testing"
)

func TestInsertionsDeletionsSubstitutions(t *testing.T) {
	target := "example"
	candidates := []string{"samples", "examples", "simple", "examine"}
	expected := []MatchResult{
		{Original: "examples", Score: 0.875, Insertions: 1, Deletions: 0, Substitutions: 0, Position: 1},
		{Original: "examine", Score: 0.7142857142857143, Insertions: 0, Deletions: 0, Substitutions: 2, Position: 3},
		{Original: "samples", Score: 0.5714285714285714, Insertions: 1, Deletions: 1, Substitutions: 1, Position: 0},
		{Original: "simple", Score: 0.5714285714285714, Insertions: 0, Deletions: 1, Substitutions: 2, Position: 2},
	}

	matches := Match(target, candidates)
	// spew.Dump(matches)

	for i, match := range matches {
		if match.Original != expected[i].Original {
			t.Errorf("expected match %d to be %s, got %s", i, expected[i].Original, match.Original)
		}
		if match.Score != expected[i].Score {
			t.Errorf("expected match %d to have score %f, got %f", i, expected[i].Score, match.Score)
		}
		if match.Insertions != expected[i].Insertions {
			t.Errorf("expected match %d to have %d insertions, got %d", i, expected[i].Insertions, match.Insertions)
		}
		if match.Deletions != expected[i].Deletions {
			t.Errorf("expected match %d to have %d deletions, got %d", i, expected[i].Deletions, match.Deletions)
		}
		if match.Substitutions != expected[i].Substitutions {
			t.Errorf("expected match %d to have %d substitutions, got %d", i, expected[i].Substitutions, match.Substitutions)
		}
		if match.Position != expected[i].Position {
			t.Errorf("expected match %d to have position %d, got %d", i, expected[i].Position, match.Position)
		}
	}
}

func TestEmptyCandidates(t *testing.T) {
	target := "example"
	candidates := []string{}
	expected := 0

	matches := Match(target, candidates)
	if len(matches) != expected {
		t.Fatalf("expected %d matches, got %d", expected, len(matches))
	}
}

func TestExactMatch(t *testing.T) {
	target := "example"
	candidates := []string{"example"}
	expected := "example"

	matches := Match(target, candidates)
	if len(matches) != 1 {
		t.Fatalf("expected 1 match, got %d", len(matches))
	}

	if matches[0].Original != expected {
		t.Errorf("expected match to be %s, got %s", expected, matches[0].Original)
	}
}
