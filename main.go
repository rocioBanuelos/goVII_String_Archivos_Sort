package main

import (
	"fmt"
	"sort"
	"os"
	"bufio"
)

func main(){
	var cantidadStrings uint64
	var stringAux string
	scanner := bufio.NewScanner(os.Stdin)
	//Creación de archivo Ascendente
	fileAsc, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileAsc.Close()
	//Creación de archivo Descendente
	fileDes, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fileDes.Close()

	fmt.Println("\n Ordenar slice de strings")
	fmt.Printf("\nIngresa la cantidad de cadenas que deseas en el slice: ")
	fmt.Scanf("%d\n", &cantidadStrings)

	sliceString := make([]string, 0, cantidadStrings)

	//Leer valores para slice
	for i := 0; i < int(cantidadStrings); i++ {
		fmt.Printf("Ingresa el texto de la cadena %d: ",i+1)
		scanner.Scan()
		stringAux = scanner.Text()
		sliceString = append(sliceString,stringAux)
	}

	fmt.Println("\nSlice de cadenas desordenado: ", sliceString)

	//Ordenar ascendente
	sort.Strings(sliceString)
	fmt.Println("\nSlice de cadenas ordenado de forma Ascendente", sliceString)
	//Escribir en archivo ascendente.txt
	for i := 0; i < int(cantidadStrings); i++ {
		fileAsc.WriteString(sliceString[i]+"\n")
	}

	//Ordenar descendente
	sort.Sort(sort.Reverse(sort.StringSlice(sliceString)))
	fmt.Println("\nSlice de cadenas ordenado de forma Descendente",sliceString)
	//Escribir en archivo descendente.txt
	for i := 0; i < int(cantidadStrings); i++ {
		fileDes.WriteString(sliceString[i]+"\n")
	}
}