package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Print("Enter the directory path: ")
	reader := bufio.NewReader(os.Stdin)
	path, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Remove the newline character at the end of the input
	path = strings.TrimSpace(path)

	fmt.Println("path: ", path)
	path = filepath.FromSlash(path)

	err = os.Chdir(path)
	if err != nil {
		fmt.Println("Error changing directory: ", err)
		return
	}

	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isVideoFile(path) {
			extractAudio(path)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error walking through the path: ", err)
	}
}

func isVideoFile(path string) bool {
	videoExtensions := []string{".mp4", ".mkv", ".avi", ".flv", ".wmv"}
	for _, ext := range videoExtensions {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}

	return false
}

func extractAudio(videopath string) {
	audioPath := strings.TrimSuffix(videopath, filepath.Ext(videopath)) + ".mp3"
	cmd := exec.Command("ffmpeg", "-i", videopath, "-vn", "-acodec", "libmp3lame", "-y", audioPath)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error extracting audio: ", err)
	}
}
