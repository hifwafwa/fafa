package main

import (
	"fmt"
	"strings"
)

const MAX_MENU = 30

type MenuItem struct {
	nama      string
	kategori  string
	harga     int
	tersedia  bool
	komposisi string
}

type ArrMenu [MAX_MENU]MenuItem

func cetakGaris() {
	fmt.Println("====================================================================================================")
}

func cetakHeaderTabel() {
	cetakGaris()
	fmt.Printf("  %-4s %-20s %-15s %10s  %-8s  %s\n", "No", "Nama Menu", "Kategori", "Harga", "Status", "Komposisi")
	fmt.Println("  ---- -------------------- --------------- ----------  --------  ----------------------------------")
}

func tampilkanSemua(tab ArrMenu, n int) {
	if n == 0 {
		fmt.Println("  [Data menu masih kosong]")
		return
	}
	cetakHeaderTabel()
	for i := 0; i < n; i++ {
		status := "Tersedia"
		if !tab[i].tersedia {
			status = "Habis"
		}
		fmt.Printf("  %-4d %-20s %-15s Rp%8d  %-8s  %s\n",
			i+1, tab[i].nama, tab[i].kategori, tab[i].harga, status, tab[i].komposisi)
	}
	cetakGaris()
}

func tambahMenu(tab *ArrMenu, n *int) {
	if *n >= MAX_MENU {
		fmt.Println("  [!] Kapasitas menu penuh.")
		return
	}

	var nama, kategori, komposisi, statusInput string
	var harga int

	fmt.Println("\n--- TAMBAH DATA MENU ---")

	fmt.Print("  Nama Menu (Tanpa Spasi) : ")
	fmt.Scan(&nama)

	for i := 0; i < *n; i++ {
		if tab[i].nama == nama {
			fmt.Println("  [!] Nama menu sudah terdaftar.")
			return
		}
	}

	fmt.Print("  Kategori (coffee/non-coffee/makanan)   : ")
	fmt.Scan(&kategori)

	fmt.Print("  Harga (Angka)           : ")
	fmt.Scan(&harga)

	if harga <= 0 {
		fmt.Println("  [!] Harga harus lebih dari 0.")
		return
	}

	fmt.Print("  Komposisi (Tanpa Spasi) : ")
	fmt.Scan(&komposisi)

	var tersedia bool
	for {
		fmt.Print("  Status (ada/habis)      : ")
		fmt.Scan(&statusInput)

		if statusInput == "ada" {
			tersedia = true
			break
		} else if statusInput == "habis" {
			tersedia = false
			break
		} else {
			fmt.Println("  [!] Maaf, silahkan input sesuai ketentuan (ada/habis).")
		}
	}

	tab[*n] = MenuItem{
		nama:      nama,
		kategori:  kategori,
		harga:     harga,
		tersedia:  tersedia,
		komposisi: komposisi,
	}

	*n = *n + 1

	fmt.Println("  [+] Menu berhasil ditambahkan!")
}

func ubahMenu(tab *ArrMenu, n int) {
	tampilkanSemua(*tab, n)
	var no int
	fmt.Print("\n  Masukkan Nomor Menu yang ingin diubah: ")
	fmt.Scan(&no)

	if no < 1 || no > n {
		fmt.Println("  [!] Nomor tidak valid.")
		return
	}

	idx := no - 1
	fmt.Println("  -- Masukkan Data Baru --")
	fmt.Print("  Harga Baru              : ")
	fmt.Scan(&tab[idx].harga)

	var statusInput string
	for {
		fmt.Print("  Status Baru (ada/habis) : ")
		fmt.Scan(&statusInput)

		if statusInput == "ada" {
			tab[idx].tersedia = true
			break
		} else if statusInput == "habis" {
			tab[idx].tersedia = false
			break
		} else {
			fmt.Println("  [!] Maaf, silahkan input sesuai ketentuan (ada/habis).")
		}
	}

	fmt.Println("  [*] Data menu berhasil diubah!")
}

func hapusMenu(tab *ArrMenu, n *int) {
	tampilkanSemua(*tab, *n)
	var no int
	fmt.Print("\n  Masukkan Nomor Menu yang ingin dihapus: ")
	fmt.Scan(&no)

	if no < 1 || no > *n {
		fmt.Println("  [!] Nomor tidak valid.")
		return
	}

	idx := no - 1
	for i := idx; i < *n-1; i++ {
		tab[i] = tab[i+1]
	}
	*n = *n - 1
	fmt.Println("  [-] Data menu berhasil dihapus!")
}

