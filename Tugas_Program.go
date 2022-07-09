package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const KAPASITAS_PRIORITAS = 36
const KAPASITAS_EKSEKUTIF = 48
const KAPASITAS_EKONOMI = 56

const HARGA_PRIORITAS = 6
const HARGA_EKSEKUTIF = 4
const HARGA_EKONOMI = 2

type penumpang struct {
	kodeTiket  int
	kelas      string
	id, nama   string
	usia       int
	nomorKursi string
	barisKursi int
}

type baris [4]penumpang

type gerbong struct {
	prioritas [9]baris
	eksekutif [12]baris
	ekonomi   [14]baris
}

type tabPenumpang struct {
	tab [140]penumpang
	N   int
}

type bayi struct {
	kodeTiket        int
	kelas            string
	idOrangTua, nama string
	usia             int
}

type tabBayi struct {
	tab [140]bayi
	N   int
}

type tarif struct {
	dewasa, anak float64
}

type rekomendasi struct {
	baris, nomor int
}

type tabRekomendasi struct {
	tab [10]rekomendasi
	N   int
}

func isFull(gerbongYangDiperiksa string, dataPenumpang tabPenumpang) bool {

	// {Mengembalikan true jika gerbong yang diperiksa penuh}

	return sisaKursi(gerbongYangDiperiksa, dataPenumpang) == 0
}

func sisaKursi(gerbongYangDiperiksa string, dataPenumpang tabPenumpang) int {

	// {Mengembalikan banyak kursi kosong dari gerbong yang diperiksa}

	var i int
	var countTerisi int
	for i = 0; i < dataPenumpang.N; i++ {
		if dataPenumpang.tab[i].kelas == gerbongYangDiperiksa {
			countTerisi++
		}
	}

	var answer int

	switch gerbongYangDiperiksa {
	case "prioritas":
		answer = KAPASITAS_PRIORITAS - countTerisi
	case "eksekutif":
		answer = KAPASITAS_EKSEKUTIF - countTerisi
	case "ekonomi":
		answer = KAPASITAS_EKONOMI - countTerisi
	}

	return answer
}

func nomorToStr(nomorInInt int) string {

	// {Konversi nomor kursi dari bilangan bulat ke string}

	var answer string

	switch nomorInInt {
	case 0:
		answer = "a"
	case 1:
		answer = "b"
	case 2:
		answer = "c"
	case 3:
		answer = "d"
	}

	return answer
}

func nomorToInt(nomorInStr string) int {

	// {Konversi nomor kursi dari string ke bilangan bulat}

	var answer int

	switch nomorInStr {
	case "a":
		answer = 0
	case "b":
		answer = 1
	case "c":
		answer = 2
	case "d":
		answer = 3
	}

	return answer
}

func statusKursi(dataPenumpang tabPenumpang, gerbongYangDipilih string, baris int, nomor int) string {

	// {mengembalikan “terisi” jika sudah ada penumpang yang memesan kursi tertentu
	// dan mengembalikan “kosong” jika belum ada penumpang yang memesan kursi tertentu}

	var i int

	for i = 0; i < dataPenumpang.N; i++ {
		if dataPenumpang.tab[i].nomorKursi == nomorToStr(nomor) && dataPenumpang.tab[i].barisKursi-1 == baris &&
			dataPenumpang.tab[i].kelas == gerbongYangDipilih {
			return "terisi"
		}
	}

	return "kosong"
}

