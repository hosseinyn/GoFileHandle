package main

import (
	"fmt"
	"strings"
)
import "os"
import "io"
import "io/ioutil"
import "bufio"

func createFile(fileName string) {
	_, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating file : ", err)
		return
	}

	fmt.Println("Created file : ", fileName)
}

func checkFile(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("File exists")
	}

	if err != nil {
		fmt.Println("Error while checking file : ", err)
		return
	}

	if fileInfo != nil {
		if fileInfo.Size() == 0 {
			fmt.Println("File is empty.")
		}
	}
}

func writeFile(text string, fileName string) {
	data := []byte(text)
	err := ioutil.WriteFile(fileName, data, 0644)

	if err != nil {
		fmt.Println("Error writing file : ", err)
	} else {
		fmt.Println("File written")
	}
}

func readFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func deleteFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println("Error deleting file : ", err)
	} else {
		fmt.Println("File deleted.")
	}
}

func renameFile(fileName string, newName string) {
	err := os.Rename(fileName, newName)
	if err != nil {
		fmt.Println("Error renaming file : ", err)
	} else {
		fmt.Println("File renamed.")
	}
}

func copyFile(destinationFileName string, sourceFileName string) {
	_, err := os.Stat(destinationFileName)
	if os.IsNotExist(err) {
		fmt.Print("File does not exist. do you want to create it? (y,n)")
		var create string
		_, err := fmt.Scanf("%s", &create)

		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}

		if create == "y" {
			createFile(destinationFileName)
		} else {
			return
		}
	}

	destinationFile, err := os.OpenFile(destinationFileName, os.O_RDWR, 0644)
	if err != nil {
	}
	if err != nil {
		fmt.Println("Error opening file : ", err)
	}

	defer destinationFile.Close()

	_, err = os.Stat(sourceFileName)
	if os.IsNotExist(err) {
		fmt.Println("Source file does not exist.")
		return
	}

	sourceFile, err := os.Open(sourceFileName)

	if err != nil {
		fmt.Println("Error opening file : ", err)
	}

	defer sourceFile.Close()

	bytes, err := io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file : ", err)
	} else {
		fmt.Println("Copied file (bytes) : ", bytes)
	}
}

func fileInformation(fileName string) {
	fileInfo, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist.")
		return
	}

	if err != nil {
		fmt.Println("Error opening file : ", err)
		return
	}

	if fileInfo != nil {
		fileType := fileInfo.Mode()
		if fileType.IsRegular() {
			fmt.Println("File is a regular file.")
		} else if fileType.IsDir() {
			fmt.Println("File is a directory.")
		}

		filePerm := fileInfo.Mode().Perm()
		fmt.Println("File permissions: ", filePerm)
	}
}

func listDirectory(directoryName string) {
	if directoryName == "" {
		directoryName = "."
	}
	fileInfos, err := ioutil.ReadDir(directoryName)
	if err != nil {
		fmt.Println("Error reading directory : ", err)
		return
	}
	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			fmt.Printf("Directory : %s | Size : %d | Permission : %s \n", fileInfo.Name(), fileInfo.Size(), fileInfo.Mode().Perm())

		} else {
			fmt.Printf("File : %s | Size : %d | Permission : %s | Created at : %s \n", fileInfo.Name(), fileInfo.Size(), fileInfo.Mode().Perm(), fileInfo.ModTime())
		}
	}
}

func main() {
	works := []string{"Create file", "Check file", "Write file", "Read file", "Delete file", "Rename file", "Copy file", "File information", "List directory"}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println()
	fmt.Println("		Golang File Handling Project		")

	fmt.Print("############ \n")
	for index, work := range works {
		fmt.Printf("%d- : %s\n", index+1, work)
	}

	fmt.Print("Select what to do : ")
	var work int
	_, err := fmt.Scanf("%d", &work)
	if err != nil {
		fmt.Println("Error while reading : ", err)
		return
	}
	_, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading : ", err)
		return
	}

	switch work {
	case 1:
		fmt.Print("Enter file name: ")
		fileName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}
		fileName = strings.TrimSpace(fileName)
		createFile(fileName)
	case 2:
		fmt.Print("Enter file name: ")
		fileName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}
		fileName = strings.TrimSpace(fileName)
		checkFile(fileName)
	case 3:
		fmt.Print("Enter file name: ")
		fileName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}
		fileName = strings.TrimSpace(fileName)
		_, err = os.Stat(fileName)
		if os.IsNotExist(err) {
			fmt.Print("File does not exist , do you want to create it? (y,n)")
			var create string
			_, err := fmt.Scanf("%s", &create)
			if err != nil {
				fmt.Println("Error while reading : ", err)
				return
			}
			if create == "y" {
				createFile(fileName)
			} else {
				return
			}
		}

		fmt.Print("Enter file content: ")
		content, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}
		content = strings.TrimSpace(content)
		writeFile(content, fileName)
	case 4:
		fmt.Print("Enter file name: ")
		fileName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}
		fileName = strings.TrimSpace(fileName)
		_, err = os.Stat(fileName)
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			return
		}

		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}

		readFile(fileName)
	case 5:
		fmt.Print("Enter file name: ")
		fileName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}
		fileName = strings.TrimSpace(fileName)
		_, err = os.Stat(fileName)
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			return
		}

		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}

		deleteFile(fileName)
	case 6:
		fmt.Print("Enter file name: ")
		fileName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}
		fileName = strings.TrimSpace(fileName)
		_, err = os.Stat(fileName)
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
			return
		}

		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}

		fmt.Print("Enter new file name: ")
		newName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}
		newName = strings.TrimSpace(newName)

		renameFile(fileName, newName)
	case 7:
		fmt.Print("Enter destination file name: ")
		destinationFileName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}
		destinationFileName = strings.TrimSpace(destinationFileName)

		fmt.Print("Enter source file name:  ")
		sourceFileName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}
		strings.TrimSpace(sourceFileName)

		copyFile(destinationFileName, sourceFileName)
	case 8:
		fmt.Print("Enter file name: ")
		fileName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}
		fileName = strings.TrimSpace(fileName)

		fileInformation(fileName)
	case 9:
		fmt.Print("Enter directory name (leave blank for current directory) : ")
		directoryName, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading : ", err)
			return
		}
		directoryName = strings.TrimSpace(directoryName)

		listDirectory(directoryName)
	}
}
