package main

import (
	//	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	userid  string
	name    string
	address string
	reason  string
}

func arrayToString(arr []string) string {

	// seperating string elements with -
	return strings.Join([]string(arr), "-")
}

func main() {
	//kalimat := "saya mau makan"
	// var caricari map[string]string
	var checking = false
	var people = []Person{
		{userid: "1", name: "Budi", address: "Jalan 1", reason: "Looking for experience"},
		{userid: "2", name: "Anto", address: "Jalan 2", reason: "Looking for job"},
		{userid: "3", name: "Arie", address: "Jalan 3", reason: "To support final project"},
		{userid: "4", name: "Bimo", address: "Jalan 4", reason: "Expand more networks"},
	}
	//	caricari := make(map[string]string)

	if len(os.Args) > 1 {
		StringArray := os.Args[1:]
		kalimat := arrayToString(StringArray)

		for _, value := range people {
			if value.userid == kalimat || value.name == kalimat {
				fmt.Println("Userid : ", value.userid)
				fmt.Println("Nama :", value.name)
				fmt.Println("Alamat :", value.address)
				fmt.Println("Alasan :", value.reason)
				checking = true
			}
		}
		if checking == false {
			fmt.Println("Data dengan nama(case sensitive)/absen tsb tidak tersedia")

		}
	} else {
		fmt.Println("Tolong masukan nama atau nomor absen")
		fmt.Println("Contoh : 'go run minitask3 Bimo' atau 'go run minitask3 3' ")

	}

}
