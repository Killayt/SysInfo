package memory

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func GetGPU() string {
	for {
		// Run the 'nvidia-smi' command to get GPU usage
		out, err := exec.Command("nvidia-smi", "--query-gpu=utilization.gpu", "--format=csv,noheader").Output()
		if err != nil {
			panic(err)
		}
		// Convert output to string
		usageStr := strings.TrimSpace(string(out))
		// Parse the usage percentage from the string (e.g. " 75 %")
		re := regexp.MustCompile(`\s*(\d+)\s*%`)
		matches := re.FindStringSubmatch(usageStr)
		if len(matches) != 2 {
			panic(fmt.Sprintf("Invalid usage string: %s", usageStr))
		}
		// Wait for 1 second
		time.Sleep(time.Second)

		return matches[1]
	}
}
