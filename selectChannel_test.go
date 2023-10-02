package main

import (
	"fmt"
	"testing"
)

func TestSelectChannel(t *testing.T) {
	chanRequestor := make(chan string, 1)   // buat variabel channel yang menampung value requestor, capacity 1 value
	chanBeneficiary := make(chan string, 1) // buat variabel channel yang menampung value beneficiary, capacity 1 value

	defer func() { // defer supaya sebeum function ini selesai, akan menjalankan function yang di dalam defer
		close(chanBeneficiary) // close channel sebelum program selesai, supaya tidak terjadi goroutine leaks
		close(chanRequestor)   // close channel sebelum program selesai, supaya tidak terjadi goroutine leaks
	}()

	go GetDDMAST(1, chanRequestor)   // jalankan goroutine, untuk get data requestor yang akan disimpan ke dalam channel chanRequestor
	go GetDDMAST(2, chanBeneficiary) // jalankan goroutine

	counter := 0 // buat counter untuk mengontrol kapan for akan berhenti
	for {        // looping terus menerus
		select { // select case channel mana yang sudah ada valuenya duluan
		case requestor := <-chanRequestor: // case jika channel chanRequestor udah ada valuenya duluan, simpan valuenya ke variabel requestor
			fmt.Println("data dari channel requestor:", requestor) // tampilkan isi dari
			counter++                                              // value counter akan bertambah 1
		case beneficiary := <-chanBeneficiary: // case jika channel chanBeneficiary udah ada valuenya duluan, simpan valuenya ke variabel beneficiary
			fmt.Println("data dari channel beneficary:", beneficiary) // tampilkan value dari variabel beneficiary
			counter++                                                 // value counter akan bertambah 1
		}

		if counter == 2 { // kondisi jika value dari counter adalah 2
			break // break (stop looping) jika value counter sudah mencapai 2
		}
	}
}

// buat function yang digunakan untuk assign value ke dalan variabel channel
func GetDDMAST(id int, chanResult chan<- string) {
	result := fmt.Sprintf("success get data rekening %v", id) // buat value dengan tipe data string, tuliskan iinformasi id yang dipilih dalam parameter
	chanResult <- result                                      // assign channel dengan value yang sudah dibuat
}
