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
	// // Check if the directory was provided
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage: go run goAudioExtractor.go <directory_path> ")
	// 	os.Exit(1)
	// }

	// // The directory is the first argument after the program name
	// dirPath := os.Args[1]

	dirPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("Failed to read directory '%s': %v", dirPath, err)
	}

	fmt.Printf("Files in directory '%s': \n", dirPath)
	for _, file := range files {
		filename := file.Name()
		fileExt := filepath.Ext(filename)
		if fileExt == ".mp4" {
			mp3Dir := createMP3directory()
			fmt.Println(filename)
			mp3Path, err := extractAudio(filename, mp3Dir)
			if err != nil {
				log.Printf("Failed to extract audio from '%s': %v", filename, err)
				continue
			}
			fmt.Printf("Extracted audio to '%s'\n", mp3Path)
		}
	}
}

func extractAudio(mp4FilePath, mp3Dir string) (string, error) {
	// check input
	if strings.ToLower(filepath.Ext(mp4FilePath)) != ".mp4" {
		return "", fmt.Errorf("input file '%s' is not an MP4 file", mp4FilePath)
	}

	// 2 - Construct the output file
	baseName := strings.TrimSuffix(mp4FilePath, filepath.Ext(mp4FilePath))
	mp3FilePath := mp3Dir + "/" + baseName + "Sound.mp3"

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

func createMP3directory() string {
	folderName := "mp3"
	// check if the directory already exists
	if _, err := os.Stat(folderName); os.IsNotExist(err) {
		// directory does not exist, create it
		err := os.Mkdir(folderName, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory '%s': '%v'", folderName, err)
		}
	} else {
		fmt.Printf("Directory '%s' already exists. Skipping creating it\n", folderName)
	}
	return folderName
}
