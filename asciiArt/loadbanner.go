package asciiArt

import (
	"bufio"
	"fmt"
	"os"
)

// LoadBannerMap loads the banner map from the file provided
func LoadBannerMap(fileName string) (map[int][]string, error) {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("Error getting file info: %q\n%v\n", fileName, err)
		os.Exit(1)
	}
	//  Checks if the file has been tampered with
	fileSize := fileInfo.Size()
	expectedSize := map[string]int64{
		"standard.txt":  6623,
		"shadow.txt":    7463,
		"thinkertoy.txt": 5558,
	}

	if size, ok := expectedSize[fileName]; ok {
		if fileSize != size {
			return nil, fmt.Errorf("%s filesize tampered: expected %d, got %d", fileName, size, fileSize)
		}
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bannerMap := make(map[int][]string)
	key := 32
	lineCount := 0
	chunk := []string{}

	for scanner.Scan() {
		lines := scanner.Text()

		if lines != "" {
			chunk = append(chunk, lines)
			lineCount++
		}

		if lineCount == 8 {
			bannerMap[key] = chunk
			key++
			chunk = []string{}
			lineCount = 0
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return bannerMap, nil
}
