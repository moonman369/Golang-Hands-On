package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Demo of Stringer interface
// Any type with methods called `String()` are of type Stringer
// Now when we create specific values of these following types and print them using pln(), they will be printed in the below formats as in their respective String() methods
type book struct {
	title  string
	author string
}

func (b book) String() string {
	return fmt.Sprint("The title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer) {
	log.Println("LOG FROM LINE 27", s.String())
}

// Writer interface
func (b book) writeOut(w io.Writer) error {
	// _, err := w.Write(append([]byte(b.title), []byte(b.author)...))
	_, err := w.Write([]byte(fmt.Sprintf("{\n\tTitle: %v\n\tAuthor: %v\n}", b.title, b.author)))
	return err
}

func main() {
	b := book{
		title:  "A Brief History of Time",
		author: "Stephen Hawking",
	}
	var c count = 42

	logInfo(b)
	logInfo(c)

	f, err := os.Create("0137-functions-contd/output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// converting string to a slice of bytes
	s := []byte("Hello Gophers!!\nThis is next line tho.\n")
	fmt.Println(s)
	_, err = f.Write(s)
	if err != nil {
		log.Fatal(err)
	}

	err = b.writeOut(f)
	if err != nil {
		log.Fatal(err)
	}

	// Writing to a bytes Buffer
	b2 := book{
		title:  "Song of Ice and Fire",
		author: "George R. R. Martin",
	}
	var buff bytes.Buffer
	err = b2.writeOut(&buff)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(buff.Bytes())
	fmt.Println(buff.String())
}
