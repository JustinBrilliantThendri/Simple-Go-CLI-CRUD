package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type siswa struct{
	nama string
	kelas string
	absen int
}

func connect() (*sql.DB, error){
	db, err := sql.Open("mysql", "root@/go_simple_database")
	if err != nil{
		return nil, err
	}
	return db, nil
}

func insert_data(db *sql.DB){
	var nama, kelas string
	var absen int
	fmt.Print("Nama : ")
	fmt.Scan(&nama)
	fmt.Print("Kelas : ")
	fmt.Scan(&kelas)
	fmt.Print("Absen : ")
	fmt.Scan(&absen)
	_, err := db.Exec("INSERT INTO tb_siswa VALUES ('', ?, ?, ?)", nama, kelas, absen)
	if err != nil{
		panic(err.Error())
	}
}

func select_data(db *sql.DB){
	rows, err := db.Query("SELECT nama, kelas, absen FROM tb_siswa")
	if err != nil{
		panic(err.Error())
	}
	result := []siswa{}
	for rows.Next(){
		each := siswa{}
		err := rows.Scan(&each.nama, &each.kelas, &each.absen)
		if err != nil{
			panic(err.Error())
		}
		result = append(result, each)
	}
	if err = rows.Err(); err != nil{
		panic(err.Error())
	}
	for idx, each := range result{
		fmt.Printf("%d. %v, %v, %d", (idx + 1), each.nama, each.kelas, each.absen)
	}
}

func main() {
	db, err := connect()
	if err != nil{
		panic(err.Error())
	}
	var pilihan int
	fmt.Print("Pilih opsi :\n1. Insert\n2. Select\nPilihan : ")
	fmt.Scan(&pilihan)
	fmt.Println()
	switch pilihan{
	case 1:
		insert_data(db)
	case 2:
		select_data(db)
	default:
		fmt.Println("Opsi tidak tersedia!")
	}
}