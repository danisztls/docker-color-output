package stdin

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func GetLines() ([]string, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		return nil, errors.New("no stdin")
	}

	var lines []string
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		if !strings.HasPrefix(line, "-") {
			lines = append(lines, s.Text())
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}