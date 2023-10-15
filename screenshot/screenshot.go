package screenshot

import (
	"github.com/go-rod/rod"
)

func MakeScreenshot(url string, path string) {
	page := rod.New().MustConnect().MustPage(url)
	page.MustWaitStable()
	page.MustScreenshotFullPage(path)
}
