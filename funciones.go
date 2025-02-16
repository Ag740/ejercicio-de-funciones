package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lectura := bufio.NewReader(os.Stdin)

	fmt.Print("Ingrese un mensaje: ")
	tex, _ := lectura.ReadString('\n')

	// Funciones de strings
	fmt.Print("                 \n")
	frase := strings.ToUpper(tex)
	fmt.Print("Nombre en mayuscula:", frase)

	texto := strings.Contains(tex, tex)
	fmt.Println("Contiene elementos especificos:", texto)

	cad := strings.Count(tex, "a")
	fmt.Println("Cuenta los caracteres de la cadena:", cad)

	fra := strings.Repeat(tex, 2)
	fmt.Print("Repiticon de la cadena:", fra)

	text := strings.Fields(tex)
	fmt.Println("Division de la cadena:", text)

	fmt.Print("                 \n")

	fmt.Print("Ingresa un número: ")
	num, _ := lectura.ReadString('\n')
	num = strings.TrimSpace(num)

	// Funciones de strconv
	fmt.Print("                 \n")
	numEnte, _ := strconv.Atoi(num)
	fmt.Println("Número entero:", numEnte)

	numDeci, _ := strconv.ParseFloat(num, 64)
	fmt.Println("Número decimal:", numDeci)

	numLogi, _ := strconv.ParseBool(num)
	fmt.Println("¿Es verdadero o falso:?", numLogi)

	numInt64, _ := strconv.ParseInt(num, 10, 64)
	fmt.Println("Número como int64:", numInt64)

	numUint, _ := strconv.ParseUint(num, 10, 64)
	fmt.Println("Número como uint64:", numUint)
}