func mencariKursi(dataPenumpang tabPenumpang, rekomendasiKursi *tabRekomendasi, gerbongYangDipilih string,
	banyakPemesan int, currentBaris, currentNomor *int) {

	// {I.S. rekomendasiKursi berfungsi untuk menyimpan kursi-kursi yang sudah ditawarkan  sebelumnya kepada pemesan
	// berdasarkan gerbongYangDipilih dengan memerhatikan kursi-kursi yang telah diisi oleh pemesan sebelumnya,
	// currentBaris dan currentNomor merupakan kursi terakhir yang telah direkomendasikan kepada pemesan terbaru}
	// F.S. rekomendasiKursi menyimpan kursi-kursi yang direkomendasikan kepada pemesan yang mengubah kursi,
	// currentBaris dan currentNomor diperbarui}

	var i, j, countRekomendasi int

	rekomendasiKursi.N = banyakPemesan

	switch gerbongYangDipilih {
	case "prioritas":

		countRekomendasi = 0
		for i = *currentBaris; i < 9 && countRekomendasi < banyakPemesan; i++ {
			for j = *currentNomor; j < 4 && countRekomendasi < banyakPemesan; j++ {
				if statusKursi(dataPenumpang, gerbongYangDipilih, i, j) == "kosong" {
					rekomendasiKursi.tab[countRekomendasi].baris = i + 1
					rekomendasiKursi.tab[countRekomendasi].nomor = j
					countRekomendasi++
				}
				*currentNomor = (*currentNomor + 1) % 4
				if j == 3 {
					*currentBaris++
				}
			}
		}

	case "eksekutif":

		countRekomendasi = 0
		for i = *currentBaris; i < 12 && countRekomendasi < banyakPemesan; i++ {
			for j = *currentNomor; j < 4 && countRekomendasi < banyakPemesan; j++ {
				if statusKursi(dataPenumpang, gerbongYangDipilih, i, j) == "kosong" {
					rekomendasiKursi.tab[countRekomendasi].baris = i + 1
					rekomendasiKursi.tab[countRekomendasi].nomor = j
					countRekomendasi++
				}
				*currentNomor = (*currentNomor + 1) % 4
				if j == 3 {
					*currentBaris++
				}
			}
		}

	case "ekonomi":

		countRekomendasi = 0
		for i = *currentBaris; i < 14 && countRekomendasi < banyakPemesan; i++ {
			for j = *currentNomor; j < 4 && countRekomendasi < banyakPemesan; j++ {
				if statusKursi(dataPenumpang, gerbongYangDipilih, i, j) == "kosong" {
					rekomendasiKursi.tab[countRekomendasi].baris = i + 1
					rekomendasiKursi.tab[countRekomendasi].nomor = j
					countRekomendasi++
				}
				*currentNomor = (*currentNomor + 1) % 4
				if j == 3 {
					*currentBaris++
				}
			}
		}
	}

}

func ubahKursi(dataPenumpang tabPenumpang, rekomendasiKursi *tabRekomendasi, gerbongYangDipilih string,
	banyakPemesan int, currentBaris, currentNomor *int) {

	// {I.S. rekomendasiKursi berfungsi untuk menyimpan kursi-kursi yang sudah ditawarkan  sebelumnya kepada pemesan
	// berdasarkan gerbongYangDipilih dengan memerhatikan kursi-kursi yang telah diisi oleh pemesan sebelumnya,
	// currentBaris dan currentNomor merupakan kursi terakhir yang telah direkomendasikan kepada pemesan terbaru}
	// F.S. rekomendasiKursi menyimpan kursi-kursi yang direkomendasikan kepada pemesan yang mengubah kursi,
	// currentBaris dan currentNomor diperbarui}

	mencariKursi(dataPenumpang, rekomendasiKursi, gerbongYangDipilih, banyakPemesan, currentBaris, currentNomor)
}

func inputKursiPenumpang(rekomendasiKursi tabRekomendasi, gerbongKereta *gerbong, gerbongYangDipilih string, dataPenumpang *tabPenumpang) {

	// {I.S. gerbongKereta dan dataPenumpang berisi data penumpang yang telah memesan sebelum pemesan terbaru
	// F.S. gerbongKereta dan dataPenumpang berisi data semua penumpang termasuk pemesan terbaru}

	var i int

	switch gerbongYangDipilih {
	case "prioritas":

		for i = 0; i < rekomendasiKursi.N; i++ {
			gerbongKereta.prioritas[rekomendasiKursi.tab[i].baris-1][rekomendasiKursi.tab[i].nomor] = dataPenumpang.tab[dataPenumpang.N-i-1]
			dataPenumpang.tab[dataPenumpang.N-i-1].barisKursi = rekomendasiKursi.tab[i].baris
			dataPenumpang.tab[dataPenumpang.N-i-1].nomorKursi = nomorToStr(rekomendasiKursi.tab[i].nomor)
		}

	case "eksekutif":

		for i = 0; i < rekomendasiKursi.N; i++ {
			gerbongKereta.eksekutif[rekomendasiKursi.tab[i].baris-1][rekomendasiKursi.tab[i].nomor] = dataPenumpang.tab[dataPenumpang.N-i-1]
			dataPenumpang.tab[dataPenumpang.N-i-1].barisKursi = rekomendasiKursi.tab[i].baris
			dataPenumpang.tab[dataPenumpang.N-i-1].nomorKursi = nomorToStr(rekomendasiKursi.tab[i].nomor)
		}

	case "ekonomi":

		for i = 0; i < rekomendasiKursi.N; i++ {
			gerbongKereta.ekonomi[rekomendasiKursi.tab[i].baris-1][rekomendasiKursi.tab[i].nomor] = dataPenumpang.tab[dataPenumpang.N-i-1]
			dataPenumpang.tab[dataPenumpang.N-i-1].barisKursi = rekomendasiKursi.tab[i].baris
			dataPenumpang.tab[dataPenumpang.N-i-1].nomorKursi = nomorToStr(rekomendasiKursi.tab[i].nomor)
		}
	}
}

