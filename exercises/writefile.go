package exercises

import (
	"bufio"
	"io/ioutil"
	"os"
)

func overWrite(file, data string) error {
	err := ioutil.WriteFile(file, []byte(data), 0666)
	return err
}

func writeln(file, data string) error {
	//TODO: write multiple lines to existing file starting from a new line
	f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	_, err = buf.WriteString(data)
	_, err = buf.WriteString("\n")
	buf.Flush()
	//output error
	return err
}
