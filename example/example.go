package example

import (
	"sync"
	"time"

	"github.com/YashMathur/spingroup"
)

func main() {
	sg := spingroup.Create(50 * time.Millisecond)

	sg.AddFunc("test 5", func(wg *sync.WaitGroup, done *bool, success *bool) {
		time.Sleep(5 * time.Second)

		*done = true
		*success = true

		defer wg.Done()
	})

	sg.AddFunc("test 2", func(wg *sync.WaitGroup, done *bool, success *bool) {
		time.Sleep(2 * time.Second)

		*done = true
		*success = true

		defer wg.Done()
	})

	sg.AddFunc("test 6", func(wg *sync.WaitGroup, done *bool, success *bool) {
		time.Sleep(6 * time.Second)

		*done = true
		*success = true

		defer wg.Done()
	})

	sg.AddFunc("test 1", func(wg *sync.WaitGroup, done *bool, success *bool) {
		time.Sleep(1 * time.Second)

		*done = true
		*success = true

		defer wg.Done()
	})

	sg.Wait()
}
