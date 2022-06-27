package dcraw

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
)

var binPath string = "dcraw-json"

// SetDCRawBinPath sets the path to the dcraw-json binary.
func SetDCRawBinPath(newBinPath string) {
	binPath = newBinPath
}

// GetImageData calls the dcraw-json binary and returns it's output as RawData.
func GetImageData(ctx context.Context, filePath string) (*RawData, error) {
	cmd := exec.CommandContext(ctx,
		binPath,
		"-v",
		"-i",
		"-J",
		filePath,
	)

	var stdOut, stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", stdErr.String(), err)
	}

	if stdErr.Len() > 0 {
		return nil, fmt.Errorf("%v", stdErr)
	}

	data := new(RawData)
	err = json.Unmarshal(stdOut.Bytes(), data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
