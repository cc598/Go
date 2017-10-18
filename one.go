package main

import (
    "fmt"
    "strconv"
    "bufio"
    "io"
    "os"
    "strings"
)

var (
    selpg, start, end, choose, choose2, temp, file string
    s_num, e_num, ch, ch_num int
)




func main() {
    ch_num = 72

    fmt.Printf("$ ")
    fmt.Scanln(&selpg, &start, &end, &choose, &choose2)

    temp = start[2:len(start)]
    temp2, err2 := strconv.Atoi(temp)
    if err2 != nil {
        panic(err2)
    }
    s_num = temp2

    temp = end[2:len(end)]
    temp3, err3 := strconv.Atoi(temp)
    if err3 != nil {
        panic(err3)
    }
    e_num = temp3

    if choose2 != "" {
        file = choose2
        if choose[1] == 'l' {
            ch = 1

	    temp = choose[2:len(end)]
	    temp4, err4 := strconv.Atoi(temp)
	    if err4 != nil {
		panic(err4)
	    }
	    ch_num = temp4
        } else {
            ch = 2
        }
    } else {
        file = choose
    }


    
    print();


    /*fmt.Println(s_num)
    fmt.Println(e_num)
    fmt.Println(ch)
    fmt.Println(ch_num)
    fmt.Println(file)*/
}




func print() error {
    if ch != 2 {
        var num, s_line, e_line int
        num = 0
        s_line = (s_num-1) * ch_num
        e_line = e_num * ch_num

        f, err := os.Open(file)
        if err != nil {
            panic(err)
        }
        buf := bufio.NewReader(f)
        for {
    	    line, err := buf.ReadString('\n')
	    line = strings.TrimSpace(line) 
            num++
            if s_line <= num && num <= e_line {
                fmt.Println(line)
            }
            if err != nil {
                if err == io.EOF {
                    return nil
                } else {
                    panic(err)
                }
            }
        }
    } else {
        var num int = 0
        f, err := os.Open(file)
        if err != nil {
            panic(err)
        }
        buf := bufio.NewReader(f)
        for {
    	    line, err := buf.ReadString('\f')
	    line = strings.TrimSpace(line) 
            num++
            if s_num <= num && num <= e_num {
                fmt.Println(line)
            }
            if err != nil {
                if err == io.EOF {
                    return nil
                } else {
                    panic(err)
                }
            }
        }
    }
    return nil
}




