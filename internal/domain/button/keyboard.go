package button

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type keyboardInputTracker struct {
	sink chan string
	id   int
}

type keyboardReader struct {
	mu    *sync.RWMutex
	sinks []keyboardInputTracker
}

var kr = (*keyboardReader)(nil)

func getKeyboardReader(id int) chan string {
	if kr == nil {
		kr = &keyboardReader{
			mu:    &sync.RWMutex{},
			sinks: make([]keyboardInputTracker, 0),
		}
		go runReader()
	}
	tracker := keyboardInputTracker{
		sink: make(chan string),
		id:   id,
	}
	kr.mu.Lock()
	kr.sinks = append(kr.sinks, tracker)
	kr.mu.Unlock()
	return tracker.sink
}

func remove(s []keyboardInputTracker, i int) []keyboardInputTracker {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func removeReader(id int) {
	kr.mu.Lock()
	removeIndex := -1
	for i, tracker := range kr.sinks {
		if tracker.id == id {
			removeIndex = i
			break
		}
	}
	if removeIndex != -1 {
		kr.sinks = remove(kr.sinks, removeIndex)
		if len(kr.sinks) == 0 {
			kr.mu.Unlock()
			kr = nil
			return
		}
	}
	kr.mu.Unlock()
}

func runReader() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if kr == nil {
			return
		}
		kr.mu.Lock()
		for _, kt := range kr.sinks {
			kt.sink <- text
		}
		kr.mu.Unlock()
	}
}

type keyboardButton struct {
	button *button
	id     int
	key    string
}

func newKeyboardButton(setup *Setup, b *button) (*keyboardButton, error) {
	log.Println("newKeyboardButton: creating")
	k := &keyboardButton{
		button: b,
		id:     setup.KeyOrGpioIn,
		key:    strconv.Itoa(setup.KeyOrGpioIn),
	}
	log.Println("newKeyboardButton: created")
	return k, nil
}

func (k *keyboardButton) runMainLoop() {
	log.Println("keyboardButton RunMainLoop: running")
	reader := getKeyboardReader(k.id)
	for {
		select {
		case _, ok := <-k.button.shutdowns:
			if !ok {
				goto finish
			}
		case text := <-reader:
			if strings.Contains(text, k.key) {
				k.button.bus.HandleButtonPush()
			}
		}
	}
finish:
	log.Println("keyboardButton RunMainLoop: closing")
	removeReader(k.id)
	k.button.wg.Done()
	log.Println("keyboardButton RunMainLoop: closed")
}
