package main

import (
	"database/sql"
	"fmt"
	"os"
	"submission_project/entity"
	"bufio"    

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbname = "Laundry"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
var choice int
var clean string 


func main() {

    for {
		fmt.Println("===============Table Database=============== \n")
        fmt.Println("1. Mst_Customer")
        fmt.Println("2. Service")
        fmt.Println("3. Transaksi")
        fmt.Println("0. Keluar")
        fmt.Print("Silakan Pilih menu (0-4): ")


        _, err := fmt.Scanln(&choice)
		fmt.Print("\n")
        if err != nil {
            fmt.Println("Masukan tidak valid. Silakan coba lagi.")
            fmt.Scanln(&clean) 
            continue			
        }

        switch choice {
        case 1:
		   fmt.Println("===============Mst_Customer=============== \n")
           fmt.Println("1.View Customer")
		   fmt.Println("2.Insert Customer")
		   fmt.Println("3.Update Customer")
		   fmt.Println("4.Delete Customer")
		   fmt.Println("88. Kembali")
		   fmt.Print("Silakan Pilih menu (0-4): ")

		   _, err := fmt.Scanln(&choice)
		   if err != nil {
			   fmt.Println("Masukan tidak valid. Silakan coba lagi.")
			   fmt.Scanln(&clean) 
			   continue
		   }
		   
		   fmt.Print("\n")

		   if choice == 1 {
			customers := getAllCustomer()
			for _, customer := range customers{
				fmt.Println(customer.Id, customer.Customer_Name, customer.Email, customer.No_Hp)
			}

			fmt.Print("\n")

			fmt.Println("0.Keluar")
			fmt.Println("1.Kembali")
			fmt.Print("Silakan Pilih menu (0-2): ")
			fmt.Scanln(&choice)
			fmt.Print("\n")

			if choice == 0 {
				fmt.Println("Terima Kasih")
				os.Exit(0)
			} else if choice == 1 {
				continue
			} else {
				fmt.Println("Masukan tidak valid. Silakan coba lagi. \n")
				continue
			}
				
		   } else if choice == 2 {

			var Id int
			var nama_customer string
			var email string
			var no_hp string
			
			for {
				fmt.Print(" Id   : ")
				_, err := fmt.Scanln(&Id)
				if err != nil {
					fmt.Println("Input Id tidak valid. Harap masukkan angka.")
					continue
				}
			
				fmt.Print("Nama Customer : ")
				_, err = fmt.Scanln(&nama_customer)
				if err != nil {
					fmt.Println("Input Nama Customer tidak valid.")
					continue
				}

				fmt.Print("Email : ")
				_, err = fmt.Scanln(&email)
				if err != nil {
					fmt.Println("Input Nama Customer tidak valid.")
					continue
				}
			
				fmt.Print("No Hp         : ")
				_, err = fmt.Scanln(&no_hp)
				if err != nil {
					fmt.Println("Input No Hp tidak valid. Harap masukkan angka.")
					continue

				}

				sqlStatment := entity.Customer{Id: Id,Customer_Name: nama_customer,  Email: email, No_Hp: no_hp}
				customerEnrollment(sqlStatment)
					
				break
			}

			fmt.Println("\n0.Keluar")
			fmt.Println("1.View All Customer")
			fmt.Println("2.Kembali")
			fmt.Print("Silakan Pilih menu (0-2): ")
			fmt.Scanln(&choice)
			fmt.Print("\n")

			if choice == 0 {
				fmt.Println("Terima Kasih")
				os.Exit(0)
			} else if choice == 1 {
				customers := getAllCustomer()
			for _, customer := range customers{
				fmt.Println(customer.Id, customer.Customer_Name, customer.Email, customer.No_Hp)} 
			} else if choice == 2{
				continue
			} else {
				fmt.Println("Masukan tidak valid. Silakan coba lagi.\n")
				continue
			}

			
		} else if choice == 3 {
			fmt.Print("Id	       : ")
			var id int
			fmt.Scan(&id)
			fmt.Print("Customer_Name : ")
			var customer_name string
			fmt.Scan(&customer_name)
			fmt.Print("Email	         : ")
			var email string
			fmt.Scan(&email)
			fmt.Print("No Hp	: ")
			var no_hp string
			fmt.Scan(&no_hp)

			sqlStatment := entity.Customer{Id: id , Customer_Name: customer_name, Email: email, No_Hp: no_hp}
			updateCustomer(sqlStatment)

			fmt.Println("\n0.Keluar")
			fmt.Println("1.View All Customer")
			fmt.Println("2.Kembali")
			fmt.Print("Silakan Pilih menu (0-2): ")
			fmt.Scanln(&choice)
			fmt.Print("\n")

			if choice == 0 {
				fmt.Println("Terima Kasih")
				os.Exit(0)
			} else if choice == 1 {
				customers := getAllCustomer()
			for _, customer := range customers{
				fmt.Println(customer.Id, customer.Customer_Name, customer.Email, customer.No_Hp)} 
			} else if choice == 2{
				continue
			} else {
				fmt.Println("Masukan tidak valid. Silakan coba lagi.\n")
				continue
			}

		} else if choice == 4  {
			fmt.Print("Delete Customer  : ")
			var delete_customer int
			fmt.Scanln(&delete_customer)
			deleteCustomer(delete_customer)

			fmt.Print("\n")

			fmt.Println("0.Keluar")
			fmt.Println("1.View All Customer")
			fmt.Println("2.Kembali")
			fmt.Print("Silakan Pilih menu (0-2): ")
			fmt.Scanln(&choice)
			fmt.Print("\n")

			if choice == 0 {
				fmt.Println("Terima Kasih")
				os.Exit(0)
			} else if choice == 1 {
				customers := getAllCustomer()
			for _, customer := range customers{
				fmt.Println(customer.Id, customer.Customer_Name, customer.Email, customer.No_Hp)} 
			} else if choice == 88 {
				continue
			} else if choice == 2{
				continue
			} else {
				fmt.Println("Input yang anda masukan tidak valid")
			}
		}  
		   return
        case 2:
			fmt.Println("===============Service=============== \n")
			fmt.Println("1.View Service")
			fmt.Println("2.Insert Service")
			fmt.Println("3.Update Service")
			fmt.Println("4.Delete Service")
			fmt.Println("88. Kembali")
			fmt.Print("Silakan Pilih menu (0-4): ")



			_, err := fmt.Scanln(&choice)
			if err != nil {
				fmt.Println("Masukan tidak valid. Silakan coba lagi.")
				fmt.Scanln(&clean) 
				continue
			}
			
			fmt.Print("\n")
 
			if choice == 1 {
			 services := getAllService()
			 for _, service := range services{
				 fmt.Println(service.Id, service.Customer_id, service.Jenis_Layanan, service.Satuan, service.Jumlah, service.Harga)
			 }
 
			 fmt.Println("\n0.Keluar")
			 fmt.Println("1.Kembali")
			 fmt.Print("Silakan Pilih menu (0-2): ")
			 fmt.Scanln(&choice)
			 fmt.Print("\n")
 
			 if choice == 0 {
				 fmt.Println("Terima Kasih")
				 os.Exit(0)
			 } else if choice == 1 {
				 continue
			 } else {
				 fmt.Println("Masukan tidak valid. Silakan coba lagi. \n")
				 continue
			 }
				 
			} else if choice == 2 {
 
				var id int
				var customer_Id int
				var jenis_layanan string
				var satuan string
				var jumlah int
				var harga int

				scanner := bufio.NewScanner(os.Stdin)
			
				for {
					fmt.Print("Id	     : ")
					_, err := fmt.Scanln(&id)
					if err != nil {
						fmt.Println("Input Service_Id tidak valid. Harap masukkan angka.")
						continue
					}

					fmt.Print("Customer_Id  : ")
					_, err = fmt.Scanln(&customer_Id)
					if err != nil {
						fmt.Println("Input customer_Id tidak valid. Harap masukkan angka.")
						continue
					}
			
					fmt.Print("Jenis Layanan: ")
					scanner.Scan()
   					jenis_layanan = scanner.Text()
					if err != nil {
						fmt.Println("Input Jenis Layanan tidak valid.")
						continue
					}

					fmt.Print("Satuan(kg/buah): ")
					_, err = fmt.Scanln(&satuan)
					if err != nil {
						fmt.Println("Input satuan tidak valid. Harap masukkan angka.")
						continue
					}

					fmt.Print("Jumlah	     : ")
					_, err = fmt.Scanln(&jumlah)
					if err != nil {
						fmt.Println("Input jumlah tidak valid. Harap masukkan angka.")
						continue
					}
			
					fmt.Print("Harga         : ")
					_, err = fmt.Scanln(&harga)
					if err != nil {
						fmt.Println("Input Harga tidak valid. Harap masukkan angka.")
						continue
					}
			
					break
				}
			
				sqlStatment := entity.Service{Id: id, Customer_id: customer_Id, Jenis_Layanan: jenis_layanan, Satuan: satuan,
				Jumlah: jumlah, Harga: harga}
				ServiceEnrollment(sqlStatment)
			
 
			 fmt.Println("0.Keluar")
			 fmt.Println("1.View All Service")
			 fmt.Println("2.Kembali")
			 fmt.Print("Silakan Pilih menu (0-2): ")
			 fmt.Scanln(&choice)
			 fmt.Print("\n")
 
			 if choice == 0 {
				 fmt.Println("Terima Kasih")
				 os.Exit(0)
			 } else if choice == 1 {
				services := getAllService()
				for _, service := range services{
					fmt.Println(service.Id, service.Customer_id, service.Jenis_Layanan, service.Satuan, service.Jumlah, service.Harga)
				}
	 		 } else if choice == 2{
				continue
			 } else {
				 fmt.Println("Masukan tidak valid. Silakan coba lagi.\n")
				 continue
			 }
		 } else if choice == 3 {
			fmt.Print("Id   	: ")
			 var id int
			 fmt.Scan(&id)
			 fmt.Print("Customer_Id   : ")
			 var customer_Id int
			 fmt.Scan(&customer_Id)
			 fmt.Print("Jenis Layanan : ")
			 var jenis_layanan string
			 fmt.Scan(&jenis_layanan)		
			 var jumlah int
			 fmt.Scan(&jumlah)
			 fmt.Print("Jumlah 		  : ") 
			 var satuan string
			 fmt.Scan(&satuan)
			 fmt.Print("Satuan 		  : ")	
			 var harga int
			 fmt.Scan(&harga)
			 fmt.Print("Harga 		  : ")
 
			 sqlStatment := entity.Service{Id: id, Customer_id: customer_Id, Jenis_Layanan: jenis_layanan, Jumlah: jumlah,
				Satuan: satuan, Harga: harga}
				updateService(sqlStatment)
 
			 fmt.Println("\n0.Keluar")
			 fmt.Println("1.View All Service")
			 fmt.Println("2.Kembali")
			 fmt.Print("Silakan Pilih menu (0-3): ")
			 fmt.Scanln(&choice)
			 fmt.Print("\n")
 
			 if choice == 0 {
				 fmt.Println("Terima Kasih")
				 os.Exit(0)
			 } else if choice == 1 {
				services := getAllService()
				for _, service := range services{
					fmt.Println(service.Id, service.Customer_id, service.Jenis_Layanan, service.Satuan, service.Jumlah, service.Harga)
				}
	 		 } else if choice == 2 {
				continue
			 } else {
				 fmt.Println("Masukan tidak valid. Silakan coba lagi.\n")
				 continue
			 }
		 } else if choice == 4  {
			 fmt.Print("Delete Service  : ")
			 var delete_service int
			 fmt.Scanln(&delete_service)
			 deleteService(delete_service)
 
			 fmt.Println("\n0.Keluar")
			 fmt.Println("1.View All Service")
			 fmt.Println(("2.Kembali"))
			 fmt.Print("Silakan Pilih menu (0-2): ")
			 fmt.Scanln(&choice)
			 fmt.Print("\n")
 
			 if choice == 0 {
				 fmt.Println("Terima Kasih")
				 os.Exit(0)
			 } else if choice == 1 {
				services := getAllService()
				for _, service := range services{
					fmt.Println(service.Id, service.Customer_id, service.Jenis_Layanan, service.Satuan, service.Jumlah, service.Harga)
				}
			 }  else if choice == 88 {
				continue
			} else if choice == 2{
				continue
			} else {
			   fmt.Println("Input yang anda masukan tidak valid")
			}
		 }
			return
        case 3:
			fmt.Println("===============Transaction=============== \n")	
			fmt.Println("1.Insert Transaksi")		
			fmt.Println("2.View Transaksi")
			fmt.Println("88. Kembali")
			fmt.Print("Silakan Pilih Menu(1-2): ")
			
			_, err := fmt.Scanln(&choice)
			if err != nil {
				fmt.Println("Masukan tidak valid. Silakan coba lagi.")
				fmt.Scanln(&clean) 
				continue
			}
				if choice == 1 {
					var (
						Id   int
						Customer_Id    int
						service_id	   int
						Tanggal_Masuk  string
						Tanggal_Keluar string
						Pelayan        string
						// total_harga	   int
					)
					

					for {
						fmt.Print("\nId 		 : ")
						_, err := fmt.Scanln(&Id)
						if err != nil {
							fmt.Println("Input id tidak valid. Harap masukkan angka.")
							continue
						}
				

						fmt.Print("Service Id 	 : ")
						_, err = fmt.Scanln(&service_id)
						if err != nil {
							fmt.Println("Input service id tidak valid.")
							continue
						}
						
						fmt.Print("Customer Id   : ")
						_, err = fmt.Scanln(&Customer_Id)
						if err != nil {
							fmt.Println("Input customer id tidak valid.")
							continue
						}
				
						fmt.Print("Tanggal Masuk : ")
						_, err = fmt.Scanln(&Tanggal_Masuk)
						if err != nil {
							fmt.Println("Input tanggal masuk tidak valid. Harap masukkan angka.")
							continue
						}

						fmt.Print("Tanggal Keluar : ")
						_, err = fmt.Scanln(&Tanggal_Keluar)
						if err != nil {
							fmt.Println("Input Tanggal keluar valid. Harap masukkan angka.")
							continue
						}

						fmt.Print("Pelayan 		  : ")
						_, err = fmt.Scanln(&Pelayan)
						if err != nil {
							fmt.Println("Input Pelayab valid. Harap masukkan angka.")
							continue
						}

						// fmt.Print("Total Harga 	   : ")
						// _, err = fmt.Scanln(&total_harga)
						// if err != nil {
						// 	fmt.Println("Input jumlah tidak valid.")
						// 	continue
						// }
				
						break
					}
				
					sqlStatment := entity.Tranction{Id: Id, Customer_Id: Customer_Id, Service_Id: service_id,
						Tanggal_Masuk: Tanggal_Masuk, Tanggal_Keluar: Tanggal_Keluar, Pelayan: Pelayan}
					transaksiEnrollment(sqlStatment)
				 
				 fmt.Println("\n0.Keluar")
				 fmt.Println("1.View Transaksi")
				 fmt.Println("2.Kembali")
				 fmt.Print("Silakan Pilih menu (0-2): ")
				 fmt.Scanln(&choice)
				 fmt.Print("\n")
	 
				 if choice == 0 {
					 fmt.Println("Terima Kasih")
					 os.Exit(0)
				 } else if choice == 1 {
					transaksi := getAllTransaksi()
					for _, transaction := range transaksi{
						fmt.Println(transaction.Id, transaction.Service_Id, transaction.Customer_Id, transaction.Tanggal_Masuk,
						transaction.Tanggal_Keluar, transaction.Pelayan, transaction.Total_Harga)
					}	
				  }  else if choice == 2{
					continue
				  } else {
					 fmt.Println("Masukan tidak valid. Silakan coba lagi.\n")
					 continue
				 }
				} else if choice == 2 {
					transaksi := getAllTransaksi()
					for _, transaction := range transaksi{
						fmt.Println(transaction.Id, transaction.Service_Id, transaction.Customer_Id, transaction.Tanggal_Masuk,
						transaction.Tanggal_Keluar, transaction.Pelayan, transaction.Total_Harga)
					}	

					fmt.Println("\n0.Keluar")
					fmt.Println("1.Kembali")
					fmt.Print("Silakan Pilih menu (0-2): ")
					fmt.Scanln(&choice)
					fmt.Print("\n")
		
					if choice == 0 {
						fmt.Println("Terima Kasih")
						os.Exit(0)
					} else if choice == 1 {
						continue
					} else {
						fmt.Println("Masukan tidak valid. Silakan coba lagi. \n")
						continue
					}

		
				} else if choice == 88 {
					continue
				} else {
					fmt.Println("Input yang anda masukan tidak valid")
				}
			
			return
        case 0:
            fmt.Println("Anda Telah Keluar")
            return 
        default:
            fmt.Println("Pilihan tidak valid. Silakan pilih 0 hingga 4.\n")
			continue
        }
    }
}

func getAllCustomer() []entity.Customer{

	db := connectDb()
	defer db.Close()

	sqlStatment := "SELECT * FROM mst_customer"

	rows, err := db.Query(sqlStatment)

	if err != nil {
		panic(err)
	} 
 
	customers := scanCustomer(rows)

	return customers
	

}

func scanCustomer(rows *sql.Rows) []entity.Customer {

	customers := []entity.Customer{}
	var err error

	for rows.Next() {
		customer := entity.Customer{}
		err := rows.Scan(&customer.Id, &customer.Customer_Name, &customer.Email, &customer.No_Hp)

		if err != nil {
			panic(err)
		}

		customers = append(customers,customer)
	}

	err = rows.Err()

	if err != nil {
		panic(err)
	}

	return customers
}

func connectDb() *sql.DB {

	db, err := sql.Open("postgres", psqlInfo)
	
		if err != nil {
			panic(err)
		}
	
		err = db.Ping()
	
		if err != nil {
			panic(err)
		} 

		return db
	}

func customerEnrollment(customerenrollment entity.Customer){
	db := connectDb()
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	insertCustomer(customerenrollment, tx)
	updateCustomer(customerenrollment)

	err = tx.Commit()
	
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Success Commited")
	}
}

