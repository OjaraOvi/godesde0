package ejercicios

import "strconv"

func ConvNumerico(text string) (int, string) {
	num, err := strconv.Atoi(text)
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}
	if num > 100 {
		return num, "Es mayor a 100"
	} else {
		return num, "Es menor a 100"
	}
}
