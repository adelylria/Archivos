package app

import (
	"bufio"
	"bytes"
	"encoding/json"
	"files/app/models"
	"fmt"
	"io/fs"
	"os"
)

// FileExists verifica si un archivo existe en la ruta especificada.
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func ReadAllFile(file string) {
	a, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	str := string(a)

	fmt.Println(str)
}

// read line by line into memory
// all file contents is stores in lines[]
func ReadLineByLine(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func RenameFile(fileName string, newName string) {
	os.Rename(fileName, newName)
}

// Refractorizado
/*func WriteFile(fileName, name string, age int) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
	}
	defer f.Close()
	u := models.NewUser(name, age)
	buf := new(bytes.Buffer)
	fileErr := json.NewEncoder(buf).Encode(u)
	if fileErr != nil {
		fmt.Println("Error al codificar en JSON:", fileErr)
	}
	_, writeErr := fmt.Fprint(f, buf.String())
	if writeErr != nil {
		fmt.Println("Error al escribir en el archivo:", writeErr)
	}
}*/

// openFile abre el archivo especificado en modo de apertura para escritura y anexar.
// Devuelve un puntero al archivo y un error en caso de que ocurra uno.
func openFile(fileName string) (*os.File, error) {
	return os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, fs.ModeAppend)
}

// encodeUser toma una estructura de usuario, la codifica en formato JSON y devuelve
// los datos codificados en bytes y un error en caso de que ocurra uno.
func encodeUser(u models.Usuario) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(u)
	return buf.Bytes(), err
}

// writeFile escribe los datos en el archivo especificado.
// Devuelve un error en caso de que ocurra uno.
func writeFile(f *os.File, data []byte) error {
	_, err := fmt.Fprint(f, string(data))
	return err
}

// WriteFile es la función principal que escribe los datos de un usuario en un archivo.
func WriteToFile(fileName, name string, age int) {
	// Abrir el archivo
	f, err := openFile(fileName)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer f.Close()

	// Crear una estructura de usuario
	u := models.NewUser(name, age)

	// Codificar el usuario en JSON
	encodedUser, err := encodeUser(u)
	if err != nil {
		fmt.Println("Error al codificar en JSON:", err)
		return
	}

	// Escribir los datos codificados en el archivo
	writeErr := writeFile(f, encodedUser)
	if writeErr != nil {
		fmt.Println("Error al escribir en el archivo:", writeErr)
	}
}
