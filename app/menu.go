package app

import (
	"fmt"
	"log"
)

func Options() {
	fmt.Println("1. Comrpobar si existe el archivo.")
	fmt.Println("2. Leer el archivo.")
	fmt.Println("3. Leer el archivo linea a linea.")
	fmt.Println("4. Renombrar un archivo.")
	fmt.Println("5. Escribir en el archivo.")
	fmt.Println("6. Salir.")
	fmt.Printf("%s", "Introduce el numero de la opcion que desees: ")
}

func Menu() {
	Ficheros()
}

func Ficheros() {
	for {
		Options()
		var opcion int
		var file string
		var newFile string
		fmt.Scan(&opcion)
		switch opcion {
		case 1:
			fmt.Print("Introduce el nombre/ruta del archivo: ")
			fmt.Scan(&file)
			fmt.Printf("File %s exists: %t \n", file, FileExists(file))
			fmt.Println()
		case 2:
			fmt.Print("Introduce el nombre/ruta del archivo: ")
			fmt.Scan(&file)
			fmt.Println()
			ReadAllFile(file)
			fmt.Println()
		case 3:
			fmt.Print("Introduce el nombre/ruta del archivo: ")
			fmt.Scan(&file)
			fmt.Println()
			lines, err := ReadLineByLine(file)
			if err != nil {
				log.Fatalf("readLines: %s", err)
			}
			// print file contents
			for i, line := range lines {
				fmt.Printf("Linea %d: %s\n", i, line)
			}
			fmt.Println()
		case 4:
			fmt.Print("Introduce el nombre/ruta del archivo: ")
			fmt.Scan(&file)
			fmt.Print("Introduce el nuevo nombre del archivo: ")
			fmt.Scan(&newFile)
			RenameFile(file, newFile)
		case 5:
			//TODO: Escribir en archivo
			fmt.Println("Por implementar ...")
		case 6:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("No existe esa opcion...")
			continue
		}
	}
}
