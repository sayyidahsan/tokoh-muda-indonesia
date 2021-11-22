// Code generated by sqlc. DO NOT EDIT.
// source: mahasiswa.sql

package db

import (
	"context"
)

const createMahasiwa = `-- name: CreateMahasiwa :exec
INSERT INTO mahasiswa (
  id, url_foto, nama_lengkap, tanggal_lahir, jenis_kelamin, asal_kampus, nim, jurusan, tahun_masuk, kota_kabupaten, id_tantangan, id_silabus, id_event
) VALUES (
  ?,?,?,?,?,?,?,?,?,?,?,?,?
)
`

type CreateMahasiwaParams struct {
	ID            int32  `json:"id"`
	UrlFoto       string `json:"url_foto"`
	NamaLengkap   string `json:"nama_lengkap"`
	TanggalLahir  string `json:"tanggal_lahir"`
	JenisKelamin  string `json:"jenis_kelamin"`
	AsalKampus    string `json:"asal_kampus"`
	Nim           string `json:"nim"`
	Jurusan       string `json:"jurusan"`
	TahunMasuk    string `json:"tahun_masuk"`
	KotaKabupaten string `json:"kota_kabupaten"`
	IDTantangan   int    `json:"id_tantangan"`
	IDSilabus     int    `json:"id_silabus"`
	IDEvent       int    `json:"id_event"`
}

func (q *Queries) CreateMahasiwa(ctx context.Context, arg CreateMahasiwaParams) error {
	_, err := q.db.ExecContext(ctx, createMahasiwa,
		arg.ID,
		arg.UrlFoto,
		arg.NamaLengkap,
		arg.TanggalLahir,
		arg.JenisKelamin,
		arg.AsalKampus,
		arg.Nim,
		arg.Jurusan,
		arg.TahunMasuk,
		arg.KotaKabupaten,
		arg.IDTantangan,
		arg.IDSilabus,
		arg.IDEvent,
	)
	return err
}