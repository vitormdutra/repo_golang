package monitoring

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const times = 5

func InitMonitoring() {
	fmt.Println("Monitoring")

	webs := readFile()
	for i := 0; i < times; i++ {
		for _, web := range webs {
			fmt.Println("Monitoring:", web)
			TestSite(web)
		}
		time.Sleep(times * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func TestSite(web string) {
	resp, _ := http.Get(web)
	if resp.StatusCode == 200 {
		fmt.Println("working")
		registerLog(web, true)
	} else {
		fmt.Println("not working")
		registerLog(web, false)
	}
}

func readFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		sites = append(sites, line)

		if err == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func registerLog(site string, status bool) {
	file, err := os.OpenFile("monitor.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		fmt.Println(err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}
