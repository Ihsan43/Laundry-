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

func ShowEmployeeMenu() {
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
			var newEmployee entity.Employee
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("Nama Employee  : ")
			scanner.Scan()
			newEmployee.Name = scanner.Text()
			fmt.Print("NoHp Employee  : ")
			scanner.Scan()
			newEmployee.NoHp = scanner.Text()
			err := service.CreateEmployee(newEmployee)
			if err != nil {
				fmt.Println(err)
			} else {
				utils.RespTemplate("Success Create New Employee")
			}
			break
		case 2:
			var employeeUpdate entity.Employee
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Masukan Id Employee (Silahkan Liat Terlebih dahulu pada view Menu) : ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			employeeUpdate.Id = id
			fmt.Print("Nama Employee  : ")
			scanner.Scan()
			employeeUpdate.Name = scanner.Text()
			fmt.Print("NoHp Employee  : ")
			scanner.Scan()
			employeeUpdate.NoHp = scanner.Text()
			err := service.UpdateEmployee(employeeUpdate)
			if err != nil {
				fmt.Println(err)
			} else {
				utils.RespTemplate("Success Updated Employee")
			}
			utils.ClearBash()
			break
		case 3:
			employee, err := service.GetListEmployee()
			if err != nil {
				fmt.Println(err)
			} else {
				utils.RespTemplate("Success Get List Employee")
				for _, dataEmployee := range employee {
					fmt.Println("Id   : ", dataEmployee.Id)
					fmt.Println("Name : ", dataEmployee.Name)
					fmt.Println("NoHp : ", dataEmployee.NoHp)
					fmt.Println(strings.Repeat("=~=", 10))
				}
			}
			break
		case 4:
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Println("Masukan Id Employee (Silahkan Liat Terlebih dahulu pada view Menu) : ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			rowsAffected, err := service.DeleteEmployee(id)
			if err != nil {
				fmt.Println(err)
			} else if rowsAffected != 1 {
				utils.RespTemplate("Failed Deleted Employee (Id Tidak ditemukan)")
			} else {
				utils.RespTemplate("Success Deleted Employee")
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