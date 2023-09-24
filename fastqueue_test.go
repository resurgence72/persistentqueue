package persistentqueue

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/cespare/xxhash/v2"
)

func TestFastQueue(t *testing.T) {
	queuePath := filepath.Join("queue", "persistent-queue", fmt.Sprintf("0_%016X", xxhash.Sum64([]byte("remoteAddr"))))
	fq := MustOpenFastQueue(queuePath, "1:secret-url", 3200, 524288000)
	dir := fq.Dirname()
	fmt.Println(dir)

	go func() {
		for i := 0; i < 100; i++ {
			fq.MustWriteBlock([]byte(fmt.Sprintf("{__name__=abc%d}", i)))
			time.Sleep(time.Millisecond * 200)
		}
		fq.MustClose()
	}()

	for {
		block, ok := fq.MustReadBlock(nil)
		if !ok {
			break
		}
		fmt.Println("block: ", string(block))
	}
}
