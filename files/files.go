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

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)

	if err != nil {
		return nil, err
	}

	return data, nil
}