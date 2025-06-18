package main

import (
	"fmt"
)

func TampilkanLaporanMining() {
	if jumlahMining == 0 {
		fmt.Println("\033[33m📭 Belum ada aktivitas mining yang dilakukan.\033[0m")
		return
	}

	fmt.Println("\n\033[1;35m📊 LAPORAN HASIL MINING\033[0m")
	fmt.Println("┌────────────────────────────────────────────────────────────────────────────┐")
	fmt.Printf("│ %-4s │ %-20s │ %-8s │ %-10s │ %-12s │\n", "No", "Nama Aset", "Power", "Waktu", "Reward")
	fmt.Println("├─────┼──────────────────────┼──────────┼────────────┼──────────────┤")

	totalReward := 0.0

	for i := 0; i < jumlahMining; i++ {
		fmt.Printf("│ %3d │ %-20s │ %8.2f │ %10.2f │ %12.4f │\n",
			i+1,
			hasilMining[i].NamaCrypto,
			hasilMining[i].PowerUsed,
			hasilMining[i].EstimasiWaktu,
			hasilMining[i].HasilReward,
		)
		totalReward += hasilMining[i].HasilReward
	}

	fmt.Println("└─────┴──────────────────────┴──────────┴────────────┴──────────────┘")

	avgReward := totalReward / float64(jumlahMining)

	fmt.Println("\n\033[1;36m📈 Rangkuman:\033[0m")
	fmt.Printf("🧮 Total Aktivitas Mining     : %d\n", jumlahMining)
	fmt.Printf("💰 Total Reward Terkumpul     : %.4f koin\n", totalReward)
	fmt.Printf("📊 Rata-rata Reward/Aktivitas : %.4f koin\n", avgReward)
}
