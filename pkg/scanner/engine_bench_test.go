package scanner

import (
	"os"
	"testing"
)

func BenchmarkScannerWorkerPool(b *testing.B) {
	// Create a dummy solidity contract for performance profiling
	tmpFile, err := os.CreateTemp("", "bench_*.sol")
	if err != nil {
		b.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	dummySol := "pragma solidity ^0.8.20;\ncontract Bench { function test() public { require(tx.origin == msg.sender); } }\n"
	if _, err := tmpFile.WriteString(dummySol); err != nil {
		b.Fatal(err)
	}
	tmpFile.Close()

	engine := NewEngine(4)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = engine.ScanDirectory(os.TempDir())
	}
}
