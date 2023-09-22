package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"submission_project/entity"
	"submission_project/service"
	"submission_project/utils"
)

func ShowCustomerMenu() {
	var choice int
	var isExit = false
	for !isExit {
		fmt.Println("=============== Selamat Datang di Aplikasi Laundry =============== ")
		fmt.Println("1. Create")
		fmt.Println("2. Update")
		fmt.Println("3. List")
		fmt.Println("4. Delete")
		fmt.Println("5. Kembali")
		fmt.Print("Silakan Pilih menu (1-4): ")
		_, err := fmt.Scanln(&choice)
		fmt.Print("\n")
		if err != nil {
			fmt.Println("Masukan tidak valid. Silakan coba lagi.")
			continue
		}
		switch choice {
		case 1:
			var newCustomer entity.Customer
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Nama Customer  : ")
			scanner.Scan()
			newCustomer.Name = scanner.Text()
			fmt.Print("Email Customer  : ")
			scanner.Scan()
			newCustomer.Email = scanner.Text()
			fmt.Print("NoHp Customer  : ")
			scanner.Scan()
			newCustomer.NoHp = scanner.Text()
			err := service.CreateCustomer(newCustomer)
			if err != nil {
				fmt.Println(err)
			} else {
				utils.RespTemplate("Success Create New Customer")
			}
			break
		case 2:
			var customerUpdated entity.Customer
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Masukan Id Customer (Silahkan Liat Terlebih dahulu pada view Menu) : ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			customerUpdated.Id = id
			fmt.Print("Nama Customer  : ")
			scanner.Scan()
			customerUpdated.Name = scanner.Text()
			fmt.Print("Email Customer  : ")
			scanner.Scan()
			customerUpdated.Email = scanner.Text()
			fmt.Print("NoHp Customer  : ")
			scanner.Scan()
			customerUpdated.NoHp = scanner.Text()
			err := service.UpdateCustomer(customerUpdated)
			if err != nil {
				fmt.Println(err)
			} else {
				utils.RespTemplate("Success Updated Customer")
			}
			utils.ClearBash()
			break
		case 3:
			customers, err := service.GetListCustomer()
			if err != nil {
				fmt.Println(err)
			} else {
				utils.RespTemplate("Success Get List Customer")
				for _, customer := range customers {
					fmt.Println("Id : ", customer.Id)
					fmt.Println("Name : ", customer.Name)
					fmt.Println("Email : ", customer.Email)
					fmt.Println("NoHp : ", customer.NoHp)
					fmt.Println(strings.Repeat("=~=", 10))
				}
			}
			break
		case 4:
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Masukan Id Customer (Silahkan Liat Terlebih dahulu pada view Menu) : ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			rowsAffected, err := service.DeleteCustomer(id)
			if err != nil {
				fmt.Println(err)
			} else if rowsAffected != 1 {
				utils.RespTemplate("Failed Deleted Customer (Id Tidak ditemukan)")
			} else {
				utils.RespTemplate("Success Deleted Customer")
			}
			break
		case 5:
			isExit = true
			utils.ClearBash()
			break
		default:
			utils.ClearBash()
			fmt.Print("Silahkan Pilih Menu dengan benar \n\n")
		}
	}
}