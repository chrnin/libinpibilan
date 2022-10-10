package libinpibilan

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

//nolint:errcheck

func TestNewBilanWithBytes(t *testing.T) {
	ass := assert.New(t)
	// WITH
	source, err := ioutil.ReadFile("test/test.xml")
	ass.Nil(err)

	// WHEN
	bilan, err := NewBilanWithBytes(source)

	// THEN
	ass.Nil(err)
	ass.Len(bilan.Lignes, 331)
	ass.Len(bilan.RapportConversion, 0)
}

func TestNewBilanWithReader(t *testing.T) {
	ass := assert.New(t)
	// WITH
	source, err := os.Open("test/test.xml")
	defer source.Close()
	ass.Nil(err)

	// WHEN
	bilan, err := NewBilanWithReader(source)

	// THEN
	ass.Nil(err)
	ass.Len(bilan.Lignes, 331)
	ass.Len(bilan.RapportConversion, 0)
}

func TestNewBilanRoutineWithBytesChan(t *testing.T) {
	ass := assert.New(t)
	// WITH
	testLength := 100
	source, err := ioutil.ReadFile("test/test.xml")
	ass.Nil(err)
	bytesChannel := make(chan []byte)
	bilanChannel := NewBilanRoutineWithBytesChan(bytesChannel, 1)

	// WHEN
	var bilans []Bilan
	for i := 0; i < testLength; i++ {
		bytesChannel <- source
		bilans = append(bilans, <-bilanChannel)
	}
	close(bytesChannel)

	// THEN
	ass.Len(bilans, testLength)
}