func totalHarga(banyakPemesan, kodeTiket int, dataPemesan tabPenumpang, gerbong string) float64 {

	// {Mengembalikan total harga pada suatu reservasi berdasarkan promo-promo yang ada}

	var i, banyakAnak, banyakDewasa int
	var totalHargaTiket, konstantaPengaliTanpaSaving, konstantaPengaliDenganSaving, konstantaPengaliAkhir float64
	banyakAnak = 0
	banyakDewasa = 0
	totalHargaTiket = 0

	for i = 0; i < banyakPemesan; i++ {
		if dataPemesan.tab[dataPemesan.N-i-1].usia >= 18 {
			banyakDewasa++
		} else if dataPemesan.tab[dataPemesan.N-i-1].usia > 2 && dataPemesan.tab[dataPemesan.N-i-1].usia < 18 {
			banyakAnak++
		}
	}

	fmt.Println("Gerbong :", gerbong)

	konstantaPengaliTanpaSaving = 0.5*float64(banyakAnak) + float64(banyakDewasa)

	if banyakPemesan < 3 {
		konstantaPengaliDenganSaving = 10
	} else if banyakPemesan == 5 {
		if banyakAnak >= 1 {
			konstantaPengaliDenganSaving = 2.5
		} else {
			konstantaPengaliDenganSaving = 3
		}
	} else if banyakPemesan == 6 {
		if banyakAnak >= 2 {
			konstantaPengaliDenganSaving = 3
		} else if banyakAnak == 1 {
			konstantaPengaliDenganSaving = 3.5
		} else {
			konstantaPengaliDenganSaving = 4
		}
	} else if banyakPemesan == 7 {
		if banyakAnak >= 3 {
			konstantaPengaliDenganSaving = 3.5
		} else {
			konstantaPengaliDenganSaving = 4
		}

	} else if banyakPemesan == 9 {
		if banyakAnak >= 1 {
			konstantaPengaliDenganSaving = 4.5
		} else {
			konstantaPengaliDenganSaving = 5
		}
	} else if banyakPemesan == 10 {
		if banyakAnak >= 2 {
			konstantaPengaliDenganSaving = 5
		} else if banyakAnak == 1 {
			konstantaPengaliDenganSaving = 5.5
		} else {
			konstantaPengaliDenganSaving = 6
		}
	} else {
		konstantaPengaliDenganSaving = 2 * float64(banyakPemesan/3)
	}

	if konstantaPengaliDenganSaving < konstantaPengaliTanpaSaving {
		konstantaPengaliAkhir = konstantaPengaliDenganSaving
	} else {
		konstantaPengaliAkhir = konstantaPengaliTanpaSaving
	}

	switch gerbong {
	case "prioritas":
		totalHargaTiket = konstantaPengaliAkhir * HARGA_PRIORITAS
		break
	case "eksekutif":
		totalHargaTiket = konstantaPengaliAkhir * HARGA_EKSEKUTIF
		break
	case "ekonomi":
		totalHargaTiket = konstantaPengaliAkhir * HARGA_EKONOMI
		break
	}

	fmt.Println("Banyak Penumpang Dewasa :", banyakDewasa)
	fmt.Println("Banyak Penumpang Anak-Anak :", banyakAnak)

	return totalHargaTiket
}

