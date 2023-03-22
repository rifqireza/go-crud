package entity

type Customer struct {
	Id           int64  `json:"id"`
	NamaCustomer string `json:"nama_customer"`
	JenisKelamin string `json:"jenis_kelamin"`
	NoHP         string `json:"no_hp"`
	Alamat       string `json:"alamat"`
	TanggalLahir string `json:"tanggal_lahir"`
}
