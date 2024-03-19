package main

import (
	"fmt"
	"os/exec"
	"time"
)

func main() {
	inputFilePath := "./public/19061743631.mp4"

	// Output frames directory
	outputDirectory := "out/"

	startTime := time.Now()
	cmd := exec.Command("ffmpeg", "-i", inputFilePath, outputDirectory+"%04d.png")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	elapsedTime := time.Since(startTime)
	fmt.Println("Execution time:", elapsedTime)

	println("Frames extracted successfully.")
}