func reservasi(kodeTiket *int, dataPenumpang *tabPenumpang, dataBayi *tabBayi, gerbongKereta *gerbong) {

	// {I.S. kodeTiket berisi kode unik yang didapatkan pemesan ketika melakukan reservasi,
	// dataBayi menyimpan data semua bayi dan dataPenumpang menyimpan data semua penumpang anak-anak sampai remaja yang sudah
	// 	melakukan reservasi, gerbongKereta digunakan untuk menyimpan kursi-kursi yang telah diisi}
	// F.S. data pemesan yang baru saja melakukan reservasi masuk ke dalam dataPenumpang atau dataBayi atau gerbongKereta,
	// dan nilai kode tiket bertambah satu}

	var banyakPemesan, usia int
	var currentBaris int = 0
	var currentNomor int = 0
	var banyakPemesanBayi int
	var gerbongYangDipilih, inginUbahKursi string
	var rekomendasiKursi tabRekomendasi
	var id, nama string

	fmt.Print("Masukkan banyak tiket yang ingin dibeli (maksimal 10) : ")
	fmt.Scan(&banyakPemesan)

	fmt.Println("Gerbong yang belum penuh : ")
	if !isFull("prioritas", *dataPenumpang) {
		fmt.Println("- Prioritas")
	}
	if !isFull("eksekutif", *dataPenumpang) {
		fmt.Println("- Eksekutif")
	}
	if !isFull("ekonomi", *dataPenumpang) {
		fmt.Println("- Ekonomi")
	}

	fmt.Print("Pilih Gerbong : ")
	fmt.Scan(&gerbongYangDipilih)
	gerbongYangDipilih = strings.ToLower(gerbongYangDipilih)

	var i int
	for i = 0; i < banyakPemesan; i++ {
		fmt.Print("Masukkan nomor identitas (nomor identitas orang tua bagi penumpang berusia kurang dari atau sama dengan dua tahun), nama, dan usia anda : ")
		fmt.Scan(&id, &nama, &usia)

		if usia > 2 {
			dataPenumpang.tab[dataPenumpang.N].id = id
			dataPenumpang.tab[dataPenumpang.N].nama = nama
			dataPenumpang.tab[dataPenumpang.N].usia = usia
			dataPenumpang.tab[dataPenumpang.N].kelas = gerbongYangDipilih
			dataPenumpang.tab[dataPenumpang.N].kodeTiket = *kodeTiket
			dataPenumpang.N++
		} else {
			dataBayi.tab[dataBayi.N].idOrangTua = id
			dataBayi.tab[dataBayi.N].nama = nama
			dataBayi.tab[dataBayi.N].usia = usia
			dataBayi.tab[dataBayi.N].kelas = gerbongYangDipilih
			dataBayi.tab[dataBayi.N].kodeTiket = *kodeTiket
			dataBayi.N++
			banyakPemesanBayi++
		}
	}

	menampilkanKetersediaanKursi(*gerbongKereta, gerbongYangDipilih)

	mencariKursi(*dataPenumpang, &rekomendasiKursi, gerbongYangDipilih, banyakPemesan-banyakPemesanBayi, &currentBaris, &currentNomor)

	fmt.Println("Rekomendasi Kursi :")
	for i = 0; i < rekomendasiKursi.N; i++ {
		fmt.Printf("(%d) baris : %d, nomor : %s\n", i+1, rekomendasiKursi.tab[i].baris, nomorToStr(rekomendasiKursi.tab[i].nomor))
	}

	fmt.Print("Ubah kursi? ya | tidak : ")
	fmt.Scan(&inginUbahKursi)
	for inginUbahKursi == "ya" {
		ubahKursi(*dataPenumpang, &rekomendasiKursi, gerbongYangDipilih, banyakPemesan-banyakPemesanBayi, &currentBaris, &currentNomor)
		fmt.Println("Rekomendasi Kursi :")
		for i = 0; i < rekomendasiKursi.N; i++ {
			fmt.Printf("(%d) baris : %d, nomor : %s\n", i+1, rekomendasiKursi.tab[i].baris, nomorToStr(rekomendasiKursi.tab[i].nomor))
		}

		fmt.Print("Ubah kursi? ya | tidak : ")
		fmt.Scan(&inginUbahKursi)
	}

	inputKursiPenumpang(rekomendasiKursi, gerbongKereta, gerbongYangDipilih, dataPenumpang)
	fmt.Println(" ")
	fmt.Println("Total harga yang harus di bayar: ", totalHarga(banyakPemesan-banyakPemesanBayi, *kodeTiket, *dataPenumpang, gerbongYangDipilih))
	*kodeTiket++
}

