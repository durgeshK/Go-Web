package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		ns := bufio.NewScanner(conn)
		for ns.Scan() {
			ln := ns.Text()
			if ln == "" {
				break
			}

			fmt.Println(ns.Text())
		}

		fmt.Fprintln(conn, "Hello User, Thanks for connecting")

		fmt.Println("Code got here.")

		conn.Close()
	}

}
