//Desafio de Conversão de escalas termométricas
// lembrando que este desafio pode ser feito de varias formas 
// objetivo - converter o ponto de ebulição de kelvin para celsius.

package main

import "fmt"

var kelvin int = 373

func main() {

	k := kelvin

	fmt.Println(
		"A minha conta Sobre a ebulição, deu o seguinte valor: ",
		"C = ",
		(k - 273),
	)

}