func cetakDatabayi(dataBayi tabBayi, dataPenumpang tabPenumpang) {

	// {I.S.-
	// F.S. Menampilkan data penumpang bayi dan tempat duduknya di tiap gerbong}

	if dataBayi.N == 0 {
		fmt.Println("Belum ada penumpang bayi di semua gerbong")
		fmt.Print("\n")
	}

	for i := 0; i < dataBayi.N; i++ {

		for j := 0; j <= dataPenumpang.N; j++ {

			if dataBayi.tab[i].idOrangTua == dataPenumpang.tab[j].id && dataBayi.tab[i].kelas == "prioritas" {
				if i == 0 {
					fmt.Println("========Penumpang Bayi di Gerbong Prioritas========")
					fmt.Print("\n")
				}
				fmt.Printf("Penumpang bayi ke-%d : \n", i+1)
				fmt.Println("Kode Tiket : ", dataBayi.tab[i].kodeTiket)
				fmt.Println("Kelas : ", dataBayi.tab[i].kelas)
				fmt.Println("Nomor Identitas Orang Tua: ", dataBayi.tab[i].idOrangTua)
				fmt.Println("Nama : ", dataBayi.tab[i].nama)
				fmt.Println("Usia : ", dataBayi.tab[i].usia)
				fmt.Println("Duduk di kursi baris ke-", dataPenumpang.tab[j].barisKursi)
				fmt.Println("Duduk di kuris nomor", dataPenumpang.tab[j].nomorKursi)
				fmt.Print("\n")
			}
		}
	}

	for i := 0; i < dataBayi.N; i++ {

		for j := 0; j <= dataPenumpang.N; j++ {

			if dataBayi.tab[i].idOrangTua == dataPenumpang.tab[j].id && dataBayi.tab[i].kelas == "eksekutif" {
				if i == 0 {
					fmt.Println("========Penumpang Bayi di Gerbong Eksekutif========")
					fmt.Print("\n")
				}
				fmt.Printf("Penumpang bayi ke-%d : \n", i+1)
				fmt.Println("Kode Tiket : ", dataBayi.tab[i].kodeTiket)
				fmt.Println("Kelas : ", dataBayi.tab[i].kelas)
				fmt.Println("Nomor Identitas Orang Tua: ", dataBayi.tab[i].idOrangTua)
				fmt.Println("Nama : ", dataBayi.tab[i].nama)
				fmt.Println("Usia : ", dataBayi.tab[i].usia)
				fmt.Println("Duduk di kursi baris ke-", dataPenumpang.tab[j].barisKursi)
				fmt.Println("Duduk di kuris nomor", dataPenumpang.tab[j].nomorKursi)
				fmt.Print("\n")
			}
		}
	}

	for i := 0; i < dataBayi.N; i++ {

		for j := 0; j <= dataPenumpang.N; j++ {

			if dataBayi.tab[i].idOrangTua == dataPenumpang.tab[j].id && dataBayi.tab[i].kelas == "ekonomi" {
				if i == 0 {
					fmt.Println("========Penumpang Bayi di Gerbong Ekonomi========")
					fmt.Print("\n")
				}
				fmt.Printf("Penumpang bayi ke-%d : \n", i+1)
				fmt.Println("Kode Tiket : ", dataBayi.tab[i].kodeTiket)
				fmt.Println("Kelas : ", dataBayi.tab[i].kelas)
				fmt.Println("Nomor Identitas Orang Tua: ", dataBayi.tab[i].idOrangTua)
				fmt.Println("Nama : ", dataBayi.tab[i].nama)
				fmt.Println("Usia : ", dataBayi.tab[i].usia)
				fmt.Println("Duduk di kursi baris ke-", dataPenumpang.tab[j].barisKursi)
				fmt.Println("Duduk di kuris nomor", dataPenumpang.tab[j].nomorKursi)
				fmt.Print("\n")
			}
		}
	}
}

