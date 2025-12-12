package logs

import (
	"fmt"
	"os"
)

func ShowLogs() {
	file, err := os.ReadFile("monitor.log")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(file))

}
