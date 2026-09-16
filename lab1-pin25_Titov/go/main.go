package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	colorRed     = "\033[91m"
	colorGreen   = "\033[92m"
	colorYellow  = "\033[93m"
	colorBlue    = "\033[94m"
	colorMagenta = "\033[95m"
	colorCyan    = "\033[96m"
	colorReset   = "\033[0m"
)

// greet формирует приветствие с текущей датой и временем (UTC).
func greet(name string) string {
	now := time.Now().UTC().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("Hello, %s! Current time (UTC): %s", name, now)
}

// getColor возвращает ANSI-цвет в зависимости от первой буквы имени.
func getColor(name string) string {
	if name == "" {
		return colorReset
	}
	letter := strings.ToLower(string(name[0]))
	switch {
	case letter <= "e":
		return colorRed
	case letter <= "j":
		return colorGreen
	case letter <= "o":
		return colorYellow
	case letter <= "t":
		return colorBlue
	case letter <= "y":
		return colorMagenta
	default:
		return colorCyan
	}
}

func main() {
	var name string

	if len(os.Args) > 1 {
		name = strings.Join(os.Args[1:], " ")
	} else {
		fmt.Print("Enter your name: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		name = strings.TrimSpace(input)
	}

	if name == "" {
		name = "World"
	}

	start := time.Now()
	message := greet(name)
	color := getColor(name)
	elapsed := time.Since(start)

	fmt.Printf("%s%s%s\n", color, message, colorReset)
	fmt.Printf("[Go] Execution time: %.4f ms\n", float64(elapsed.Nanoseconds())/1e6)
}
