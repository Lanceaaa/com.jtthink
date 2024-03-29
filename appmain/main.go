package main

import (
    "fmt"
    _ "com.jtthink/services"
	// . "com.jtthink/core"
	_ "github.com/go-sql-driver/mysql"
	// "database/sql"
	// "os"
	_ "com.jtthink/models"
	// "time"
	// "io/ioutil"
	// "net/http"
	// "bytes"
	"io"
	"os"
	"bufio"
	"sync"
)

type all interface {

}

func test() {
	panic("exp")
}

func sum(min, max int, c chan int)  {
	result := 0
	for i := min; i <= max; i++ {
		result = result + i
	}
	c <- result
}

func main()  {
	// fmt.Println(GetService().Get(123))

	var wg sync.WaitGroup
	var mutex sync.Mutex
	file, _ := os.OpenFile("./test/test", os.O_RDONLY, 666)
	defer file.Close()
	fw := bufio.NewReader(file)
	for i := 1; i <= 2; i++ {
		go func (index int)  {
			defer wg.Done()
			mutex.Lock()
			for {
				str, err := fw.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						break
					}
					fmt.Println(err.Error())
				}
			    fmt.Printf("协程%d:%s", index, str)
			}
			mutex.Unlock()
		} (i)
	}
	
	wg.Add(2)
	wg.Wait()
	fmt.Println("读取成功")

	// num := 100000000
	// start := time.Now()
	// c := make(chan int, 2)
	// go sum(1, num/2, c)
	// go sum(num/2+1, num, c)
	// ret1, ret2 := <- c, <- c

	// users := strings.Split("shenyi,zhangsan,lisi,wangwu", ",")
	// ages := strings.Split("17,67,5,8,6", ",")
	
	// c1, c2 := make(chan bool), make(chan bool)
	// ret := make([]string, 0)

	// go func ()  {
	// 	for _, v := range users {
	// 		<- c1 
	// 		ret = append(ret, v)
	// 		c2 <- true
	// 	}
	// } ()
	// go func ()  {
	// 	for _, v := range ages {
	// 		<- c2
	// 		ret = append(ret, v)
	// 		c1 <- true
	// 	}
	// } ()
	// c1 <- true
	// fmt.Println(ret)

	// url := "https://news.cnblogs.com/n/page/%d/"

	// c := make(chan map[int][]byte)
	// for i := 1; i <= 3; i++ {
	// 	go func (index int)  {
	// 	    url := fmt.Sprintf(url, index)
	//         res, _ := http.Get(url)
	// 	    cnt, _ := ioutil.ReadAll(res.Body)
	// 	    c <- map[int][]byte{index:cnt}
	// 	} (i)
	// }
	// for getcnt := range c {
	// 	for k, v := range getcnt {
	// 	    ioutil.WriteFile(fmt.Sprintf("./files/%d", k), v, 666)
	// 	}
	// }

	// end := time.Now()
	// fmt.Println(end.Sub(start), ret1 + ret2)

	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4")
	// if err != nil {
	// 	fmt.Println("Failed to connect to mysql, err:"+ err.Error())
	// 	os.Exit(1)
	// }

	// rows, err := db.Query("select user_name, signup_at from tbl_user")
	// if err != nil {
	// 	fmt.Println("Failed to select to mysql, err:"+ err.Error())
	// 	return
	// }

	// columns, _ := rows.Columns()

	// allRows := make([]interface{}, 0)
	// for rows.Next() {
	//     fileMap := make(map[string]interface{})
	// 	row := make([]interface{}, len(columns))
	// 	scanRow := make([]interface{}, len(columns))
	// 	for i, _ := range row {
	// 		scanRow[i] = &row[i]
	// 	}
	// 	// rows.Scan(&row[0], &row[1])
	// 	rows.Scan(scanRow...)

	// 	for k, val := range row {
	// 		v, ok := val.([]byte)
	// 		if ok {
	// 			// row[k] = string(v)
	// 			fileMap[columns[k]] = string(v)
	// 		}
	// 	}

	// 	allRows = append(allRows, fileMap)
	// }

	// for _, v := range allRows {
	// 	fmt.Println(v)
	// }
}