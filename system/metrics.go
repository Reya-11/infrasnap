package system

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type Metrics struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
	Disk   string `json:"disk"`
	Uptime string `json:"uptime"`
}

func GetMetrics() (Metrics, error) {
	cpu := getCPUUsage()
	mem := getMemoryUsage()
	disk := getDiskUsage()
	uptime := getUptime()

	metrics := Metrics{
		CPU:    cpu,
		Memory: mem,
		Disk:   disk,
		Uptime: uptime,
	}
	return metrics, nil
}

// -------------------------
// Each helper function below uses shell commands

func getCPUUsage() string {
	out, _ := exec.Command("sh", "-c", `top -bn1 | grep "Cpu(s)"`).Output()
	fields := strings.Fields(string(out))
	if len(fields) < 8 {
		return "N/A"
	}
	idleStr := fields[7] // idle CPU percent
	idle, _ := strconv.ParseFloat(strings.TrimSuffix(idleStr, "%id,"), 64)
	usage := 100.0 - idle
	return fmt.Sprintf("%.2f%%", usage)
}

func getMemoryUsage() string {
	out, _ := exec.Command("sh", "-c", `free | grep Mem`).Output()
	fields := strings.Fields(string(out))
	if len(fields) < 3 {
		return "N/A"
	}
	total, _ := strconv.ParseFloat(fields[1], 64)
	used, _ := strconv.ParseFloat(fields[2], 64)
	percent := (used / total) * 100
	return fmt.Sprintf("%.2f%%", percent)
}

func getDiskUsage() string {
	out, _ := exec.Command("sh", "-c", `df / | tail -1`).Output()
	fields := strings.Fields(string(out))
	if len(fields) < 5 {
		return "N/A"
	}
	return fields[4] // already in %
}

func getUptime() string {
	out, _ := exec.Command("sh", "-c", `uptime -p`).Output()
	return strings.TrimSpace(string(out))
}
