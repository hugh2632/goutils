package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"sync"
	"testing"
)

func BenchmarkTest(b *testing.B) {
	var i = 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			resp, err := http.Get("http://192.168.33.156:10085/api/v1/GetPdfBytes?http://192.168.33.156:8089/Reports/pdfExport/Report_vul.aspx?tid=20671098-b9c1-4620-9003-0000796efe25&uid=219a9680-8140-43a9-9ac8-ce9227517226")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			a, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			var aaa A

			// unmarschal JSON
			err = json.Unmarshal(a, &aaa)
			if err != nil {
				panic(err)
			}
			decoded, err := base64.StdEncoding.DecodeString(aaa.Data)
			if err != nil {
				panic(err)
			}
			f, err := os.Create("D:\\a" + strconv.Itoa(i) + ".pdf")
			if err != nil {
				panic(err)
			}
			_, err = f.Write(decoded)
			if err != nil {
				panic(err)
			}
			err = f.Close()
			if err != nil {
				panic(err)
			}
			i++
		}
	})
}

func TestA(t *testing.T) {
	var g = sync.WaitGroup{}
	g.Add(300)
	for i := 0; i < 300; i++ {
		go func(j int){
			resp, err := http.Get("http://192.168.33.156:10085/api/v1/GetPdfBytes?http://192.168.33.156:8089/Reports/pdfExport/Report_vul.aspx?tid=20671098-b9c1-4620-9003-0000796efe25&uid=219a9680-8140-43a9-9ac8-ce9227517226")
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			a, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}
			var aaa A
			// unmarschal JSON
			err = json.Unmarshal(a, &aaa)
			if err != nil {
				panic(err)
			}
			decoded, err := base64.StdEncoding.DecodeString(aaa.Data)
			if err != nil {
				panic(err)
			}
			f, err := os.Create("D:\\a" + strconv.Itoa(j)  + ".pdf")
			if err != nil {
				panic(err)
			}
			_, err = f.Write(decoded)
			if err != nil {
				panic(err)
			}
			err = f.Close()
			if err != nil {
				panic(err)
			}
			g.Done()
		}(i)
	}
	g.Wait()
}

type A struct {
	Iserror bool
	Msg     string
	Data    string
}

func TestB(t *testing.T){
	var i = 0
	t.Log("D:\\a" + strconv.Itoa(i) + ".pdf")
}