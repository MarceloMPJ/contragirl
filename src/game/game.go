package game

import (
	"fmt"
	"time"

	"github.com/MarceloMPJ/contragirl/src/keyboard"
	keyboardlib "github.com/eiannone/keyboard"
)

var playerone player

type player struct {
	name string
	hp   int
	posx int
	posy int
}

func init() {
	fmt.Println("Carregando recursos ...")
	playerone = player{name: "Joaquim", hp: 100, posx: 0, posy: 0}
	keyboard.AddObserver(ProcessComand)
	go keyboard.ListenKeyboard()
}

// Start é a função usada para iniciar o jogo
func Start() {
	fmt.Println("Iniciando o Jogo")

	for {
		time.Sleep(10 * time.Second)
		fmt.Println(playerone)
	}
	fmt.Println("Finalizando o Jogo")
}

// ProcessComand é a função usada para atualizar o jogo, através do parametro key
func ProcessComand(key keyboardlib.Key) {
	fmt.Println("Tecla apertada: ", key)
}
