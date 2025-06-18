package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Reader global biar bisa dipakai semua
var inputReader = bufio.NewReader(os.Stdin)

// Fungsi untuk input integer dengan validasi
func InputInt(prompt string) int {
	for {
		fmt.Print(prompt)
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("⚠️  Input bukan angka. Coba lagi.")
		} else {
			return value
		}
	}
}

// Fungsi untuk input float dengan validasi
func InputFloat(prompt string) float64 {
	for {
		fmt.Print(prompt)
		input, _ := inputReader.ReadString('\n')
		input = strings.TrimSpace(input)
		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("⚠️  Input bukan angka desimal. Coba lagi.")
		} else {
			return value
		}
	}
}

// Fungsi pause (tekan Enter untuk lanjut)
func Pause() {
	fmt.Println("\nTekan Enter untuk kembali ke menu...")
	inputReader.ReadString('\n')
}

// Fungsi clear screen (cross-platform friendly, but not perfect)
func ClearScreen() {
	fmt.Print("\033[H\033[2J") // ANSI escape code
}
