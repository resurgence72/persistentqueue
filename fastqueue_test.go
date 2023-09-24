package persistentqueue

import (
	"fmt"
	"testing"
)

func TestFastQueue(t *testing.T) {
	fq := MustOpenFastQueue("./queue", "1:secret-url", 50, 524288000)
	dir := fq.Dirname()
	fmt.Println(dir)

	fq.MustWriteBlock([]byte("{__name__=abc1}"))
	fq.MustWriteBlock([]byte("{__name__=abc2}"))
	fq.MustWriteBlock([]byte("{__name__=abc3}"))
	fq.MustWriteBlock([]byte("{__name__=abc4}"))
	fq.MustWriteBlock([]byte("{__name__=abc5}"))

	for {
		block, ok := fq.MustReadBlock(nil)
		if !ok {
			break
		}
		fmt.Println("block: ", string(block))
	}
}
