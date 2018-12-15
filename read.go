package go_1w

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func Read(file_path string) float32 {
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data := make([]byte, 100)
	_, err = file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	fields := strings.Split(lines[1], " ")
	temperature_field := fields[len(fields)-1]
	s_temp := strings.Split(temperature_field, "=")[1]

	i_temp, err := strconv.Atoi(strings.Trim(s_temp, "\n"))
	if err != nil {
		log.Fatal(err)
	}

	return float32(i_temp) / 1000
}
