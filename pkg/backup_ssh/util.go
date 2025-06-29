package backup_ssh

import (
	"strconv"
	"strings"
)

func getFreeDiskSpaceInKB(output string) (int64, error) {
	availableKB := strings.TrimSpace(string(output))
	available, err := strconv.ParseInt(availableKB, 10, 64)
	if err != nil {
		return 0, err
	}

	return available, nil
}
