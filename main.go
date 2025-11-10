package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// Check if the directory was provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run goAudioExtractor.go <directory_path> ")
		os.Exit(1)
	}

	// The directory is the first argument after the program name
	dirPath := os.Args[1]

	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Failed to read directory '%s': %v", dirPath, err)
	}

	fmt.Printf("Files in directory '%s': \n", dirPath)
	for _, file := range files {
		filename := file.Name()
		fileExt := filepath.Ext(filename)
		if fileExt == ".mp4" {
			fmt.Println(filename)
			mp3Path, err := extractAudio(filename)
			if err != nil {
				log.Printf("Failed to extract audio from '%s': %v", filename, err)
				continue
			}
			fmt.Printf("Extracted audio to '%s'\n", mp3Path)
		}
	}
}

func extractAudio(mp4FilePath string) (string, error) {
	// check input
	if strings.ToLower(filepath.Ext(mp4FilePath)) != ".mp4" {
		return "", fmt.Errorf("input file '%s' is not an MP4 file", mp4FilePath)
	}

	// 2 - Construct the output file
	baseName := strings.TrimSuffix(mp4FilePath, filepath.Ext(mp4FilePath))
	mp3FilePath := baseName + "Sound.mp3"

	// 3 - Prepare ffmpeg command
	cmd := exec.Command("ffmpeg", "-i", mp4FilePath, "-q:a", "0", "-map", "a", mp3FilePath)

	//4 Run the command and capture output
	fmt.Printf("Running command: %s\n", cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute FFmpeg command: %v\nFFmpeg output: %s", err, string(output))
	}

	return mp3FilePath, nil
}
