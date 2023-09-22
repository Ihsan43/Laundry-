package service

import (
	"fmt"
	"strings"
	"submission_project/entity"
	"submission_project/utils"
)

const INSERT_CUSTOMER = "INSERT INTO mst_customer(name,email,no_hp) VALUES($1,$2,$3)"
const UPDATE_CUSTOMER = "UPDATE mst_customer SET name=$1, email=$2, no_hp=$3 where id=$4"
const GET_CUSTOMERS = "SELECT * FROM mst_customer"
const DELETE_CUSTOMER = "DELETE FROM mst_customer WHERE id = $1"

func CreateCustomer(newCustomer entity.Customer) error {
	connect := utils.GetConnection()
	_, err := connect.Exec(INSERT_CUSTOMER,newCustomer.Name,newCustomer.Email, newCustomer.NoHp)
	if err != nil{
		fmt.Println("Error Insert Customer ",err)
		return err
	}
	return nil
}

func UpdateCustomer(newCustomer entity.Customer) error {
    connect := utils.GetConnection()
    var customerCount int
    err := connect.QueryRow("SELECT COUNT(*) FROM mst_customer WHERE id = $1", newCustomer.Id).Scan(&customerCount)
    if err != nil {
        return err
    }
    if customerCount == 0 {
        return fmt.Errorf("Customer dengan ID %d tidak ditemukan", newCustomer.Id)
    }  
    _, err = connect.Exec(UPDATE_CUSTOMER, newCustomer.Name, newCustomer.Email, newCustomer.NoHp, newCustomer.Id)
    if err != nil {
        fmt.Println("Error Update Customer ", err)
        return err
    }
    
    return nil
}

func GetListCustomer() ([]entity.Customer, error) {
	connect := utils.GetConnection()
	rows, err := connect.Query(GET_CUSTOMERS)
	if err != nil {
		return nil, err
	}

	var customers []entity.Customer
	for rows.Next() {
		var customer entity.Customer
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.NoHp)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return  customers, nil
}


func DeleteCustomer(id int) (int64, error) {
    connect := utils.GetConnection()
	
    var customerCount int
    err := connect.QueryRow("SELECT COUNT(*) FROM mst_customer WHERE id = $1", id).Scan(&customerCount)
    if err != nil {
        return 0, err
    }

    if customerCount == 0 {
        return 0, fmt.Errorf("Customer dengan ID %d tidak ditemukan", id)
    } 
    result, err := connect.Exec(DELETE_CUSTOMER, id)
    if err != nil {
        return 0, err
    }
    num, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }
    
    return num, nil
}

const INSERT_EMPLOYEE = "INSERT INTO mst_employee(name,no_hp) VALUES($1,$2)"
const UPDATE_EMPLOYEE = "UPDATE mst_employee SET name=$1, no_hp=$2 where id=$3"
const GET_EMPLOYEE = "SELECT * FROM mst_employee"
const DELETE_EMPLOYEE = "DELETE FROM mst_employee WHERE id = $1"

func CreateEmployee(newEmployee entity.Employee) error {
	connect := utils.GetConnection()
	_, err := connect.Exec(INSERT_EMPLOYEE, newEmployee.Name, newEmployee.NoHp)
	if err != nil{
		fmt.Println("Error Insert Employee ",err)
		return err
	}
	return nil
}

func UpdateEmployee(newEmployee entity.Employee) error {
    connect := utils.GetConnection()
    
    var employeeCount int
    err := connect.QueryRow("SELECT COUNT(*) FROM mst_employee WHERE id = $1", newEmployee.Id).Scan(&employeeCount)
    if err != nil {
        return err
    }
    
    if employeeCount == 0 {
        return fmt.Errorf("Karyawan dengan ID %d tidak ditemukan", newEmployee.Id)
    }
    
    _, err = connect.Exec(UPDATE_EMPLOYEE, newEmployee.Name, newEmployee.NoHp, newEmployee.Id)
    if err != nil {
        fmt.Println("Error Update Product ", err)
        return err
    }
    
    return nil
}

