package Models

import (
	DB "go-gin-mysql/Database"
	"log"
	"time"
)

type RiwayatUpdateNilaiPembimbing struct {
	Nama    string `json:"nama" gorm:"column:nama"`
	Kelas   string `json:"kelas"`
	Jurusan string `json:"jurusan"`
	Rombel  string `json:"rombel"`
	NIS     string `json:"nis"`

	NilaiSoftskillIndustri   float32   `json:"nilai_softskill_industri"`
	NilaiHardskillPembimbing float32   `json:"nilai_hardskill_pembimbing"`
	NilaiHardskillIndustri   float32   `json:"nilai_hardskill_industri"`
	NilaiPengujianPembimbing float32   `json:"nilai_pengujian_pembimbing"`
	NamaIndustri             string    `json:"nama_industri"  gorm:"column:nama_industri"`
	AlamatPerusahaan         string    `json:"alamat_industri"  gorm:"column:alamat_industri"`
	NamaPembimbing           string    `json:"nama_pembimbing"`
	UpdatedAtNilaiPembimbing time.Time `json:"update_at_nilai_pembimbing"`
}

type RiwayatUpdateNilaiFasilitator struct {
	Nama    string `json:"nama"`
	Kelas   string `json:"kelas"`
	Jurusan string `json:"jurusan"`
	Rombel  string `json:"rombel"`
	NIS     string `json:"nis"`

	NilaiKemandirianFasilitator float32   `json:"nilai_kemandirian_fasilitator"`
	NilaiSoftskillFasilitator   float32   `json:"nilai_softskill_fasilitator"`
	NamaIndustri                string    `json:"nama_industri"`
	NamaFasilitator             string    `json:"nama_fasilitator"`
	UpdatedAtNilaiFasilitator   time.Time `json:"update_at_nilai_fasilitator"`
}

type JumlahPemangku struct {
	NamaPemangku string `json:"nama" gorm:"column:nama"`
	Total        int    `json:"total"`
}

type JumlahIndustri struct {
	NamaIndustri string `json:"nama"`
	Total        int    `json:"total"`
}

type JumlahSiswaPklJurusan struct {
	Jurusan string `json:"jurusan" gorm:"column:jurusan"`
	Total   int    `json:"total"`
}

func GetRiwayatNilaiPembimbing() ([]RiwayatUpdateNilaiPembimbing, error) {

	var dataRiwayat []RiwayatUpdateNilaiPembimbing

	query := `SELECT data_siswa.nama as nama, data_siswa.kelas as kelas, data_siswa.jurusan as jurusan, data_siswa.rombel, data_siswa.nis as nis,
	nilai_softskill_industri,nilai_hardskill_pembimbing, nilai_hardskill_industri, nilai_pengujian_pembimbing,
	industri.nama as nama_industri,industri.alamat as alamat_industri, pegawai.nama as nama_pembimbing, updated_at_nilai_pembimbing
	FROM data_siswa
	JOIN pegawai on pegawai.id = data_siswa.fk_id_pembimbing
	JOIN industri on industri.id = data_siswa.fk_id_industri
	WHERE updated_at_nilai_pembimbing IS NOT NULL
	ORDER BY updated_at_nilai_pembimbing DESC`

	rows := DB.Database.Raw(query).Scan(&dataRiwayat)
	return dataRiwayat, rows.Error

}

func GetRiwayatNilaiFasilitator() ([]RiwayatUpdateNilaiFasilitator, error) {

	var dataRiwayat []RiwayatUpdateNilaiFasilitator

	query := `SELECT data_siswa.nama as nama, data_siswa.kelas as kelas, data_siswa.jurusan as jurusan, data_siswa.rombel, data_siswa.nis as nis,
	nilai_kemandirian_fasilitator, nilai_softskill_fasilitator,
	industri.nama as nama_industri,industri.alamat as alamat_industri, pegawai.nama as nama_fasilitator, updated_at_nilai_fasilitator
	FROM data_siswa
	JOIN pegawai on pegawai.id = data_siswa.fk_id_fasilitator
	JOIN industri on industri.id = data_siswa.fk_id_industri
	WHERE updated_at_nilai_fasilitator IS NOT NULL
	ORDER BY updated_at_nilai_fasilitator DESC`

	rows := DB.Database.Raw(query).Scan(&dataRiwayat)
	return dataRiwayat, rows.Error

}

