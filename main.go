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
		if val == query {
			return true
		}
	}
	return false
}

func gatherFiles() []string {
	// todo: make this configurable
	toSkip := []string{"node_modules/", ".git"}
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
				// todo: trim line and remove leading "todo:"
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
		fmt.Printf("%s:\n", result.filename)
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
