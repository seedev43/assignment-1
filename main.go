package main

import (
	"fmt"
	"os"
)

type Peserta struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

func dataPeserta() []Peserta {
	return []Peserta{
		{nama: "Budi", alamat: "Semarang", pekerjaan: "Pegawai kantor", alasan: "Lorem ipsum"},
		{nama: "Ahmad", alamat: "Madiun", pekerjaan: "Programmer", alasan: "Alasannya adalah"},
		{nama: "Saiful", alamat: "Kediri", pekerjaan: "Pengusaha", alasan: "Dolor sit amet"},
		{nama: "Wahyu", alamat: "Lamongan", pekerjaan: "Polisi", alasan: "ada alasan"},
	}

}

func cariData(data []Peserta, nama string) []int {
	input := os.Args[1]

	var simpanIndex []int

	for i, v := range data {
		if input == v.nama {
			simpanIndex = append(simpanIndex, i)
		}
	}
	return simpanIndex
}

func hasilData(index []int, data []Peserta) {
	if len(index) > 0 {
		for _, i := range index {
			peserta := data[i]
			fmt.Printf("ID: %d\nNama: %s\nAlamat: %s\nPekerjaan: %s\nAlasan: %s\n", i, peserta.nama, peserta.alamat, peserta.pekerjaan, peserta.alasan)
		}
	} else {
		fmt.Println("Nama", os.Args[1], "tidak ada")
	}
}

func main() {
	data := dataPeserta()
	cari := cariData(data, os.Args[1])
	hasilData(cari, data)

}