func insertCustomer(customer entity.Customer, tx *sql.Tx){

	insertCustomer := "INSERT INTO mst_customer (id, customer_name, email, no_hp) VALUES ($1, $2, $3, $4)"

	_, err := tx.Exec(insertCustomer, customer.Id, 
		customer.Customer_Name, customer.Email, customer.No_Hp)

		validate(err, "Insert", tx)

}

func updateCustomer (updateCustomer entity.Customer) {
	db := connectDb()
	defer db.Close()

    sqlStatment :=  "UPDATE mst_customer SET customer_name = $2, email = $3, no_hp = $4 WHERE id = $1"

	_,err := db.Exec(sqlStatment, updateCustomer.Id, updateCustomer.Customer_Name, 
	updateCustomer.Email, updateCustomer.No_Hp)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succsefully Update Customer")
	}
}

func validate (err error, messege string, tx *sql.Tx){
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
	} else {
		fmt.Println("Succesfully " + messege + "Data!")
	}
}


func deleteCustomer (id int){

	db := connectDb()
	defer db.Close()

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM mst_customer WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		panic(err)
	}

	if !exists {
		fmt.Println("Customer dengan Id", id, "tidak ditemukan dalam database.")
		return
	}

	sqlStatment := "DELETE FROM mst_customer WHERE id = $1"
	_, err = db.Exec(sqlStatment, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Delete customer")
	}


}
	
