package main

import (
	"fmt"
)

const NMAX = 100 //kapasitas maksimal aplikasi

// ==========================================
// DEKLARASI STRUCT & ARRAY
// ==========================================
type lapangan struct { //baris 10-13, membuat tipe data struct untuk
	IDlapang, harga int // mengelompokkan detail informasi dari sebuah lapangan
	nama            string
}
type dataPenyewa struct { //baris 14-17, struct ini sebagai cetakan profil pelanggan
	Nama, NoTlp string //cetakan ini menyimpan ID penyewa serta Nama dan No Telepon
	IDpenyewa   int
}
type jadwalDANtransaksi struct { //baris 18-22, untuk struk atau nota pembayaran
	IDtransaksi, IDlapang, IDpenyewa int //struct ini menyimpan relasi antar tabel
	bulan, tanggal, jamMulai         int
	durasi, total                    int
}
type jadwalKosong struct { //baris 23-26, wadah data sementara, untuk menyimpan detailJam, harga dan namaLapanagan
	jam, harga int //secara lebnih spesifik saat program sedang menjalankan fitur mencari dan mengurutkan jadwal kosong
	namaLap    string
}
type arrLapangan [NMAX]lapangan   //baris 27-29, agar menyimpan banyak struct sekaligus
type arrPenyewa [NMAX]dataPenyewa //ukurannya dipatok mentok sampai [NMAX] alias maksimal 100 slot data.
type arrTransaksi [NMAX]jadwalDANtransaksi

// baris 32, Array jadwal bisa lebih besar karena 1 lapangan bisa punya banyak jam kosong
type arrJadwal [NMAX * 24]jadwalKosong

// ==========================================
// FUNGSI BANTUAN
// ==========================================
func bersihkanLayar() { //baris 37-39, Kegunaannya untuk membersihkan semua teks yang ada di layar terminal (Clear Screen)
	fmt.Print("\033[H\033[2J") //dan agar tampilan program nya selalu rapi
}
func tahanLayar() { //baris 40-44, untuk menahan atau nge-pause jalannya program sementara waktu
	var tahan string
	fmt.Print("\nKetik apa saja lalu Enter untuk kembali...")
	fmt.Scan(&tahan)
}

