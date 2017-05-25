package dcraw

import (
	"encoding/json"
	"errors"
	"os/exec"
)

var ErrBinNotFound error = errors.New("dcraw-json bin not found")

var binPath string = "dcraw-json"

func SetDCRawBinPath(newBinPath string) {
	binPath = newBinPath
}

func GetImageData(filePath string) (data *RawData, err error) {
	cmd := exec.Command(
		binPath,
		"-v",
		"-i",
		"-J",
		filePath,
	)

	outputBuf, err := cmd.Output()
	if err != nil {
		if err == exec.ErrNotFound {
			err = ErrBinNotFound
		}
		return
	}

	// log.Debug(string(outputBuf))

	data = &RawData{}
	err = json.Unmarshal(outputBuf, data)
	if err != nil {
		panic(err)
		return
	}
	return
}
