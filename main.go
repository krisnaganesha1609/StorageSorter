package main

import "fmt"

// ----> Kamus Global <-----

type Produk struct {
	namaProduk string
	merek      string
	jenis      string
	harga      float64
	stok       int
}

const MAXPRODUCT int = 1024

type Data [MAXPRODUCT]Produk

//---------------------------------------------------------------

// ----> Main Program <----
func main() {
	// Variabel untuk menentukan pilihan menu dan mencegah infinite-loop
	var determinator int

	// Variabel yang akan dijadikan penyimpanan data produk
	var dataProduk Data

	// Variabel yang akan digunakan untuk membatasi iterasi ketika membaca data dan menampilkan data
	var nData int

	// Start Menu
	menuStart()
	// Lifecycle Main Feature
	for determinator != 7 {
		menuProcess()
		fmt.Println("Masukkan nomor menu:")
		fmt.Print(">>>>")
		fmt.Scan(&determinator)
		if determinator == 1 {
			isiDataProduk(&dataProduk, &nData)
		} else if determinator == 2 {
			tampilSemuaDataProduk(dataProduk, nData)
		}
	}
	menuEnd()
}

//TODO: Rencanakan Tipe Data, Fungsi dan Prosedur Yang Akan Dipakai

func isiDataProduk(data *Data, n *int) {
	var det int
	fmt.Println("Instruksi")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Kembali ke Menu Utama")
	fmt.Println("Please enter menu: ")
	fmt.Scan(&det)
	for det != 2 {
		if det == 1 && *n < MAXPRODUCT {
			fmt.Println("---------------------------")
			fmt.Println("P E N G I S I A N   D A T A")
			fmt.Println("---------------------------")
			fmt.Println("Nama Produk: ")
			fmt.Print(">>>>")
			fmt.Scan(&data[*n].namaProduk)
			for data[*n].namaProduk == "" {
				fmt.Print("Nama produk kosong! Mohon masukkan nama produk: ")
				fmt.Print(">>>>")
				fmt.Scan(&data[*n].namaProduk)
			}
			fmt.Println("---------------------------")
			fmt.Println("Merek Produk: ")
			fmt.Print(">>>>")
			fmt.Scan(&data[*n].merek)
			for data[*n].merek == "" {
				fmt.Print("Merek produk kosong! Mohon masukkan merek produk: ")
				fmt.Print(">>>>")
				fmt.Scan(&data[*n].merek)
			}
			fmt.Println("---------------------------")
			fmt.Println("Jenis Produk: ")
			fmt.Print(">>>>")
			fmt.Scan(&data[*n].jenis)
			for data[*n].jenis == "" {
				fmt.Print("Jenis produk kosong! Mohon masukkan jenis produk: ")
				fmt.Print(">>>>")
				fmt.Scan(&data[*n].jenis)
			}
			fmt.Println("---------------------------")
			fmt.Println("Harga Produk: ")
			fmt.Print(">>>>")
			fmt.Scan(&data[*n].harga)
			for data[*n].harga == 0 {
				fmt.Print("Harga produk kosong! Mohon masukkan harga produk: ")
				fmt.Print(">>>> Rp.")
				fmt.Scan(&data[*n].harga)
			}
			fmt.Println("---------------------------")
			fmt.Println("Stok Produk: ")
			fmt.Print(">>>>")
			fmt.Scan(&data[*n].stok)
			for data[*n].stok == -1 {
				fmt.Print("Stok produk kosong! Mohon masukkan stok produk: ")
				fmt.Print(">>>>")
				fmt.Scan(&data[*n].stok)
			}
			fmt.Println("---------------------------")
			fmt.Println("Data Ke-", *n+1, "Berhasil Diinput")
			*n++
		} else if det == 1 && *n >= MAXPRODUCT {
			fmt.Println("Memori telah habis!")
		}
		fmt.Println("Pilihan:")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Kembali ke Menu Utama")
		fmt.Println("Please enter menu: ")
		fmt.Scan(&det)
	}

}

func tampilSemuaDataProduk(data Data, n int) {
	fmt.Println("|------------------------------------------------------------------------------|")
	fmt.Println("|                            T A B E L   D A T A                               |")
	fmt.Println("|------------------------------------------------------------------------------|")
	fmt.Println("| No. | Nama Produk | Merek Produk | Jenis Produk | Harga Produk | Stok Produk |")
	for i := 0; i < n; i++ {
		fmt.Printf("| %d | %s | %s | %s | %.2f | %d |\n", i+1, data[i].namaProduk, data[i].merek, data[i].jenis, data[i].harga, data[i].stok)
	}
	fmt.Println("|------------------------------------------------------------------------------|")
}

// -----> Menampilkan menu secara estetik pada CLI <----------
func menuStart() {
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("-----------------------------------------------")
	fmt.Println("                 STORAGE SORTER                ")
	fmt.Println("-----------------------------------------------")
	fmt.Println(" Aplikasi Manajemen Inventaris Toko Elektronik ")
	fmt.Println("-----------------------------------------------")
	fmt.Println("-----------------------------------------------")
	fmt.Println("                 Developed By:                 ")
	fmt.Println("                  Kelompok 1                   ")
	fmt.Println("-----------------------------------------------")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
}

func menuProcess() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("               M E N U   F I T U R             ")
	fmt.Println("-----------------------------------------------")
	fmt.Println("1. Isi Data Produk")
	fmt.Println("2. Tampilkan Semua Data Produk")
	fmt.Println("3. ....")
	fmt.Println("3. ....")
	fmt.Println("4. ....")
	fmt.Println("5. ....")
	fmt.Println("6. ....")
	fmt.Println("7. Exit Program")
	fmt.Println("-----------------------------------------------")
}

func menuEnd() {
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
	fmt.Println("-----------------------------------------------")
	fmt.Println("                Program Selesai                ")
	fmt.Println("-----------------------------------------------")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
}

// ---------------------------------------------------------------
