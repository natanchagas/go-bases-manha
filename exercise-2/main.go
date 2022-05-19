package main

import "fmt"

var (
	temperatura        int8    = 11
	umidade            float32 = 76.2
	pressaoAtmosferica int16   = 1016
)

func main() {
	fmt.Println("Temperatura:", temperatura, "graus")
	fmt.Println("Umidade do ar:", umidade, "%")
	fmt.Println("Pressão atmosférica:", pressaoAtmosferica, "hPa")
}
