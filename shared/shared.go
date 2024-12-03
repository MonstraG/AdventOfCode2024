package shared

import "os"

func ReadFile() string {
	bytes, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
