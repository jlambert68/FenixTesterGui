package memoryUsage

import (
	sharedCode "FenixTesterGui/common_code"
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/skoona/sknlinechart"
	"runtime"
	"sync"
	"time"
)

//go:embed graphics/graph_84x48.png
var memoryUsage []byte

var memoryUsageWindow fyne.Window

// GenerateMemoryUsageIcon
// Genrate the UI-component to be used at the bottom of the Fenix-UI, to show ongoing probes(pigs)
func GenerateMemoryUsageIcon() (
	memoryUsageContainer *ClickableImageStruct) {

	// turn the raw bytes into a Fyne resource
	var memoryUsageResource fyne.Resource
	memoryUsageResource = fyne.NewStaticResource("picture.png", memoryUsage)

	// create an image from that resource
	var memoryUsageImage *canvas.Image
	memoryUsageImage = canvas.NewImageFromResource(memoryUsageResource)

	// Set size on image
	memoryUsageImage.SetMinSize(fyne.NewSize(84, 48))

	var clickableContainer *ClickableImageStruct
	clickableContainer = NewClickableImage(memoryUsageImage, openStatisticsWindowFunction)

	return clickableContainer
}

// Open up the statistics window
func openStatisticsWindowFunction(clickableContainer *ClickableImageStruct) {

	if clickableContainer.AlreadyOpen == false {

		// Initiate new window
		memoryUsageWindow = (*sharedCode.FenixAppPtr).NewWindow("Memory Usage")
		clickableContainer.AlreadyOpen = true

	} else {
		// Set focus on the window
		memoryUsageWindow.RequestFocus()

		return
	}

	var wg sync.WaitGroup
	var waitingToCloseDown bool = false

	var stopOneMinuteTicker chan struct{}
	stopOneMinuteTicker = make(chan struct{})

	// We have 2 flows
	wg.Add(2)

	memoryUsageWindow.SetCloseIntercept(func() {

		// Initiate close down
		waitingToCloseDown = true

		// Wait for go routines to finish
		wg.Wait()

		// Remove the intercept so Close() will actually close
		memoryUsageWindow.SetCloseIntercept(nil)

		defer func() {
			clickableContainer.AlreadyOpen = false
		}()

		memoryUsageWindow.Close()

	})

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
		// Stop Window to be double opened

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

			if waitingToCloseDown == false {
				fyne.Do(func() {

					chart.ApplyDataPoint("Every Second", &pt)
					//chart.ApplyDataPoint("Every Second", &pt2)
					//chart.ApplyDataPoint("Every Second", &pt3)
					chart.Refresh()

				})
			} else {

				wg.Done()

				// Stop 1 minute ticker
				stopOneMinuteTicker <- struct{}{}
				return
			}

		}
	}()

	// 4) Goroutine for 1-minute sampling
	go func() {

		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()

		for {
			select {
			case t := <-ticker.C:

				runtime.ReadMemStats(&ms)
				mib := float32(ms.Alloc) / (1024 * 1024)
				pt := sknlinechart.NewChartDatapoint(mib, theme.ColorRed, t.Format("15:04"))

				if waitingToCloseDown == false {

					fyne.Do(func() {
						chart.ApplyDataPoint("Every Minute", &pt)
						chart.Refresh()
					})
				} else {

					wg.Done()
					return
				}

			case <-stopOneMinuteTicker:

				wg.Done()
				return
			}

		}
	}()

	// 5) Show the window
	memoryUsageWindow.SetContent(container.NewStack(chart))
	memoryUsageWindow.Resize(fyne.NewSize(620, 420))
	memoryUsageWindow.Show()
}
