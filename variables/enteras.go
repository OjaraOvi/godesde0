package variables

import "fmt"

func ShowEnteros() {
	var intComun int
	intDe32 := int32(10)
	intDe64 := int64(100)
	fmt.Println("int comun = ", intComun)
	fmt.Println("int 32 = ", intDe32)
	fmt.Println("int 64 = ", intDe64)
}