func ServiceEnrollment(serviceenrollment entity.Service){
	db := connectDb()
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	insertService(serviceenrollment, tx)

	total_harga := getSumTotalHargaTransaksii(serviceenrollment.Customer_id, tx)
	
	updateTransacion(total_harga, serviceenrollment.Customer_id, tx)

	err = tx.Commit()
	
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Success Commited")
	}
}

func updateTransacion (total_harga int, service_id int, tx *sql.Tx){
	updateStudent := "UPDATE transaction SET total_harga = $1 WHERE id = $2"

	_, err := tx.Exec(updateStudent, total_harga, service_id)

	validate(err, "Update", tx)
}

func getSumTotalHargaTransaksii(id int, tx *sql.Tx) int {
	sumTotalHarga := "SELECT SUM(harga * jumlah) FROM mst_service WHERE customer_id = $1"

	total_harga := 0
	err := tx.QueryRow(sumTotalHarga,id).Scan(&total_harga)

	validate(err, "Select", tx)

	return total_harga

}
func insertService(service entity.Service, tx *sql.Tx){

	insertservice := "INSERT INTO mst_service (id, customer_id, jenis_layanan, jumlah, satuan , harga) VALUES ($1, $2, $3, $4, $5, $6)"

	_, err := tx.Exec(insertservice, service.Id, service.Customer_id, 
		service.Jenis_Layanan, service.Jumlah, service.Satuan, service.Harga)

		validate(err, "Insert", tx)
}
func updateService (updateService entity.Service, ){

	db := connectDb()
	defer db.Close()

    sqlStatment :=  "UPDATE mst_service SET customer_id = $2, jenis_layanan = $3, satuan = $4, jumlah = $5, harga = $6 WHERE id = $1"

	_, err := db.Exec(sqlStatment, updateService.Id, updateService.Customer_id, updateService.Jenis_Layanan, updateService.Satuan ,
		updateService.Jumlah, updateService.Harga)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succsefully Update Service")
	}
}

