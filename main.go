package main

//Configure how to use os import to clear terminal everytime user hit "back to main menu" option
import (
	//"os"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

// ----> Kamus Global <-----

type Produk struct {
	namaProduk string
	merek      string
	jenis      string
	harga      float64
	stok       int
}

type Transaksi struct {
	pembeli          string
	barangTerjual    Produk
	jumlahDibeli     int
	tanggalTransaksi time.Time
	subtotal         float64
}

const MAXPRODUCT int = 1024

type Data [MAXPRODUCT]Produk

type Buku [MAXPRODUCT]Transaksi

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
	// Lifecycle dari Main Feature
	for determinator != 3 {
		menuProcess()
		fmt.Println("Masukkan nomor menu:")
		fmt.Print(">>>> ")
		fmt.Scanln(&determinator)
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		if determinator == 1 {
			konfigurasiDataProduk(&dataProduk, &nData)
		} else if determinator == 2 {
			pencatatanTransaksi()
		}
	}
	// End Program
	menuEnd()
}

// ----> Prosedur Untuk Menu Utama <----
func konfigurasiDataProduk(data *Data, n *int) {
	var det int
	// menuHeaderKonfigurasiDataProduk()
	menuOptionsKonfigurasiDataProduk()
	fmt.Println("Masukkan menu: ")
	fmt.Print(">>>> ")
	fmt.Scanln(&det)
	for det != 3 {
		if det == 1 && *n < MAXPRODUCT {
			inputDataProduk(data, n)
		} else if det == 1 && *n >= MAXPRODUCT {
			fmt.Println("Memori telah habis!")
		} else if det == 2 {
			tampilSemuaDataProduk(data, n)
		}
		// menuHeaderKonfigurasiDataProduk()
		menuOptionsKonfigurasiDataProduk()
		fmt.Println("Masukkan menu: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&det)
	}

}

// TODO: Buat Prosedur untuk Fitur Pencatatan Transaksi beserta logic programnya
func pencatatanTransaksi() {

}

// Input data produk
func inputDataProduk(data *Data, n *int) {
	fmt.Println("---------------------------")
	fmt.Println("P E N G I S I A N   D A T A")
	fmt.Println("---------------------------")
	fmt.Println("Nama Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[*n].namaProduk)
	for data[*n].namaProduk == "" {
		fmt.Println("Nama produk kosong! Mohon masukkan nama produk: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&data[*n].namaProduk)
	}
	fmt.Println("---------------------------")
	fmt.Println("Merek Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[*n].merek)
	for data[*n].merek == "" {
		fmt.Println("Merek produk kosong! Mohon masukkan merek produk: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&data[*n].merek)
	}
	fmt.Println("---------------------------")
	fmt.Println("Jenis Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[*n].jenis)
	for data[*n].jenis == "" {
		fmt.Println("Jenis produk kosong! Mohon masukkan jenis produk: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&data[*n].jenis)
	}
	fmt.Println("---------------------------")
	fmt.Println("Harga Produk:")
	fmt.Print(">>>> Rp.")
	fmt.Scanln(&data[*n].harga)
	for data[*n].harga == 0 {
		fmt.Println("Harga produk kosong! Mohon masukkan harga produk: ")
		fmt.Print(">>>> Rp.")
		fmt.Scanln(&data[*n].harga)
	}
	fmt.Println("---------------------------")
	fmt.Println("Stok Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[*n].stok)
	for data[*n].stok == 0 {
		fmt.Println("Stok produk kosong! Mohon masukkan stok produk: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&data[*n].stok)
	}
	fmt.Println("---------------------------")
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	fmt.Println("Data Ke-", *n+1, "Berhasil Diinput")
	*n++
}

//TODO: Buat Prosedur dan Function untuk Sub-Menu dari Konfigurasi Data Produk

func tampilSemuaDataProduk(data *Data, n *int) {
	var det int
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	//menuHeaderKonfigurasiDataProduk()
	showAllProduct(*data, *n)
	for det != 5 {
		menuOptionsShowAllProduct()
		fmt.Println("Masukkan menu: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&det)
		if det == 1 {
			editProductData(data, *n)
		} else if det == 2 {
			deleteProductData(data, n)
		} else if det == 3 {
			showSearchedProduct(*data, *n)
		}
	}
}

func deleteProductData(data *Data, n *int) {
	var x string
	fmt.Println("Masukkan nama/jenis/merek produk yang ingin dihapus secara lengkap: ")
	fmt.Print(">>>> ")
	fmt.Scanln(&x)
	index := findSingleData(*data, *n, x)
	if index == -1 {
		fmt.Println("Data Tidak Ditemukan!")
	} else {
		i := index
		for i <= *n-2 {
			data[i] = data[i+1]
			i++
		}
		*n = *n - 1
		fmt.Println("Data Berhasil Dihapus!")
	}
}

// Tampilkan data
// TODO: need a better placement for the interface header
// TODO_2: is calling tampilSemuaDataProduk() is necessary?
func showAllProduct(data Data, n int) {
	fmt.Println()
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("No.", "Nama Produk", "Merek Produk", "Jenis Produk", "Harga Produk", "Stok Produk")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for i := 0; i < n; i++ {
		tbl.AddRow(i+1, data[i].namaProduk, data[i].merek, data[i].jenis, data[i].harga, data[i].stok)
	}
	tbl.Print()
	fmt.Println()
	//tampilSemuaDataProduk(data, n)
}

func showSearchedProduct(data Data, n int) {
	var x string
	fmt.Println("Masukkan nama/jenis/merek produk yang ingin dicari secara lengkap: ")
	fmt.Print(">>>> ")
	fmt.Scanln(&x)
	index := findSingleData(data, n, x)
	fmt.Println()

	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("No.", "Nama Produk", "Merek Produk", "Jenis Produk", "Harga Produk", "Stok Produk")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	tbl.AddRow(index+1, data[index].namaProduk, data[index].merek, data[index].jenis, data[index].harga, data[index].stok)

	tbl.Print()
	fmt.Println()
}

// Edit Data
/*
 * TODO: asumsi user mengira index 0 dari data adalah Data pertama
 * atau 1, akan akan aneh jika user mengedit data 1 dan yang teredit
 * menjadi data kedua.
 * TODO_2: asumsi user adalah tinkerer dan mencoba untuk input 0 untuk
 * select kolom data, kita buat user untuk input ulang dengan benar.
 * extra untuk mengecek apakah user meng-select kolom data lebih dari
 * MAXPRODUCT.
 * TODO_3: do we need to use by-pointers for the variables?
 */
func editProductData(data *Data, x int) {
	var n int
	var namaTemp, merekTemp, jenisTemp string
	var hargaTemp float64
	var stokTemp int
	fmt.Print("Masukkan kolom data yang akan diedit: ")
	fmt.Scanln(&n)
	namaTemp = data[n-1].namaProduk
	merekTemp = data[n-1].merek
	jenisTemp = data[n-1].jenis
	hargaTemp = data[n-1].harga
	stokTemp = data[n-1].stok
	for n == 0 && n > MAXPRODUCT {
		fmt.Println("Masukkan kolom data yang benar!")
		fmt.Print("Masukkan kolom data yang akan diedit: ")
		fmt.Scanln(&n)
		namaTemp = data[n-1].namaProduk
		merekTemp = data[n-1].merek
		jenisTemp = data[n-1].jenis
		hargaTemp = data[n-1].harga
		stokTemp = data[n-1].stok
	}
	fmt.Println("Nama Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[n-1].namaProduk)
	if data[n-1].namaProduk == "\n" {
		data[n-1].namaProduk = namaTemp
	}
	fmt.Println("---------------------------")
	fmt.Println("Merek Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[n-1].merek)
	if data[n-1].merek == "\n" {
		data[n-1].merek = merekTemp
		fmt.Println(merekTemp)
	}
	fmt.Println("---------------------------")
	fmt.Println("Jenis Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[n-1].jenis)
	if data[n-1].jenis == "\n" {
		data[n-1].jenis = jenisTemp
		fmt.Println(jenisTemp)
	}
	fmt.Println("---------------------------")
	fmt.Println("Harga Produk:")
	fmt.Print(">>>> Rp.")
	fmt.Scanln(&data[n-1].harga)
	if data[n-1].harga == 0 {
		data[n-1].harga = hargaTemp
		fmt.Println(hargaTemp)
	}
	fmt.Println("---------------------------")
	fmt.Println("Stok Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[n-1].stok)
	if data[n-1].stok == stokTemp {
		data[n-1].stok = stokTemp
	} else if data[n-1].stok == 0 {
		data[n-1].stok = 0
	}
	fmt.Println("---------------------------")
	fmt.Println("Data Ke-", n, "Berhasil Diedit")
	showAllProduct(*data, x)

}

// -----> Find Data Function using Sequential Search Algorithm <----

func findSingleData(data Data, n int, x string) int {
	var k int
	var idx int = -1
	for idx == -1 && k < n {
		if data[k].namaProduk == x || data[k].merek == x || data[k].jenis == x {
			idx = k
		}
		k++
	}
	return idx
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
}

func menuProcess() {
	fmt.Println("|-----------------------------------------------|")
	fmt.Println("|               M E N U   U T A M A             |")
	fmt.Println("|-----------------------------------------------|")
	fmt.Println("1. Konfigurasi Data Produk")
	fmt.Println("2. Pencatatan Transaksi")
	fmt.Println("3. Exit Program")
	fmt.Println("|-----------------------------------------------|")
}

func menuEnd() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("                Program Selesai                ")
	fmt.Println("-----------------------------------------------")
	fmt.Println("$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$")
}

func menuOptionsShowAllProduct() {
	menuHeaderKonfigurasiDataProduk()
	fmt.Println("1. Edit Data")
	fmt.Println("2. Hapus Data")
	fmt.Println("3. Cari Data")
	fmt.Println("4. Urutkan Data")
	fmt.Println("5. Kembali Ke Menu Utama")
}

func menuHeaderKonfigurasiDataProduk() {
	fmt.Println("|---------------------------------|")
	fmt.Println("| K O N F I G U R A S I - D A T A |")
	fmt.Println("|---------------------------------|")
}

func menuOptionsKonfigurasiDataProduk() {
	fmt.Println("|---------------------------------|")
	fmt.Println("|         C O O K - D A T A       |")
	fmt.Println("|---------------------------------|")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Tampilkan Semua Data")
	fmt.Println("3. Kembali ke Menu Utama")
}

// ---------------------------------------------------------------
