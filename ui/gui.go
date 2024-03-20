package ui

import (
	"fmt"
	"github.com/hakmkoyan/metrics/memory"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	memoryInfo *memory.MemoryInfo
	interval int
	content *fyne.Container
)

func GetMetrics() map[string]int {
	memoryInfo = &memory.MemoryInfo{}
	memInfoMap := memory.GetMemoryInfo(memoryInfo)

	return memInfoMap
}

func MetricsDashboard(window fyne.Window, interval int) {
	MemTotal := widget.NewLabel("")
	MemFree := widget.NewLabel("")
	MemAvailable := widget.NewLabel("")
	Buffers := widget.NewLabel("")
	Cached := widget.NewLabel("")
	Active := widget.NewLabel("")
	Inactive := widget.NewLabel("")
	SwapCached := widget.NewLabel("")
	SwapTotal := widget.NewLabel("")
	SwapFree := widget.NewLabel("")

	content := container.NewVBox(
		MemTotal,
		MemFree,
		MemAvailable,
		Buffers,
		Cached,
		Active,
		Inactive,
		SwapCached,
		SwapTotal,
		SwapFree,
	)

	go func() {
		for {
			metrics := GetMetrics()
			MemTotal.TextStyle = fyne.TextStyle{Bold: true}
			memTotal := fmt.Sprintf("Total Memory: %f", metrics["MemTotal"])
			MemTotal.SetText(memTotal)
			MemFree.TextStyle = fyne.TextStyle{Bold: true}
			memFree := fmt.Sprintf("Free Memory: %.2f%%\n", metrics["MemFree"])
			MemFree.SetText(memFree)
			MemAvailable.TextStyle = fyne.TextStyle{Bold: true}
			memAvailable := fmt.Sprintf("Available Memory: %.2f%%\n", metrics["MemAvailable"])
			MemAvailable.SetText(memAvailable)
			Buffers.TextStyle = fyne.TextStyle{Bold: true}
			buffers := fmt.Sprintf("Buffers: %.2f%%\n", metrics["Buffers"])
			Buffers.SetText(buffers)
			Cached.TextStyle = fyne.TextStyle{Bold: true}
			cached := fmt.Sprintf("Cached: %.2f%%\n", metrics["Cached"])
			Cached.SetText(cached)
			Active.TextStyle = fyne.TextStyle{Bold: true}
			active := fmt.Sprintf("Active: %.2f%%\n", metrics["Active"])
			Active.SetText(active)
			Inactive.TextStyle = fyne.TextStyle{Bold: true}
			inactive := fmt.Sprintf("Inactive: %.2f%%\n", metrics["Inactive"])
			Inactive.SetText(inactive)
			SwapCached.TextStyle = fyne.TextStyle{Bold: true}
			swapCached := fmt.Sprintf("Swap Cached: %.2f%%\n", metrics["SwapCached"])
			SwapCached.SetText(swapCached)
			SwapTotal.TextStyle = fyne.TextStyle{Bold: true}
			swapTotal := fmt.Sprintf("Swap Total: %.2f%%\n", metrics["SwapTotal"])
			SwapTotal.SetText(swapTotal)
			SwapFree.TextStyle = fyne.TextStyle{Bold: true}
			swapFree := fmt.Sprintf("Swap Free: %.2f%%\n", metrics["SwapFree"])
			SwapFree.SetText(swapFree)
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}()
	
	window.SetContent(content)
}

func StartDashboard(window fyne.Window) {
	intervalInput := widget.NewEntry()
	intervalInput.SetPlaceHolder("Interval")

	content = container.NewVBox(
		intervalInput,
		widget.NewButton("Start", func() {
			interval, _ = strconv.Atoi(string(intervalInput.Text))
			MetricsDashboard(window, interval)
		}),
	)

	window.SetContent(content)
}

func MainDashboard() {
	metricsApp := app.New()
	window := metricsApp.NewWindow("Metrifics")
	
	StartDashboard(window)

	window.Resize(fyne.Size{
		Width: 350,
		Height: 150,
	})
	window.ShowAndRun()
}