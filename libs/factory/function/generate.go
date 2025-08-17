package function

import "fmt"

func GenerateCode(prefix string, counter int64) (string, error) {
	if counter < 0 {
		return "", fmt.Errorf("counter must be a non-negative integer, got %d", counter)
	}

	return fmt.Sprintf("%s%04d", prefix, counter), nil
}
