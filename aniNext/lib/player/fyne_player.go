// startVideoPlayback takes over the window to play the selected stream.
func startVideoPlayback(win fyne.Window, streamURL string) {
	videoCanvas := canvas.NewImageFromImage(nil)
	videoCanvas.FillMode = canvas.ImageFillContain
	videoCanvas.SetMinSize(fyne.NewSize(1, 1))

	win.SetContent(videoCanvas)

	// Use fixed directory instead of temp
	tempDir := `C:\Users\ayush\AppData\Local\Temp\test-frames`
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		log.Fatalf("Failed to create test-frames dir: %v", err)
	}
	log.Println("Decoding frames to:", tempDir)

	// Test writing a file to verify access
	testFile := filepath.Join(tempDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		log.Fatalf("Failed to write test file to test-frames dir: %v", err)
	}

	go func() {
		outputPath := filepath.Join(tempDir, "frame-%06d.png")
		log.Printf("FFmpeg output path: %s", outputPath)

		// Create buffers to capture FFmpeg stdout and stderr
		var stdout, stderr bytes.Buffer

		// Build the FFmpeg command
		stream := ffmpeg.Input(streamURL, ffmpeg.KwArgs{
			"referer":    hardcodedReferer,
			"user_agent": "Mozilla/5.0",
		}).
			Output(outputPath, ffmpeg.KwArgs{"fps_mode": "vfr"}).
			OverWriteOutput().
			WithOutput(&stdout, &stderr)

		// Log the full FFmpeg command
		log.Printf("FFmpeg command: %s", stream.String())

		// Run the FFmpeg command
		err := stream.Run()

		if err != nil {
			log.Printf("FFmpeg command failed: %v", err)
			log.Printf("FFmpeg stdout: %s", stdout.String())
			log.Printf("FFmpeg stderr: %s", stderr.String())
		} else {
			log.Println("FFmpeg decoding process finished successfully.")
			log.Printf("FFmpeg stdout: %s", stdout.String())
		}
	}()

	go func() {
		ticker := time.NewTicker(time.Second / 30) // ~30 FPS
		defer ticker.Stop()
		frameNumber := 1
		for range ticker.C {
			framePath := filepath.Join(tempDir, fmt.Sprintf("frame-%06d.png", frameNumber))
			if _, err := os.Stat(framePath); err != nil {
				if os.IsNotExist(err) {
					continue
				}
				log.Printf("Error accessing frame %s: %v", framePath, err)
			}

			videoCanvas.File = framePath
			videoCanvas.Refresh()
			frameNumber++
		}
	}()
}