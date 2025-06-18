package main

import (
	"fmt"
	"strings"
)

// === SEQUENTIAL SEARCH ===
func SequentialSearch() {
	if jumlahCrypto == 0 {
		fmt.Println("\033[33m📭 Belum ada data aset kripto.\033[0m")
		return
	}

	fmt.Println("\n\033[1;36m🔍 Pencarian Aset Kripto (Sequential Search)\033[0m")
	fmt.Print("🔎 Masukkan nama aset yang dicari: ")
	nama, _ := inputReader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	found := false
	for i := 0; i < jumlahCrypto; i++ {
		if strings.EqualFold(daftarCrypto[i].Nama, nama) {
			TampilkanCrypto(i)
			found = true
		}
	}

	if !found {
		fmt.Println("\033[31m❌ Aset tidak ditemukan.\033[0m")
	}
}

// === BINARY SEARCH ===
func BinarySearch() {
	if jumlahCrypto == 0 {
		fmt.Println("\033[33m📭 Belum ada data aset kripto.\033[0m")
		return
	}

	fmt.Println("\n\033[1;36m🔍 Pencarian Aset Kripto (Binary Search)\033[0m")
	fmt.Print("🔎 Masukkan nama aset yang dicari: ")
	nama, _ := inputReader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	kiri := 0
	kanan := jumlahCrypto - 1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		compare := strings.Compare(strings.ToLower(daftarCrypto[tengah].Nama), strings.ToLower(nama))

		if compare == 0 {
			TampilkanCrypto(tengah)
			return
		} else if compare < 0 {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	fmt.Println("\033[31m❌ Aset tidak ditemukan.\033[0m")
}

// === DISPLAY DETAIL ASET ===
func TampilkanCrypto(i int) {
	fmt.Println("\n\033[1;35m📄 Detail Aset Kripto\033[0m")
	fmt.Println("┌────────────────────────────────────────┐")
	fmt.Printf("│ Nama       : %-26s │\n", daftarCrypto[i].Nama)
	fmt.Printf("│ Algoritma  : %-26s │\n", daftarCrypto[i].Algoritma)
	fmt.Printf("│ Difficulty : %-26d │\n", daftarCrypto[i].Difficulty)
	fmt.Printf("│ Reward     : %-26.4f │\n", daftarCrypto[i].Reward)
	fmt.Println("└────────────────────────────────────────┘")
}

// === WRAPPER UNTUK MAIN MENU ===
func PencarianAset() {
	fmt.Println("\n\033[1;34m🔎 PENCARIAN ASET\033[0m")
	fmt.Println("1. 🔍 Sequential Search")
	fmt.Println("2. 🔎 Binary Search (pastikan sudah diurutkan ASCENDING!)")
	metode := InputInt("\nPilih metode pencarian: ")

	switch metode {
	case 1:
		SequentialSearch()
	case 2:
		BinarySearch()
	default:
		fmt.Println("\033[33m⚠️ Pilihan tidak valid.\033[0m")
	}
}
