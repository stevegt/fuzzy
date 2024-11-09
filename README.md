# Fuzzy

Fuzzy is a Go library for performing fuzzy matching on strings. The
primary function provided by this library is `Match()`, which compares
a target string against a list of candidate strings and returns
matches sorted by their similarity to the target string.

## Features

- Compute similarity scores for string matches.
- Return matches sorted by similarity score.
- Provide detailed information for each match, including:
  - The original string from the list.
  - The similarity score.
  - The number of insertions, deletions, and substitutions needed.
  - The position of the match in the provided list.

## Usage

```go
package main

import (
    "fmt"
    "github.com/stevegt/fuzzy"
)

func main() {
    target := "example"
    candidates := []string{"samples", "examples", "simple", "examine"}
    
    matches := fuzzy.Match(target, candidates)
    
    for _, match := range matches {
        fmt.Printf("String: %s, Score: %f, Insertions: %d, Deletions: %d, Substitutions: %d, Position: %d\n",
            match.Original, match.Score, match.Insertions, match.Deletions, match.Substitutions, match.Position)
    }
}
```

