package main

import (
	"bufio"
	"fmt"
	// "io"
	"io/ioutil"
	"regexp"
	"strings"
	"os"
	"log"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("./journal.txt")	
	check(err)
	// fmt.Println(string(dat))
	strJournal := string(dat)
	validDateTime := regexp.MustCompile(`\[([0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2})\]`)
	dateTime := validDateTime.FindAllStringSubmatch(strJournal, -1)

	// for i := 0; i < 8; i++ {
	// 	// fmt.Println(i)
	// 	fmt.Println(dateTime[i][1])
	// 	fmt.Println("=======")
	// }

	jrnls := validDateTime.Split(strJournal, -1)

	for i, j := range jrnls {
		fmt.Println(dateTime[i][1])
		fmt.Println(j)
	}

	// for i, j := range jrnls {
	// 	fmt.Println(dateTime[j])
	// 	fmt.Println(j)
	// 	fmt.Println(i)
	// 	fmt.Println("=========================")
	// }


	// journals := strings.Split(strJournal, "\n\n")
	// for _, jrnl := range journals {
	// 	jrnlToBlog(jrnl)
	// 	fmt.Println("=========================")
	// }
}

func jrnlToBlog(jrnl string) {
	validDateTime := regexp.MustCompile(`\[([0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2})\]`)

	dateTime := validDateTime.FindStringSubmatch(jrnl)[1]
	date := strings.Split(dateTime, " ")[0]
	// time := strings.Split(dateTime, " ")[1]
	// fmt.Println(date, time)

	title := strings.Split(jrnl, "\n")[0]
	validTitle := regexp.MustCompile(`\[[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}\] (.*)`)
	t := validTitle.FindStringSubmatch(title)[1]
	// fmt.Println(t)


	descSlice := strings.Split(jrnl, "\n")[1:]
	description := strings.Join(descSlice, "\n")

	fmt.Println(description)


	// fileHandle, _ := os.OpenFile("./out.txt", os.O_APPEND, 0666)
	fileHandle, err := os.OpenFile("./"+date+"-"+t+".md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(fileHandle)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	fmt.Fprintln(writer, "---")
	fmt.Fprintln(writer, "layout: post")
	fmt.Fprintln(writer, "title: "+ t)
	fmt.Fprintln(writer, "date: "+ date)
	fmt.Fprintln(writer, "---")
	fmt.Fprintln(writer, "\n")

	fmt.Fprintln(writer, description)
	writer.Flush()
}