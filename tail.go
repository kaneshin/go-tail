package tail

import (
	"bufio"
	"fmt"
	"io"
	"time"
)

func Exec(r io.Reader) error {

	ch := make(chan string, 10)

	sc := bufio.NewScanner(r)
	for sc.Scan() {
		if len(ch) >= 10 {
			<-ch
		}
		ch <- sc.Text()
	}

	go func() {
		for {
			t, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(t)
		}
	}()

	go func() {
		for {
			sc := bufio.NewScanner(r)
			for sc.Scan() {
				ch <- sc.Text()
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	time.Sleep(time.Minute)

	return nil
}
