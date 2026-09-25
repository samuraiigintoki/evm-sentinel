package scanner

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

type Rule struct {
	ID          string
	Name        string
	Severity    Severity
	Regex       *regexp.Regexp
	Remediation string
}

type Engine struct {
	rules   []Rule
	workers int
}

func NewEngine(workers int) *Engine {
	return &Engine{
		rules:   DefaultRules(),
		workers: workers,
	}
}

func (e *Engine) ScanDirectory(rootDir string) (*ScanResult, error) {
	startTime := time.Now()

	var files []string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".sol") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	filesChan := make(chan string, len(files))
	resultsChan := make(chan []Issue, len(files))
	var wg sync.WaitGroup

	for i := 0; i < e.workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range filesChan {
				resultsChan <- e.scanFile(file)
			}
		}()
	}

	for _, file := range files {
		filesChan <- file
	}
	close(filesChan)

	wg.Wait()
	close(resultsChan)

	var allIssues []Issue
	for issues := range resultsChan {
		allIssues = append(allIssues, issues...)
	}

	return &ScanResult{
		TotalFilesScanned: len(files),
		DurationMs:        time.Since(startTime).Milliseconds(),
		Issues:            allIssues,
	}, nil
}

func (e *Engine) scanFile(filePath string) []Issue {
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()

	var issues []Issue
	scanner := bufio.NewScanner(file)
	lineNum := 1

	for scanner.Scan() {
		line := scanner.Text()
		for _, rule := range e.rules {
			if rule.Regex.MatchString(line) {
				issues = append(issues, Issue{
					RuleID:      rule.ID,
					RuleName:    rule.Name,
					Severity:    rule.Severity,
					File:        filePath,
					Line:        lineNum,
					Snippet:     strings.TrimSpace(line),
					Remediation: rule.Remediation,
				})
			}
		}
		lineNum++
	}

	return issues
}