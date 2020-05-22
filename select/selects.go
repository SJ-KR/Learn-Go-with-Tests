package selects

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	_, _ = http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	_, _ = http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		fmt.Println(a)
		fmt.Println(aDuration)
		return a
	}
	fmt.Println(b)
	fmt.Println(bDuration)
	return b
}
