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

func UpdateSetDownload1s(msg *FilSwan, gid, drive, table string) (err error) { // 下载中
	sqlx := `UPDATE  ` + table + ` set  file_active=1  , locked=1  ,gid='` + gid + `',create_times=now()   where download_url='` + msg.DownloadUrl + `' AND  drive = '` + drive + `'`
	log.Debug(sqlx)
	if err = orm.Eloquent.Exec(sqlx).Error; err != nil {
		return
	}
	return
}

func UpdateSetDownload2s(msg *FilSwan, gid string, path, drive, table string, size int64) (err error) {
	sqlx := `UPDATE  ` + table + ` set  file_active=2 ,locked=0 ,update_times=now() ,local_path='` + path + `' ,file_size=%d   where download_url='` + msg.DownloadUrl + `' AND gid = '` + gid + `' AND  drive = '` + drive + `'`
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

type Dw struct {
	Downloading string `gorm:"column:downloading" json:"downloading"`
	Downloaded  string `gorm:"column:downloaded" json:"downloaded"`
	Total       string `gorm:"column:total" json:"total"`
	Atcv        string `gorm:"column:actv" json:"actv"`
	Successful  string `gorm:"column:success" json:"success"`
}

func GetCount(table string) (ret []*Dw, err error) {

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