func cetakSemuaData(dataBayi tabBayi, dataPenumpang tabPenumpang) {

	// {I.S. -
	// F.S. menampilkan semua data penumpang dan tempat duduknya terurut dari usia yang paling muda}

	var pass1, pass2, iMin1, iMin2, i1, i2 int
	var temp1 bayi
	var temp2 penumpang

	for pass1 = 1; pass1 <= dataBayi.N-1; pass1++ {
		iMin1 = pass1 - 1
		for i1 = pass1; i1 <= dataBayi.N-1; i1++ {
			if dataBayi.tab[iMin1].usia > dataBayi.tab[i1].usia {
				iMin1 = i1
			}
		}
		temp1 = dataBayi.tab[iMin1]
		dataBayi.tab[iMin1] = dataBayi.tab[pass1-1]
		dataBayi.tab[pass1-1] = temp1
	}

	for pass2 = 1; pass2 <= dataPenumpang.N-1; pass2++ {
		iMin2 = pass2 - 1
		for i2 = pass2; i2 <= dataPenumpang.N-1; i2++ {
			if dataPenumpang.tab[iMin2].usia > dataPenumpang.tab[i2].usia {
				iMin2 = i2
			}
		}
		temp2 = dataPenumpang.tab[iMin2]
		dataPenumpang.tab[iMin2] = dataPenumpang.tab[pass2-1]
		dataPenumpang.tab[pass2-1] = temp2
	}

	if dataBayi.N == 0 {
		fmt.Println("Belum ada penumpang bayi")
		fmt.Print("\n")
	} else {
		fmt.Println("========Data Bayi========")
		fmt.Print("\n")
		for i := 0; i < dataBayi.N; i++ {
			fmt.Println("Penumpang bayi ke-" + strconv.Itoa(i+1))
			fmt.Println("Kode Tiket : ", dataBayi.tab[i].kodeTiket)
			fmt.Println("Kelas : ", dataBayi.tab[i].kelas)
			fmt.Println("Nomor Identitas Orang Tua: ", dataBayi.tab[i].idOrangTua)
			fmt.Println("Nama : ", dataBayi.tab[i].nama)
			fmt.Println("Usia : ", dataBayi.tab[i].usia)
			fmt.Print("\n")
		}
	}

	fmt.Print("\n")

	if dataPenumpang.N == 0 {
		fmt.Println("Belum ada penumpang")
		fmt.Print("\n")
	} else {
		fmt.Println("========Data Penumpang Anak-Anak atau Dewasa========")
		fmt.Print("\n")
		for i := 0; i < dataPenumpang.N; i++ {
			fmt.Println("Penumpang anak-anak atau dewasa ke-" + strconv.Itoa(i+1))
			fmt.Println("Kode Tiket : ", dataPenumpang.tab[i].kodeTiket)
			fmt.Println("Kelas : ", dataPenumpang.tab[i].kelas)
			fmt.Println("Nomor Identitas : ", dataPenumpang.tab[i].id)
			fmt.Println("Nama : ", dataPenumpang.tab[i].nama)
			fmt.Println("Usia : ", dataPenumpang.tab[i].usia)
			fmt.Printf("Duduk di kursi baris ke-%d\n", dataPenumpang.tab[i].barisKursi)
			fmt.Println("Duduk di kursi nomor", dataPenumpang.tab[i].nomorKursi)
			fmt.Print("\n")
		}
	}

}

func catatData(dataBayi tabBayi, dataPenumpang tabPenumpang) {

	// {I.S. -
	// F.S. Mencatat data penumpang pada sebuah file teks}

	var pass1, pass2, iMin1, iMin2, i1, i2 int
	var temp1 bayi
	var temp2 penumpang

	for pass1 = 1; pass1 <= dataBayi.N-1; pass1++ {
		iMin1 = pass1 - 1
		for i1 = pass1; i1 <= dataBayi.N-1; i1++ {
			if dataBayi.tab[iMin1].usia > dataBayi.tab[i1].usia {
				iMin1 = i1
			}
		}
		temp1 = dataBayi.tab[iMin1]
		dataBayi.tab[iMin1] = dataBayi.tab[pass1-1]
		dataBayi.tab[pass1-1] = temp1
	}

	for pass2 = 1; pass2 <= dataPenumpang.N-1; pass2++ {
		iMin2 = pass2 - 1
		for i2 = pass2; i2 <= dataPenumpang.N-1; i2++ {
			if dataPenumpang.tab[iMin2].usia > dataPenumpang.tab[i2].usia {
				iMin2 = i2
			}
		}
		temp2 = dataPenumpang.tab[iMin2]
		dataPenumpang.tab[iMin2] = dataPenumpang.tab[pass2-1]
		dataPenumpang.tab[pass2-1] = temp2
	}

	var file, _ = os.OpenFile("Data Penumpang.txt", os.O_RDWR, 0644)
	defer file.Close()

	if dataBayi.N == 0 {
		file.WriteString("Belum ada penumpang bayi\n")
		file.WriteString("\n")
	} else {
		file.WriteString("========Data Bayi========\n")
		file.WriteString("\n")
		for i := 0; i < dataBayi.N; i++ {
			file.WriteString("Penumpang bayi ke-" + strconv.Itoa(i+1) + "\n")
			file.WriteString("Kode Tiket : " + strconv.Itoa(dataBayi.tab[i].kodeTiket) + "\n")
			file.WriteString("Kelas : " + dataBayi.tab[i].kelas + "\n")
			file.WriteString("Nomor Identitas Orang Tua :" + dataBayi.tab[i].idOrangTua + "\n")
			file.WriteString("Nama : " + dataBayi.tab[i].nama + "\n")
			file.WriteString("Usia : " + strconv.Itoa(dataBayi.tab[i].usia) + "\n")
			file.WriteString("\n")
		}
	}

	if dataPenumpang.N == 0 {
		file.WriteString("Belum ada penumpang\n")
		file.WriteString("\n")
	} else {
		file.WriteString("========Data Penumpang Anak-Anak atau Dewasa========\n")
		file.WriteString("\n")
		for i := 0; i < dataPenumpang.N; i++ {
			file.WriteString("Penumpang Anak-anak atau Dewasa ke-" + strconv.Itoa(i+1) + "\n")
			file.WriteString("Kode Tiket : " + strconv.Itoa(dataPenumpang.tab[i].kodeTiket) + "\n")
			file.WriteString("Kelas : " + dataPenumpang.tab[i].kelas + "\n")
			file.WriteString("Nomor Identitas : " + dataPenumpang.tab[i].id + "\n")
			file.WriteString("Nama : " + dataPenumpang.tab[i].nama + "\n")
			file.WriteString("Usia : " + strconv.Itoa(dataPenumpang.tab[i].usia) + "\n")
			file.WriteString("Baris Kursi : " + strconv.Itoa(dataPenumpang.tab[i].barisKursi) + "\n")
			file.WriteString("Nomor Kursi : " + dataPenumpang.tab[i].nomorKursi + "\n")
			file.WriteString("\n")
		}
	}

	file.Sync()
}

