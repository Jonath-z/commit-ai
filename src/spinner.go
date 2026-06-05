package src

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Spinner struct {
	stop chan struct{}
	done chan struct{}
	once sync.Once
}

var spinnerFrames = []rune{'⠋', '⠙', '⠹', '⠸', '⠼', '⠴', '⠦', '⠧', '⠇', '⠏'}

func stderrIsTTY() bool {
	fi, err := os.Stderr.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

func StartSpinner(msg string) *Spinner {
	s := &Spinner{
		stop: make(chan struct{}),
		done: make(chan struct{}),
	}
	if !stderrIsTTY() {
		fmt.Fprintf(os.Stderr, "%s...\n", msg)
		close(s.done)
		return s
	}
	go func() {
		defer close(s.done)
		ticker := time.NewTicker(80 * time.Millisecond)
		defer ticker.Stop()
		fmt.Fprint(os.Stderr, "\x1b[?25l")
		defer fmt.Fprint(os.Stderr, "\x1b[?25h\r\x1b[K")
		i := 0
		for {
			select {
			case <-s.stop:
				return
			case <-ticker.C:
				fmt.Fprintf(os.Stderr, "\r\x1b[K%c %s", spinnerFrames[i%len(spinnerFrames)], msg)
				i++
			}
		}
	}()
	return s
}

func (s *Spinner) Stop() {
	s.once.Do(func() {
		select {
		case <-s.done:
			return
		default:
		}
		close(s.stop)
		<-s.done
	})
}