func selectionSortHarga(tab ArrMenu, n int) ArrMenu {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if tab[j].harga < tab[minIdx].harga {
				minIdx = j
			}
		}
		tab[i], tab[minIdx] = tab[minIdx], tab[i]
	}
	return tab
}

func insertionSortHarga(tab ArrMenu, n int) ArrMenu {
	for i := 1; i < n; i++ {
		key := tab[i]
		j := i - 1
		for j >= 0 && tab[j].harga > key.harga {
			tab[j+1] = tab[j]
			j--
		}
		tab[j+1] = key
	}
	return tab
}

func sequentialSearchKategori(tab ArrMenu, n int, kat string) {
	fmt.Printf("\n--- HASIL PENCARIAN (Sequential): Kategori '%s' ---\n", kat)
	ada := false
	cetakHeaderTabel()
	for i := 0; i < n; i++ {
		if strings.ToLower(tab[i].kategori) == strings.ToLower(kat)  {
			ada = true
			status := "Tersedia"
			if !tab[i].tersedia {
				status = "Habis"
			}
			fmt.Printf("  %-4d %-20s %-15s Rp%8d  %-8s  %s\n",
				i+1, tab[i].nama, tab[i].kategori, tab[i].harga, status, tab[i].komposisi)
		}
	}
	if !ada {
		fmt.Println("  (Tidak ada menu ditemukan)")
	}
	cetakGaris()
}

func sortKategori(tab *ArrMenu, n int) {
	for i := 1; i < n; i++ {
		key := tab[i]
		j := i - 1
		for j >= 0 && strings.ToLower(tab[j].kategori) > strings.ToLower(key.kategori) {
			tab[j+1] = tab[j]
			j--
		}
		tab[j+1] = key
	}
}

