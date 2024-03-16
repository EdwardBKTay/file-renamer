package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type file struct {
	oldName string
	newName string
}

func main() {
	filename := flag.String("file", "", "CSV file containing old and new file names")
	folder := flag.String("folder", "", "Folder containing files to rename")

	flag.Parse()

	records, err := readCSVFile(*filename)
	if err != nil {
		fmt.Printf("Error in CSV file: %v\n", err)
		os.Exit(1)
	}

	for _, record := range records {
		err := filepath.WalkDir(*folder, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.Name() == record.oldName {

				err := os.Rename(path, filepath.Join(filepath.Dir(path), record.newName))
				if err != nil {
					fmt.Printf("Error renaming file: %v\n", err)
				}
			}
			return nil
		})

		if err != nil {
			fmt.Printf("Error walking directory: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Println("Files renamed successfully")
}

func readCSVFile(filename string) ([]file, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var files []file
	for _, record := range records {
		if len(record) < 2 {
			return nil, errors.New("not enough rows in csv file")
		}
		files = append(files, file{oldName: record[0], newName: record[1]})
	}

	for i := 0; i < len(files); i++ {
		for j := i + 1; j < len(files); j++ {
			if files[i].oldName == files[j].oldName {
				return nil, errors.New("duplicate oldName in csv file")
			}
		}
	}

	return files, nil
}
