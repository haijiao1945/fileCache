package fileCache

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

type Driver struct {
}

func (driver Driver) Read(key string) (string, error) {
	key = driver.cleanKey(key)
	path := config.SavePath + key + config.extName
	stat, err := os.Stat(path + "t")
	if err == nil {
		expire, _ := ioutil.ReadFile(path + "t")
		oldtime := stat.ModTime().Unix()
		nowtime := time.Now().Unix()
		_sexpire := string(expire)
		_expire, _ := strconv.Atoi(_sexpire)
		_i64expire := int64(_expire)
		if (nowtime - _i64expire) > oldtime {
			os.Remove(path)
			os.Remove(path + "t")
			return "", nil
		}
	}

	res, err := ioutil.ReadFile(path)

	return string(res), err
}

func (driver Driver) Write(key string, value string, expire int) (bool, error) {
	key = driver.cleanKey(key)
	path := config.SavePath + key + config.extName
	ishave, _ := driver.pathExist(path)
	if !ishave {
		file, _ := os.Create(path)
		defer file.Close()
	}
	err := ioutil.WriteFile(path, []byte(value), 0644)
	if expire > 0 {
		_expire := strconv.Itoa(expire)
		ioutil.WriteFile(path+"t", []byte(_expire), 0644)
	}
	res := true
	if err != nil {
		res = false
	}
	return res, err
}

func (driver Driver) cleanKey(key string) string {
	m := md5.New()
	m.Write([]byte(key))
	return hex.EncodeToString(m.Sum(nil))
}

func (driver Driver) pathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	res := false
	if err != nil {
		res = true
	}
	return res, err
}