// ==========================================
// PROGRAM UTAMA
// ==========================================
func main() { //baris 49-60, untuk mendeklarasikan variabel-variabel lokal utama
	var dataL arrLapangan //berguna untuk mendeklarasikan variabel-variabel lokal utama
	var nLap int = 0

	var dataP arrPenyewa
	var nPenyewa int = 0

	var dataT arrTransaksi
	var nTrans int = 0

	var pilihan int
	aplikasiJalan := true

	// DATA DUMMY (Untuk keperluan testing)
	// ==========================================

	// 1. Data Dummy Lapangan
	dataL[0] = lapangan{1, 40000, "lapangan_1"}
	dataL[1] = lapangan{2, 60000, "lapangan_2"}
	dataL[2] = lapangan{3, 80000, "lapangan_3"}
	dataL[3] = lapangan{4, 100000, "lapangan_4"}
	dataL[4] = lapangan{5, 120000, "lapangan_5"}
	nLap = 5

	// 2. Data Dummy Penyewa
	dataP[0] = dataPenyewa{"Budi", "08111222333", 1}
	dataP[1] = dataPenyewa{"Andi", "08222333444", 2}
	dataP[2] = dataPenyewa{"Citra", "08333444555", 3}
	dataP[3] = dataPenyewa{"Dedi", "08444555666", 4}
	dataP[4] = dataPenyewa{"Eka", "08555666777", 5}
	nPenyewa = 5

	// 3. Data Dummy Transaksi (Bulan 5 / Mei)
	// Format struct: IDtransaksi, IDlapang, IDpenyewa, bulan, tanggal, jamMulai, durasi, total
	dataT[0] = jadwalDANtransaksi{101, 1, 1, 5, 10, 10, 2, 80000}  // Budi sewa lapangan_1 (40rb x 2 jam)
	dataT[1] = jadwalDANtransaksi{102, 2, 2, 5, 10, 14, 1, 60000}  // Andi sewa lapangan_2 (60rb x 1 jam)
	dataT[2] = jadwalDANtransaksi{103, 3, 3, 5, 11, 19, 2, 160000} // Citra sewa lapangan_3 (80rb x 2 jam)
	dataT[3] = jadwalDANtransaksi{104, 4, 4, 5, 12, 20, 2, 200000} // Dedi sewa lapangan_4 (100rb x 2 jam)
	dataT[4] = jadwalDANtransaksi{105, 5, 5, 5, 12, 16, 1, 120000} // Eka sewa lapangan_5 (120rb x 1 jam)
	nTrans = 5

	for aplikasiJalan { //baris 90-104, ini mmebuat perulangan dan menampilkan antarmuka. Kegunaannya, selama "saklar" tadi masih bernilai true
		bersihkanLayar() //program akan membersihkan layar dan mencetak teks Menu Utama (1-7) ke terminal
		fmt.Println("==========================================")
		fmt.Println("        Penyewaan Lapangan Futsal         ")
		fmt.Println("==========================================")
		fmt.Println("1. Kelola Data Lapangan")
		fmt.Println("2. Kelola Data Penyewa")
		fmt.Println("3. Catat Transaksi Sewa")
		fmt.Println("4. Cari Data Penyewa (Sequential & Binary)")
		fmt.Println("5. Urutkan Jadwal Kosong (Selection & Insertion)")
		fmt.Println("6. Statistik Pendapatan & Jam Terlaris")
		fmt.Println("7. Keluar")
		fmt.Println("==========================================")
		fmt.Print("Pilih menu (1-7): ")
		fmt.Scan(&pilihan)

		if pilihan == 1 { //baris 106-117, untuk mengatur rute navigasi menggunakan percabangan
			menuLapangan(&dataL, &nLap) //berguna untuk mengecek angka yang diketik pengguna untuk memanggil fungsi
		} else if pilihan == 2 { //dan pada variabel yang dikirim, itu menandakan penggunaan pointer
			menuPenyewa(&dataP, &nPenyewa) //kegunaannya agar perubahan data di dalam menu tersebut benar-benar tersimpan ke wadah utama)
		} else if pilihan == 3 {
			tambahTransaksi(&dataT, &nTrans, dataP, nPenyewa, dataL, nLap)
		} else if pilihan == 4 {
			cariPenyewa(&dataP, nPenyewa)
		} else if pilihan == 5 {
			urutJadwalKosong(dataL, nLap, dataT, nTrans)
		} else if pilihan == 6 {
			tampilStatistik(dataT, nTrans)
		} else if pilihan == 7 { //baris 118-121, untuk mengeksekusi fitur Keluar (opsi 7)
			aplikasiJalan = false //Kegunaannya "saklar" aplikasiJalan akan diubah menjadi false
			bersihkanLayar()
			fmt.Println("Terima kasih telah menggunakan Futsal-Book!")
		} else { //baris 122-125, untuk penanganan error dan berguna saat  mengetik angka selain 1 sampai 7, program akan masuk ke blok else ini
			fmt.Println("\nPilihan tidak valid!") //Program akan menampilkan peringatan "Pilihan tidak valid!" dan memanggil fungsi
			tahanLayar()
		}
	}
}

