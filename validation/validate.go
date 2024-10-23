package validation

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
)

// ValidateMigrationFiles validates the files in the given directory.
// Files must follow the pattern <number>_<name>.sql and the numbers must be
// monotonically increasing without gaps.
func ValidateMigrationFiles(dir string) (bool, error) {
	// Regular expression to match files in the format <number>_<name>.sql
	filePattern := regexp.MustCompile(`^(\d+)_.*\.sql$`)

	// Slice to store all the file numbers
	var fileNumbers []int

	// Walk through the directory and collect the file numbers
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filename := filepath.Base(path)
			matches := filePattern.FindStringSubmatch(filename)
			if matches == nil {
				return fmt.Errorf("invalid file format: %s", filename)
			}
			// Convert the number part of the filename to an integer
			number, err := strconv.Atoi(matches[1])
			if err != nil {
				return fmt.Errorf("invalid number in file name: %s", filename)
			}
			fileNumbers = append(fileNumbers, number)
		}
		return nil
	})

	if err != nil {
		return false, err
	}

	// If no valid files are found
	if len(fileNumbers) == 0 {
		return false, fmt.Errorf("no valid SQL migration files found in directory")
	}

	// Sort the file numbers
	sort.Ints(fileNumbers)

	// Check if the numbers start at 0 and are sequential
	for i, num := range fileNumbers {
		if num != i {
			return false, fmt.Errorf("file number sequence is broken. Expected %d but got %d", i, num)
		}
	}

	// All checks passed, return true
	return true, nil
}
