package game

import (
	"fmt"

	"github.com/MarceloMPJ/contragirl/src/keyboard"
	keyboardlib "github.com/eiannone/keyboard"
)

func init() {
	fmt.Println("Carregando recursos ...")
	keyboard.AddObserver(ProcessComand)
}

// Start é a função usada para iniciar o jogo
func Start() {
	fmt.Println("Iniciando o Jogo")
	keyboard.ListenKeyboard()
	fmt.Println("Finalizando o Jogo")
}

// ProcessComand é a função usada para atualizar o jogo, através do parametro key
func ProcessComand(key keyboardlib.Key) {
	fmt.Println("Tecla apertada: ", key)
}
