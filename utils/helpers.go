package utils

import (
	"bufio"
	"encoding/base64"
	"os"
	"strconv"
)

// Map an alias for map[string]interface{}
type Map map[string]interface{}

// Pager ...
func Pager(pageIndex, pageSize int64) (offset, limit int64) {
	if pageIndex <= 0 {
		pageIndex = 1
	}

	limit = pageSize
	if pageSize <= 0 {
		limit = 15
	}

	offset = (pageIndex - 1) * limit
	return
}

// DelSliceItem delete an item from a []string
func DelSliceItem(s *[]string, i int) {
	copy((*s)[i:], (*s)[i+1:]) // Shift a[i+1:] left one index.
	(*s)[len(*s)-1] = ""       // Erase last element (write zero value).
	*s = (*s)[:len(*s)-1]
}

// Atoi Converts a string to an int, returns zero if string is empty or contains
// an invalid int
func Atoi(val string) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}

	return i
}

// Atoi64 Converts a string to an int64, returns zero if string is empty or contains
// an invalid int64
func Atoi64(val string) int64 {
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return int64(0)
	}

	return i
}

// Atoui64 Converts a string to uint64, returns zero if string is empty or contains
// an invalid uint64
func Atoui64(val string) uint64 {
	i, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return uint64(0)
	}

	return i
}

// RecToPageCount converts as record count to page count based on the provided limit
func RecToPageCount(recCount, limit uint64) (pageCount uint64) {
	pageCount = recCount / limit
	if pageCount%limit != 0 {
		pageCount++
	}

	return
}

func OpenandReadFile(path string) ([]byte, error) {

	fileEntity, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer fileEntity.Close()

	// create a new buffer base on file size
	fInfo, _ := fileEntity.Stat()

	var size int64 = fInfo.Size()
	buf := make([]byte, size)

	// read file content into buffer
	fReader := bufio.NewReader(fileEntity)
	fReader.Read(buf)

	return buf, nil
}

// EncodeImageBase64 ...
func EncodeImageBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
