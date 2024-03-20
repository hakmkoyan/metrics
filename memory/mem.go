package memory

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/goccy/go-json"
)

type MemoryInfo struct {
	MemTotal     int    `json: "memTotal"`
	MemFree      int    `json: "memFree"`
	MemAvailable int    `json: "memAvailable"`
	Buffers      int    `json: "buffers"`
	Cached       int    `json: "cached"`
	Active       int    `json: "active"`
	Inactive     int    `json: "inactive"`
	SwapCached   int    `json: "swapCached"`
	SwapTotal    int    `json: "swapTotal"`
	SwapFree     int    `json: "swapFree"`
}

func OpenMemInfoFile(open bool) *os.File {
	var procMemInfo *os.File

	if open {
		procMemInfo, err := os.Open("/proc/meminfo")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(2)
		}
		return procMemInfo
	}
	procMemInfo.Close()
	return nil
}

func GetMemoryInfo(memoryInfo *MemoryInfo) map[string]int {
	memoryInfoContent := OpenMemInfoFile(true)
	r := bufio.NewReader(memoryInfoContent)

	for {
		infoLine, _, err := r.ReadLine()
		if err != nil {
			break
		}

		info := strings.Split(string(infoLine), ":")
		key := strings.TrimSpace(info[0])
		value := strings.TrimSpace(string(info[1]))
		value = strings.TrimRight(value, " kB")

		switch key {
		case "MemTotal":
		  statNum, err := strconv.Atoi(value)
	      if err != nil {
			os.Exit(1)
		  }
		memoryInfo.MemTotal = statNum
		case "MemFree":
		  statNum, err := strconv.Atoi(value)
		  if err != nil {
		  	os.Exit(1)
		  }
		  memoryInfo.MemFree = statNum
		case "MemAvailable":
		  statNum, err := strconv.Atoi(value)
		  if err != nil {
		  	os.Exit(1)
		  }
		  memoryInfo.MemAvailable = statNum
		case "Buffers":
		  statNum, err := strconv.Atoi(value)
		  if err != nil {
		  	os.Exit(1)
		  }
		  memoryInfo.Buffers = statNum
		case "Cached":
		  statNum, err := strconv.Atoi(value)
		  if err != nil {
		  	os.Exit(1)
		  }
		  memoryInfo.Cached = statNum
		case "Active":
		  statNum, err := strconv.Atoi(value)
		  if err != nil {
		  	os.Exit(1)
		  }
		  memoryInfo.Active = statNum
		case "Inactive":
		  statNum, err := strconv.Atoi(value)
		  if err != nil {
		  	os.Exit(1)
		  }
		  memoryInfo.Inactive = statNum
		case "SwapCached":
		  statNum, err := strconv.Atoi(value)
		  if err != nil {
		  	os.Exit(1)
		  }
		  memoryInfo.SwapCached = statNum
		case "SwapTotal":
		  statNum, err := strconv.Atoi(value)
		  if err != nil {
		  	os.Exit(1)
		  }
		  memoryInfo.SwapTotal = statNum
		case "SwapFree":
		  statNum, err := strconv.Atoi(value)
		  if err != nil {
		  	os.Exit(1)
		  }
		  memoryInfo.SwapFree = statNum
		}
	}

	memoryInfoContent.Close()

	var memInfoMap map[string]int
	memInfoBytes, _ := json.Marshal(memoryInfo)
	
	err := json.Unmarshal(memInfoBytes, &memInfoMap)
	if err != nil {
		os.Exit(1)
	}

	return memInfoMap
}