func binarySearchKategori(tab ArrMenu, n int, kat string) {
	sortKategori(&tab, n)

	left := 0
	right := n - 1
	foundIdx := -1
	target := strings.ToLower(kat)

	for left <= right {
		mid := (left + right) / 2
		currentKat := strings.ToLower(tab[mid].kategori)

		if currentKat == target {
			foundIdx = mid
			break
		} else if currentKat < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Printf("\n--- HASIL PENCARIAN (Binary): Kategori '%s' ---\n", kat)
	if foundIdx != -1 {
		leftBound := foundIdx
		for leftBound > 0 && strings.ToLower(tab[leftBound-1].kategori) == target {
			leftBound--
		}
		rightBound := foundIdx
		for rightBound < n-1 && strings.ToLower(tab[rightBound+1].kategori) == target {
			rightBound++
		}

		cetakHeaderTabel()
		for i := leftBound; i <= rightBound; i++ {
			status := "Tersedia"
			if !tab[i].tersedia {
				status = "Habis"
			}
			fmt.Printf("  %-4d %-20s %-15s Rp%8d  %-8s  %s\n",
				i+1, tab[i].nama, tab[i].kategori, tab[i].harga, status, tab[i].komposisi)
		}
		cetakGaris()
	} else {
		fmt.Println("  (Tidak ada menu ditemukan)")
	}
}

func tampilStatistik(tab ArrMenu, n int) {
	if n == 0 {
		fmt.Println("  [Data menu masih kosong]")
		return
	}

	totalHarga := 0
	var katUnik [MAX_MENU]string
	var jumlahKat [MAX_MENU]int
	jumUnik := 0

	for i := 0; i < n; i++ {
		
		totalHarga += tab[i].harga
		ketemu := false
		for j := 0; j < jumUnik; j++ {
			if katUnik[j] == tab[i].kategori {
				jumlahKat[j]++
				ketemu = true
				break
			}
		}
		if !ketemu {
			katUnik[jumUnik] = tab[i].kategori
			jumlahKat[jumUnik] = 1
			jumUnik++
		}
	}

	rataRata := float64(totalHarga) / float64(n)

	fmt.Println("\n=== STATISTIK CAFE ===")
	fmt.Printf("  Rata-rata Harga Seluruh Menu : Rp%.2f\n", rataRata)
	fmt.Println("  Jumlah Menu per Kategori:")
	for i := 0; i < jumUnik; i++ {
		fmt.Printf("    - %-15s : %d item\n", katUnik[i], jumlahKat[i])
	}
	cetakGaris()
}

func main() {
	var menu ArrMenu
	n := 0

	menu[0] = MenuItem{"Espresso", "coffee", 15000, true, "Ekstrak_Kopi_Murni"}
	menu[1] = MenuItem{"Americano", "coffee", 18000, true, "Espresso_dan_Air"}
	menu[2] = MenuItem{"Cappuccino", "coffee", 22000, true, "Espresso_Susu_Foam"}
	menu[3] = MenuItem{"Cafe_Latte", "coffee", 22000, true, "Espresso_Susu_Steamed"}
	menu[4] = MenuItem{"Mochaccino", "coffee", 25000, true, "Espresso_Susu_Cokelat"}
	menu[5] = MenuItem{"Caramel_Macchiato", "coffee", 28000, true, "Espresso_Susu_Karamel"}
	menu[6] = MenuItem{"Cold_Brew", "coffee", 25000, true, "Kopi_Seduh_Dingin"}

	menu[7] = MenuItem{"Matcha_Latte", "non-coffee", 24000, true, "Bubuk_Matcha_Susu"}
	menu[8] = MenuItem{"Taro_Latte", "non-coffee", 24000, true, "Bubuk_Taro_Susu"}
	menu[9] = MenuItem{"Red_Velvet", "non-coffee", 24000, true, "Red_Velvet_Susu"}
	menu[10] = MenuItem{"Lemon_Tea", "non-coffee", 15000, true, "Teh_dan_Sirup_Lemon"}
	menu[11] = MenuItem{"Lychee_Tea", "non-coffee", 18000, true, "Teh_dan_Buah_Leci"}

	menu[12] = MenuItem{"Nasi_Goreng_Cafe", "makanan", 30000, true, "Nasi_Telur_Ayam"}
	menu[13] = MenuItem{"Mie_Goreng", "makanan", 28000, true, "Mie_Sosis_Telur"}
	menu[14] = MenuItem{"Kentang_Goreng", "makanan", 18000, true, "Kentang_Saus_Sambal"}
	menu[15] = MenuItem{"Snack_Platter", "makanan", 35000, true, "Sosis_Nugget_Kentang"}
	menu[16] = MenuItem{"Spaghetti_Meat", "makanan", 35000, false, "Pasta_Saus_Daging"}

	menu[17] = MenuItem{"Croissant", "dessert", 20000, true, "Pastry_Mentega"}
	menu[18] = MenuItem{"Waffle_Ice_Cream", "dessert", 25000, true, "Waffle_dan_Es_Krim"}
	menu[19] = MenuItem{"Cheesecake", "dessert", 30000, true, "Kue_Keju_Lumer"}

	n = 20

	pilihan := 0
	for pilihan != 9 {
		fmt.Println("\n===========================================")
		fmt.Println("    SISTEM DIGITAL CAFE MENU")
		fmt.Println("===========================================")
		fmt.Println(" -- ADMIN --")
		fmt.Println("  1. Tambah Data Menu")
		fmt.Println("  2. Ubah Data Menu")
		fmt.Println("  3. Hapus Data Menu")
		fmt.Println(" -- PELANGGAN / UMUM --")
		fmt.Println("  4. Tampilkan Seluruh Menu")
		fmt.Println("  5. Urutkan Menu (Termurah)")
		fmt.Println("  6. Cari Menu by Kategori (Sequential)")
		fmt.Println("  7. Cari Menu by Kategori (Binary)")
		fmt.Println("  8. Lihat Statistik Cafe")
		fmt.Println("  9. Keluar")
		fmt.Print("  Pilih menu (1-9): ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahMenu(&menu, &n)
		case 2:
			ubahMenu(&menu, n)
		case 3:
			hapusMenu(&menu, &n)
		case 4:
			tampilkanSemua(menu, n)
		case 5:
			var metode int
			fmt.Println("  Pilih metode urut: 1. Selection Sort  2. Insertion Sort")
			fmt.Print("  Pilih (1/2): ")
			fmt.Scan(&metode)
			var hasilUrut ArrMenu
			if metode == 1 {
				hasilUrut = selectionSortHarga(menu, n)
			} else {
				hasilUrut = insertionSortHarga(menu, n)
			}
			fmt.Println("\n--- MENU TERURUT (TERMURAH - TERMAHAL) ---")
			tampilkanSemua(hasilUrut, n)
		case 6:
			var kat string
			fmt.Print("  Masukkan Kategori dicari (misal: coffee): ")
			fmt.Scan(&kat)
			sequentialSearchKategori(menu, n, kat)
		case 7:
			var kat string
			fmt.Print("  Masukkan Kategori dicari (misal: coffee): ")
			fmt.Scan(&kat)
			binarySearchKategori(menu, n, kat)
		case 8:
			tampilStatistik(menu, n)
		case 9:
			fmt.Println("  Terima kasih telah menggunakan aplikasi Cafe-Menu!")
		default:
			fmt.Println("  [!] Pilihan tidak valid.")
		}
	}

}
