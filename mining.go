package main

import (
	"fmt"
	"strings"
)

type MiningResult struct {
	NamaCrypto    string
	PowerUsed     float64
	EstimasiWaktu float64
	HasilReward   float64
}

const MAX_MINING = 100

var hasilMining [MAX_MINING]MiningResult
var jumlahMining int = 0

func SimulasiMining() {
	if jumlahCrypto == 0 {
		fmt.Println("\033[33m⚠️ Belum ada aset kripto yang tersedia.\033[0m")
		return
	}

	fmt.Println("\n\033[1;34m🔧 SIMULASI MINING\033[0m")
	fmt.Print("💡 Masukkan nama aset kripto yang ingin ditambang: ")
	nama, _ := inputReader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	// Cari index
	index := -1
	for i := 0; i < jumlahCrypto; i++ {
		if strings.EqualFold(daftarCrypto[i].Nama, nama) {
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("\033[31m❌ Aset tidak ditemukan.\033[0m")
		return
	}

	power := InputFloat("⚙️  Masukkan daya komputasi (power): ")
	if power <= 0 {
		fmt.Println("\033[33m⚠️  Daya harus lebih dari 0.\033[0m")
		return
	}

	crypto := daftarCrypto[index]

	// Rumus mining
	estimasiWaktu := float64(crypto.Difficulty) * 100 / power
	hasilReward := crypto.Reward * (power / (float64(crypto.Difficulty) * 10))

	if jumlahMining >= MAX_MINING {
		fmt.Println("\033[31m❌ Slot hasil mining penuh.\033[0m")
		return
	}

	hasilMining[jumlahMining] = MiningResult{
		NamaCrypto:    crypto.Nama,
		PowerUsed:     power,
		EstimasiWaktu: estimasiWaktu,
		HasilReward:   hasilReward,
	}
	jumlahMining++

	// Tampilkan hasil simulasi
	fmt.Println("\n\033[1;32m🟢 HASIL SIMULASI MINING\033[0m")
	fmt.Println("┌──────────────────────────────────────────────┐")
	fmt.Printf("│ Nama Aset      : %-28s │\n", crypto.Nama)
	fmt.Printf("│ Algoritma      : %-28s │\n", crypto.Algoritma)
	fmt.Printf("│ Difficulty     : %-28d │\n", crypto.Difficulty)
	fmt.Printf("│ Power Digunakan: %-28.2f │\n", power)
	fmt.Printf("│ Estimasi Waktu : %-28.2f │\n", estimasiWaktu)
	fmt.Printf("│ Reward Diterima: %-28.4f │\n", hasilReward)
	fmt.Println("└──────────────────────────────────────────────┘")
}
