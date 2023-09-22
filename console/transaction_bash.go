package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"submission_project/entity"
	"submission_project/service"
	"submission_project/utils"
	"submission_project/validasi"

)

func ShowTransaksiMenu() {
	var choice int
	var isExit = false
	for !isExit {
		fmt.Println("=============== Selamat Datang di Aplikasi Laundry =============== ")
		fmt.Println("1. Create")
		fmt.Println("2. List")
		fmt.Println("0. Kembali")
		fmt.Print("Silakan Pilih menu (1-2): ")
		_, err := fmt.Scanln(&choice)
		fmt.Print("\n")
		if err != nil {
			fmt.Println("Masukan tidak valid. Silakan coba lagi.")
			continue
		}
		switch choice {
		case 1:
			var newTransactionDetail entity.TransactionDetail
			var newTransaction entity.Transaction
		
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Id              : ")
			scanner.Scan()
			transactionIDStr := scanner.Text()
			transactionID, err := strconv.Atoi(transactionIDStr)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			newTransaction.Id = transactionID
		
			newTransaction.TanggalMasuk = validasi.ValidateDateInput("Tanggal Masuk")
			newTransaction.TanggalSelesai = validasi.ValidateDateInput("Tanggal Selesai")
		
			employeeIDStr := validasi.ValidateNumericInput("Employee Id", isExit)
			employeeID, err := strconv.Atoi(employeeIDStr)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			newTransaction.EmployeeId = employeeID
		
			customerIDStr := validasi.ValidateNumericInput("Customer Id", isExit)
			customerId, err := strconv.Atoi(customerIDStr)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			newTransaction.CustomerId = customerId
		
			DetailIDStr := validasi.ValidateMaxNumber("Transaction Detail Id")
			DetailID, err := strconv.Atoi(DetailIDStr)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			newTransactionDetail.Id = DetailID
		
			ProductIDStr := validasi.ValidateNumericInput("Product Id", isExit)
			ProductID, err := strconv.Atoi(ProductIDStr)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			newTransactionDetail.ProductId = ProductID
		
			QuantityStr := validasi.ValidateNumericInput("Quantity", isExit)
			quantity, err := strconv.Atoi(QuantityStr)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			newTransactionDetail.Quantity = quantity
		
			err = service.CreateTransaction(newTransaction, newTransactionDetail)
		
			if err != nil {
				fmt.Println(err)
			} else {
				utils.RespTemplate("Success Create New Transaksi")
			}
			break
		
		case 2:
		service.DisplayTransactionDetails()
			break
		case 0:
			isExit = true
			utils.ClearBash()
			break
		default:
			utils.ClearBash()
			fmt.Print("Silahkan Pilih Menu dengan benar \n\n")
		}
	}
}