func getAllService() []entity.Service {

	db := connectDb()
	defer db.Close()

	sqlStatment := "SELECT * FROM mst_service"

	rows, err := db.Query(sqlStatment)

	if err != nil {
		panic(err)
	} 
 
	services := scanServices(rows)

	return services

}

func scanServices(rows *sql.Rows) []entity.Service {

	services := []entity.Service{}
	var err error

	for rows.Next() {
		service := entity.Service{}
		err := rows.Scan(&service.Id, &service.Customer_id, &service.Jenis_Layanan, &service.Jumlah, &service.Satuan, &service.Harga)

		if err != nil {
			panic(err)
		}

		services = append(services,service)
	}

	err = rows.Err()

	if err != nil {
		panic(err)
	}

	return services
}


func deleteService (id int){

	db := connectDb()
	defer db.Close()

	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM mst_service WHERE id = $1)", id).Scan(&exists)
	if err != nil {
		panic(err)
	}

	if !exists {
		fmt.Println("Service dengan Id", id, "tidak ditemukan dalam database.")
		return
	}

	sqlStatment := "DELETE FROM mst_service WHERE id = $1"
	_, err = db.Exec(sqlStatment, id)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Delete Service")
	}

}

func transaksiEnrollment(transaksiEnrollmentent entity.Tranction){

	db := connectDb()
	defer db.Close()

	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	insertTransaksi(transaksiEnrollmentent, tx)

	err = tx.Commit()
	
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Success Commited")
	}
}

