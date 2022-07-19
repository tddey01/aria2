package models

import orm "github.com/tddey01/aria2/drive/mysql"

type FilSwan struct {
	FileName    string `gorm:"column:file_name" json:"file_name"`
	FileSize    uint64 `gorm:"column:file_size" json:"file_size"`
	PieceCid    string `gorm:"column:piece_cid" json:"piece_cid"`
	DataCid     string `gorm:"column:data_cid" json:"data_cid"`
	DownloadUrl string `gorm:"column:download_url" json:"download_url"`
	FileActive  string `gorm:"column:file_active" json:"file_active"`
}

func GetAll() (ret []*FilSwan, err error) {
	sqlx := `select  * from  filswan where file_active=0`
	if err = orm.Eloquent.Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}
	return
}
func UpdateSetDownload1(msg *FilSwan) (err error) {
	sqlx := `UPDATE  filswan set  file_active=1 where data_cid='` + msg.DataCid + `'`
	if err = orm.Eloquent.Exec(sqlx).Error; err != nil {
		return
	}
	return
}

func UpdateSetDownload2(msg *FilSwan) (err error) {
	sqlx := `UPDATE  filswan set  file_active=2 where data_cid='` + msg.DataCid + `'`
	if err = orm.Eloquent.Exec(sqlx).Error; err != nil {
		return
	}
	return
}

func UpdateSetDownload3(msg *FilSwan) (err error) {
	sqlx := `UPDATE  filswan set  file_active=3 where data_cid='` + msg.DataCid + `'`
	if err = orm.Eloquent.Exec(sqlx).Error; err != nil {
		return
	}

	return
}

func GetFindOne() (ret *FilSwan, err error) {
	sqlx := `select * from filswan where file_active=0 limit 0,1`
	if err = orm.Eloquent.Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}
	return
}
