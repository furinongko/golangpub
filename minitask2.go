package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
    //kalimat := "saya mau makan"
   // var caricari map[string]string
    caricari := make(map[string]string)

    fmt.Println("Masukan kalimat : ")

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan() // use `for scanner.Scan()` to keep reading
    kalimat := scanner.Text()

    
    chars := []rune(kalimat)
    for i := 0; i < len(chars); i++ {
        char := string(chars[i])
        println(char)
        jumlah := strings.Count(string(kalimat),char)
        hcount := strconv.Itoa(jumlah)
        //fmt.Println(hcount) // 4
        caricari[char]=hcount
//        caricari[hcount]=hcount
    }
//    fmt.Println(caricari)
    fmt.Println(strings.Repeat("-",20))
    for key, val:= range caricari {
        fmt.Print(key, ":", val)
        fmt.Print(" ")
    }
}