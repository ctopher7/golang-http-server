package model

import "time"

type Berat struct {
	Id        int32     `json:"id,omitempty"`
	Tanggal   time.Time `json:"tanggal"`
	Max       float64   `json:"max"`
	Min       float64   `json:"min"`
	Perbedaan float64   `json:"perbedaan"`
}

type GetBeratRes struct {
	RataRata Berat   `json:"rata_rata"`
	Data     []Berat `json:"data"`
}

type GetBeratReq struct {
	Id int32 `json:"id"`
}
