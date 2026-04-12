package main
// import 
import "file-transfer/internal/compression"

import (
	"file-transfer/internal/fileio"
	"file-transfer/internal/protocol"
	"file-transfer/internal/transport"
	"fmt"
	"net"
	"os"
	"strings"
	"path/filepath"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	var file *os.File
	var useGzip bool  // ✅ FIXED

	buffer := make([]byte, 0)
	tmp := make([]byte, 64*1024)

	for {
		n, err := conn.Read(tmp)
		if err != nil {
			break
		}

		buffer = append(buffer, tmp[:n]...)

		for {
			packet, remaining, err := protocol.TryDecode(buffer)
			if err != nil {
				break
			}

			buffer = remaining

			switch packet.Type {

			case protocol.MSG_METADATA:

				meta := string(packet.Data)
				parts := strings.Split(meta, "|")

				filename := parts[0]

				if len(parts) > 1 && parts[1] == "gzip" {
					useGzip = true
				}

				filename = filepath.Base(filename)
				finalName := "received_" + filename

				file, err = os.Create(finalName)
				if err != nil {
					fmt.Println("File create error:", err)
					return
				}

				fmt.Println("Receiving:", finalName)

			case protocol.MSG_CHUNK:

				if useGzip {
					packet.Data, err = compression.DecompressChunk(packet.Data)
					if err != nil {
						fmt.Println("Decompression error:", err)
						return
					}
				}

				fileio.WriteChunk(file, packet.Data)

			case protocol.MSG_END:

				file.Close()
				fmt.Println("File received successfully")
				return
			}
		}
	}
}


func main() {
	ln, err := transport.StartServer("8080")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	fmt.Println("Listening on port 8080")

	for {
		conn, _ := ln.Accept()
		//var useGzip bool
		go handleConnection(conn)
	}
}