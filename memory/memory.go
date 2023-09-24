package memory

import (
	"sync"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/flagutil"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/logger"
)

var (
	allowedPercent float64
	allowedBytes   = flagutil.NewBytes("", 0, "")
)

var (
	allowedMemory   int
	remainingMemory int
	memoryLimit     int
)
var once sync.Once

func initOnce() {
	memoryLimit = sysTotalMemory()

	allowedPercent = 60
	if allowedBytes.N <= 0 {
		if allowedPercent < 1 || allowedPercent > 100 {
			logger.Fatalf("FATAL: -memory.allowedPercent must be in the range [1...100]; got %g", allowedPercent)
		}
		percent := allowedPercent / 100
		allowedMemory = int(float64(memoryLimit) * percent)
		remainingMemory = memoryLimit - allowedMemory
		logger.Infof("limiting caches to %d bytes, leaving %d bytes to the OS according to -memory.allowedPercent=%g", allowedMemory, remainingMemory, allowedPercent)
	} else {
		allowedMemory = allowedBytes.IntN()
		remainingMemory = memoryLimit - allowedMemory
		logger.Infof("limiting caches to %d bytes, leaving %d bytes to the OS according to -memory.allowedBytes=%s", allowedMemory, remainingMemory, allowedBytes.String())
	}
}

// Allowed returns the amount of system memory allowed to use by the app.
//
// The function must be called only after flag.Parse is called.
func Allowed() int {
	once.Do(initOnce)
	return allowedMemory
}
