package screenshot

import (
	"github.com/go-rod/rod"
)

func MakeScreenshot(url string) {
	page := rod.New().MustConnect().MustPage(url)
	page.MustWaitStable()

	page.MustScreenshotFullPage("tmp/a.png")
}
