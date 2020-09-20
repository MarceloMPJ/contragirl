package keyboard

import (
	"github.com/eiannone/keyboard"
	keyboardlib "github.com/eiannone/keyboard"
)

var observers []func(keyboardlib.Key)

// ListenKeyboard é uma função para escutar todos os comandos vindo do teclado
func ListenKeyboard() {
	openKeyboard()
	defer closeKeyboard()
	getKey()
}

func openKeyboard() {
	if err := keyboardlib.Open(); err != nil {
		panic(err)
	}
}

func closeKeyboard() {
	_ = keyboardlib.Close()
}

func getKey() {
	for {
		_, key, err := keyboardlib.GetKey()
		if err != nil {
			panic(err)
		}
		if key == keyboardlib.KeyEsc {
			break
		}
		notifyAllObservers(key)
	}
}

// AddObserver é a função utilizada para adicionar um novo observerFunction para o slice observers
func AddObserver(observerFunction func(keyboard.Key)) {
	observers = append(observers, observerFunction)
}

func notifyAllObservers(key keyboardlib.Key) {
	for _, observer := range observers {
		observer(key)
	}
}
