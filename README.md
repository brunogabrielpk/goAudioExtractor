# Go Audio Extractor

This is a simple Go application to extract audio from all MP4 files in a directory and save them as `.mp3` files.

## Prerequisites

- [Go](https://golang.org/) installed on your system.
- [FFmpeg](https://ffmpeg.org/) installed and available in your system's PATH.

## Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/goAudioExtractor.git
   ```
2. Navigate to the project directory:
   ```bash
   cd goAudioExtractor
   ```
3. Run the application with the path to the directory containing your MP4 files:
   ```bash
   go run main.go /path/to/your/videos
   ```

The application will process all `.mp4` files in the specified directory, and for each file, it will create a corresponding `.mp3` file with the audio content.

## Build

To build the executable, run the following command:
```bash
go build .
```

This will create an executable file named `goAudioExtractor` (or `goAudioExtractor.exe` on Windows) in the project directory.

You can then run the executable directly:
```bash
./goAudioExtractor /path/to/your/videos
```

## Compiling for Windows

To compile the application for Windows, you can use the following command:
```bash
GOOS=windows GOARCH=amd64 go build -o goAudioExtractor.exe .
```
