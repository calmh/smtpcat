package main

import (
	"flag"
	"io/ioutil"
	"net"
	"net/smtp"
	"os"
)

func main() {
	from := flag.String("from", "user@example.com", "From address")
	to := flag.String("to", "user@example.com", "To address")
	server := flag.String("smtp", "127.0.0.1:25", "SMTP server address")
	flag.Parse()

	if flag.NFlag() != 3 {
		flag.Usage()
		os.Exit(1)
	}

	msg, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", *server)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	client, err := smtp.NewClient(conn, *server)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	err = client.Mail(*from)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	err = client.Rcpt(*to)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	wr, err := client.Data()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	_, err = wr.Write(msg)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	err = wr.Close()
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