// ==========================================
// FUNGSI MENU CRUD (Disatukan Biar Rapi)
// ==========================================
func menuLapangan(L *arrLapangan, n *int) { //baris 131-142, untuk mendeklarasikan fungsi dan menampilkan data (fitur Read)
	bersihkanLayar()                            // program akan mengecek apakah jumlah data saat ini masih nol
	fmt.Println("--- KELOLA DATA LAPANGAN ---") // Jika "ya", program mencetak "Data lapangan masih kosong".
	if *n == 0 {                                //jika "tidak" program akan melakukan perulangan for untuk mencetak seluruh dataLapangan
		fmt.Println("Data lapangan masih kosong.")
	} else {
		i := 0
		for i < *n {
			fmt.Printf("%d. ID: %d | Nama: %s | Harga/Jam: Rp%d\n", i+1, L[i].IDlapang, L[i].nama, L[i].harga)
			i++
		}
	}
	fmt.Println("\nAksi:")        //baris 143-152, untuk  Jika ya, program mencetak "Data lapangan masih kosong"
	fmt.Println("1. Tambah Data") //berguna untk memberikan pilihan kepada admin (1. Tambah, 2. Ubah, 3. Hapus, 4. Kembali),
	fmt.Println("2. Ubah Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Kembali")
	fmt.Print("Pilih aksi (1-4): ")

	var aksi int
	fmt.Scan(&aksi)

	if aksi == 1 { //baris 152-166, dibuat untk mengeksekusi proses Tambah Data (fitur Create)
		if *n < NMAX { //program akan mengecek terlebih dahulu apakah kapasitas array masih tersedia
			fmt.Print("\nID Lapangan Baru: ") //Jika masih memukinkan program meminta input ID, Nama, dan Harga untuk dimasukkan
			fmt.Scan(&L[*n].IDlapang)         //ke array indeks (*n), jika berhasil nilai total data (*n = *n + 1) akan bertambah
			fmt.Print("Nama Lapangan: ")
			fmt.Scan(&L[*n].nama)
			fmt.Print("Harga Sewa Per Jam: Rp ")
			fmt.Scan(&L[*n].harga)
			*n = *n + 1
			fmt.Println("Data berhasil ditambahkan!")
		} else {
			fmt.Println("Kapasitas penuh!")
		}
		tahanLayar()
	} else if aksi == 2 { //baris 167-190, untuk mengeksekusi proses Ubah Data (fitur Update), berguna jika admin memilih aksi "== 2"
		var cariID int
		fmt.Print("\nMasukkan ID Lapangan yang diubah: ")
		fmt.Scan(&cariID) //program akan meminta input ID lapangan yang ingin diedit
		ketemu := false   //program akan melakukan pencarian (Sequential Search) dari indeks ke-0
		j := 0            //Jika ID (ketemu = true) maka program akan meminta input Nama dan Harga baru untuk menimpa data lama pada posisi indeks (idx)
		idx := 0
		for j < *n && !ketemu {
			if L[j].IDlapang == cariID {
				ketemu = true
				idx = j
			}
			j++
		}
		if ketemu {
			fmt.Print("Nama Baru: ")
			fmt.Scan(&L[idx].nama)
			fmt.Print("Harga Baru: Rp ")
			fmt.Scan(&L[idx].harga)
			fmt.Println("Data berhasil diubah!")
		} else {
			fmt.Println("ID tidak ditemukan.")
		}
		tahanLayar()
	} else if aksi == 3 { //baris 191-217, untuk mengeksekusi proses (fitur Delete), berguna jika admin memilih aksi == 3
		var cariID int //program akan mencari ID target menggunakan Sequential Search.
		fmt.Print("\nMasukkan ID Lapangan yang dihapus: ")
		fmt.Scan(&cariID)
		ketemu := false //program menggunakan "Geser Kiri" pada perulangan "for k < *n-1. Logika L[k] = L[k+1]"
		j := 0          //data di sebelah kanan akan digeser ke kiri untuk menghapus data yang ditargetkan
		idx := 0        // jumlah total data dikurangi satu "*n = *n - 1"
		for j < *n && !ketemu {
			if L[j].IDlapang == cariID {
				ketemu = true
				idx = j
			}
			j++
		}
		if ketemu {
			k := idx
			for k < *n-1 {
				L[k] = L[k+1]
				k++
			}
			*n = *n - 1
			fmt.Println("Data berhasil dihapus!")
		} else {
			fmt.Println("ID tidak ditemukan.")
		}
		tahanLayar()
	}
}
func menuPenyewa(P *arrPenyewa, n *int) { //baris 2019-230, mendeklarasikan fungsi dan menampilkan data pelanggan "fitur Read"
	bersihkanLayar() //program akan mengecek apakah jumlah data penyewa saat ini masih nol "*n == 0"
	fmt.Println("--- KELOLA DATA PENYEWA ---")
	if *n == 0 { //jika "true" program akan melakukan perulangan "for" untuk mencetak seluruh rincian identitas penyewa
		fmt.Println("Data penyewa masih kosong.")
	} else {
		i := 0
		for i < *n {
			fmt.Printf("%d. ID: %d | Nama: %s | No.Tlp: %s\n", i+1, P[i].IDpenyewa, P[i].Nama, P[i].NoTlp)
			i++
		}
	}
	fmt.Println("\nAksi:") //baris 231-240, menampilkan opsi menu manajemen penyewa ke layar,
	fmt.Println("1. Tambah Data")
	fmt.Println("2. Ubah Data")
	fmt.Println("3. Hapus Data")
	fmt.Println("4. Kembali")
	fmt.Print("Pilih aksi (1-4): ")

	var aksi int
	fmt.Scan(&aksi)

	if aksi == 1 { //baris 241-254, mengeksekusi proses Tambah Data "fitur Create"
		if *n < NMAX { //jika ""== 1", program melakukan validasi batas maksimal data "*n < NMAX"
			fmt.Print("\nID Penyewa Baru: ")
			fmt.Scan(&P[*n].IDpenyewa)
			fmt.Print("Nama Penyewa (Tanpa Spasi): ")
			fmt.Scan(&P[*n].Nama)
			fmt.Print("Nomor Telepon: ")
			fmt.Scan(&P[*n].NoTlp)
			*n = *n + 1
			fmt.Println("Data berhasil ditambahkan!")
		} else {
			fmt.Println("Kapasitas penuh!")
		}
		tahanLayar()
	} else if aksi == 2 { //baris 255-278, mengeksekusi proses Ubah Data "fitur Update"
		var cariID int //program meminta admin menginput ID penyewa yang ingin diperbarui datanya
		fmt.Print("\nMasukkan ID Penyewa yang diubah: ")
		fmt.Scan(&cariID) //program mencari posisi data menggunakan "Sequential Search"
		ketemu := false   //jika ditemukan program akan meminta input "Nama" dan "Nomor Telepon" yang baru
		j := 0
		idx := 0
		for j < *n && !ketemu {
			if P[j].IDpenyewa == cariID {
				ketemu = true
				idx = j
			}
			j++
		}
		if ketemu {
			fmt.Print("Nama Baru: ")
			fmt.Scan(&P[idx].Nama)
			fmt.Print("No Tlp Baru: ")
			fmt.Scan(&P[idx].NoTlp)
			fmt.Println("Data berhasil diubah!")
		} else {
			fmt.Println("ID tidak ditemukan.")
		}
		tahanLayar()
	} else if aksi == 3 { //baris 279-306, mengeksekusi proses Hapus Data "fitur Delete"
		var cariID int //program akan melacak posisi menggunakan "Sequential Search"
		fmt.Print("\nMasukkan ID Penyewa yang dihapus: ")
		fmt.Scan(&cariID) //jika ditemukan, program melakukan proses penghapusan menggunakan "Geser Kiri" ("P[k] = P[k+1]")
		ketemu := false   //seluruh data pelanggan yang berada di sebelah kanan posisi target akan disuruh bergeser ke kiri
		j := 0
		idx := 0
		for j < *n && !ketemu {
			if P[j].IDpenyewa == cariID {
				ketemu = true
				idx = j
			}
			j++
		}
		if ketemu {
			k := idx
			for k < *n-1 {
				P[k] = P[k+1]
				k++
			}
			*n = *n - 1
			fmt.Println("Data berhasil dihapus!")
		} else {
			fmt.Println("ID tidak ditemukan.")
		}
		tahanLayar()
	}
}

