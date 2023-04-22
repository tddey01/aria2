package model

import (
	"fmt"
	orm "github.com/tddey01/aria2/drive/mysql"
)

type FilSwan struct {
	FileName    string `gorm:"column:file_name" json:"file_name"`
	FileSize    uint64 `gorm:"column:file_size" json:"file_size"`
	DownloadUrl string `gorm:"column:download_url" json:"download_url"`
	FileActive  string `gorm:"column:file_active" json:"file_active"`
	Locked      string `gorm:"column:locked" json:"locked"`
	GId         string `gorm:"column:gid" json:"gid"`
	CreateTimes string `gorm:"column:create_times" json:"create_times"`
	UpdateTimes string `gorm:"column:update_times" json:"update_times"`
	LocalPath   string `gorm:"column:local_path" json:"local_path"`
	Success     string `gorm:"column:success" json:"success"`
	DiskPath    string `gorm:"column:disk_path" json:"disk_path"`
	Drive       string `gorm:"column:drive" json:"drive"`
}

func GetAll(table string) (ret []*FilSwan, err error) {
	sqlx := `select  * from  ` + table + ` where file_active=0   limit 0,1`
	log.Debug(sqlx)
	if err = orm.Eloquent.Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}
	return
}

func UpdateSetDownload1(msg *FilSwan, gid, drive, table, fileName string) (err error) { // 下载中
	sqlx := `UPDATE  ` + table + ` set   file_name='` + fileName + `',  file_active=1 ,locked=1  ,gid='` + gid + `',create_times=now()   where  download_url='` + msg.DownloadUrl + `' AND  drive = '` + drive + `' `
	log.Debug(sqlx)
	if err = orm.Eloquent.Exec(sqlx).Error; err != nil {
		return
	}
	return
}

func UpdateSetDownload2(msg *FilSwan, gid string, path, drive, table string, size int64) (err error) {
	sqlx := `UPDATE  ` + table + ` set  file_active=2 ,locked=0 ,update_times=now() ,local_path='` + path + `' ,file_size=%d  where download_url='` + msg.DownloadUrl + `' AND gid = '` + gid + `' AND  drive = '` + drive + `'`
	sqlx = fmt.Sprintf(sqlx, size)
	log.Debug(sqlx)
	if err = orm.Eloquent.Debug().Exec(sqlx).Error; err != nil {
		return
	}
	return
}

func UpdateSetDownload3(msg *FilSwan, fileName, path, drive, table string, size int64) (err error) {
	sqlx := `UPDATE  ` + table + ` set file_name='` + fileName + `', file_active=2 ,locked=0 ,update_times=now() ,local_path='` + path + `' ,file_size=%d  where download_url='` + msg.DownloadUrl + `'AND  drive = '` + drive + `'`
	sqlx = fmt.Sprintf(sqlx, size)
	log.Debug(sqlx)
	if err = orm.Eloquent.Debug().Exec(sqlx).Error; err != nil {
		return
	}
	return
}

func GetFindOne(table string) (*FilSwan, error) {

	sk := FilSwan{}
	sqlx := `select * from   ` + table + `  where file_active=0     limit 0,1`
	log.Debug(sqlx)
	if err := orm.Eloquent.Raw(sqlx).Scan(&sk).Error; err != nil {
		return nil, nil
	}
	return &sk, nil
}

func GeTGId(table string) (ret []*FilSwan, err error) {

	sqlx := `select  * from    ` + table + ` where locked=1  `
	log.Debug(sqlx)
	if err = orm.Eloquent.Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}
	return
}

func GeTLocked(table string) (ret []*FilSwan, err error) {

	sqlx := `select  * from    ` + table + ` where locked=1  `
	log.Debug(sqlx)
	if err = orm.Eloquent.Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}
	return
}
