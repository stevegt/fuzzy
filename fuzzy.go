package fuzzy

import (
	"sort"
	"unicode/utf8"
)

// MatchResult represents a single match, including similarity details.
type MatchResult struct {
	Original      string
	Score         float64
	Insertions    int
	Deletions     int
	Substitutions int
	Position      int
}

// Match takes a target string and a list of candidate strings and returns all matches
// sorted by their similarity score.
func Match(target string, candidates []string) []MatchResult {
	var matches []MatchResult

	for i, candidate := range candidates {
		insertions, deletions, substitutions := levenshteinDistance(target, candidate)
		score := calculateSimilarityScore(insertions, deletions, substitutions, len(target), len(candidate))
		matches = append(matches, MatchResult{
			Original:      candidate,
			Score:         score,
			Insertions:    insertions,
			Deletions:     deletions,
			Substitutions: substitutions,
			Position:      i,
		})
	}

	// Sort by descending score
	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Score > matches[j].Score
	})

	return matches
}

// levenshteinDistance calculates the number of insertions, deletions, and substitutions
// required to change source into target using the Levenshtein distance algorithm.
func levenshteinDistance(source, target string) (int, int, int) {
	sLen := utf8.RuneCountInString(source)
	tLen := utf8.RuneCountInString(target)

	d := make([][]int, sLen+1)
	for i := range d {
		d[i] = make([]int, tLen+1)
	}

	for i := 0; i <= sLen; i++ {
		d[i][0] = i
	}
	for j := 0; j <= tLen; j++ {
		d[0][j] = j
	}

	for i := 1; i <= sLen; i++ {
		for j := 1; j <= tLen; j++ {
			if rune(source[i-1]) == rune(target[j-1]) {
				d[i][j] = d[i-1][j-1] // no operation needed
			} else {
				deletion := d[i-1][j] + 1
				insertion := d[i][j-1] + 1
				substitution := d[i-1][j-1] + 1
				d[i][j] = minInt(deletion, insertion, substitution)
			}
		}
	}

	i, j := sLen, tLen
	insertions, deletions, substitutions := 0, 0, 0

	for i > 0 || j > 0 {
		if i > 0 && j > 0 && d[i][j] == d[i-1][j-1] && rune(source[i-1]) == rune(target[j-1]) {
			i--
			j--
		} else if i > 0 && d[i][j] == d[i-1][j]+1 {
			deletions++
			i--
		} else if j > 0 && d[i][j] == d[i][j-1]+1 {
			insertions++
			j--
		} else if i > 0 && j > 0 && d[i][j] == d[i-1][j-1]+1 {
			substitutions++
			i--
			j--
		}
	}

	return insertions, deletions, substitutions
}

// calculateSimilarityScore computes a similarity score based on the edit distances.
func calculateSimilarityScore(insertions, deletions, substitutions, sLen, tLen int) float64 {
	totalOperations := float64(insertions + deletions + substitutions)
	maxLen := float64(maxInt(sLen, tLen))
	if maxLen == 0 {
		return 1 // if both strings are empty, they are considered the same
	}
	return 1 - (totalOperations / maxLen)
}

// minInt calculates the minimum of three integers.
func minInt(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// maxInt calculates the maximum of two integers.
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
