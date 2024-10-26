package main

import (
	"archive/zip"
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// Структура для работы с JSON и XML
type Example struct {
	Name  string `json:"name" xml:"name"`
	Value int    `json:"value" xml:"value"`
}

func waitForEnter() {
	fmt.Println("\nНажмите Enter для продолжения...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func listDisks() {
	fmt.Println("Информация о логических дисках:")
	switch runtime.GOOS {
	case "windows":
		cmd := "wmic logicaldisk get caption,volumename,size,filesystem"
		output, err := exec.Command("cmd", "/C", cmd).Output()
		if err != nil {
			fmt.Println("Ошибка получения информации о дисках:", err)
			return
		}
		fmt.Println(string(output))
	case "darwin":
		cmd := "df -h"
		output, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println("Ошибка получения информации о дисках:", err)
			return
		}
		fmt.Println(string(output))
	default:
		fmt.Println("Функция не поддерживается для этой операционной системы")
	}
}

func createFile(fileName string) {
	fmt.Println("Введите строку для записи в файл:")
	reader := bufio.NewReader(os.Stdin)
	content, _ := reader.ReadString('\n')

	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
	} else {
		fmt.Println("Файл успешно создан:", fileName)
	}
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	fmt.Println("Содержимое файла:", string(data))
}

func deleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Ошибка удаления файла:", err)
	} else {
		fmt.Println("Файл успешно удалён:", fileName)
	}
}

func createJSON(fileName string) {
	fmt.Println("Введите имя:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Введите значение:")
	var value int
	fmt.Scanln(&value)

	exampleJSON := Example{Name: name, Value: value}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка создания JSON файла:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(exampleJSON)
	if err != nil {
		fmt.Println("Ошибка записи в JSON файл:", err)
	} else {
		fmt.Println("JSON файл успешно создан:", fileName)
	}
}

func readJSON(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения JSON файла:", err)
		return
	}
	fmt.Println("Содержимое JSON файла:", string(data))
}

func createXMLFromFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка создания XML файла:", err)
		return
	}
	defer file.Close()

	fmt.Println("Введите имя:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Введите значение:")
	var value int
	fmt.Scanln(&value)

	exampleXML := Example{Name: name, Value: value}

	encoder := xml.NewEncoder(file)
	err = encoder.Encode(exampleXML)
	if err != nil {
		fmt.Println("Ошибка записи в XML файл:", err)
	} else {
		fmt.Println("XML файл успешно создан:", fileName)
	}
}

func readXML(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения XML файла:", err)
		return
	}
	fmt.Println("Содержимое XML файла:", string(data))
}

func createZipArchive(zipFileName string, fileName string) {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		fmt.Println("Ошибка создания ZIP файла:", err)
		return
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fileToZip, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		fmt.Println("Ошибка получения информации о файле:", err)
		return
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		fmt.Println("Ошибка создания заголовка архива:", err)
		return
	}

	header.Name = fileName
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		fmt.Println("Ошибка создания записи в архиве:", err)
		return
	}

	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		fmt.Println("Ошибка копирования содержимого файла в архив:", err)
		return
	}

	fmt.Println("ZIP архив успешно создан:", zipFileName)
}

func unzipArchive(zipFileName string, destDir string) {
	zipReader, err := zip.OpenReader(zipFileName)
	if err != nil {
		fmt.Println("Ошибка открытия архива:", err)
		return
	}
	defer zipReader.Close()

	for _, file := range zipReader.File {
		path := filepath.Join(destDir, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
		} else {
			if err := os.MkdirAll(filepath.Dir(path), file.Mode()); err != nil {
				fmt.Println("Ошибка создания директории для файла:", err)
				return
			}

			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
			if err != nil {
				fmt.Println("Ошибка создания файла для извлечения:", err)
				return
			}
			defer outFile.Close()

			rc, err := file.Open()
			if err != nil {
				fmt.Println("Ошибка открытия файла из архива:", err)
				return
			}
			defer rc.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				fmt.Println("Ошибка извлечения файла из архива:", err)
				return
			}

			fmt.Println("Файл успешно извлечён из архива:", path)
		}
	}
}

func main() {
	// Информация о дисках (только для Windows)
	listDisks()
	waitForEnter()

	// Работа с файлами
	fmt.Println("\nРабота с файлами:")
	fileName := "example.txt"
	createFile(fileName)
	waitForEnter()

	readFile(fileName)
	waitForEnter()

	deleteFile(fileName)
	waitForEnter()

	// Работа с JSON
	fmt.Println("\nРабота с JSON:")
	jsonFileName := "example.json"
	createJSON(jsonFileName)
	waitForEnter()

	readJSON(jsonFileName)
	waitForEnter()

	deleteFile(jsonFileName)
	waitForEnter()

	// Работа с XML
	fmt.Println("\nРабота с XML:")
	xmlFileName := "example.xml"
	createXMLFromFile(xmlFileName)
	waitForEnter()

	readXML(xmlFileName)
	waitForEnter()

	deleteFile(xmlFileName)
	waitForEnter()

	// Работа с ZIP архивом
	fmt.Println("\nРабота с ZIP архивом:")
	fileToArchive := "file_for_zip.txt"
	createFile(fileToArchive)
	waitForEnter()

	zipFileName := "archive.zip"
	createZipArchive(zipFileName, fileToArchive)
	waitForEnter()

	unzipArchive(zipFileName, "./unzipped")
	waitForEnter()

	deleteFile(fileToArchive)
	deleteFile(zipFileName)
	waitForEnter()
}
