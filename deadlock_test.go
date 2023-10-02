package main

import (
	"fmt"
	"sync"
	"testing"
)

type UserAccount struct {
	Name string
	CBal int
}

// method update balance
func (u *UserAccount) Update(amount int) {
	u.CBal += amount
}

// function to transfer
func Transfer(id int, mu *sync.Mutex, wg *sync.WaitGroup, requestor *UserAccount, beneficiary *UserAccount, amount int) {
	defer func() {
		mu.Unlock()
		wg.Done()
	}()

	mu.Lock()

	requestor.Update(-amount)
	fmt.Println(fmt.Sprintf("ada di goroutine %v, success debet requestor cbal", id))

	beneficiary.Update(amount)
	fmt.Println(fmt.Sprintf("ada di goroutine %v, success credit beneficiary cbal", id))
}

// buat function test
func TestDeadLock(t *testing.T) {
	// buat object untuk requestor
	requestor := &UserAccount{
		Name: "reo",
		CBal: 1000000,
	}

	// buat object untuk beneficiarynya
	beneficiary := &UserAccount{
		Name: "sahobby",
		CBal: 1000000,
	}

	// buat variabel waitgroup untuk menunggu semua gorooutine yang dijalankan, dan memastikan semuanya selesai
	wg := &sync.WaitGroup{}

	// buat variabel mutex untuk lock goroutine yang sedang jalan
	mu := &sync.Mutex{}

	wg.Add(2)
	go Transfer(1, mu, wg, requestor, beneficiary, 100000)
	go Transfer(2, mu, wg, beneficiary, requestor, 100000)

	wg.Wait()
	fmt.Println("requestor saldonya menjadi :", requestor.CBal)
	fmt.Println("beneficiary saldonya menjadi :", beneficiary.CBal)
}
