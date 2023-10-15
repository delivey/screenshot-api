package screenshot

import "os"

func CleanupFromPath(path string) {
	os.Remove(path)
}
