package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func contains(toSearch []string, query string) bool {
	for _, val := range toSearch {
		if strings.Contains(query, val) {
			return true
		}
	}
	return false
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func getIgnoredFiles() []string {
	ignoredList := []string{}
	// check if gitignore exists and read that
	if fileExists(".gitignore") {
		raw, _ := os.ReadFile(".gitignore")
		contents := string(raw)
		for _, filename := range strings.Split(contents, "\n") {
			if len(strings.TrimSpace(filename)) == 0 || filename[0] == '#' {
				continue
			}
			ignoredList = append(ignoredList, strings.Trim(filename, "/\\"))
		}
	}

	// todo: check for ~/.config/todo.conf and load that

	return ignoredList
}

func gatherFiles() []string {
	toSkip := []string{".git"}
	toSkip = append(toSkip, getIgnoredFiles()...)
	cwd, _ := os.Getwd()
	toScan := []string{}
	_ = filepath.Walk(cwd, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && contains(toSkip, info.Name()) {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			toScan = append(toScan, path)
		}
		return nil
	})
	return toScan

}

type Todo struct {
	content string
	lineNum string
}

type FileResult struct {
	filename string
	todos    []Todo
}

func scan(files []string) []FileResult {
	pattern := `(?i)^\s*(//|\*|#)\s*todo:`
	re := regexp.MustCompile(pattern)
	results := []FileResult{}
	for _, filename := range files {
		shouldCapture := false
		result := FileResult{
			filename: filename,
			todos:    []Todo{},
		}
		dat, _ := os.ReadFile(filename)
		contents := string(dat)
		for idx, line := range strings.Split(contents, "\n") {
			lower := strings.ToLower(line)
			if re.MatchString(lower) {
				tidiedTodo := re.ReplaceAllString(line, "")
				tidiedTodo = strings.TrimSpace(tidiedTodo)
				result.todos = append(result.todos, Todo{
					content: tidiedTodo,
					lineNum: strconv.Itoa(idx + 1),
				})
				shouldCapture = true
			}
		}

		if shouldCapture {
			results = append(results, result)
		}
	}
	return results
}

func printResults(results []FileResult) {
	for _, result := range results {
		fmt.Fprintf(os.Stdout, "\033[1m%s\033[0m:\n", result.filename)
		for _, todo := range result.todos {
			fmt.Printf("  %s: %s\n", todo.lineNum, todo.content)
		}
	}
}

func main() {
	files := gatherFiles()
	results := scan(files)
	printResults(results)
}
