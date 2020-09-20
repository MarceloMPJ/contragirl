package keyboard

import (
	"fmt"

	keyboard "github.com/eiannone/keyboard"
)

// Input é uma struct para guardar o valor lido pelo teclado
type Input struct {
	key string
}

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
