package main
import "file-transfer/internal/compression"

import (
	"file-transfer/internal/fileio"
	"file-transfer/internal/protocol"
	"file-transfer/internal/transport"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: sender <file> <addr:port>")
		return
	}

	// ✅ Define filePath HERE
	filePath := os.Args[1]
	addr := os.Args[2]

	conn, err := transport.Connect(addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// ✅ Extract only filename (IMPORTANT FIX)
	filename := filepath.Base(filePath)
	useGzip := compression.ShouldCompress(filename)
	algo := "none"
	if useGzip {
	algo = "gzip"
	}

meta := filename + "|" + algo

	// Send metadata
	metaPacket := protocol.Packet{
		Type: protocol.MSG_METADATA,
		Data: []byte(meta),
	}
	data, _ := protocol.Encode(metaPacket)
	conn.Write(data)

	// Read file in chunks
	chunks, err := fileio.ReadChunks(filePath, 64*1024)
	if err != nil {
		panic(err)
	}

	for chunk := range chunks {
		p := protocol.Packet{
			Type: protocol.MSG_CHUNK,
			Data: chunk,
		}
		data, _ := protocol.Encode(p)
		conn.Write(data)
	}

	// Send END
	endPacket := protocol.Packet{Type: protocol.MSG_END}
	data, _ = protocol.Encode(endPacket)
	conn.Write(data)

	fmt.Println("File sent successfully")
}