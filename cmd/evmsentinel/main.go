package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"evm-sentinel/pkg/formatter"
	"evm-sentinel/pkg/scanner"
)

func main() {
	dirPtr := flag.String("dir", ".", "Directory path containing Solidity contracts")
	jsonPtr := flag.Bool("json", false, "Output results as JSON for CI integration")
	workersPtr := flag.Int("workers", 4, "Number of concurrent scanning workers")
	flag.Parse()

	engine := scanner.NewEngine(*workersPtr)
	result, err := engine.ScanDirectory(*dirPtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running scan: %v\n", err)
		os.Exit(1)
	}

	if *jsonPtr {
		data, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(data))
		return
	}

	formatter.RenderConsoleReport(result)
}