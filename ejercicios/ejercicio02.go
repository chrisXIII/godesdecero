package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func TabladeMultiplicar() string {
	scan := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingresar Número: ")

		if scan.Scan() {
			numero, err = strconv.Atoi(scan.Text())

			if err != nil {
				fmt.Println("Error: Número incorrecto")
				continue
			}
		}
		break
	}

	for i := 1; i <= 10; i++ {
		num := numero * i
		texto += fmt.Sprintf("%d x %d = %d \n",numero,i,num)
	}

	return texto
}
