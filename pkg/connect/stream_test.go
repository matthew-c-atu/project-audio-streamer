package connect_test

import (
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/matthew-c-atu/project-audio-streamer/pkg/connect"
	"github.com/stretchr/testify/assert"
)

const (
	testFilePathName = "testFilePath"
	testFilePath     = "music/test/test_song.wav"
	testFile         = "test_song.wav"
	headerSize       = 44
)

func TestSkipHeader(t *testing.T) {
	dir, _ := os.Getwd()
	println("WORKING DIR:")
	println(dir)
	joined := filepath.Join(dir, testFile)
	println("FILE LOCATION:")
	println(joined)

	f, err := os.Open(joined)
	if err != nil {
		slog.Info("Failed to open file", "joined path", joined)
	}
	stat, err := f.Stat()
	if err != nil {
		slog.Info("Failed to get stats on file", "joined path", joined)
	}
	expectedSize := int(stat.Size() - headerSize)
	remain, err := connect.SkipHeader(f)
	actualSize := len(remain)
	slog.Info("expected remaining size:", "expectedSize", expectedSize)
	slog.Info("actual remaining size:", "actualSize", actualSize)
	assert.Equal(t, expectedSize, actualSize)

}
