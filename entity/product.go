package entity

type Product struct {
	Id               int64  `json:"id"`
	Nama             string `json:"nama"`
	Harga            string `json:"harga"`
	TanggalPembelian string `json:"tanggal_pembelian"`
	Deskripsi        string `json:"deskripsi"`
}