func cetakStatistik(dataPenumpang tabPenumpang) {

	//	{I.S. -
	//	F.S. Mencetak prosentasi okupansi masing-masing gerbong}

	var isiGerbong int
	var persentase float64
	var stringPersentase string

	isiGerbong = KAPASITAS_PRIORITAS - sisaKursi("prioritas", dataPenumpang)
	persentase = (float64(isiGerbong) * 100) / KAPASITAS_PRIORITAS
	fmt.Print("Persentase okupansi gerbong prioritas : ")
	stringPersentase = fmt.Sprintf("%.2f", persentase)
	fmt.Print(stringPersentase, "%\n")

	isiGerbong = KAPASITAS_EKSEKUTIF - sisaKursi("eksekutif", dataPenumpang)
	persentase = (float64(isiGerbong) * 100) / KAPASITAS_EKSEKUTIF
	fmt.Print("Persentase okupansi gerbong eksekutif : ")
	stringPersentase = fmt.Sprintf("%.2f", persentase)
	fmt.Print(stringPersentase, "%\n")

	isiGerbong = KAPASITAS_EKONOMI - sisaKursi("ekonomi", dataPenumpang)
	persentase = (float64(isiGerbong) * 100) / KAPASITAS_EKONOMI
	fmt.Print("Persentase okupansi gerbong ekonomi : ")
	stringPersentase = fmt.Sprintf("%.2f", persentase)
	fmt.Print(stringPersentase, "%\n")

}

func cariOrangTuaBayi(dataBayi tabBayi, namaBayi string, kodeTiket int) string {

	// {Mengembalikan nomor identitas orang tua bayi berdasarkan nama bayi dan kode tiket bayi}

	var i int

	for i = 0; i < dataBayi.N; i++ {
		if dataBayi.tab[i].nama == namaBayi && dataBayi.tab[i].kodeTiket == kodeTiket {
			return dataBayi.tab[i].idOrangTua
		}
	}
	return "Data Bayi Tidak Ditemukan"
}

func menampilkanKetersediaanKursi(gerbongKereta gerbong, gerbongYangDipilih string) {

	// 	{I.S. -
	// 	F.S. Menampilkan kursi-kursi yang tersedia pada gerbong yang dipilih}

	switch gerbongYangDipilih {
	case "prioritas":
		fmt.Println("========Gerbong Prioritas========")
		for i := 0; i < 9; i++ {
			fmt.Printf("Baris %d : ", i+1)
			for j := 0; j < 4; j++ {
				if gerbongKereta.prioritas[i][j].kodeTiket > 0 {
					fmt.Print("Terisi ")
				} else {
					fmt.Print("Kosong ")
				}
				if j == 3 {
					fmt.Print("\n")
				}
			}
		}
		break
	case "eksekutif":
		fmt.Println("========Gerbong Eksekutif========")
		for i := 0; i < 12; i++ {
			fmt.Printf("Baris %d : ", i+1)
			for j := 0; j < 4; j++ {
				if gerbongKereta.eksekutif[i][j].kodeTiket > 0 {
					fmt.Print("Terisi ")
				} else {
					fmt.Print("Kosong ")
				}
				if j == 3 {
					fmt.Print("\n")
				}
			}
		}
		break
	case "ekonomi":
		fmt.Println("========Gerbong Ekonomi========")
		for i := 0; i < 14; i++ {
			fmt.Printf("Baris %d : ", i+1)
			for j := 0; j < 4; j++ {
				if gerbongKereta.ekonomi[i][j].kodeTiket > 0 {
					fmt.Print("Terisi ")
				} else {
					fmt.Print("Kosong ")
				}
				if j == 3 {
					fmt.Print("\n")
				}
			}
		}
		break
	}

}

