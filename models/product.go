package models

import (
	"database/sql"

	"github.com/jeypc/lat1/config"
	"github.com/jeypc/lat1/entity"
)

type ProductModel struct {
	conn *sql.DB
}

func NewProductModel() *ProductModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &ProductModel{
		conn: conn,
	}
}

func (p *ProductModel) FindAll() ([]entity.Product, error) {
	query := "select * from product"
	rows, err := p.conn.Query(query)
	if err != nil {
		return nil, err
	}

	var products []entity.Product
	for rows.Next() {
		var product entity.Product
		rows.Scan(
			&product.Id,
			&product.Nama,
			&product.Harga,
			&product.TanggalPembelian,
			&product.Deskripsi,
		)

		products = append(products, product)
	}

	return products, nil
}
