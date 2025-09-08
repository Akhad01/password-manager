package files

import (
	"fmt"
	"os"
)

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	if _, err = file.Write(content); err != nil {
		fmt.Println(err)
		return 
	}

	fmt.Println("Запись успешно")
}

func ReadFile() {
	data, err := os.ReadFile("files.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}