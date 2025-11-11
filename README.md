# Go Audio Extractor

This is a simple Go application to extract audio from all MP4 files in the current directory and save them as `.mp3` files.

## Prerequisites

- [Go](https://golang.org/) installed on your system.
- [FFmpeg](https://ffmpeg.org/) installed and available in your system's PATH.

## Usage

1. Place the executable (or the `main.go` script) in the directory containing your MP4 files.
2. Run the application:
   ```bash
   go run main.go
   ```
3. The script will automatically create a directory named `mp3` and save the extracted audio files inside it.

## Build

To build the executable, run the following command:
```bash
go build .
```

This will create an executable file named `goAudioExtractor` (or `goAudioExtractor.exe` on Windows) in the project directory.

You can then run the executable directly in the folder with your videos:
```bash
./goAudioExtractor
```

## Compiling for Windows

To compile the application for Windows, you can use the following command:
```bash
GOOS=windows GOARCH=amd64 go build -o goAudioExtractor.exe .
```
