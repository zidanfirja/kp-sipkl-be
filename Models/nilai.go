package Models

import (
	"fmt"
	DB "go-gin-mysql/Database"
	"strings"
	"time"
)

type IndustriPembimbingFasil struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type NilaiSiswaPkl struct {
	NIS     string `json:"nis"`
	Nama    string `json:"nama"`
	Kelas   string `json:"kelas"`
	Jurusan string `json:"jurusan"`
	Rombel  string `json:"rombel"`

	NilaiSoftskillFasilitator   float32 `json:"nilai_softskill_fasilitator" gorm:"type:float"`
	NilaiSoftskillIndustri      float32 `json:"nilai_softskill_industri" gorm:"type:float"`
	NilaiHardskillPembimbing    float32 `json:"nilai_hardskill_pembimbing" gorm:"type:float"`
	NilaiHardskillIndustri      float32 `json:"nilai_hardskill_industri" gorm:"type:float"`
	NilaiKemandirianFasilitator float32 `json:"nilai_kemandirian_fasilitator" gorm:"type:float"`
	NilaiPengujianPembimbing    float32 `json:"nilai_pengujian_pembimbing" gorm:"type:float"`

	TanggalMasuk  *time.Time `json:"tanggal_masuk" gorm:"date"`
	TanggalKeluar *time.Time `json:"tanggal_keluar" gorm:"date"`

	NamaIndustri string `json:"nama_industri" gorm:"type:varchar(255)" `
	Alamat       string `json:"alamat" `
}

type NilaiSiswaPklPembimbing struct {
	NIS                      string `json:"nis"`
	Nama                     string `json:"nama"`
	Kelas                    string `json:"kelas"`
	Jurusan                  string `json:"jurusan"`
	Rombel                   string `json:"rombel"`
	NilaiSoftskillIndustri   int    `json:"nilai_softskill_industri" gorm:"type:float"`
	NilaiHardskillIndustri   int    `json:"nilai_hardskill_industri" gorm:"type:float"`
	NilaiHardskillPembimbing int    `json:"nilai_hardskill_pembimbing" gorm:"type:float"`
	NilaiPengujianPembimbing int    `json:"nilai_pengujian_pembimbing" gorm:"type:float"`
}

type NilaiSiswaPklFasilitator struct {
	NIS                         string  `json:"nis"`
	Nama                        string  `json:"nama"`
	Kelas                       string  `json:"kelas"`
	Jurusan                     string  `json:"jurusan"`
	Rombel                      string  `json:"rombel"`
	NilaiSoftskillFasilitator   float32 `json:"nilai_softskill_fasilitator" gorm:"type:float"`
	NilaiKemandirianFasilitator float32 `json:"nilai_kemandirian_fasilitator" gorm:"type:float"`
}

type IndustriForNilai struct {
	ID            int       `json:"id_perusahaan" gorm:"column:id_perusahaan"`
	Nama          string    `json:"nama_perusahaan" gorm:"column:nama_perusahaan"`
	Alamat        string    `json:"alamat_perusahaan" gorm:"column:alamat_perusahaan"`
	TanggalMasuk  time.Time `json:"tanggal_masuk" gorm:"column:tanggal_masuk"`
	TanggalKeluar time.Time `json:"tanggal_keluar" gorm:"column:tanggal_keluar"`
}

type ReqUpdateNilaiPembimbing struct {
	NIS                      string  `json:"nis"`
	NilaiSoftskillIndustri   float32 `json:"nilai_softskill_industri" gorm:"type:float"`
	NilaiHardskillIndustri   float32 `json:"nilai_hardskill_industri" gorm:"type:float"`
	NilaiHardskillPembimbing float32 `json:"nilai_hardskill_pembimbing" gorm:"type:float"`
	NilaiPengujianPembimbing float32 `json:"nilai_pengujian_pembimbing" gorm:"type:float"`
}

