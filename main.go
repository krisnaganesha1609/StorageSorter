package main

import (
	"fmt"
	// time module/package untuk mendapatkan tipe data time.Time di struct 'Transaksi'
	"time"
	// strings module/package untuk menambah akurasi ketika melakukan pencarian data
	"strings"
	// bufio dan os module/package untuk menginput string yang mengandung spasi ke dalam variabel namaProduk
	"bufio"
	"os"

	// third-party package untuk membuat tabel untuk CLI yang memudahkan pengguna untuk membaca data yang ditampilkan
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

/*
Variabel Global
*/
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

const MAXPRODUCT int = 25
const MAXTRANSACTION int = 1000

type Data [MAXPRODUCT]Produk
type CatatanTransaksi [MAXTRANSACTION]Transaksi

/*
Program Utama
*/
func main() {
	// Variabel untuk menentukan pilihan menu dan mencegah infinite-loop
	var determinator int

	// Variabel yang akan dijadikan penyimpanan data produk
	var dataProduk Data

	// Variabel yang akan digunakan untuk membatasi iterasi ketika membaca data dan menampilkan data
	var nData int

	// Variabel yang akan dijadikan penyimpanan data transaksi
	var dataTransaksi CatatanTransaksi

	// Varabel yang akan digunakan untuk membatasi iterasi ketika membaca data transaksi dan menampilkan data transaksi
	var nTransaksi int

	// Variabel untuk mendeteksi apakah data sudah terurut atau belum
	var isSorted string

	// Start Menu
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	menuStart()
	time.Sleep(3 * time.Second)
	dataProduk[0].namaProduk = "iPhone 14"
	dataProduk[0].merek = "Smartphone"
	dataProduk[0].jenis = "Apple"
	dataProduk[0].harga = 13999.99
	dataProduk[0].stok = 50
	dataProduk[1].namaProduk = "Galaxy S22"
	dataProduk[1].merek = "Samsung"
	dataProduk[1].jenis = "Smartphone"
	dataProduk[1].harga = 12999.99
	dataProduk[1].stok = 70
	dataProduk[2].namaProduk = "XPS 13"
	dataProduk[2].merek = "Dell"
	dataProduk[2].jenis = "Laptop"
	dataProduk[2].harga = 15999.99
	dataProduk[2].stok = 30
	dataProduk[3].namaProduk = "Thinkpad X1 Carbon"
	dataProduk[3].merek = "Lenovo"
	dataProduk[3].jenis = "Laptop"
	dataProduk[3].harga = 17999.99
	dataProduk[3].stok = 20
	dataProduk[4].namaProduk = "WH-100XM4"
	dataProduk[4].merek = "Sony"
	dataProduk[4].jenis = "Headphone"
	dataProduk[4].harga = 3999.99
	dataProduk[4].stok = 100
	dataProduk[5].namaProduk = "AirPods Pro"
	dataProduk[5].merek = "Apple"
	dataProduk[5].jenis = "Earbud"
	dataProduk[5].harga = 2499.99
	dataProduk[5].stok = 200
	dataProduk[6].namaProduk = "OLED55CXPUA"
	dataProduk[6].merek = "LG"
	dataProduk[6].jenis = "TV"
	dataProduk[6].harga = 19999.99
	dataProduk[6].stok = 15
	dataProduk[7].namaProduk = "PS5"
	dataProduk[7].merek = "Sony"
	dataProduk[7].jenis = "Konsol Game"
	dataProduk[7].harga = 7999.99
	dataProduk[7].stok = 40
	dataProduk[8].namaProduk = "MacBook Pro 14''"
	dataProduk[8].merek = "Apple"
	dataProduk[8].jenis = "Laptop"
	dataProduk[8].harga = 24999.99
	dataProduk[8].stok = 25
	dataProduk[9].namaProduk = "Surface Pro 8"
	dataProduk[9].merek = "Microsoft"
	dataProduk[9].jenis = "Tablet"
	dataProduk[9].harga = 13999.99
	dataProduk[9].stok = 35
	nData = 10
	fmt.Print("\033[2J")
	fmt.Print("\033[H")

	// Lifecycle dari Main Feature
	for determinator != 3 {
		menuProcess()
		menuOptionsProcess()
		fmt.Println("Masukkan nomor menu:")
		fmt.Print(">>>> ")
		fmt.Scanln(&determinator)
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		if determinator == 1 {
			// Jalankan konfigurasi data produk
			konfigurasiDataProduk(&dataProduk, &nData, &isSorted)
		} else if determinator == 2 {
			// Jalankan pencatatan transaksi
			pencatatanTransaksi(&dataTransaksi, &nTransaksi, &dataProduk, nData)
		}
	}
	// Akhiri program
	menuEnd()
	time.Sleep(3 * time.Second)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

/*
Prosedur untuk Menu pertama, yaitu
Konfigurasi Data Produk.
*/
func konfigurasiDataProduk(data *Data, n *int, isSorted *string) {
	var det int
	menuHeaderKonfigurasiDataProduk()
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
			tampilSemuaDataProduk(data, n, isSorted)
		}
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		menuHeaderKonfigurasiDataProduk()
		menuOptionsKonfigurasiDataProduk()
		fmt.Println("Masukkan menu: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&det)
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

/*
Prosedur untuk Menu kedua, yaitu
Pencatatan transaksi.
*/
func pencatatanTransaksi(transaksi *CatatanTransaksi, nT *int, produk *Data, nD int) {
	var det int
	showAllTransaction(*transaksi, *nT)
	menuHeaderPencatatanTransaksi()
	menuOptionsPencatatanTransaksi()
	fmt.Println("Masukkan menu: ")
	fmt.Print(">>>> ")
	fmt.Scanln(&det)
	for det != 2 {
		if det == 1 {
			inputDataTransaksi(transaksi, nT, produk, nD)
			fmt.Print("\033[2J")
			fmt.Print("\033[H")
		}
		showAllTransaction(*transaksi, *nT)
		menuHeaderPencatatanTransaksi()
		menuOptionsPencatatanTransaksi()
		fmt.Println("Masukkan menu: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&det)
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

/*
Prosedur untuk menu pertama di dalam
Pencatatan transaksi, yaitu
Input Data Transaksi.
*/
func inputDataTransaksi(transaksi *CatatanTransaksi, nT *int, produk *Data, nD int) {
	var beli int
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	menuHeaderInputDataProduk()
	fmt.Println("Nama Pembeli:")
	fmt.Print(">>>> ")
	fmt.Scanln(&transaksi[*nT].pembeli)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	for transaksi[*nT].pembeli == "" {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		menuHeaderInputDataProduk()
		fmt.Println("Nama pembeli kosong! Mohon masukkan nama pembeli: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&transaksi[*nT].pembeli)
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showAllProduct(*produk, nD)
	menuHeaderInputDataProduk()
	fmt.Println("Masukkan nomor produk yang dibeli:")
	fmt.Print(">>>> ")
	fmt.Scanln(&beli)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	for beli <= 0 {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		showAllProduct(*produk, nD)
		menuHeaderInputDataProduk()
		fmt.Println("Produk yang dibeli tidak ditemukan/tidak boleh kosong! Mohon masukkan nomor produk yang dibeli: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&beli)
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}
	transaksi[*nT].barangTerjual.namaProduk = produk[beli-1].namaProduk
	transaksi[*nT].barangTerjual.merek = produk[beli-1].merek
	transaksi[*nT].barangTerjual.jenis = produk[beli-1].jenis
	transaksi[*nT].barangTerjual.harga = produk[beli-1].harga
	transaksi[*nT].barangTerjual.stok = produk[beli-1].stok
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	menuHeaderInputDataProduk()
	fmt.Println("Jumlah Pembelian:")
	fmt.Print(">>>> ")
	fmt.Scanln(&transaksi[*nT].jumlahDibeli)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	for transaksi[*nT].jumlahDibeli == 0 {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		menuHeaderInputDataProduk()
		fmt.Println("Jumlah pembelian kosong! Mohon masukkan jumlah pembelian: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&transaksi[*nT].jumlahDibeli)
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}
	transaksi[*nT].tanggalTransaksi = time.Now()
	transaksi[*nT].subtotal = float64(transaksi[*nT].jumlahDibeli) * transaksi[*nT].barangTerjual.harga
	produk[beli-1].stok = produk[beli-1].stok - transaksi[*nT].jumlahDibeli
	transaksi[*nT].barangTerjual.stok = produk[beli-1].stok
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	menuHeaderInputDataProduk()
	fmt.Println("Transaksi Ke -", *nT+1, "Berhasil Diinput.")
	fmt.Printf("Stok produk berkurang sebanyak %d buah\n", transaksi[*nT].jumlahDibeli)
	*nT++
	time.Sleep(3 * time.Second)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

/*
Prosedur untuk menu pertama di dalam
Konfigurasi Data Produk, yaitu
Tambah Data.
*/
func inputDataProduk(data *Data, n *int) {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	menuHeaderInputDataProduk()
	fmt.Println("Nama Produk:")
	fmt.Print(">>>> ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		data[*n].namaProduk = scanner.Text()
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	for data[*n].namaProduk == "" {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		menuHeaderInputDataProduk()
		fmt.Println("Nama produk kosong! Mohon masukkan nama produk: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&data[*n].namaProduk)
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	menuHeaderInputDataProduk()
	fmt.Println("Merek Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[*n].merek)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	for data[*n].merek == "" {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		menuHeaderInputDataProduk()
		fmt.Println("Merek produk kosong! Mohon masukkan merek produk: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&data[*n].merek)
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	menuHeaderInputDataProduk()
	fmt.Println("Jenis Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[*n].jenis)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	for data[*n].jenis == "" {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		menuHeaderInputDataProduk()
		fmt.Println("Jenis produk kosong! Mohon masukkan jenis produk: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&data[*n].jenis)
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	menuHeaderInputDataProduk()
	fmt.Println("Harga Produk:")
	fmt.Print(">>>> Rp.")
	fmt.Scanln(&data[*n].harga)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	for data[*n].harga == 0 {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		menuHeaderInputDataProduk()
		fmt.Println("Harga produk kosong! Mohon masukkan harga produk: ")
		fmt.Print(">>>> Rp. ")
		fmt.Scanln(&data[*n].harga)
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	menuHeaderInputDataProduk()
	fmt.Println("Stok Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[*n].stok)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	for data[*n].stok == 0 {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		menuHeaderInputDataProduk()
		fmt.Println("Stok produk kosong! Mohon masukkan stok produk: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&data[*n].stok)
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	menuHeaderInputDataProduk()
	fmt.Println("Data Ke -", *n+1, "Berhasil Diinput.")
	*n++
	time.Sleep(3 * time.Second)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

/*
Prosedur untuk menu kedua di dalam
Konfigurasi Data Produk, yaitu
Utilitas Data.

Sebelumnya kami menggunakan menu kedua
ini untuk menampilkan semua data
yang tersedia, tetapi semenjak penamaan
nya tidak terlalu sesuai dengan aslinya,
kami ubah yang sebelumnya "Tampilkan Semua
Data Produk" menjadi "Utilitas Data".
*/
func tampilSemuaDataProduk(data *Data, n *int, isSorted *string) {
	var det int
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showAllProduct(*data, *n)
	for det != 5 {
		menuHeaderTampilSemuaDataProduk()
		menuOptionsTampilSemuaDataProduk()
		fmt.Println("Masukkan menu: ")
		fmt.Print(">>>> ")
		fmt.Scanln(&det)
		if det == 1 {
			editProductData(data, *n)
		} else if det == 2 {
			deleteProductData(data, n)
		} else if det == 3 {
			if *isSorted != "" {
				var det2 int
				fmt.Println("1. Cari Berdasarkan Harga")
				fmt.Println("2. Cari Berdasarkan Stok")
				fmt.Println("3. Cari Berdasarkan Nama Produk/Merk Produk/Jenis Produk")
				fmt.Println("Masukkan menu: ")
				fmt.Print(">>>> ")
				fmt.Scanln(&det2)
				if det2 == 1 {
					showHargaSearchedData(*data, *n, *isSorted)
				} else if det2 == 2 {
					showStokSearchedData(*data, *n, *isSorted)
				} else if det2 == 3 {
					showSearchedProduct(*data, *n)
				}
			} else {
				showSearchedProduct(*data, *n)
			}
		} else if det == 4 {
			SortProduct(data, *n, isSorted)
		}
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

/*
Prosedur untuk menu kedua di dalam
Utilitas Data, yaitu Penghapusan Data.
*/
func deleteProductData(data *Data, n *int) {
	var x int
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showAllProduct(*data, *n)
	menuHeaderDeleteData()
	fmt.Println("Masukkan nomor produk yang ingin dihapus: ")
	fmt.Print(">>>> ")
	fmt.Scanln(&x)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	if x <= 0 {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		showAllProduct(*data, *n)
		menuHeaderDeleteData()
		fmt.Println("Data Tidak Ditemukan!")
		time.Sleep(3 * time.Second)
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	} else {
		i := x - 1
		for i <= *n-2 {
			data[i] = data[i+1]
			i++
		}
		*n = *n - 1
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		showAllProduct(*data, *n)
		menuHeaderDeleteData()
		fmt.Println("Data Berhasil Dihapus!")
		time.Sleep(3 * time.Second)
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}
	showAllProduct(*data, *n)
}

/*
Prosedur untuk menu kedua di dalam
Utilitas Data, yaitu menampilkan semua
tabel data yang sudah masuk di atas header.
Dipakai saat menampilkan menu:
- Edit data
- Hapus data
- Cari data
- Urutkan data
- Tambah transaksi
*/
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
}

/*
Prosedur di dalam Pencatatan Transaksi,
yaitu menampilkan semua transaksi data yang
sudah masuk di atas header.
Dipakai saat menampilkan menu:
- Pencatatan Transaksi
*/
func showAllTransaction(transaksi CatatanTransaksi, n int) {
	fmt.Println()
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("No.", "Tgl Transaksi", "Pembeli", "Nama Produk", "Jenis Produk", "Merek Produk", "Jumlah", "Subtotal")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for i := 0; i < n; i++ {
		tbl.AddRow(i+1, transaksi[i].tanggalTransaksi.Format(time.UnixDate), transaksi[i].pembeli, transaksi[i].barangTerjual.namaProduk, transaksi[i].barangTerjual.jenis, transaksi[i].barangTerjual.merek, transaksi[i].jumlahDibeli, transaksi[i].subtotal)
	}
	tbl.Print()
	fmt.Println()
}

/*
Prosedur untuk menu ketiga di dalam
Utilitas Data, yaitu menampilkan transaksi data
yang dicari berdasarkan kata kunci yang di masukkan.
Dipakai saat menampilkan menu:
- Cari Data
*/
//Sequential Search
func showSearchedProduct(data Data, n int) {
	defer timeElapsed(time.Now(), "Sequential")
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showAllProduct(data, n)
	menuHeaderShowSearchedProduct()
	var x string
	fmt.Println("Masukkan nama/jenis/merek produk yang ingin dicari secara lengkap: ")
	fmt.Print(">>>> ")
	fmt.Scanln(&x)
	fmt.Println()
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("No.", "Nama Produk", "Merek Produk", "Jenis Produk", "Harga Produk", "Stok Produk")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	var k int
	var index int = -1
	for k < n {
		if strings.HasPrefix(strings.ToLower(data[k].namaProduk), strings.ToLower(x)) || strings.HasPrefix(strings.ToLower(data[k].merek), strings.ToLower(x)) || strings.HasPrefix(strings.ToLower(data[k].jenis), strings.ToLower(x)) {
			index = k
			tbl.AddRow(index+1, data[index].namaProduk, data[index].merek, data[index].jenis, data[index].harga, data[index].stok)
		}
		k++
	}
	if index != -1 {
		fmt.Println()
		fmt.Println("Data ditemukan!")
		fmt.Println()
		tbl.Print()
	} else {
		fmt.Println()
		fmt.Println("Data tidak ditemukan!")
		showAllProduct(data, n)
	}
	fmt.Println()
}

// Binary Search
func showHargaSearchedData(data Data, n int, isSorted string) {
	defer timeElapsed(time.Now(), "Binary")
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showAllProduct(data, n)
	menuHeaderShowSearchedProduct()
	var x int
	fmt.Println("Masukkan harga produk yang ingin dicari (syarat: data harus TERURUT): ")
	fmt.Print(">>>> ")
	fmt.Scanln(&x)
	fmt.Println()

	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("No.", "Nama Produk", "Merek Produk", "Jenis Produk", "Harga Produk", "Stok Produk")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	var mid, left, right, i int
	i = 1
	left = 0
	right = n - 1
	if isSorted == "asc" {
		for left <= right {
			mid = (left + right) / 2
			if int(data[mid].harga) < x {
				left = mid + 1
			} else if int(data[mid].harga) > x {
				right = mid
			} else {
				tbl.AddRow(i, data[mid].namaProduk, data[mid].merek, data[mid].jenis, data[mid].harga, data[mid].stok)
				i++
			}
		}
	} else if isSorted == "desc" {
		for left <= right {
			mid = (left + right) / 2
			if x > int(data[mid].harga) {
				right = mid - 1
			} else if x < int(data[mid].harga) {
				left = mid + 1
			} else {
				tbl.AddRow(i, data[mid].namaProduk, data[mid].merek, data[mid].jenis, data[mid].harga, data[mid].stok)
				i++
			}
		}
	}
	if mid == 0 {
		fmt.Println("Data ditemukan!")
		tbl.Print()
	} else {
		fmt.Println("Data tidak ditemukan!")
	}

	fmt.Println()
}

// Binary Search
func showStokSearchedData(data Data, n int, isSorted string) {
	defer timeElapsed(time.Now(), "Binary")
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showAllProduct(data, n)
	menuHeaderShowSearchedProduct()
	var x int
	fmt.Println("Masukkan stok produk yang ingin dicari (syarat: data harus TERURUT): ")
	fmt.Print(">>>> ")
	fmt.Scanln(&x)
	fmt.Println()

	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("No.", "Nama Produk", "Merek Produk", "Jenis Produk", "Harga Produk", "Stok Produk")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	var mid, left, right, i int
	i = 1
	left = 0
	right = n - 1
	if isSorted == "asc" {
		for left <= right {
			mid = (left + right) / 2
			if data[mid].stok < x {
				left = mid + 1
			} else if data[mid].stok > x {
				right = mid
			} else {
				tbl.AddRow(i, data[mid].namaProduk, data[mid].merek, data[mid].jenis, data[mid].harga, data[mid].stok)
				i++
			}
		}
	} else if isSorted == "desc" {
		for left <= right {
			mid = (left + right) / 2
			if x > data[mid].stok {
				right = mid - 1
			} else if x < data[mid].stok {
				left = mid + 1
			} else {
				tbl.AddRow(i, data[mid].namaProduk, data[mid].merek, data[mid].jenis, data[mid].harga, data[mid].stok)
				i++
			}
		}
	}
	if mid == 0 {
		fmt.Println("Data ditemukan!")
		tbl.Print()
	} else {
		fmt.Println("Data tidak ditemukan!")
	}

	fmt.Println()
}

/*
Prosedur untuk menu pertama di dalam
Utilitas Data, yaitu menampilkan transaksi data
yang lagi di edit dengan tabel yang ingin di edit.
Dipakai saat menampilkan menu:
- Edit Data
*/
func showSelectedProduct(data Data, n int) {
	fmt.Println()
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("No.", "Nama Produk", "Merek Produk", "Jenis Produk", "Harga Produk", "Stok Produk")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	tbl.AddRow(n, data[n-1].namaProduk, data[n-1].merek, data[n-1].jenis, data[n-1].harga, data[n-1].stok)

	tbl.Print()
	fmt.Println()
}

/*
Prosedur untuk menu pertama di dalam
Utilitas Data, yaitu Edit Data.
*/
func editProductData(data *Data, x int) {
	var n int
	var namaTemp, merekTemp, jenisTemp string
	var hargaTemp float64
	var stokTemp int
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showAllProduct(*data, x)
	menuHeaderEditData()
	fmt.Print("Masukkan kolom data yang akan diedit: ")
	fmt.Scanln(&n)
	namaTemp = data[n-1].namaProduk
	merekTemp = data[n-1].merek
	jenisTemp = data[n-1].jenis
	hargaTemp = data[n-1].harga
	stokTemp = data[n-1].stok
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	for n == 0 && n > MAXPRODUCT {
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
		showAllProduct(*data, x)
		menuHeaderEditData()
		fmt.Println("Masukkan kolom data yang benar!")
		fmt.Print("Masukkan kolom data yang akan diedit: ")
		fmt.Scanln(&n)
		namaTemp = data[n-1].namaProduk
		merekTemp = data[n-1].merek
		jenisTemp = data[n-1].jenis
		hargaTemp = data[n-1].harga
		stokTemp = data[n-1].stok
		fmt.Print("\033[2J")
		fmt.Print("\033[H")
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showSelectedProduct(*data, n)
	menuHeaderEditData()
	fmt.Println("Nama Produk:")
	fmt.Print(">>>> ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		data[n-1].namaProduk = scanner.Text()
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	if data[n-1].namaProduk == "\n" {
		data[n-1].namaProduk = namaTemp
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showSelectedProduct(*data, n)
	menuHeaderEditData()
	fmt.Println("Merek Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[n-1].merek)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	if data[n-1].merek == "\n" {
		data[n-1].merek = merekTemp
		fmt.Println(merekTemp)
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showSelectedProduct(*data, n)
	menuHeaderEditData()
	fmt.Println("Jenis Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[n-1].jenis)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	if data[n-1].jenis == "\n" {
		data[n-1].jenis = jenisTemp
		fmt.Println(jenisTemp)
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showSelectedProduct(*data, n)
	menuHeaderEditData()
	fmt.Println("Harga Produk:")
	fmt.Print(">>>> Rp. ")
	fmt.Scanln(&data[n-1].harga)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	if data[n-1].harga == 0 {
		data[n-1].harga = hargaTemp
		fmt.Println(hargaTemp)
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showSelectedProduct(*data, n)
	menuHeaderEditData()
	fmt.Println("Stok Produk:")
	fmt.Print(">>>> ")
	fmt.Scanln(&data[n-1].stok)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	if data[n-1].stok == stokTemp {
		data[n-1].stok = stokTemp
	} else if data[n-1].stok == 0 {
		data[n-1].stok = 0
	}
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showSelectedProduct(*data, n)
	menuHeaderEditData()
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	fmt.Println("Data Ke -", n, "Berhasil Diedit")
	showAllProduct(*data, x)

}

func SortProduct(data *Data, n int, isSorted *string) {
	var det int
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	showAllProduct(*data, n)
	menuHeaderSortProduct()
	menuOptionsUrutkanDataProduk()
	fmt.Println("Masukkan menu: ")
	fmt.Print(">>>> ")
	fmt.Scanln(&det)
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
	if det == 1 {
		sortByAscending(data, n, "nama")
		fmt.Println()
		fmt.Println("Data berhasil di urutkan")
		*isSorted = "asc"
	} else if det == 2 {
		sortByDescending(data, n, "nama")
		fmt.Println()
		fmt.Println("Data berhasil di urutkan")
		*isSorted = "desc"
	} else if det == 3 {
		sortByAscending(data, n, "harga")
		fmt.Println()
		fmt.Println("Data berhasil di urutkan")
		*isSorted = "asc"
	} else if det == 4 {
		sortByDescending(data, n, "harga")
		fmt.Println()
		fmt.Println("Data berhasil di urutkan")
		*isSorted = "desc"
	} else if det == 5 {
		sortByAscending(data, n, "stok")
		fmt.Println()
		fmt.Println("Data berhasil di urutkan")
		*isSorted = "asc"
	} else if det == 6 {
		sortByDescending(data, n, "stok")
		fmt.Println()
		fmt.Println("Data berhasil di urutkan")
		*isSorted = "desc"
	} else if det == 7 {

	} else {
		fmt.Println("Mohon masukkan pengurutan data yang benar!")
	}
	showAllProduct(*data, n)
}

/*
Prosedur untuk menu ke-tiga dan empat di dalam
Utilitas Data, yaitu Cari Data & Urutkan Data.
Berdasarkan nama produk, Ascending.
*/
func findMinByString(data Data, n, pass int) int {
	var idx int = pass - 1
	var k int = pass
	for k < n {
		if data[idx].namaProduk[0] > data[k].namaProduk[0] {
			idx = k
		} else if data[idx].namaProduk[0] == data[k].namaProduk[0] {
			if data[idx].namaProduk[1] > data[k].namaProduk[1] {
				idx = k
			} else if data[idx].namaProduk[1] == data[k].namaProduk[1] {
				if data[idx].namaProduk[2] > data[k].namaProduk[2] {
					idx = k
				}
			}
		}
		k++
	}
	return idx
}

/*
Prosedur untuk menu ke-tiga dan empat di dalam
Utilitas Data, yaitu Cari Data & Urutkan Data.
Berdasarkan harga, Ascending.
*/
func findMinByHarga(data Data, n, pass int) int {
	var idx int = pass - 1
	var k int = pass
	for k < n {
		if data[idx].harga > data[k].harga {
			idx = k
		}
		k++
	}
	return idx
}

/*
Prosedur untuk menu ke-tiga dan empat di dalam
Utilitas Data, yaitu Cari Data & Urutkan Data.
Berdasarkan stok tersedia, Ascending.
*/
func findMinByStok(data Data, n, pass int) int {
	var idx int = pass - 1
	var k int = pass
	for k < n {
		if data[idx].stok > data[k].stok {
			idx = k
		}
		k++
	}
	return idx
}

/*
Prosedur untuk menu ke-tiga dan empat di dalam
Utilitas Data, yaitu Cari Data & Urutkan Data.
Berdasarkan nama produk, Descending.
*/
// func findMaxByString(data Data, n, pass int) int {
// 	var idx int = pass - 1
// 	var k int = pass
// 	for k < n {
// 		if data[idx].namaProduk[0] < data[k].namaProduk[0] {
// 			idx = k
// 		} else if data[idx].namaProduk[0] == data[k].namaProduk[0] {
// 			if data[idx].namaProduk[1] < data[k].namaProduk[1] {
// 				idx = k
// 			} else if data[idx].namaProduk[1] == data[k].namaProduk[1] {
// 				if data[idx].namaProduk[2] < data[k].namaProduk[2] {
// 					idx = k
// 				}
// 			}
// 		}
// 		k++
// 	}
// 	return idx
// }

/*
Prosedur untuk menu ke-tiga dan empat di dalam
Utilitas Data, yaitu Cari Data & Urutkan Data.
Berdasarkan harga, Descending.
*/
// func findMaxByHarga(data Data, n, pass int) int {
// 	var idx int = pass - 1
// 	var k int = pass
// 	for k < n {
// 		if data[idx].harga < data[k].harga {
// 			idx = k
// 		}
// 		k++
// 	}
// 	return idx
// }

/*
Prosedur untuk menu ke-tiga dan empat di dalam
Utilitas Data, yaitu Cari Data & Urutkan Data.
Berdasarkan stok tersedia, Descending.
*/
// func findMaxByStok(data Data, n, pass int) int {
// 	var idx int = pass - 1
// 	var k int = pass
// 	for k < n {
// 		if data[idx].stok < data[k].stok {
// 			idx = k
// 		}
// 		k++
// 	}
// 	return idx
// }

/*
Prosedur untuk empat di dalam
Utilitas Data, yaitu Urutkan Data.
Logic utama, Ascending.
SELECTION SORT
*/
func sortByAscending(data *Data, n int, det string) {
	defer timeElapsed(time.Now(), "Selection")
	var pass, idx int
	var temp Produk
	pass = 1
	for pass <= n-1 {
		if det == "nama" {
			idx = findMinByString(*data, n, pass)
		} else if det == "harga" {
			idx = findMinByHarga(*data, n, pass)
		} else if det == "stok" {
			idx = findMinByStok(*data, n, pass)
		}

		temp = data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp
		pass++
	}
}

/*
Prosedur untuk empat di dalam
Utilitas Data, yaitu Urutkan Data.
Logic utama, Descending.
INSERTION SORT
*/
func sortByDescending(data *Data, n int, det string) {
	defer timeElapsed(time.Now(), "Insertion")
	var pass, j int
	var temp Produk
	pass = 1
	for pass <= n-1 {
		j = pass
		temp = data[j]
		if det == "nama" {
			for j > 0 && temp.namaProduk > data[j-1].namaProduk {
				data[j] = data[j-1]
				j--
			}
			data[j] = temp
			pass++
		} else if det == "harga" {
			for j > 0 && temp.harga > data[j-1].harga {
				data[j] = data[j-1]
				j--
			}
			data[j] = temp
			pass++
		} else if det == "stok" {
			for j > 0 && temp.stok > data[j-1].stok {
				data[j] = data[j-1]
				j--
			}
			data[j] = temp
			pass++
		}
	}
}

// Fitur Eksplorasi
func timeElapsed(start time.Time, name string) func() {
	elapsed := time.Since(start).Milliseconds()
	return func() {
		fmt.Println()
		fmt.Printf("%s time elapsed: %dms \n", name, elapsed)
	}
}

/*
Printlines estetik pada program.
Program startup.
*/
func menuStart() {
	fmt.Println(" ____  _                                ")
	fmt.Println("/ ___|| |_ ___  _ __ __ _  __ _  ___    ")
	fmt.Println("\\___ \\| __/ _ \\| '__/ _` |/ _` |/ _ \\")
	fmt.Println(" ___) | || (_) | | | (_| | (_| |  __/   ")
	fmt.Println("|____/ \\__\\___/|_|  \\__,_|\\__, |\\___| ")
	fmt.Println("                          |___/        Aplikasi      ")
	fmt.Println(" ____             _                    Manajemen Inventaris   ")
	fmt.Println("/ ___|  ___  _ __| |_ ___ _ __         Toko Elektronik   ")
	fmt.Println("\\___ \\ / _ \\| '__| __/ _ \\ '__|        ")
	fmt.Println(" ___) | (_) | |  | ||  __/ |           Dibuat oleh:")
	fmt.Println("|____/ \\___/|_|   \\__\\___|_|           Kelompok 1")
}

/*
Printlines estetik pada program.
Header Menu Utama.
*/
func menuProcess() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("\x1b[7;37m               M E N U - U T A M A             \x1b[0;37m")
	fmt.Println("-----------------------------------------------")
}

/*
Printlines estetik pada program.
Opsi Menu Utama.
*/
func menuOptionsProcess() {
	fmt.Println(" ")
	fmt.Println("1. Konfigurasi Data Produk")
	fmt.Println("2. Pencatatan Transaksi")
	fmt.Println("3. Exit Program")
	fmt.Println(" ")
}

/*
Printlines estetik pada program.
Program shut down.
*/
func menuEnd() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("\x1b[7;37m            S A M P A I - J U M P A            \x1b[0;37m")
	fmt.Println("-----------------------------------------------")
}

/*
Printlines estetik pada program.
Header Utlitas Data.
*/
func menuHeaderTampilSemuaDataProduk() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("\x1b[7;37m           U T I L I T A S - D A T A           \x1b[0;37m")
	fmt.Println("-----------------------------------------------")
}

/*
Printlines estetik pada program.
Opsi Utilitas Data.
*/
func menuOptionsTampilSemuaDataProduk() {
	fmt.Println(" ")
	fmt.Println("1. Edit Data")
	fmt.Println("2. Hapus Data")
	fmt.Println("3. Cari Data")
	fmt.Println("4. Urutkan Data")
	fmt.Println("5. Kembali Ke Menu Utama")
	fmt.Println(" ")
}

/*
Printlines estetik pada program.
Header Konfigurasi Data Produk.
*/
func menuHeaderKonfigurasiDataProduk() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("\x1b[7;37m K O N F I G U R A S I - D A T A - P R O D U K \x1b[0;37m")
	fmt.Println("-----------------------------------------------")
}

/*
Printlines estetik pada program.
Header Pengisisan Data.
*/
func menuHeaderInputDataProduk() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("\x1b[7;37m          P E N G I S I A N - D A T A          \x1b[0;37m")
	fmt.Println("-----------------------------------------------")
	fmt.Println(" ")
}

/*
Printlines estetik pada program.
Header Pencatatan Transaksi
*/
func menuHeaderPencatatanTransaksi() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("\x1b[7;37m    P E N C A T A T A N - T R A N S A K S I    \x1b[0;37m")
	fmt.Println("-----------------------------------------------")
}

/*
Printlines estetik pada program.
Opsi Pencatatan Transaksi.
*/
func menuOptionsPencatatanTransaksi() {
	fmt.Println(" ")
	fmt.Println("1. Tambah Transaksi")
	fmt.Println("2. Kembali ke Menu Utama")
	fmt.Println(" ")
}

/*
Printlines estetik pada program.
Opsi Konfigurasi Data Produk.
*/
func menuOptionsKonfigurasiDataProduk() {
	fmt.Println(" ")
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Utilitas Data")
	fmt.Println("3. Kembali ke Menu Utama")
	fmt.Println(" ")
}

/*
Printlines estetik pada program.
Opsi Urutkan Data.
*/
func menuOptionsUrutkanDataProduk() {
	fmt.Println(" ")
	fmt.Println("1. Urutkan Berdasarkan Nama Produk Secara Ascending")
	fmt.Println("2. Urutkan Berdasarkan Nama Produk Secara Descending")
	fmt.Println("3. Urutkan Berdasarkan Harga Secara Ascending")
	fmt.Println("4. Urutkan Berdasarkan Harga Secara Descending")
	fmt.Println("5. Urutkan Berdasarkan Stok Secara Ascending")
	fmt.Println("6. Urutkan Berdasarkan Stok Secara Descending")
	fmt.Println("7. Kembali ke Menu Utama")
	fmt.Println(" ")
}

/*
Printlines estetik pada program.
Header Edit Data.
*/
func menuHeaderEditData() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("\x1b[7;37m               E D I T - D A T A               \x1b[0;37m")
	fmt.Println("-----------------------------------------------")
	fmt.Println(" ")
}

/*
Printlines estetik pada program.
Header Hapus Data.
*/
func menuHeaderDeleteData() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("\x1b[7;37m              H A P U S - D A T A              \x1b[0;37m")
	fmt.Println("-----------------------------------------------")
	fmt.Println(" ")
}

/*
Printlines estetik pada program.
Header Cari Data.
*/
func menuHeaderShowSearchedProduct() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("\x1b[7;37m               C A R I - D A T A               \x1b[0;37m")
	fmt.Println("-----------------------------------------------")
	fmt.Println(" ")
}

/*
Printlines estetik pada program.
Header Urutkan Data.
*/
func menuHeaderSortProduct() {
	fmt.Println("-----------------------------------------------")
	fmt.Println("\x1b[7;37m            U R U T K A N - D A T A            \x1b[0;37m")
	fmt.Println("-----------------------------------------------")
	fmt.Println(" ")
}

// ---------------------------------------------------------------
