package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "file_rename_test")
	if err != nil {
		t.Fatalf("Error creating temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	createTestFiles(tempDir, t)

	csvContent :=
		`test_001.txt, test_001_renamed.txt
test_002.txt, test_002_renamed.txt
test_003.txt, test_003_renamed.txt
test_004.txt, test_004_renamed.txt
test_005.txt, test_005_renamed.txt`

	expectedNames := []string{
		"test_001_renamed.txt",
		"test_002_renamed.txt",
		"test_003_renamed.txt",
		"test_004_renamed.txt",
		"test_005_renamed.txt",
	}

	csvFile := filepath.Join(tempDir, "test.csv")
	err = os.WriteFile(csvFile, []byte(csvContent), 0644)
	if err != nil {
		t.Fatalf("Error creating test CSV file: %v", err)
	}

	os.Args = []string{"file-renamer", "-file", csvFile, "-folder", tempDir}
	main()

	checkRenamedFile(tempDir, expectedNames, t)

}

func createTestFiles(dir string, t *testing.T) {
	for i := 1; i <= 5; i++ {
		fileName := filepath.Join(dir, fmt.Sprintf("test_00%d.txt", i))
		err := os.WriteFile(fileName, []byte("test"), 0644)
		if err != nil {
			t.Fatalf("Error creating test file: %v", err)
		}
	}
}

func checkRenamedFile(dir string, expectedNames []string, t *testing.T) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatalf("Error reading directory: %v", err)
	}

	var actualNames []string
	for _, file := range entries {
		actualNames = append(actualNames, strings.TrimSpace(file.Name()))
	}

	for _, expectedName := range expectedNames {
		found := false
		for _, actualName := range actualNames {
			if expectedName == actualName {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected file %s not found", expectedName)
		}
	}
}
