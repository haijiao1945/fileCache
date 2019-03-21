# fileCache

拿go练手学习用

# 咋用这个玩意儿?

import ("github.com/haijiao1945/fileCache")

cache := new(fileCache.Cache)
db := cache.DB()
db.SetConfig("")

res, err := db.SetEx("111", "22222", 5) #设置有过期时间的
res, err := db.Set("111", "22222") #设置没有过期时间的
val, err := db.Get("111") #读取
