package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaDeMultiplicar() {
	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error
	for {
		fmt.Println("Ingresar numero : ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}

	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d X %d = %d \n", numero, i, i*numero)

	}
}
