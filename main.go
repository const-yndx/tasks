package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type File interface {
	Name() string
	Path() string
	// TODO дописать свои методы
}

func main() {
	currentFile := loadFiles()

	for {
		fmt.Printf("[%s] %s $ ", time.Now().Format("02.01 15:04"), currentFile.Path())

		cmd, arg, err := readCommand(os.Stdin)
		if err != nil {
			fmt.Printf("Ошибка ввода: %v", err)
			continue
		}

		var result string
		result, currentFile = doCommand(cmd, arg, currentFile)
		fmt.Print(result)
	}
}

// loadFiles загружает дерево файлов из data.go
// Возвращает назад корневой файл rootFile — директорию /
func loadFiles() (rootFile File) {
	return nil // TODO реализовать загрузку дерева файлов
}

// doCommand выполняет команду cmd с аргументом arg (может быть пустым) над текущим файлом currentFile
// Возвращает назад сообщение о результате result (может быть пустым) и текущий после выполнения команды файл newCurrentFile
func doCommand(cmd, arg string, currentFile File) (result string, newCurrentFile File) {
	switch {
	case cmd == "cd" && arg == "":
		return "", currentFile
	case cmd == "cd" && arg == "..":
		return "", loadFiles()
	case cmd == "cd":
		return "", loadFiles()
	case (cmd == "exit" || cmd == "quit") && arg == "":
		os.Exit(0)
		return "", loadFiles()
	case cmd == "ls" && arg == "":
		return "", currentFile
	case cmd == "pwd" && arg == "":
		return "", currentFile
	default:
		return "Неизвестная команда: " + cmd, currentFile
	}
	// TODO реализовать выполнение команд
	/*| Команда | Аргумент       | Описание                                                                        |
	|---------|----------------|---------------------------------------------------------------------------------|
	| `cd`    | нет            | Переключение текущей директории. Без аргумента ничего не делает                 |
	| `cd`    | `..`           | Переключение текущей директории на родительскую                                 |
	| `cd`    | имя директории | Переключение текущей директории на вложенную, соответствующую переданному имени |
	| `exit`  | нет            | Выход из системы                                                                |
	| `ls`    | нет            | Вывод списка вложенных файлов и директорий                                      |
	| `pwd`   | нет            | Вывод полного пути текущей директории, начиная от корневой                      |
	| `quit`  | нет            | то же, что и `exit`
	*/
}

// readCommand читает пользовательский ввод через reader
// Возвращает назад команду cmd, ее аргумент arg (может быть пустым) и ошибку (не пустая, если сломался ввод)
func readCommand(reader io.Reader) (cmd string, arg string, err error) {
	value, err := bufio.NewReader(reader).ReadString('\n')
	if err != nil {
		return
	}

	value = value[:len(value)-1]

	parts := strings.Split(value, " ")
	switch len(parts) {
	case 0:
		return
	case 1:
		cmd = parts[0]
	default:
		cmd = parts[0]
		arg = value[len(parts[0])+1:]
	}

	return
}