func GetListEmployee() ([]entity.Employee, error) {
	connect := utils.GetConnection()
	rows, err := connect.Query(GET_EMPLOYEE)
	if err != nil {
		return nil, err
	}

	var employee []entity.Employee
	for rows.Next() {
		var dataEmployee entity.Employee
		err := rows.Scan(&dataEmployee.Id, &dataEmployee.Name, &dataEmployee.NoHp)
		if err != nil {
			return nil, err
		}
		employee = append(employee, dataEmployee)
	}
	return  employee, nil
}

func DeleteEmployee(id int) (int64, error) {
    connect := utils.GetConnection()

    var employeeCount int
    err := connect.QueryRow("SELECT COUNT(*) FROM mst_employee WHERE id = $1", id).Scan(&employeeCount)
    if err != nil {
        return 0, err
    }

    if employeeCount == 0 {
        return 0, fmt.Errorf("Karyawan dengan ID %d tidak ditemukan", id)
    } 
    result, err := connect.Exec(DELETE_EMPLOYEE, id)
    if err != nil {
        return 0, err
    }
    num, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }
    
    return num, nil
}


const INSERT_PRODUCT = "INSERT INTO mst_product(name, price, satuan) VALUES($1, $2, $3)"
const UPDATE_PRODUCT = "UPDATE mst_product SET name=$1, price=$2, satuan=$3 WHERE id=$4"
const GET_PRODUCT = "SELECT * FROM mst_product"
const DELETE_PRODUCT = "DELETE FROM mst_product WHERE id = $1"

func CreateProduct(newProduct entity.Product) error {
	connect := utils.GetConnection()
	_, err := connect.Exec(INSERT_PRODUCT, newProduct.Name, newProduct.Price, newProduct.Satuan)
	if err != nil{
		fmt.Println("Error Insert Product ",err)
		return err
	}
	return nil
}

func UpdateProduct(newProduct entity.Product) error {
    connect := utils.GetConnection()
    
    var productCount int
    err := connect.QueryRow("SELECT COUNT(*) FROM mst_product WHERE id = $1", newProduct.Id).Scan(&productCount)
    if err != nil {
        return err
    }
    
    if productCount == 0 {
        return fmt.Errorf("Produk dengan ID %d tidak ditemukan", newProduct.Id)
    }
    
    _, err = connect.Exec(UPDATE_PRODUCT, newProduct.Name, newProduct.Price, newProduct.Satuan, newProduct.Id)
    if err != nil {
        fmt.Println("Error Update Product ", err)
        return err
    }
    
    return nil
}

func GetListProduct() ([]entity.Product, error) {
	connect := utils.GetConnection()
	rows, err := connect.Query(GET_PRODUCT)
	if err != nil {
		return nil, err
	}

	var product []entity.Product
	for rows.Next() {
		var dataProduct entity.Product
		err := rows.Scan(&dataProduct.Id, &dataProduct.Name, &dataProduct.Price, &dataProduct.Satuan)
		if err != nil {
			return nil, err
		}
		product = append(product, dataProduct)
	}
	return  product, nil
}

func DeleteProduct(id int) (int64, error) {
    connect := utils.GetConnection()
	
    var productCount int
    err := connect.QueryRow("SELECT COUNT(*) FROM mst_product WHERE id = $1", id).Scan(&productCount)
    if err != nil {
        return 0, err
    }
    if productCount == 0 {
        return 0, fmt.Errorf("Product dengan ID %d tidak ditemukan", id)
    } 
    result, err := connect.Exec(DELETE_PRODUCT, id)
    if err != nil {
        return 0, err
    }
    num, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }
    return num, nil
}

