package model

import (
	"fmt"
	"github.com/tddey01/aria2/config"
	orm "github.com/tddey01/aria2/drive/mysql"
)

type FilSwan struct {
	FileName    string `gorm:"column:file_name" json:"file_name"`
	FileSize    uint64 `gorm:"column:file_size" json:"file_size"`
	PieceCid    string `gorm:"column:piece_cid" json:"piece_cid"`
	DataCid     string `gorm:"column:data_cid" json:"data_cid"`
	DownloadUrl string `gorm:"column:download_url" json:"download_url"`
	FileActive  string `gorm:"column:file_active" json:"file_active"`
	//FileError   string `gorm:"column:file_error" json:"file_error"`
	Locked          string `gorm:"column:locked" json:"locked"`
	GId             string `gorm:"column:gid" json:"gid"`
	Import          string `gorm:"column:import_successful" json:"import_successful"`
	CreateTimes     string `gorm:"column:create_times" json:"create_times"`
	UpdateTimes     string `gorm:"column:update_times" json:"update_times"`
	LocalPath       string `gorm:"column:local_path" json:"local_path"`
	Successful      string `gorm:"column:successful" json:"successful"`
	TimesSuccessful string `gorm:"column:times_successful" json:"times_successful"`
}

func GetAll() (ret []*FilSwan, err error) {
	table := config.GetConfig().Mysql.Table
	sqlx := `select  * from  ` + table + ` where file_active=0  limit 0,1`
	log.Debug(sqlx)
	if err = orm.Eloquent.Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}
	return
}
func UpdateSetDownload1(msg *FilSwan, gid string) (err error) { // 下载中
	table := config.GetConfig().Mysql.Table
	sqlx := ``
	switch {
	case config.GetConfig().Typeof.FilSwan:
		sqlx = `UPDATE  ` + table + ` set  file_active=1 ,locked=1  ,gid='` + gid + `',create_times=now()   where data_cid='` + msg.DataCid + `'`

	case config.GetConfig().Typeof.BiGd:
		sqlx = `UPDATE  ` + table + ` set  file_active=1 ,locked=1  ,gid='` + gid + `',create_times=now()   where download_url='` + msg.DownloadUrl + `'`
	}
	log.Debug(sqlx)
	if err = orm.Eloquent.Exec(sqlx).Error; err != nil {
		return
	}
	return
}

func UpdateSetDownload2(msg *FilSwan, gid string, path string, size int64) (err error) {
	table := config.GetConfig().Mysql.Table

	sqlx := ``
	switch {
	case config.GetConfig().Typeof.FilSwan:
		sqlx = `UPDATE  ` + table + `  set  file_active=2 ,locked=0 ,update_times=now() ,local_path='` + path + `' ,file_size=%d  where data_cid='` + msg.DataCid + `' AND gid = '` + gid + `'`

	case config.GetConfig().Typeof.BiGd:
		sqlx = `UPDATE  ` + table + ` set  file_active=2 ,locked=0 ,update_times=now() ,local_path='` + path + `' ,file_size=%d  where download_url='` + msg.DownloadUrl + `' AND gid = '` + gid + `'`
	}
	sqlx = fmt.Sprintf(sqlx, size)
	log.Debug(sqlx)
	if err = orm.Eloquent.Debug().Exec(sqlx).Error; err != nil {
		return
	}
	return
}

func UpdateSetDownload1s(msg *FilSwan, gid string) (err error) { // 下载中
	table := config.GetConfig().Mysql.Table
	sqlx := ``
	switch {
	case config.GetConfig().Typeof.FilSwan:
		sqlx = `UPDATE  ` + table + ` set  file_active=1  ,locked=1 , gid='` + gid + `',create_times=now()   where data_cid='` + msg.DataCid + `'`

	case config.GetConfig().Typeof.BiGd:
		sqlx = `UPDATE  ` + table + ` set  file_active=1  , locked=1  ,gid='` + gid + `',create_times=now()   where download_url='` + msg.DownloadUrl + `'`
	}
	log.Debug(sqlx)
	if err = orm.Eloquent.Exec(sqlx).Error; err != nil {
		return
	}
	return
}

func UpdateSetDownload2s(msg *FilSwan, gid string, path string, size int64) (err error) {
	table := config.GetConfig().Mysql.Table

	sqlx := ``
	switch {
	case config.GetConfig().Typeof.FilSwan:
		sqlx = `UPDATE  ` + table + `  set  file_active=2 ,locked=0 ,update_times=now() ,local_path='` + path + `' ,file_size=%d    where data_cid='` + msg.DataCid + `' AND gid = '` + gid + `'`

	case config.GetConfig().Typeof.BiGd:
		sqlx = `UPDATE  ` + table + ` set  file_active=2 ,locked=0 ,update_times=now() ,local_path='` + path + `' ,file_size=%d   where download_url='` + msg.DownloadUrl + `' AND gid = '` + gid + `'`
	}
	sqlx = fmt.Sprintf(sqlx, size)
	log.Debug(sqlx)
	if err = orm.Eloquent.Debug().Exec(sqlx).Error; err != nil {
		return
	}
	return
}