func GetJumlahPembimbing() (int, error) {
	var jumlah int64
	query := `SELECT COUNT(kr.id) as jumlah
	FROM konfigurasi_roles kr
	JOIN role r ON r.id = kr.fk_id_role
	WHERE lower(nama) = 'pembimbing'`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return int(jumlah), err
}

func GetJumlahFasilitator() (int, error) {
	var jumlah int64
	query := `SELECT COUNT(kr.id) as jumlah
	FROM konfigurasi_roles kr
	JOIN role r ON r.id = kr.fk_id_role
	WHERE lower(nama) = 'fasilitator'`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return int(jumlah), err
}

func GetJumlahHubin() (int, error) {
	var jumlah int64
	query := `SELECT COUNT(kr.id) as jumlah
	FROM konfigurasi_roles kr
	JOIN role r ON r.id = kr.fk_id_role
	WHERE lower(nama) = 'hubin'`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return int(jumlah), err
}

func GetTotalPemangku() (int, error) {
	var jumlah int64
	query := `SELECT COUNT(DISTINCT(fk_id_pegawai)) as total_pemangku FROM konfigurasi_roles`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return int(jumlah), err
}

func GetTotalIndustri() (int, error) {
	var jumlah int64
	query := `SELECT COUNT(DISTINCT(id)) as total_industri FROM industri`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return int(jumlah), err
}

func GetTotalJurusan() (int, error) {
	var jumlah int64
	query := `SELECT COUNT(DISTINCT(jurusan)) as jurusan FROM data_siswa`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return int(jumlah), err
}

func GetJumlahSiswaPkl() (int64, error) {
	var jumlah int64
	query := `SELECT COUNT(nis) FROM data_siswa WHERE aktif = true`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return jumlah, err
}

func GetJumlahWakel() (int, error) {
	var jumlah int64
	query := `SELECT COUNT(kr.id) as jumlah
	FROM konfigurasi_roles kr
	JOIN role r ON r.id = kr.fk_id_role
	WHERE lower(nama) = 'wali kelas'`

	err := DB.Database.Raw(query).Scan(&jumlah).Error
	return int(jumlah), err
}

func GetJumlahSetiapPemangku() ([]JumlahPemangku, int, error) {

	totalAllPemangku, err := GetTotalPemangku()
	if err != nil {
		return nil, 0, err

	}

	var data []JumlahPemangku
	query := `
	SELECT role.nama AS nama, COUNT(konfigurasi_roles.fk_id_pegawai) AS total FROM konfigurasi_roles 
	JOIN role ON role.id = konfigurasi_roles.fk_id_role 
	GROUP BY role.nama;
	`

	log.Println(data)

	rows := DB.Database.Raw(query).Scan(&data)
	if rows.Error != nil {
		return nil, 0, rows.Error
	}

	return data, totalAllPemangku, nil

}

func GetJumlahSetiapIndustri() ([]JumlahIndustri, int, error) {

	totalAllIndustri, err := GetTotalIndustri()
	if err != nil {
		return nil, 0, err

	}

	var data []JumlahIndustri
	query := `SELECT industri.nama as nama_industri, COUNT(data_siswa.jurusan) as total FROM industri
	LEFT JOIN data_siswa on data_siswa.fk_id_industri = industri.id
	GROUP BY industri.nama`

	log.Println(data)

	rows := DB.Database.Raw(query).Scan(&data)
	if rows.Error != nil {
		return nil, 0, rows.Error
	}

	return data, totalAllIndustri, nil

}

func GetTotalSiswaPklJurusan() ([]JumlahSiswaPklJurusan, int, error) {

	totalAlljurusan, err := GetTotalJurusan()
	if err != nil {
		return nil, 0, err

	}

	var data []JumlahSiswaPklJurusan
	query := `
	SELECT jurusan as jurusan, COUNT(nis) as total from data_siswa GROUP BY jurusan`

	rows := DB.Database.Raw(query).Scan(&data)
	if rows.Error != nil {
		return nil, 0, rows.Error
	}

	return data, totalAlljurusan, nil

}