func insertTransaksi(transaksi entity.Tranction, tx *sql.Tx) {

	insertTransaksi := "INSERT INTO transaction (id, service_id, customer_Id, tanggal_masuk, tanggal_keluar, pelayan, total_harga ) VALUES ($1, $2, $3, $4, $5 ,$6, $7)"

	Total_Harga := 0
	_, err := tx.Exec(insertTransaksi, transaksi.Id, transaksi.Service_Id, transaksi.Customer_Id, 
		transaksi.Tanggal_Masuk, transaksi.Tanggal_Keluar, transaksi.Pelayan, Total_Harga)

		isCustomerExist(transaksi.Customer_Id, tx)

		isServiceExist(transaksi.Service_Id, tx)

		validate(err, "Insert", tx)
}

func getAllTransaksi() []entity.Tranction {

	db := connectDb()
	defer db.Close()

	sqlStatment := "SELECT * FROM transaction"

	rows, err := db.Query(sqlStatment)

	if err != nil {
		panic(err)
	} 
 
	transaksi := scanTransaksi(rows)

	return transaksi

}

func scanTransaksi(rows *sql.Rows) []entity.Tranction {

	Transaksi := []entity.Tranction{}
	var err error

	for rows.Next() {
		transaksi := entity.Tranction{}
		err := rows.Scan(&transaksi.Id, &transaksi.Service_Id, &transaksi.Customer_Id, &transaksi.Tanggal_Masuk, &transaksi.Tanggal_Keluar, 
			&transaksi.Pelayan, &transaksi.Total_Harga)

		if err != nil {
			panic(err)
		}

		Transaksi = append(Transaksi,transaksi)
	}

	err = rows.Err()

	if err != nil {
		panic(err)
	}

	return Transaksi
}


func isCustomerExist(customerId int, tx *sql.Tx) bool {
    query := "SELECT COUNT(*) FROM mst_customer WHERE id = $1"
    var count int
    err := tx.QueryRow(query, customerId).Scan(&count)
    if err != nil {
        return false
    }
    return count > 0
}

func isServiceExist(serviceId int, tx *sql.Tx) bool {
    query := "SELECT COUNT(*) FROM mst_service WHERE id = $1"
    var count int
    err := tx.QueryRow(query, serviceId).Scan(&count)
    if err != nil {
        return false
    }
    return count > 0
}