func GetFindOne() (*FilSwan, error) {
	table := config.GetConfig().Mysql.Table
	sk := FilSwan{}
	sqlx := `select * from   ` + table + `  where file_active=0   limit 0,1`
	log.Debug(sqlx)
	if err := orm.Eloquent.Raw(sqlx).Scan(&sk).Error; err != nil {
		return nil, nil
	}
	return &sk, nil
}

func GeTGId() (ret []*FilSwan, err error) {
	table := config.GetConfig().Mysql.Table
	sqlx := `select  * from    ` + table + ` where locked=1 `
	log.Debug(sqlx)
	if err = orm.Eloquent.Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}
	return
}

func GeTLocked() (ret []*FilSwan, err error) {
	table := config.GetConfig().Mysql.Table
	sqlx := `select  * from    ` + table + ` where locked=1 `
	log.Debug(sqlx)
	if err = orm.Eloquent.Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}
	return
}

type Dw struct {
	Downloading string `gorm:"column:downloading" json:"downloading"`
	Downloaded  string `gorm:"column:downloaded" json:"downloaded"`
	Total       string `gorm:"column:total" json:"total"`
	Atcv        string `gorm:"column:actv" json:"actv"`
	Successful  string `gorm:"column:success" json:"success"`
}

func GetCount() (ret []*Dw, err error) {
	table := config.GetConfig().Mysql.Table
	//sqlx := `select sum(skt) as downloading  ,sum(stk) as downloaded , sum(toal) total   from  (
	//select  count(data_cid) as  skt ,0 as stk , 0 as toal  from   ` + table + ` where  locked=1
	//union  all
	//select  0 as skt , count(data_cid) as stk  ,0 as toal from  ` + table + ` where file_active=2
	//union  all
	//select    0 as skt , 0 as stk , count(data_cid) as toal from  ` + table + `
	//)a`
	//sqlx := `select sum(skt) as downloading  ,sum(stk) as downloaded ,sum(toal) total   from  (
	//select  count(data_cid) as  skt ,0 as stk , 0 as toal  from  filswan where  locked=1
	//union  all
	//select  count(data_cid) as  skt ,0 as stk , 0 as toal  from  filswan197 where  locked=1
	//union  all
	//select  0 as skt , count(data_cid) as stk  ,0 as toal from filswan    where file_active=2 AND import_successful<>1
	//union  all
	//select    0 as skt , 0 as stk , count(data_cid) as toal from filswan  where import_successful<>1
	//union  all
	//select  0 as skt , count(data_cid) as stk  ,0 as toal from filswan197    where file_active=2 AND import_successful<>1
	//union  all
	//select    0 as skt , 0 as stk , count(data_cid) as toal from filswan197  where import_successful<>1
	//)a`
	//sqlx := `select sum(skt) as downloading  ,sum(stk) as downloaded ,sum(toal) total  ,sum(av)  as actv from  (
	//select  count(data_cid) as  skt ,0 as stk , 0 as toal ,0 as av  from  filswan where  locked=1
	//union  all
	//select  count(data_cid) as  skt ,0 as stk , 0 as toal,0 as av  from  filswan197 where  locked=1
	//union  all
	//select  0 as skt , count(data_cid) as stk  ,0 as toal,0 as av from filswan    where file_active=2 AND import_successful<>1
	//union  all
	//select    0 as skt , 0 as stk , count(data_cid) as toal,0 as av from filswan  where import_successful<>1
	//union  all
	//select  0 as skt , count(data_cid) as stk  ,0 as toal,0 as av from filswan197   where file_active=2 AND import_successful<>1
	//union  all
	//select    0 as skt , 0 as stk , count(data_cid) as toal ,0 as av  from filswan197  where import_successful<>0
	//union  all
	//select    0 as skt , 0 as stk , 0 as toal ,count(data_cid) as av  from filswan  where import_successful<>0
	//)a`
	sqlx := `select sum(skt) as downloading  ,sum(stk) as downloaded ,sum(toal) total  ,sum(av)   as actv ,sum(suc) as success from  (
		select  count(data_cid) as  skt ,0 as stk , 0 as toal ,0 as av ,0 as suc from  '` + table + `' where  locked=1
		union  all
		select  0 as skt , count(data_cid) as stk  ,0 as toal,0 as av,0 as suc  from '` + table + `'    where file_active=2
		union  all
		select    0 as skt , 0 as stk , count(data_cid) as toal,0 as av,0 as suc  from '` + table + `'
		union  all
		select    0 as skt , 0 as stk , 0 as toal ,count(data_cid) as av,0 as suc   from '` + table + `'  where import_successful<>0
		union  all
		select    0 as skt , 0 as stk , 0 as toal ,0 as av,count(data_cid) as suc  from '` + table + `'  where successful<>0
	)a`
	log.Debug(sqlx)
	if err = orm.Eloquent.Raw(sqlx).Scan(&ret).Error; err != nil {
		return
	}
	return
}