// ==========================================
// TRANSAKSI SEWA (Poin b)
// ==========================================
func tambahTransaksi(T *arrTransaksi, nT *int, P arrPenyewa, nP int, L arrLapangan, nL int) {
	bersihkanLayar() //baris 311-317, mendeklarasikan fungsi dan melakukan validasi syarat awal
	fmt.Println("--- CATAT TRANSAKSI SEWA ---")
	if *nT < NMAX { //sebelum transaksi diproses program akan mengecek papakah kapasitas memori transaksi
		if nP == 0 || nL == 0 {
			fmt.Println("Data Lapangan dan Penyewa harus ada minimal 1!")
		} else { //dan program memastikan minimal ada 1 data penyewa dan 1 data lapangan yang terdaftar di sistem
			fmt.Print("ID Transaksi: ")   //baros 318-331, melakukan proses input rincian sewa, program meminta admin
			fmt.Scan(&T[*nT].IDtransaksi) //mengetikkan data transaksi secara lengkap, semua data ini langsung disimpan ke dalam array transaksi
			fmt.Print("ID Lapangan yang disewa: ")
			fmt.Scan(&T[*nT].IDlapang)
			fmt.Print("ID Penyewa: ")
			fmt.Scan(&T[*nT].IDpenyewa)
			fmt.Print("Bulan (1-12): ")
			fmt.Scan(&T[*nT].bulan)
			fmt.Print("Tanggal (1-31): ")
			fmt.Scan(&T[*nT].tanggal)
			fmt.Print("Jam Mulai (contoh 14 untuk 14:00): ")
			fmt.Scan(&T[*nT].jamMulai)
			fmt.Print("Durasi Sewa (Jam): ")
			fmt.Scan(&T[*nT].durasi)

			// Cek Harga Lapangan
			hargaPerJam := 0   //baris 334-343, melakukan pencarian harga lapangan "Sequential Search"
			ketemuLap := false //program mencari berapa harga sewa lapangan tersebut
			i := 0             //dan program akan menelusuri array data lapangan satu per satu
			for i < nL && !ketemuLap {
				if L[i].IDlapang == T[*nT].IDlapang {
					hargaPerJam = L[i].harga
					ketemuLap = true
				}
				i++
			}
			if ketemuLap { //baris 344-351, mengeksekusi perhitungan total biaya dan finalisasi
				T[*nT].total = T[*nT].durasi * hargaPerJam
				*nT = *nT + 1 //jika lapangan valid dan ditemuka program akan menghitung secara otomatis total tagihan
				fmt.Printf("\nTransaksi Berhasil! Total Harga: Rp%d\n", T[*nT-1].total)
			} else { //dengan cara mengalikan "durasi sewa" dengan "hargaPerJam"
				fmt.Println("ID Lapangan tidak ditemukan, transaksi gagal.")
			} //data transaksi ditambah satu dan program mencetak pesan "Transaksi Berhasil!" dan total harganya ke layar
		}
	} else { //baris 352-356, program akan mencetak "Kapasitas transaksi penuh!"
		fmt.Println("Kapasitas transaksi penuh!")
	} //fungsi "tahanLayar()"" dipanggil u=agar nota "total harga"
	tahanLayar() //atau pesan error tidak langsung hilang dari layar sebelum dibaca admin
}

