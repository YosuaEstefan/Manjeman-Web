package models

type Belanja struct {
	ID     int    `json:"id"`
	Name   string `json:"nama"`
	Jenis  string `json:"jenis"`
	Jumlah int    `json:"jumlah"`
}

var DaftarBelanja []Belanja

func init() {
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 1, Name: "baju", Jenis: "pakaian", Jumlah: 3})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 2, Name: "Sabun", Jenis: "Peralatan Mandi", Jumlah: 2})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 3, Name: "Beras", Jenis: "Makanan", Jumlah: 4})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 4, Name: "obeng", Jenis: "Peralatan Mobil", Jumlah: 1})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 5, Name: "dongkrak", Jenis: "Peralatan Mobil", Jumlah: 1})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 6, Name: "Tisu", Jenis: "Kosmetik", Jumlah: 10})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 7, Name: "Gula", Jenis: "Makanan", Jumlah: 10})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 8, Name: "Shampo", Jenis: "Peralatan Mandi", Jumlah: 6})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 9, Name: "Bedak", Jenis: "Kosmetik", Jumlah: 6})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 10, Name: "Sikat Gigi", Jenis: "Peralatan Mandi", Jumlah: 3})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 11, Name: "Ikan", Jenis: "Makanan", Jumlah: 30})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 12, Name: "Sabun cuci", Jenis: "Peralatan Mandi", Jumlah: 11})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 13, Name: "Sendol", Jenis: "Peralatan Makanan", Jumlah: 11})
	DaftarBelanja = append(DaftarBelanja, Belanja{ID: 14, Name: "Ember", Jenis: "Peralatan Mandi", Jumlah: 5})
}
