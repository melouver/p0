// Implementation of a KeyValueServer. Students should write their code in this file.

package p0

import (
	"fmt"
	"net"
	"strconv"
	"errors"
	"os"
	"strings"
)
type myConn struct {
	conn net.Conn
	req string
}

type keyValueServer struct {
	connections []myConn
	// TODO: implement this!
}

// New creates and returns (but does not start) a new KeyValueServer.
func New() KeyValueServer {
	// TODO: implement this!
	init_db()
	srv := keyValueServer{}
	return &srv
}


func (kvs *keyValueServer) Start(port int) error {
	ln, err := net.Listen("tcp", ":" + strconv.Itoa(port))
	if err != nil {
		fmt.Println("Error on listen: ", err);
		return errors.New("Error on listening")
	}
	for {
		fmt.Println("waiting for a connection via Accept")
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error on accpet: ", err)
			os.Exit(-1)
		}
		mconn := &myConn{
			conn:conn,
			req:string(""),
		}
		append(kvs.connections, conn)
		go hanldeConn(mconn, kvs)
	}
	fmt.Println("Server exiting")
	return nil
}

func (kvs *keyValueServer) Close() {
	// TODO: implement this!
}

func (kvs *keyValueServer) Count() int {
	// TODO: implement this!
	return -1
}

// TODO: add additional methods/functions below!
func hanldeConn(mconn *myConn, kvs *keyValueServer) {
	fmt.Println("reading once from connection")

	var buf [1024]byte
	_, err := mconn.conn.Read(buf[:])
	if err != nil {
		fmt.Println("error on read")
		os.Exit(-1)
	}
	v := strings.Split(string(buf[:]), ",")
	if len(v) == 2 {
		// get
		t := strings.Split(v[1], "\n")
		v[1] = t[0]
		val := get(v[1])
		for _, elem := range kvs.connections {
			elem.conn.Write(val)
		}
	} else if len(v) == 3 {
		// put
		t := strings.Split(v[2], "\n")
		v[2] = t[0]
		put(v[1], []byte(v[2]))
	} else {
		fmt.Println("arg cnt error")
		os.Exit(-1)
	}


}







