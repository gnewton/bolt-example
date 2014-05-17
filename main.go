package main

// Copyright 2014 Glen Newton
// Apache2 license
//
// Growing memory in bolt example

import (
	"github.com/boltdb/bolt"
	"strconv"
	"fmt"
	"time"
	"runtime"
)


const BUCKET = "widgets"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	db, err := bolt.Open("./test.bolt.db2", 0666)
	defer db.Close()
	if err != nil{
		fmt.Println(err)
		return
	}

	tx, err := db.Begin(true)
	if err != nil{
		fmt.Println(err)
		return
	}

	commitSize := 100000
	infoSize := 100000

	b, err := tx.CreateBucket([]byte(BUCKET))
	if err != nil{
		fmt.Println(err)
		return
	}
	var startTime = time.Now()
	for i:=0; i<1000000000; i++{
		key := "foo" + strconv.Itoa(i) + "M"
		b.Put([]byte(key), []byte("bar"))
		key = "baz" + strconv.Itoa(i) + "Z"
		b.Put([]byte(key), []byte("bat"))
		//b.Delete([]byte("foo"))
		if i%commitSize== 0 && i != 0{
			err = tx.Commit()
			if err != nil{
				fmt.Println(err)
				return
			}
			tx, err = db.Begin(true)
			b = tx.Bucket([]byte(BUCKET))
			if err != nil{
				fmt.Println(err)
				return
			}
		}
		if i%infoSize == 0 && i != 0{
			fmt.Println(i, time.Since(startTime))
			startTime = time.Now()
		}
	}



}