type ReqUpdateNilaiFasilitator struct {
	NIS                         string  `json:"nis"`
	NilaiSoftskillFasilitator   float32 `json:"nilai_softskill_fasilitator" gorm:"type:float"`
	NilaiKemandirianFasilitator float32 `json:"nilai_kemandirian_fasilitator" gorm:"type:float"`
}

func GetIndustriPembimbing(id int) ([]IndustriPembimbingFasil, error) {
	var data []IndustriPembimbingFasil

	query := `SELECT fk_id_industri AS id, industri.nama AS nama
	FROM data_siswa 
	JOIN pegawai ON pegawai.id = data_siswa.fk_id_pembimbing
	JOIN industri ON industri.id = data_siswa.fk_id_industri
	WHERE fk_id_pembimbing = ?
	GROUP BY fk_id_industri, industri.nama;`

	rows := DB.Database.Raw(query, id).Scan(&data)
	if rows.Error != nil {
		return nil, rows.Error
	}

	return data, nil
}

func GetIndustriFasilitator(id int) ([]IndustriPembimbingFasil, error) {
	var data []IndustriPembimbingFasil

	query := `SELECT fk_id_industri AS id, industri.nama AS nama
	FROM data_siswa 
	JOIN pegawai ON pegawai.id = data_siswa.fk_id_fasilitator
	JOIN industri ON industri.id = data_siswa.fk_id_industri
	WHERE fk_id_fasilitator = ?
    GROUP BY fk_id_industri,  industri.nama `

	rows := DB.Database.Raw(query, id).Scan(&data)
	if rows.Error != nil {
		return nil, rows.Error
	}

	return data, nil
}

func GetIndustri(id_industri int) (IndustriForNilai, error) {
	var dataIndustri IndustriForNilai

	query := `SELECT industri.id AS id_perusahaan, industri.nama AS nama_perusahaan,industri.alamat AS alamat_perusahaan, data_siswa.tanggal_masuk, data_siswa.tanggal_keluar from data_siswa 
    JOIN pegawai on pegawai.id = data_siswa.fk_id_pembimbing
    JOIN industri on industri.id = data_siswa.fk_id_industri
    WHERE fk_id_industri = ?
    LIMIT 1`

	rows := DB.Database.Raw(query, id_industri).Scan(&dataIndustri)
	if rows.Error != nil {
		return dataIndustri, rows.Error
	}

	return dataIndustri, nil

}

func GetNilaiByPemb(id_pembimbing, id_industri int) ([]NilaiSiswaPklPembimbing, error) {

	var nilai []NilaiSiswaPklPembimbing

	query := ` SELECT 
    nis,nama,kelas,jurusan,rombel,
    nilai_softskill_industri,
    nilai_hardskill_industri,
    nilai_hardskill_pembimbing,
    nilai_pengujian_pembimbing
    FROM data_siswa
    WHERE fk_id_pembimbing = ? AND fk_id_industri = ?`

	rows := DB.Database.Raw(query, id_pembimbing, id_industri).Scan(&nilai)
	if rows.Error != nil {
		return nilai, rows.Error
	}

	return nilai, nil

}

func GetNilaiByFasil(id_fasil, id_industri int) ([]NilaiSiswaPklFasilitator, error) {

	var nilai []NilaiSiswaPklFasilitator

	query := `SELECT 
    nis,nama,kelas,jurusan,rombel,
    nilai_softskill_fasilitator,
    nilai_kemandirian_fasilitator
    FROM data_siswa
    WHERE fk_id_fasilitator = ? AND fk_id_industri = ?`

	rows := DB.Database.Raw(query, id_fasil, id_industri).Scan(&nilai)
	if rows.Error != nil {
		return nil, rows.Error
	}

	return nilai, nil

}

