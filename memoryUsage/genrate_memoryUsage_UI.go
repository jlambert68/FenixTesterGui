package memoryUsage

import (
	sharedCode "FenixTesterGui/common_code"
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/skoona/sknlinechart"
	"runtime"
	"time"
)

//go:embed graphics/graph_84x48.png
var memoryUsage []byte

// Genrate the UI-component to be used at the bottom of the Fenix-UI, to show ongoing probes(pigs)
func GenerateMemoryUsageIcon() (
	memoryUsageContainer *fyne.Container) {

	// turn the raw bytes into a Fyne resource
	var memoryUsageResource fyne.Resource
	memoryUsageResource = fyne.NewStaticResource("picture.png", memoryUsage)

	// create an image from that resource
	var memoryUsageImage *canvas.Image
	memoryUsageImage = canvas.NewImageFromResource(memoryUsageResource)

	// Set size on image
	memoryUsageImage.SetMinSize(fyne.NewSize(84, 48))

	// Prepare and start collecting data
	var memoryUsageWindow fyne.Window
	memoryUsageWindow = (*sharedCode.FenixAppPtr).NewWindow("Memory Usage")
	// 1) Take an initial sample so series aren't empty
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	now := time.Now()
	ptSec0 := sknlinechart.NewChartDatapoint(
		float32(ms.Alloc)/(1024*1024), theme.ColorBlue, now.Format("15:04:05"),
	)
	ptMin0 := sknlinechart.NewChartDatapoint(
		float32(ms.Alloc)/(1024*1024), theme.ColorRed, now.Format("15:04"),
	)

	// 2) Build the chart, seeding both series
	opts := sknlinechart.NewChartOptions(
		sknlinechart.WithTitle("Heap Alloc (MiB)"),
		sknlinechart.WithLeftScaleLabel("MiB"),
		// ← add this to get numeric Y-axis ticks every 10 MiB
		sknlinechart.WithYScaleFactor(20),
		sknlinechart.WithBottomLeftLabel("Time"),
		sknlinechart.WithColorLegend(true),
		sknlinechart.WithDataPoints(map[string][]*sknlinechart.ChartDatapoint{
			"Every Second": {&ptSec0},
			"Every Minute": {&ptMin0},
		}),
	)

	chart, err := sknlinechart.NewWithOptions(opts)

	if err != nil {
		panic(err)
	}
	chart.SetMinSize(fyne.NewSize(600, 400))

	// 3) Goroutine for 1-second sampling

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for t := range ticker.C {
			runtime.ReadMemStats(&ms)
			mib := float32(ms.Alloc) / (1024 * 1024)
			pt := sknlinechart.NewChartDatapoint(mib, theme.ColorBlue, t.Format("15:04:05"))

			//sysMb := float32(ms.HeapSys) / (1024 * 1024)
			//pt2 := sknlinechart.NewChartDatapoint(sysMb, theme.ColorBlue, t.Format("15:04:05"))

			//objs := float32(ms.HeapObjects)
			//pt3 := sknlinechart.NewChartDatapoint(objs, theme.ColorBlue, t.Format("15:04:05"))

			// marshal back to the UI thread
			fyne.Do(func() {
				chart.ApplyDataPoint("Every Second", &pt)
				//chart.ApplyDataPoint("Every Second", &pt2)
				//chart.ApplyDataPoint("Every Second", &pt3)
				chart.Refresh()
			})
		}
	}()

	// 4) Goroutine for 1-minute sampling
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for t := range ticker.C {
			runtime.ReadMemStats(&ms)
			mib := float32(ms.Alloc) / (1024 * 1024)
			pt := sknlinechart.NewChartDatapoint(mib, theme.ColorRed, t.Format("15:04"))

			fyne.Do(func() {
				chart.ApplyDataPoint("Every Minute", &pt)
				chart.Refresh()
			})
		}
	}()

	btn := widget.NewButtonWithIcon("", memoryUsageResource, func() {

		// Open new window with statistics
		// 5) Show the window
		memoryUsageWindow.SetContent(container.NewStack(chart))
		memoryUsageWindow.Resize(fyne.NewSize(620, 420))
		memoryUsageWindow.Show()
	})

	btn.Resize(fyne.NewSize(84, 48))

	memoryUsageContainer = container.NewStack(btn)

	return memoryUsageContainer
}

// Open up the statistics window
func openStatisticsWindow() {

	var memoryUsageWindow fyne.Window
	memoryUsageWindow = (*sharedCode.FenixAppPtr).NewWindow("Memory Usage")
	// 1) Take an initial sample so series aren't empty
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	now := time.Now()
	ptSec0 := sknlinechart.NewChartDatapoint(
		float32(ms.Alloc)/(1024*1024), theme.ColorBlue, now.Format("15:04:05"),
	)
	ptMin0 := sknlinechart.NewChartDatapoint(
		float32(ms.Alloc)/(1024*1024), theme.ColorRed, now.Format("15:04"),
	)

	// 2) Build the chart, seeding both series
	opts := sknlinechart.NewChartOptions(
		sknlinechart.WithTitle("Heap Alloc (MiB)"),
		sknlinechart.WithLeftScaleLabel("MiB"),
		// ← add this to get numeric Y-axis ticks every 10 MiB
		sknlinechart.WithYScaleFactor(20),
		sknlinechart.WithBottomLeftLabel("Time"),
		sknlinechart.WithColorLegend(true),
		sknlinechart.WithDataPoints(map[string][]*sknlinechart.ChartDatapoint{
			"Every Second": {&ptSec0},
			"Every Minute": {&ptMin0},
		}),
	)

	chart, err := sknlinechart.NewWithOptions(opts)

	if err != nil {
		panic(err)
	}
	chart.SetMinSize(fyne.NewSize(600, 400))

	// 3) Goroutine for 1-second sampling

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for t := range ticker.C {
			runtime.ReadMemStats(&ms)
			mib := float32(ms.Alloc) / (1024 * 1024)
			pt := sknlinechart.NewChartDatapoint(mib, theme.ColorBlue, t.Format("15:04:05"))

			//sysMb := float32(ms.HeapSys) / (1024 * 1024)
			//pt2 := sknlinechart.NewChartDatapoint(sysMb, theme.ColorBlue, t.Format("15:04:05"))

			//objs := float32(ms.HeapObjects)
			//pt3 := sknlinechart.NewChartDatapoint(objs, theme.ColorBlue, t.Format("15:04:05"))

			// marshal back to the UI thread
			fyne.Do(func() {
				chart.ApplyDataPoint("Every Second", &pt)
				//chart.ApplyDataPoint("Every Second", &pt2)
				//chart.ApplyDataPoint("Every Second", &pt3)
				chart.Refresh()
			})
		}
	}()

	// 4) Goroutine for 1-minute sampling
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for t := range ticker.C {
			runtime.ReadMemStats(&ms)
			mib := float32(ms.Alloc) / (1024 * 1024)
			pt := sknlinechart.NewChartDatapoint(mib, theme.ColorRed, t.Format("15:04"))

			fyne.Do(func() {
				chart.ApplyDataPoint("Every Minute", &pt)
				chart.Refresh()
			})
		}
	}()

	// 5) Show the window
	memoryUsageWindow.SetContent(container.NewStack(chart))
	memoryUsageWindow.Resize(fyne.NewSize(620, 420))
	memoryUsageWindow.Show()
}