// ==========================================
// SEARCHING (Poin c)
// ==========================================
func cariPenyewa(P *arrPenyewa, n int) { //baris 358-367, mendeklarasikan fungsi dan menampilkan menu awal pencarian
	bersihkanLayar() //memberikan dua opsi metode, mencari berdasarkan Nama dan berdasarkan Nomor HP
	fmt.Println("--- CARI PENYEWA ---")
	fmt.Println("1. Berdasarkan Nama (Sequential Search)")
	fmt.Println("2. Berdasarkan No HP (Binary Search)")
	fmt.Print("Pilih (1/2): ")

	var pilCari int    //baris 368-386, mengeksekusi percabangan "Sequential Search", jika memilih mencari nama
	fmt.Scan(&pilCari) //program akan meminta input "cariNama", menelusuri data penyewa satu per satu dari urutan paling atas

	if pilCari == 1 {
		var cariNama string
		fmt.Print("Masukkan Nama: ")
		fmt.Scan(&cariNama)

		ketemu := false
		i := 0
		for i < n && !ketemu {
			if P[i].Nama == cariNama {
				fmt.Printf("\nDitemukan! ID: %d | Nama: %s | Tlp: %s\n", P[i].IDpenyewa, P[i].Nama, P[i].NoTlp)
				ketemu = true
			}
			i++
		}
		if !ketemu {
			fmt.Println("Nama tidak ditemukan.")
		}

	} else if pilCari == 2 { //baris 389-400, melakukan pengurutan data menggunakan "Insertion Sort"
		var cariNo string             //menyusun seluruh data pelanggan di dalam sistem
		fmt.Print("Masukkan No HP: ") //berdasarkan Nomor HP dari yang terkecil ke terbesar (Ascending)
		fmt.Scan(&cariNo)

		// Insertion Sort Ascending berdasarkan NoTlp untuk syarat Binary Search
		i := 1
		for i < n {
			temp := P[i]
			j := i - 1
			for j >= 0 && P[j].NoTlp > temp.NoTlp {
				P[j+1] = P[j]
				j-- //baris 401-430, mengeksekusi pencarian "Binary Search"
			} //program mencari posisi data dengan cara membelah daftar nomor HP yang sudah urut
			P[j+1] = temp
			i++ //Jika nomor HP target lebih kecil, program langsung membuang belahan kanan dan fokus ke belahan kiri, begitu pula sebaliknya
		}

		// Binary Search
		kiri := 0
		kanan := n - 1
		ketemu := false

		for kiri <= kanan && !ketemu {
			tengah := (kiri + kanan) / 2
			if P[tengah].NoTlp == cariNo {
				fmt.Printf("\nDitemukan! ID: %d | Nama: %s | Tlp: %s\n", P[tengah].IDpenyewa, P[tengah].Nama, P[tengah].NoTlp)
				ketemu = true
			} else if P[tengah].NoTlp < cariNo {
				kiri = tengah + 1
			} else {
				kanan = tengah - 1
			}
		}
		if !ketemu {
			fmt.Println("No HP tidak ditemukan.")
		}
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
	tahanLayar()
}
// ==========================================
// SORTING JADWAL KOSONG (Poin d)
// ==========================================
func urutJadwalKosong(L arrLapangan, nLap int, T arrTransaksi, nTrans int) {
	bersihkanLayar() //baris 453-446, mendeklarasikan fungsi dan persiapan awal
	fmt.Println("--- URUTKAN JADWAL KOSONG ---")
	//program mencetak judul menu meminta admin memasukkan Tanggal dan Bulan yang ingin dicek
	var cariBulan, cariTanggal int //untuk menyiapkan sebuah array sementara
	fmt.Print("Cek jadwal kosong untuk Bulan (1-12): ")
	fmt.Scan(&cariBulan)
	fmt.Print("Cek jadwal kosong untuk Tanggal (1-31): ")
	fmt.Scan(&cariTanggal)

	var J arrJadwal
	nJadwal := 0

	// Buat jadwal untuk semua lapangan, dari jam 08:00 sd 22:00
	iLap := 0         //baris 449-473, mengeksekusi logika penyaringan "Filtering" menggunakan perulangan bersarang
	for iLap < nLap { //program menyimulasikan jam operasional dari semua lapangan dari (jam := 8 sampai 22)
		jam := 8        //program mencocokkan dengan data transaksi, jika pada tgl, bulan dan lapangan yang sama
		for jam <= 22 { //jam tersebut berada di dalam rentang waktu sewa pelanggan lain, statusnya akan = "true"
			dipakai := false
			// Cek apakah jam ini bentrok dengan transaksi
			iT := 0
			for iT < nTrans && !dipakai {
				if T[iT].bulan == cariBulan && T[iT].tanggal == cariTanggal && T[iT].IDlapang == L[iLap].IDlapang {
					if jam >= T[iT].jamMulai && jam < (T[iT].jamMulai+T[iT].durasi) {
						dipakai = true
					}
				}
				iT++
			}
			if !dipakai {
				J[nJadwal].jam = jam
				J[nJadwal].namaLap = L[iLap].nama
				J[nJadwal].harga = L[iLap].harga
				nJadwal++
			}
			jam++
		}
		iLap++
	}
	fmt.Println("\nPilih Metode Pengurutan:")
	fmt.Println("1. Berdasarkan Jam Mulai (Selection Sort - Ascending)")
	fmt.Println("2. Berdasarkan Harga (Insertion Sort - Descending)")
	fmt.Print("Pilih (1/2): ")

	var pilUrut int
	fmt.Scan(&pilUrut)

	if pilUrut == 1 { //baris 482-498, mengeksekusi algoritma "Selection Sort"
		// Selection Sort berdasarkan Jam
		i := 0 //program akan mengurutkan jadwal berdasarkan jam paling awal (Ascending)
		for i < nJadwal-1 {
			idxMin := i //Program mencari nilai jam terkecil dan menukar data terkecil
			j := i + 1  //ke urutan paling depan menggunakan bantuan variabel "temp"
			for j < nJadwal {
				if J[j].jam < J[idxMin].jam {
					idxMin = j
				}
				j++
			}
			temp := J[i]
			J[i] = J[idxMin]
			J[idxMin] = temp
			i++
		}
	} else if pilUrut == 2 { //baris 499-511, mengeksekusi algoritma Insertion Sort
		// Insertion Sort berdasarkan Harga
		i := 1            //program akan mengurutkan jadwal dari harga termurah dan program mengambil satu harga
		for i < nJadwal { //dan membandingkannya dengan data di sebelah kiri,
			temp := J[i] //jika harga lebih mahal di kiri data tersebut digeser ke kanan
			j := i - 1
			for j >= 0 && J[j].harga < temp.harga {
				J[j+1] = J[j]
				j--
			}
			J[j+1] = temp
			i++
		}
	}
	fmt.Println("\n--- HASIL JADWAL KOSONG ---") //baris 513-520, menampilkan hasil akhir
	i := 0                                       //program melakukan perulangan untuk mencetak seluruh jadwal
	for i < nJadwal {                            //yang sudah diurutkan tadi layar dan ditutup dengan fungsi "tahanLayar()"
		fmt.Printf("- Pukul %02d:00 | %s | Rp%d\n", J[i].jam, J[i].namaLap, J[i].harga)
		i++
	}
	tahanLayar()
}

// ==========================================
// STATISTIK (Poin e)
// ==========================================
func tampilStatistik(T arrTransaksi, n int) { //baris 525-531, mendeklarasikan fungsi dan menyiapkan variabel awal
	bersihkanLayar() //program membersihkan layar dan mencetak judul dan meminta admin memasukkan angka "Bulan" yg ingin dianalisis
	fmt.Println("--- STATISTIK PENDAPATAN ---")

	var cariBulan int //program juga menyiapkan variabel "totalPendapatan"  dan sebuah array "frekJam" berkapasitas 24 slot
	fmt.Print("Masukkan Bulan (1-12) untuk melihat statistik: ")
	fmt.Scan(&cariBulan) //untuk merecap frekuensi "berapa kali" setiap jam operasional disewa.

	totalPendapatan := 0 //baris 533-543, mengeksekusi perhitungan pendapatan dan perekaman jam sewa
	var frekJam [24]int  //Array untuk melacak jam berapa saja yang dipesan

	i := 0      //program menelusuri seluruh data transaksi menggunakan perulangan
	for i < n { //Jika bulan transaksinya sama dengan bulan yang dicari, maka nilai transaksinya ditambahkan ke "totalPendapatan"
		if T[i].bulan == cariBulan {
			totalPendapatan += T[i].total
			//perulangan di di "for k < T[i].durasi" yang bertugas mencatat jam penyewaan.
			jam := T[i].jamMulai
			k := 0 //
			for k < T[i].durasi {
				if (jam + k) < 24 { //baris 544, mengeksekusi algoritma pencarian nilai maksimum "Find Maximum"
					frekJam[jam+k]++
				}
				k++
			}
		}
		i++
	}
	jamTerlaris := 0
	maxOrder := 0
	i = 0
	for i < 24 { //baris 55-564, menampilkan hasil akhir laporan statistik
		if frekJam[i] > maxOrder {
			maxOrder = frekJam[i] //program mencetak Total Pendapatan Rupiah ke layar
			jamTerlaris = i       //program mengecek apakah ada penyewaan "jamTerlaris != -1" Jika ada
		} //program menampilkan Pukul berapa yang menjadi "Jam Paling Laris" beserta jumlah sewanya
		i++ //Jika tidak ada transaksi sama sekali di bulan tersebut, program mencetak pesan "Belum ada data penyewaan"
	} //dipanggil "tahanLayar()"" agar laporan tidak langsung tertutup sebelum dibaca
	fmt.Println("\n--- HASIL STATISTIK ---")
	fmt.Printf("Total Pendapatan Bulan %d : Rp %d\n", cariBulan, totalPendapatan)
	if maxOrder > 0 {
		fmt.Printf("Jam Paling Laris Bulan %d : Pukul %02d:00 (Disewa sebanyak %d jam)\n", cariBulan, jamTerlaris, maxOrder)
	} else {
		fmt.Println("Belum ada data transaksi di bulan ini.")
	}
	tahanLayar()
}

//====================================================================================================================================================================================
// | = ====================================================================== =|
// |                             // DATA DUMMY                                 |
// | = ====================================================================== =|
//nLap = 6
//dataL[0] = lapangan{IDlapang: 0, harga: 50000, nama: "lap0"}
//dataL[1] = lapangan{IDlapang: 1, harga: 100000, nama: "lap1"}
//dataL[2] = lapangan{IDlapang: 2, harga: 120000, nama: "lap2"}
//dataL[3] = lapangan{IDlapang: 3, harga: 150000, nama: "lap3"}
//dataL[4] = lapangan{IDlapang: 4, harga: 170000, nama: "lap4"}
//dataL[5] = lapangan{IDlapang: 5, harga: 200000, nama: "lap5"}

//nPenyewa = 6
//dataP[0] = dataPenyewa{IDpenyewa: 100, Nama: "andi", NoTlp: "0810"}
//dataP[1] = dataPenyewa{IDpenyewa: 101, Nama: "budi", NoTlp: "0811"}
//dataP[2] = dataPenyewa{IDpenyewa: 102, Nama: "cici", NoTlp: "0812"}
//dataP[3] = dataPenyewa{IDpenyewa: 103, Nama: "asep", NoTlp: "0813"}
//dataP[4] = dataPenyewa{IDpenyewa: 104, Nama: "siti", NoTlp: "0814"}
//dataP[5] = dataPenyewa{IDpenyewa: 105, Nama: "toni", NoTlp: "0815"}

//nTrans = 6
//dataT[0] = jadwalDANtransaksi{IDtransaksi: 0, IDlapang: 0, IDpenyewa: 100, bulan: 1, tanggal: 2, jamMulai: 8, durasi: 2, total: 100000}
//dataT[1] = jadwalDANtransaksi{IDtransaksi: 1, IDlapang: 1, IDpenyewa: 101, bulan: 2, tanggal: 4, jamMulai: 10, durasi: 1, total: 100000}
//dataT[2] = jadwalDANtransaksi{IDtransaksi: 2, IDlapang: 2, IDpenyewa: 102, bulan: 3, tanggal: 6, jamMulai: 11, durasi: 2, total: 240000}
//dataT[0] = jadwalDANtransaksi{IDtransaksi: 3, IDlapang: 3, IDpenyewa: 103, bulan: 4, tanggal: 8, jamMulai: 14, durasi: 2, total: 300000}
//dataT[1] = jadwalDANtransaksi{IDtransaksi: 4, IDlapang: 4, IDpenyewa: 104, bulan: 5, tanggal: 10, jamMulai: 15, durasi: 3, total: 510000}
//dataT[2] = jadwalDANtransaksi{IDtransaksi: 5, IDlapang: 5, IDpenyewa: 105, bulan: 6, tanggal: 3, jamMulai: 18, durasi: 2, total: 400000}
//====================================================================================================================================================================================