func UpdateNilaiPembimbing(data *[]ReqUpdateNilaiPembimbing) error {

	var listNis []string
	var caseNilaiSoftskillIndustri, caseNilaiHardskillIndustri, caseNilaiHardskillPembimbing, caseNilaiPengujianPembimbing string

	for _, dataNilai := range *data {
		listNis = append(listNis, fmt.Sprintf("'%s'", dataNilai.NIS))
		caseNilaiSoftskillIndustri += fmt.Sprintf("WHEN '%s' THEN %f ", dataNilai.NIS, dataNilai.NilaiSoftskillIndustri)
		caseNilaiHardskillIndustri += fmt.Sprintf("WHEN '%s' THEN %f ", dataNilai.NIS, dataNilai.NilaiHardskillIndustri)
		caseNilaiHardskillPembimbing += fmt.Sprintf("WHEN '%s' THEN %f ", dataNilai.NIS, dataNilai.NilaiHardskillPembimbing)
		caseNilaiPengujianPembimbing += fmt.Sprintf("WHEN '%s' THEN %f ", dataNilai.NIS, dataNilai.NilaiPengujianPembimbing)
	}

	query := fmt.Sprintf(`
	UPDATE data_siswa
	SET 
	nilai_softskill_industri = CASE nis %s END,
	nilai_hardskill_industri = CASE nis %s END,
	nilai_hardskill_pembimbing = CASE nis %s END,
	nilai_pengujian_pembimbing = CASE nis %s END,
	updated_at_nilai_pembimbing = NOW()
	WHERE nis IN (%s);
`, caseNilaiSoftskillIndustri, caseNilaiHardskillIndustri, caseNilaiHardskillPembimbing, caseNilaiPengujianPembimbing, strings.Join(listNis, ", "))

	if err := DB.Database.Exec(query).Error; err != nil {
		return err
	}

	return nil
}

func UpdateNilaiFasilitator(data *[]ReqUpdateNilaiFasilitator) error {

	var listNis []string
	var caseNilaiSoftskillFasilitator, caseNilaiKemandirianFasilitator string

	// NilaiSoftskillFasilitator   float64 `json:"nilai_softskill_fasilitator"`
	// NilaiKemandirianFasilitator float64 `json:"nilai_kemandirian_fasilitator"`

	for _, dataNilai := range *data {
		listNis = append(listNis, fmt.Sprintf("'%s'", dataNilai.NIS))
		caseNilaiSoftskillFasilitator += fmt.Sprintf("WHEN '%s' THEN %f ", dataNilai.NIS, dataNilai.NilaiSoftskillFasilitator)
		caseNilaiKemandirianFasilitator += fmt.Sprintf("WHEN '%s' THEN %f ", dataNilai.NIS, dataNilai.NilaiKemandirianFasilitator)

	}

	query := fmt.Sprintf(`
	UPDATE data_siswa
	SET 
	nilai_softskill_fasilitator = CASE nis %s END,
	nilai_kemandirian_fasilitator = CASE nis %s END,
	updated_at_nilai_fasilitator = NOW()
	WHERE nis IN (%s);
`, caseNilaiSoftskillFasilitator, caseNilaiKemandirianFasilitator, strings.Join(listNis, ", "))

	if err := DB.Database.Exec(query).Error; err != nil {
		return err
	}

	return nil
}

func GetNilaiWakel(kelas, jurusan, rombel string) ([]NilaiSiswaPkl, error) {
	var dataNilai []NilaiSiswaPkl

	query := `
	SELECT nis, data_siswa.nama as nama, kelas, data_siswa.jurusan as  jurusan, rombel, 
	nilai_softskill_fasilitator,nilai_softskill_industri, 
	nilai_hardskill_pembimbing, nilai_hardskill_industri,
	nilai_kemandirian_fasilitator, nilai_pengujian_pembimbing,
    tanggal_masuk, tanggal_keluar,
    industri.nama as nama_industri, industri.alamat as alamat
	FROM data_siswa
    JOIN industri on industri.id  = data_siswa.fk_id_industri
	WHERE data_siswa.kelas = ? AND data_siswa.jurusan = ? AND rombel = ?`

	rows := DB.Database.Raw(query, kelas, jurusan, rombel).Scan(&dataNilai)
	if rows.Error != nil {
		return nil, rows.Error
	}

	return dataNilai, nil
}
