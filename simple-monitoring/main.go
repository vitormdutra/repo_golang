package main

import (
	"fmt"
	"os"

	"simple-monitoring/logs"
	"simple-monitoring/monitoring"
)

var command = 0

func main() {
	for {
		initSoftware()
		result := readCommand()

		selectCommand(result)
	}
}

func initSoftware() {
	fmt.Println("Type what you want")

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")
}

func readCommand() int {
	fmt.Scanf("%d", &command)
	return command
}

func selectCommand(command int) {
	switch command {
	case 1:
		monitoring.InitMonitoring()
	case 2:
		logs.ShowLogs()
	case 3:
		os.Exit(0)
	default:
		os.Exit(-1)
	}
}
