package models

import (
	"database/sql"
	"fmt"

	"github.com/jeypc/lat1/config"
	"github.com/jeypc/lat1/entity"
)

type CustomerModel struct {
	conn *sql.DB
}

func NewCustomerModel() *CustomerModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &CustomerModel{
		conn: conn,
	}
}

func (c *CustomerModel) FindAll() ([]entity.Customer, error) {
	rows, _ := c.conn.Query("select * from customers")

	var dataCustomer []entity.Customer

	for rows.Next() {
		var customer entity.Customer
		rows.Scan(
			&customer.Id,
			&customer.NamaCustomer,
			&customer.JenisKelamin,
			&customer.NoHP,
			&customer.Alamat,
			&customer.TanggalLahir,
		)
		dataCustomer = append(dataCustomer, customer)
	}

	return dataCustomer, nil
}

func (c *CustomerModel) FindById(id string) (*entity.Customer, error) {
	query := fmt.Sprintf("select * from customers where id = %s", id)
	row, err := c.conn.Query(query)
	if err != nil {
		return nil, err
	}

	if !row.Next() {
		return nil, fmt.Errorf("error: User with ID %s not found", id)
	}

	var customer entity.Customer
	row.Scan(
		&customer.Id,
		&customer.NamaCustomer,
		&customer.JenisKelamin,
		&customer.NoHP,
		&customer.Alamat,
		&customer.TanggalLahir,
	)
	return &customer, nil
}

func (c *CustomerModel) Create(customer entity.Customer) bool {
	query := fmt.Sprintf("insert into customers (nama_customer, jenis_kelamin, no_hp, alamat, tanggal_lahir) values('%s','%s','%s','%s','%s')", customer.NamaCustomer, customer.JenisKelamin, customer.NoHP, customer.Alamat, customer.TanggalLahir)

	result, err := c.conn.Exec(query)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (c *CustomerModel) DeleteByID(id string) (string, error) {
	query := fmt.Sprintf("delete from customers where id = %s", id)
	_, err := c.conn.Query(query)
	if err != nil {
		return "error", err
	}

	return "Data Berhasil Dihapus", nil
}
