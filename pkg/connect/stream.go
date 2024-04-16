package connect

import (
	// "bytes"
	// "fmt"
	"bufio"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

const clearLine = "\033[2K"

func ReadWav(fp string) ([]byte, error) {

	file, err := os.Open(fp)
	if filepath.Ext(fp) != ".wav" {
		slog.Info(filepath.ErrBadPattern.Error(), "extension", filepath.Ext(fp))
		log.Fatal(filepath.ErrBadPattern)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	println("contents:")
	// fmt.Printf("%v", contents)
	println(contents)

	// For recording position of where we are in header - need to read 'data'

	var foundBytes []byte
	file.Seek(0, 0)
	reader := bufio.NewReader(file)
	// buffer := make([]byte, 1)

	var i, j, nRead int
	reset := true
	found := false
	for {
		if j >= 3 {
			break
		}

		var dots string
		i++
		switch i {
		case 1:
			dots = "..."
		case 2:
			dots = "......"
		case 3:
			dots = "........."
		default:
			i = 0
		}
		fmt.Printf("reading%v", dots)
		print(clearLine)
		print("\r")
		b, err := reader.ReadByte()
		nRead++
		if err != nil {
			break
		}

		switch len(foundBytes) {
		case 0:
			if b == 'd' {
				reset = false
			}
		case 1:
			if b == 'a' {
				reset = false
			}
		case 2:
			if b == 't' {
				reset = false
			}
		case 3:
			if b == 'a' {
				reset = false
				found = true
				n, err := reader.Discard(4)
				println("discarded:", n)
				if err != nil {
					slog.Info("Failed to discard remainder of header:", "nRead", nRead)
				}
				nRead += n
				break
			}
		default:
			continue
		}

		if !reset {
			foundBytes = append(foundBytes, b)
			fmt.Printf("%c\n", foundBytes)
		}
		reset = true
		if found {
			break
		}
	}
	println("done!")
	slog.Info("Size of header:", "nRead", nRead)
	stat, err := file.Stat()
	if err != nil {
		log.Fatal("Could not get file stats")
	}
	totalFileSize := stat.Size()
	remaining := totalFileSize - int64(nRead)
	slog.Info("Size of WAV file", "totalFileSize", totalFileSize)
	slog.Info("Size of remaining data portion of WAV:", "remaining", remaining)

	// for err := file.Read(buffer) {
	// 	if err != nil {
	// 		break
	// 	}
	//
	// }
	// tempfile := bytes.NewReader()
	// var audioData []byte
	// // Load entire file first...
	// r := bytes.Reader{}
	// r.Read()
	return nil, nil
	// return audioData, nil
}

func skipHeader(wavBytes []byte) ([]byte, error) {
	return nil, nil
}
