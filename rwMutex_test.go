package main

import (
	"fmt"
	"sync"
	"testing"
)

type BankAccount struct { // buat sebuah struct yang merepresentasikan object data rekening dan saldo
	AccountNumber string        // property value nomor rekeking
	CBal          int           // property value  saldo yang ada
	Mtx           *sync.RWMutex // property rwmutex untuk locing ketika get dan update balance
}

// method untuk melakukan update saldo
func (b *BankAccount) UpdateBalance(inputAmount int) {
	defer b.Mtx.Unlock() // diakhir sebelum mehtod ini selesai dijalankan, jangan lupa unlock mutexnya

	b.Mtx.Lock()          // lock mute`xnya sebelum operasi update data saldo dilakukan
	b.CBal += inputAmount // update data saldo dengan menjumlahkan saldo sekarang + amount
}

// method untuk
func (b *BankAccount) GetCBal() int {
	defer b.Mtx.RUnlock() // diakhir sebelum method selesai dijalankan, jangan lupa untuk unloc mutex read nya

	b.Mtx.RLock()  // lakukan lock pada mutex read sebelum proses read variabel cbal dilakukan
	cbal := b.CBal // get/read data property cbal
	return cbal    // returnkan value dari cbal
}

// function test untuk menggunakan RWMutex
func TestRWMutex(t *testing.T) {
	myAccount := &BankAccount{ // buat object dari struct BankAccount dengan nama myAccount
		AccountNumber: "045202000001809", // isi property value nomor rekening
		CBal:          0,                 // isi property value saldo
		Mtx:           &sync.RWMutex{},   // isi property value remutex
	}

	wg := &sync.WaitGroup{} // buat variabel waitgroup untuk menunggu semua goroutine yang sedang berjalan, dan memastikan semuanya sudah selesai dijalankan

	for i := 0; i < 1000; i++ { // buat looping sebanyak 1000 kali perulangan
		wg.Add(1)   // beri tahu waitgroup bawah ada 1 goroutine yang bertambah di setiap perulangan
		go func() { // jalankan goroutine menggunakan anonymous function, bakal ada 1000 goroutine yang jalan
			defer wg.Done() // diakhir sebelum anonymous function selesai dijalankan, beri tahu ke waitgroup bahwa goroutine ini sudah selesai dijalankan

			myAccount.UpdateBalance(100)                        // lakukan update pada saldo dengan menambah saldo saat ini + 100
			fmt.Println("saldo sesudah :", myAccount.GetCBal()) // tampilkan info saldo setelah diupdate
		}()
	}

	wg.Wait()                                        // tunggu waitgroup menerima informasi bahwa semua goroutine sudah selesai dikerjakan
	fmt.Println("total saldo:", myAccount.GetCBal()) // tampilkan info total saldo setelah diupdate, harusnya 100000
}

func TestLocker(t *testing.T) {
	mtx := &sync.RWMutex{} // buat variabel rw mutex
	//lck := mtx.RLocker()    // buat variabel locker dari rw mutex yang sudah ada
	wg := &sync.WaitGroup{} // buat variabel waitgroup untuk menunggu semua goroutine selesai dijalankan

	x := 0 // buat variabel x dengan value awal 0

	for i := 0; i < 1000000; i++ { // buat perulangan sebanyak 1juta kali
		// beri tahu waigroup bahwa ada goroutine baru di setiap perulangan, jadi nanti ada 1juta goroutine yang akan dijalankan
		wg.Add(1)
		go func() { // jalankan goroutine dengan anonymous function di setiap perulangan.
			defer wg.Done()    // diakhir sebelum goroutine ini selesai dijalankan, beri tahu waitgroup bahwa goroutine ini sudah selesai dijalankan
			defer mtx.Unlock() // diakhir sebelum function ini selesai dijalankan, unlock mutexnya

			mtx.Lock()
			x++
		}()
	}

	wg.Wait()
	fmt.Println("hasil akhir x:", x)
}
