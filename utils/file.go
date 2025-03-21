package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func IsFileExists(filePath, fileName string) bool {
	fileFullPath := filepath.Join(filePath, fileName)
	_, err := os.Stat(fileFullPath)

	if err != nil {
		log.Info(err)
		return false
	}

	return true
}

func IsFileExistsFullPath(fileFullPath string) bool {
	_, err := os.Stat(fileFullPath)

	if err != nil {
		log.Info(err)
		return false
	}

	return true
}

func IsPathFile(dirFullPath string) (*bool, error) {
	fi, err := os.Stat(dirFullPath)

	if err != nil {
		log.Info(err)
		return nil, err
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		isFile := false
		return &isFile, nil
	case mode.IsRegular():
		isFile := true
		return &isFile, nil
	default:
		err := fmt.Errorf("unknown path type")
		log.Error(err)
		return nil, err
	}
}

func GetPathType(dirFullPath string) int {
	fi, err := os.Stat(dirFullPath)

	if err != nil {
		log.Info(err)
		return PATH_TYPE_NOT_EXIST
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		return PATH_TYPE_DIR
	case mode.IsRegular():
		return PATH_TYPE_FILE
	default:
		return PATH_TYPE_UNKNOWN
	}
}

const (
	PATH_TYPE_NOT_EXIST = 0 //this path not exists
	PATH_TYPE_FILE      = 1 //file
	PATH_TYPE_DIR       = 2 //directory
	PATH_TYPE_UNKNOWN   = 3 //unknown path type
)

func RemoveFile(filePath, fileName string) {
	fileFullPath := filepath.Join(filePath, fileName)
	err := os.Remove(fileFullPath)
	if err != nil {
		log.Error(err.Error())
	}
}

func GetFileSize(fileFullPath string) int64 {
	fi, err := os.Stat(fileFullPath)
	if err != nil {
		log.Info(err)
		return -1
	}

	return fi.Size()
}

func GetFileSize2(dir, fileName string) int64 {
	fileFullPath := filepath.Join(dir, fileName)
	fi, err := os.Stat(fileFullPath)
	if err != nil {
		log.Info(err)
		return -1
	}

	return fi.Size()
}

func CopyFile(srcFilePath, destFilePath string) (int64, error) {
	sourceFileStat, err := os.Stat(srcFilePath)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		err = errors.New(srcFilePath + " is not a regular file")
		log.Error(err)
		return 0, err
	}

	source, err := os.Open(srcFilePath)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	defer source.Close()

	destination, err := os.Create(destFilePath)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	defer destination.Close()

	nBytes, err := io.Copy(destination, source)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return nBytes, err
}

func CreateFileWithContents(filefullpath string, lines []string) (int, error) {
	f, err := os.Create(filefullpath)

	if err != nil {
		log.Error(err)
		return 0, nil
	}

	defer f.Close()

	bytesWritten := 0
	for _, line := range lines {
		bytesWritten1, err := f.WriteString(line + "\n")
		if err != nil {
			log.Error(err)
			return 0, nil
		}
		bytesWritten = bytesWritten + bytesWritten1
	}

	if err != nil {
		log.Error(err)
		return 0, nil
	}

	log.Info(filefullpath, " generated.")
	return bytesWritten, nil
}

func CreateFileWithByteContents(filefullpath string, contents []byte) (int, error) {
	f, err := os.Create(filefullpath)

	if err != nil {
		log.Error(err)
		return 0, nil
	}

	defer f.Close()

	bytesWritten, err := f.Write(contents)
	if err != nil {
		log.Error(err)
		return 0, nil
	}

	log.Info(filefullpath, " generated.")
	return bytesWritten, nil
}

func ReadAllLines(dir, filename string) ([]string, error) {
	fileFullPath := filepath.Join(dir, filename)

	file, err := os.Open(fileFullPath)

	if err != nil {
		log.Error("failed opening file: ", fileFullPath)
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func ReadFile(filePath string) (string, []byte, error) {
	sourceFileStat, err := os.Stat(filePath)
	if err != nil {
		log.Error(err)
		return "", nil, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		err = errors.New(filePath + " is not a regular file")
		log.Error(err)
		return "", nil, err
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error("failed reading data from file: ", filePath)
		log.Error(err)
		return "", nil, err
	}

	return sourceFileStat.Name(), data, nil
}

func IsDirExists(dir string) bool {
	if IsStrEmpty(&dir) {
		err := fmt.Errorf("dir is not provided")
		log.Error(err)
		return false
	}

	if GetPathType(dir) != PATH_TYPE_DIR {
		return false
	}

	return true
}

func CreateDir(dir string) error {
	if len(dir) == 0 {
		err := fmt.Errorf("dir is not provided")
		log.Info(err)
		return err
	}

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		err := fmt.Errorf("%s, failed to create output dir:%s", err.Error(), dir)
		log.Error(err)
		return err
	}

	return nil
}

func GenerateFile(filepath, filename string, filesize int64) {
	filefullpath := filepath + "/" + filename
	file, err := os.Create(filefullpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Info("start to generate file:", filefullpath, ", target size:", filesize, "GB")

	filesizeInByte := filesize * 100000000
	var i int64
	for i = 0; i < filesizeInByte; i++ {
		_, err := file.WriteString("Hello World")
		if err != nil {
			fmt.Println(err)
			file.Close()
			return
		}
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Info("file:", filefullpath, " generated, size:", filesize, "GB")
}

func CreateDirIfNotExists(dir, dirName string) error {
	if IsStrEmpty(&dir) {
		err := fmt.Errorf("%s directory is required", dirName)
		log.Info(err)
		return err
	}

	if IsDirExists(dir) {
		return nil
	}

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		err := fmt.Errorf("failed to create %s directory:%s,%s", dirName, dir, err.Error())
		log.Error(err)
		return err
	}

	log.Info(dirName, " directory: ", dir, " created")
	return nil
}

func CheckDirExists(dir, dirName string) error {
	if IsStrEmpty(&dir) {
		err := fmt.Errorf("%s directory is required", dirName)
		log.Error(err)
		return err
	}

	if !IsDirExists(dir) {
		err := fmt.Errorf("%s directory:%s not exists", dirName, dir)
		log.Error(err)
		return err
	}

	return nil
}

func GetFilesSize(dir string) (*int64, error) {
	srcFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	srcFilesSize := int64(0)
	for _, srcFile := range srcFiles {
		srcFilesSize = srcFilesSize + srcFile.Size()
	}

	return &srcFilesSize, nil
}
