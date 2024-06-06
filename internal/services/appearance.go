package services

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
)

func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[H\033[2J")
	}
}

func PrintColorText(colorText color.Attribute, text string) {
	color.Set(colorText)

	fmt.Print(text)

	color.Unset()
}

func ReadInput() string {
	var x string
	reader := bufio.NewReader(os.Stdin)
	fmt.Scan(&x)
	text, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Ошибка ввода: ", err)
    }
	result := strings.ReplaceAll(x + " " + text, "\r\n", "")
	return result
}


