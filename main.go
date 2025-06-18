package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		ClearScreen()
		fmt.Println("\033[1;34m┌──────────────────────────────────────────────┐")
		fmt.Println("│              CRYPTO MINING SIMULATOR         │")
		fmt.Println("└──────────────────────────────────────────────┘\033[0m")

		fmt.Println("\033[1;36m[0]\033[0m Tambah Dummy Data")
		fmt.Println("\033[1;36m[1]\033[0m Tambah Aset Kripto")
		fmt.Println("\033[1;36m[2]\033[0m Edit Aset Kripto")
		fmt.Println("\033[1;36m[3]\033[0m Hapus Aset Kripto")
		fmt.Println("\033[1;36m[4]\033[0m Lihat Daftar Aset")
		fmt.Println("\033[1;36m[5]\033[0m Simulasi Mining")
		fmt.Println("\033[1;36m[6]\033[0m Pencarian Aset Kripto")
		fmt.Println("\033[1;36m[7]\033[0m Pengurutan Aset")
		fmt.Println("\033[1;36m[8]\033[0m Laporan Hasil Mining")
		fmt.Println("\033[1;31m[9]\033[0m Keluar")
		fmt.Print("\n\033[1;33mPilih menu [0-9]: \033[0m")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "0":
			TambahDummyData()
		case "1":
			TambahCrypto()
		case "2":
			EditCrypto()
		case "3":
			HapusCrypto()
		case "4":
			LihatDaftarCrypto()
		case "5":
			SimulasiMining()
		case "6":
			PencarianAset()
		case "7":
			PengurutanAset()
		case "8":
			TampilkanLaporanMining()
		case "9":
			fmt.Println("\n\033[1;32mTerima kasih! Sampai jumpa 👋\033[0m")
			return
		default:
			fmt.Println("\033[31mPilihan tidak valid!\033[0m")
		}

		Pause()
	}
}
