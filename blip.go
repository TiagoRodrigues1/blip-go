package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Repo struct {
	Name          string
	Commits       int
	Files         int
	Lines         int
	ActivityScore float64
}

func main() {

	// Grab all commits from csv file
	commits := readCsvFile("./commits.csv")

	// Parse the data
	parsedData := parseDataToMap(commits)

	// Weights for each metric
	weightCommits, weightFiles, weightLines := 1.0, 0.5, 0.01

	for _, repo := range parsedData {
		repo.ActivityScore = weightCommits*float64(repo.Commits) + weightFiles*float64(repo.Files) + weightLines*float64(weightLines)
	}

	var sorted []*Repo
	for _, repo := range parsedData {
		sorted = append(sorted, repo)
	}

	// Sort per activity score
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].ActivityScore > sorted[j].ActivityScore
	})

	for _, repo := range sorted {
		fmt.Printf("%s: %.2f\n", repo.Name, repo.ActivityScore)
	}

}

func readCsvFile(filepath string) [][]string {
	f, err := os.Open(filepath)

	if err != nil {
		log.Fatal("Can't open file")
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	commits, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("Can't parse csv: "+filepath, err)
	}

	return commits
}

func parseDataToMap(commits [][]string) map[string]*Repo {
	data := make(map[string]*Repo)
	for _, commit := range commits[1:] {
		// Skip when no username
		if commit[1] == "" {
			continue
		}

		repoName := commit[2]
		files, _ := strconv.Atoi(commit[3])
		additions, _ := strconv.Atoi(commit[4])
		deletions, _ := strconv.Atoi(commit[5])

		linesChanged := additions + deletions

		// Create map entry if it dones't exist
		if _, exists := data[repoName]; !exists {
			data[repoName] = &Repo{Name: repoName}
		}

		repo := data[repoName]
		repo.Commits++
		repo.Files += files
		repo.Lines += linesChanged
	}

	return data
}
