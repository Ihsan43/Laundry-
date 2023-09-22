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

func ShowProductMenu() {
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
			var newProduct entity.Product
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Nama Product   : ")
			scanner.Scan()
			newProduct.Name = scanner.Text()
			fmt.Print("Price Product  : ")
			scanner.Scan()
			newProduct.Price = scanner.Text()
			fmt.Print("Satuan Product : ")
			scanner.Scan()
			newProduct.Satuan = scanner.Text()
			err := service.CreateProduct(newProduct)
			if err != nil {
				fmt.Println(err)
			} else {
				utils.RespTemplate("Success Create New Product")
			}
			break
		case 2:
			var productUpdate entity.Product
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Masukan Id Product (Silahkan Liat Terlebih dahulu pada view Menu) : ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			productUpdate.Id = id
			fmt.Print("Nama Product   : ")
			scanner.Scan()
			productUpdate.Name = scanner.Text()
			fmt.Print("Price Product  : ")
			scanner.Scan()
			productUpdate.Price = scanner.Text()
			fmt.Print("Satuan Product : ")
			scanner.Scan()
			productUpdate.Satuan = scanner.Text()
			err := service.UpdateProduct(productUpdate)
			if err != nil {
				fmt.Println(err)
			} else {
				utils.RespTemplate("Success Updated Product")
			}
			utils.ClearBash()
			break
		case 3:
			Product, err := service.GetListProduct()
			if err != nil {
				fmt.Println(err)
			} else {
				utils.RespTemplate("Success Get List Product")
				for _, dataProduct := range Product {
					fmt.Println("Id     	   : ", dataProduct.Id)
					fmt.Println("Product Name   : ", dataProduct.Name)
					fmt.Println("Product Price  : ", dataProduct.Price)
					fmt.Println("Product Satuan : ", dataProduct.Satuan)
					fmt.Println(strings.Repeat("=~=", 10))
				}
			}
			break
		case 4:
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Masukan Id Product (Silahkan Liat Terlebih dahulu pada view Menu) : ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			rowsAffected, err := service.DeleteProduct(id)
			if err != nil {
				fmt.Println(err)
			} else if rowsAffected != 1 {
				utils.RespTemplate("Failed Deleted Product (Id Tidak ditemukan)")
			} else {
				utils.RespTemplate("Success Deleted Product")
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