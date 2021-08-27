package main
import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:4242")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err :=l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go copyToStderr(conn)

	}
}

func copyToStderr(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(5*time.Second))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Printf("Finished with err = %v", err)
			return
		}
		os.Stderr.Write(buf[:n])
	}
	n, err := io.Copy(os.Stderr, conn)
	log.Printf("copied %d bytes; finished with wrr= %v", n, err)
}