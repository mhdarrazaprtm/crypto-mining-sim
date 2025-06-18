package main

import (
	"fmt"
	"strings"
)

type Crypto struct {
	Nama       string
	Algoritma  string
	Difficulty int
	Reward     float64
}

const MAX_CRYPTO = 100

var daftarCrypto [MAX_CRYPTO]Crypto
var jumlahCrypto int = 0

func TambahCrypto() {
	if jumlahCrypto >= MAX_CRYPTO {
		fmt.Println("\033[31m❌ Data aset sudah penuh.\033[0m")
		return
	}

	fmt.Println("\n\033[1;34m📥 Tambah Aset Kripto\033[0m")
	fmt.Print("🔹 Nama Aset Kripto: ")
	nama, _ := inputReader.ReadString('\n')
	nama = strings.TrimSpace(nama)

	fmt.Print("🔹 Algoritma (misal: SHA-256, Ethash): ")
	algo, _ := inputReader.ReadString('\n')
	algo = strings.TrimSpace(algo)

	difficulty := InputInt("🔹 Tingkat Kesulitan (angka): ")
	reward := InputFloat("🔹 Reward (angka desimal): ")

	daftarCrypto[jumlahCrypto] = Crypto{
		Nama:       nama,
		Algoritma:  algo,
		Difficulty: difficulty,
		Reward:     reward,
	}
	jumlahCrypto++

	fmt.Println("\033[32m✅ Aset kripto berhasil ditambahkan!\033[0m")
}

func EditCrypto() {
	fmt.Println("\n\033[1;34m✏️ Edit Aset Kripto\033[0m")
	fmt.Print("🔍 Masukkan nama aset yang ingin diubah: ")
	nama, _ := inputReader.ReadString('\n')
	nama = strings.TrimSpace(nama)

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

	fmt.Print("🔄 Algoritma baru: ")
	algo, _ := inputReader.ReadString('\n')
	algo = strings.TrimSpace(algo)

	difficulty := InputInt("🔄 Tingkat Kesulitan baru: ")
	reward := InputFloat("🔄 Reward baru: ")

	daftarCrypto[index].Algoritma = algo
	daftarCrypto[index].Difficulty = difficulty
	daftarCrypto[index].Reward = reward

	fmt.Println("\033[32m✅ Aset berhasil diperbarui!\033[0m")
}

func HapusCrypto() {
	fmt.Println("\n\033[1;34m🗑️ Hapus Aset Kripto\033[0m")
	fmt.Print("🔍 Masukkan nama aset yang ingin dihapus: ")
	nama, _ := inputReader.ReadString('\n')
	nama = strings.TrimSpace(nama)

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

	for i := index; i < jumlahCrypto-1; i++ {
		daftarCrypto[i] = daftarCrypto[i+1]
	}
	jumlahCrypto--

	fmt.Println("\033[32m✅ Aset berhasil dihapus!\033[0m")
}

func LihatDaftarCrypto() {
	if jumlahCrypto == 0 {
		fmt.Println("\033[33m📭 Belum ada aset kripto.\033[0m")
		return
	}

	fmt.Println("\n\033[1;35m📋 Daftar Aset Kripto\033[0m")
	fmt.Println("┌────┬────────────────────┬──────────────┬────────────┬──────────┐")
	fmt.Println("│ No │ Nama               │ Algoritma    │ Difficulty │ Reward   │")
	fmt.Println("├────┼────────────────────┼──────────────┼────────────┼──────────┤")

	for i := 0; i < jumlahCrypto; i++ {
		fmt.Printf("│ %2d │ %-18s │ %-12s │ %10d │ %8.2f │\n",
			i+1, daftarCrypto[i].Nama, daftarCrypto[i].Algoritma,
			daftarCrypto[i].Difficulty, daftarCrypto[i].Reward)
	}

	fmt.Println("└────┴────────────────────┴──────────────┴────────────┴──────────┘")
}

func TambahDummyData() {
	dummy := []Crypto{
		{"Bitcoin", "SHA-256", 9, 6.25},
		{"Ethereum", "Ethash", 8, 3.50},
		{"Litecoin", "Scrypt", 6, 12.5},
		{"Dogecoin", "Scrypt", 4, 25.0},
		{"Monero", "RandomX", 7, 1.26},
		{"Zcash", "Equihash", 5, 2.50},
		{"Ravencoin", "KawPow", 6, 5.00},
		{"Dash", "X11", 5, 2.00},
		{"Ethereum Classic", "Etchash", 7, 3.20},
		{"Beam", "BeamHash", 3, 8.00},
	}

	for i := 0; i < len(dummy) && jumlahCrypto < MAX_CRYPTO; i++ {
		daftarCrypto[jumlahCrypto] = dummy[i]
		jumlahCrypto++
	}

	fmt.Println("\033[32m✅ Dummy data aset kripto berhasil ditambahkan.\033[0m")
}
