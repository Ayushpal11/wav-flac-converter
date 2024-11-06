package ffmpeg

import (
	"bytes"
	"os/exec"
)

func ConvertToFLAC(wavData []byte) ([]byte, error) {
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "flac", "pipe:1")
	cmd.Stdin = bytes.NewReader(wavData)
	flacData, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	return flacData, nil
}