const INSERT_TRANSACTION = "INSERT INTO transaction(id, tanggal_masuk, tanggal_selesai, employee_id, customer_id) VALUES($1, $2, $3, $4, $5)"
const INSERT_TRANSACTION_DETAIL = "INSERT INTO transaction_detail(id, product_id, quantity) VALUES($1, $2, $3)"
const GET_TRANSACTION = "SELECT transaction.id, transaction.tanggal_masuk, transaction.tanggal_selesai, transaction.employee_id, transaction.customer_id, transaction_detail.product_id, transaction_detail.quantity, transaction_detail.total FROM transaction JOIN transaction_detail ON transaction.id = transaction_detail.id"
const SELECT_TOTAL_PRICE = " SELECT SUM(mst_product.price * transaction_detail.quantity) FROM transaction_detail INNER JOIN mst_product ON transaction_detail.product_id = mst_product.id WHERE transaction_detail.id = $1"
const UPDATE_TOTAL_PRICE = "UPDATE transaction_detail SET total = mst_product.price * transaction_detail.quantity FROM mst_product WHERE transaction_detail.product_id = mst_product.id AND transaction_detail.id = $1 RETURNING total;"


func CreateTransaction(transaction entity.Transaction, transactionDetail entity.TransactionDetail) error {
    connect := utils.GetConnection()

    tx, err := connect.Begin()
    if err != nil {
        fmt.Println("Error starting transaction: ", err)
        return err
    }

    _, err = tx.Exec(INSERT_TRANSACTION, transaction.Id, transaction.TanggalMasuk, transaction.TanggalSelesai,
        transaction.EmployeeId, transaction.CustomerId)
    if err != nil {
        fmt.Println("Error Insert Transaction: ", err)
        tx.Rollback() 
        return err
    }

  
    _, err = tx.Exec(INSERT_TRANSACTION_DETAIL, transactionDetail.Id, transactionDetail.ProductId, transactionDetail.Quantity)
    if err != nil {
        fmt.Println("Error Insert Transaction Detail: ", err)
        tx.Rollback()
        return err
    }
    _, err = tx.Exec(UPDATE_TOTAL_PRICE, transactionDetail.Id)
    if err != nil {
        fmt.Println("Error updating total price: ", err)
        tx.Rollback() 
        return err
    }

    err = tx.Commit()
    if err != nil {
        fmt.Println("Error committing transaction: ", err)
        tx.Rollback() 
        return err
    }

    return nil
}

func GetListTransaction() ([]entity.Transaction, []entity.TransactionDetail, error) {
	connect := utils.GetConnection()
	rows, err := connect.Query(GET_TRANSACTION)
	if err != nil {
		panic(err)
	}

	 var transaction []entity.Transaction
	 var transactionDetail []entity.TransactionDetail
	for rows.Next() {
		var dataTransaction entity.Transaction
		var dataTransactionDetail entity.TransactionDetail
		err := rows.Scan(
            &dataTransaction.Id,
            &dataTransaction.TanggalMasuk,
            &dataTransaction.TanggalSelesai,
            &dataTransaction.EmployeeId,
            &dataTransaction.CustomerId,
            &dataTransactionDetail.ProductId,
            &dataTransactionDetail.Quantity,
            &dataTransactionDetail.Total)
		if err != nil {
			panic(err)
		}
		transaction = append(transaction, dataTransaction)
		transactionDetail = append(transactionDetail, dataTransactionDetail)

	}

	return  transaction, transactionDetail, nil
}

func DisplayTransactionDetails() {
    viewTransaction, viewTransactionDetail, err := GetListTransaction()
    if err != nil {
        fmt.Println(err)
    } else {
        utils.RespTemplate("Success Get List Transaksi")
        for i := 0; i < len(viewTransaction); i++ {
            transaction := viewTransaction[i]
            transactionDetail := viewTransactionDetail[i]

            fmt.Println("Id            : ", transaction.Id)
            fmt.Println("Tanggal Masuk : ", transaction.TanggalMasuk)
            fmt.Println("Tanggal Selesai: ", transaction.TanggalSelesai)
            fmt.Println("Employee Id   : ", transaction.EmployeeId)
            fmt.Println("Customer Id   : ", transaction.CustomerId)
            fmt.Println("Product Id    : ", transactionDetail.ProductId)
            fmt.Println("Quantity      : ", transactionDetail.Quantity)
            fmt.Println("Total         : ", transactionDetail.Total)
            fmt.Println(strings.Repeat("=~=", 10))
        }
    }
}


