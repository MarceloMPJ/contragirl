package keyboard

import (
	"fmt"

	keyboard "github.com/eiannone/keyboard"
)

var observers []func(keyboard.Key)

// ListenKeyboard é uma função para escutar todos os comandos vindo do teclado
func ListenKeyboard() {
	openKeyboard()
	defer closeKeyboard()
	getKey()
}

func openKeyboard() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
}

func closeKeyboard() {
	_ = keyboard.Close()
}

func getKey() {
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyEsc {
			break
		}
		fmt.Println("Key:", key)
	}
}

func addObserver(observerFunction func(keyboard.Key)) {
	observers = append(observers, observerFunction)
}

func notifyAllObservers(key keyboard.Key) {
	for _, observer := range observers {
		observer(key)
	}
}
