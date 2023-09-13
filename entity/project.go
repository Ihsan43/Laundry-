package entity

type Customer struct {
	Id   int
	Customer_Name string
	Email 		  string
	No_Hp         string
}

type Service struct {
	Id 		 		int
	Customer_id 	int
	Jenis_Layanan   string
	Satuan			string
	Jumlah 			int
	Harga  	  		 int
}

type Tranction struct {
	Id   int
	Service_Id    int
	Customer_Id     int
	Tanggal_Masuk  string
	Tanggal_Keluar string
	Pelayan        string
	Total_Harga    int
}
