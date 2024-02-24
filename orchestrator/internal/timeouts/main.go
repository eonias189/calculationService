package timeouts

import (
	c "backend/contract"
	"encoding/json"
	"io"
	"os"
)

const (
	Path = "timeouts.json"
)

func GetTimeouts() (*c.Timeouts, error) {
	timeouts := c.Timeouts{Add: 0, Substract: 0, Divide: 0, Multiply: 0}
	file, err := os.Open(Path)
	if err != nil {
		return &timeouts, err
	}

	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return &timeouts, err
	}
	err = json.Unmarshal(data, &timeouts)
	return &timeouts, err
}

func SetTimeouts(timeouts *c.Timeouts) error {
	file, err := os.OpenFile(Path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	data, err := json.Marshal(timeouts)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	return err

}