func main() {
	var kodeTiket int = 1
	var dataPenumpang tabPenumpang
	var dataBayi tabBayi
	var gerbongKereta gerbong
	var menuPilihan int
	var berjalan bool
	var ulang int

	berjalan = true
	for berjalan {
		fmt.Println("Selamat datang di Program Kereta Bandung-Jakarta")
		fmt.Println("1 : Reservasi")
		fmt.Println("2 : Cek persentase okupansi gerbong")
		fmt.Println("3 : Menampilkan semua penumpang terurut berdasarkan usia dari yang paling muda")
		fmt.Println("4 : Menampilkan semua data bayi pada masing-masing gerbong")
		fmt.Println("5 : Mencari orang tua bayi")
		fmt.Println("6 : Mencatat data ke file teks")
		fmt.Println("0 : Keluar")
		fmt.Print("Masukkan nomor di atas untuk memilih menu : ")
		fmt.Scan(&menuPilihan)
		fmt.Println(" ")
		for menuPilihan < 0 || menuPilihan > 6 {
			fmt.Println("Anda memilih nomor yang salah, coba lagi")
			fmt.Println("1 : Reservasi")
			fmt.Println("2 : Cek persentase okupansi gerbong")
			fmt.Println("3 : Menampilkan semua penumpang terurut berdasarkan usia dari yang paling muda per kategori (bayi atau bukan bayi)")
			fmt.Println("4 : Menampilkan semua data bayi pada masing-masing gerbong")
			fmt.Println("5 : Mencari orang tua bayi")
			fmt.Println("6 : Mencatat data ke file teks")
			fmt.Println("0 : Keluar")
			fmt.Print("Masukkan nomor di atas untuk memilih menu : ")
			fmt.Scan(&menuPilihan)
			fmt.Println(" ")
		}

		if menuPilihan == 1 {
			reservasi(&kodeTiket, &dataPenumpang, &dataBayi, &gerbongKereta)
		} else if menuPilihan == 2 {
			cetakStatistik(dataPenumpang)
		} else if menuPilihan == 3 {
			cetakSemuaData(dataBayi, dataPenumpang)
		} else if menuPilihan == 4 {
			cetakDatabayi(dataBayi, dataPenumpang)
		} else if menuPilihan == 5 {
			var namaBayi string
			var kodeTiketBayi int

			fmt.Print("Masukkan nama bayi : ")
			fmt.Scan(&namaBayi)
			fmt.Print("Masukkan kode tiket : ")
			fmt.Scan(&kodeTiketBayi)
			fmt.Println("Nomor Identitas Orang Tua Bayi :", cariOrangTuaBayi(dataBayi, namaBayi, kodeTiketBayi))

			for i := 0; i < dataPenumpang.N; i++ {
				if dataPenumpang.tab[i].id == cariOrangTuaBayi(dataBayi, namaBayi, kodeTiketBayi) {
					fmt.Println("Nama Orang Tua Bayi :", dataPenumpang.tab[i].nama)
				}
			}
		} else if menuPilihan == 6 {
			catatData(dataBayi, dataPenumpang)
		} else {
			fmt.Println("Terima kasih")
			fmt.Print("\n")
			catatData(dataBayi, dataPenumpang)
			break
		}

		fmt.Println(" ")
		fmt.Println("Ketik 1 untuk kembali ke menu awal")
		fmt.Println("Ketik 0 untuk keluar")
		fmt.Print("Pilihan anda: ")
		fmt.Scan(&ulang)
		if ulang == 0 {
			fmt.Println("Terima Kasih")
			catatData(dataBayi, dataPenumpang)
			berjalan = false
		}
		fmt.Print("\n")
	}
}
