package main

import (
	"fmt"
	"submission_project/console"
	"submission_project/service"
	"submission_project/utils"
)

func main() {
	showMainMenu()
}

func showMainMenu(){
	var choice int
	var isExit = false
	for !isExit {		
		fmt.Println("=============== Selamat Datang di Aplikasi Laundry =============== ")		
		fmt.Println("1. Buat Pesanan")
		fmt.Println("2. Manage Customer")
		fmt.Println("3. Manage Product")
		fmt.Println("4. Manage Karyawan")
		fmt.Println("5. Lihat Semua Transaksi")
		fmt.Println("0. Keluar")
		fmt.Print("Silakan Pilih menu (1-5): ")
		_, err := fmt.Scanln(&choice)
		fmt.Print("\n")
		if err != nil {
			fmt.Println("Masukan tidak valid. Silakan coba lagi.")
			continue
		}
		switch choice {

		case 1:
			utils.ClearBash()
			console.ShowTransaksiMenu()
			break
		case 2:
			utils.ClearBash()
			console.ShowCustomerMenu()			
			break
		case 3:
			utils.ClearBash()
			console.ShowProductMenu()			
			break
		case 4:
			utils.ClearBash()
			console.ShowEmployeeMenu()			
			break
		case 5:
			utils.ClearBash()
			service.DisplayTransactionDetails()
			break
		case 0 : 
			isExit = true
			break			
		default : 
			utils.ClearBash()
			fmt.Print("Silahkan Pilih Menu dengan benar \n\n")
		}
	}
}




