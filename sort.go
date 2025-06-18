package main

import (
	"fmt"
	"strings"
)

// Fungsi tukar elemen
func swap(i, j int) {
	temp := daftarCrypto[i]
	daftarCrypto[i] = daftarCrypto[j]
	daftarCrypto[j] = temp
}

// === SELECTION SORT ===
func SelectionSortByDifficulty(ascending bool) {
	for i := 0; i < jumlahCrypto-1; i++ {
		idx := i
		for j := i + 1; j < jumlahCrypto; j++ {
			if (ascending && daftarCrypto[j].Difficulty < daftarCrypto[idx].Difficulty) ||
				(!ascending && daftarCrypto[j].Difficulty > daftarCrypto[idx].Difficulty) {
				idx = j
			}
		}
		swap(i, idx)
	}
}

func SelectionSortByReward(ascending bool) {
	for i := 0; i < jumlahCrypto-1; i++ {
		idx := i
		for j := i + 1; j < jumlahCrypto; j++ {
			if (ascending && daftarCrypto[j].Reward < daftarCrypto[idx].Reward) ||
				(!ascending && daftarCrypto[j].Reward > daftarCrypto[idx].Reward) {
				idx = j
			}
		}
		swap(i, idx)
	}
}

// === INSERTION SORT ===
func InsertionSortByDifficulty(ascending bool) {
	for i := 1; i < jumlahCrypto; i++ {
		temp := daftarCrypto[i]
		j := i - 1
		for j >= 0 && ((ascending && daftarCrypto[j].Difficulty > temp.Difficulty) ||
			(!ascending && daftarCrypto[j].Difficulty < temp.Difficulty)) {
			daftarCrypto[j+1] = daftarCrypto[j]
			j--
		}
		daftarCrypto[j+1] = temp
	}
}

func InsertionSortByReward(ascending bool) {
	for i := 1; i < jumlahCrypto; i++ {
		temp := daftarCrypto[i]
		j := i - 1
		for j >= 0 && ((ascending && daftarCrypto[j].Reward > temp.Reward) ||
			(!ascending && daftarCrypto[j].Reward < temp.Reward)) {
			daftarCrypto[j+1] = daftarCrypto[j]
			j--
		}
		daftarCrypto[j+1] = temp
	}
}

func InsertionSortByName(ascending bool) {
	for i := 1; i < jumlahCrypto; i++ {
		temp := daftarCrypto[i]
		j := i - 1
		for j >= 0 && ((ascending && strings.ToLower(daftarCrypto[j].Nama) > strings.ToLower(temp.Nama)) ||
			(!ascending && strings.ToLower(daftarCrypto[j].Nama) < strings.ToLower(temp.Nama))) {
			daftarCrypto[j+1] = daftarCrypto[j]
			j--
		}
		daftarCrypto[j+1] = temp
	}
}

// === WRAPPER MENU SORT ===
func PengurutanAset() {
	fmt.Println("\n\033[1;34m📊 PENGURUTAN ASET KRIPTO\033[0m")
	fmt.Println("┌────────────────────────────────────────────┐")
	fmt.Println("│ 1. 📈 Difficulty (Selection Sort)          │")
	fmt.Println("│ 2. 💰 Reward (Selection Sort)              │")
	fmt.Println("│ 3. 📈 Difficulty (Insertion Sort)          │")
	fmt.Println("│ 4. 💰 Reward (Insertion Sort)              │")
	fmt.Println("│ 5. 🔤 Nama Aset (Insertion Sort)           │")
	fmt.Println("└────────────────────────────────────────────┘")

	metode := InputInt("🔧 Pilih metode sorting (1-5): ")

	fmt.Print("🔃 Urutan Ascending (true/false): ")
	ascInput, _ := inputReader.ReadString('\n')
	ascInput = strings.TrimSpace(strings.ToLower(ascInput))
	ascending := ascInput == "true"

	switch metode {
	case 1:
		SelectionSortByDifficulty(ascending)
	case 2:
		SelectionSortByReward(ascending)
	case 3:
		InsertionSortByDifficulty(ascending)
	case 4:
		InsertionSortByReward(ascending)
	case 5:
		InsertionSortByName(ascending)
	default:
		fmt.Println("\033[33m⚠️  Pilihan tidak valid.\033[0m")
		return
	}

	fmt.Println("\n\033[32m✅ Data berhasil diurutkan. Berikut hasilnya:\033[0m")
	LihatDaftarCrypto()
}
