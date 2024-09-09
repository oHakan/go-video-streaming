package helpers

import (
	"bytes"
	"fmt"
	"os/exec"
)

func RunCommand(command *exec.Cmd) (error, bytes.Buffer) {
	err := command.Run()

	var out bytes.Buffer
	var stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr

	if err != nil {
		return err, stderr
	}

	return nil, out
}

func GenerateFFMPEGCommand(fileDestination string, fileName string) *exec.Cmd {
	return exec.Command("ffmpeg", "-i", fmt.Sprintf("%s/%s", fileDestination, fileName), "-c:v", "libx264", "-preset", "veryfast", "-g", "48", "-keyint_min", "48", "-sc_threshold", "0", "-b:v", "2500k", "-maxrate", "2500k", "-bufsize", "5000k", "-c:a", "aac", "-b:a", "128k", "-hls_time", "10", "-hls_playlist_type", "vod", "-hls_segment_filename", fmt.Sprintf("%s/output%%03d.ts", fileDestination), fmt.Sprintf("%s/playlist.m3u8", fileDestination))
}
