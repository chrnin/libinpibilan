package libinpibilan

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

var SCHEMA = BuildCodesInpi()

// NewBilanWithReader convertir une source XML en libinpibilan.Bilan
func NewBilanWithReader(reader io.Reader) (Bilan, error) {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return Bilan{}, ErreurLectureSourceImpossible{wrap(err)}
	}
	var xmlBilans XMLBilans
	err = xml.Unmarshal(data, &xmlBilans)
	if err != nil {
		return Bilan{}, ErreurConversionImpossible{wrap(err)}
	}
	bilan := xmlBilans.BuildBilan()
	return bilan, nil
}

// NewBilanWithBytes convertir une source XML en libinpibilan.Bilan
func NewBilanWithBytes(data []byte) (Bilan, error) {
	var xmlBilans XMLBilans
	err := xml.Unmarshal(data, &xmlBilans)
	if err != nil {
		return Bilan{}, ErreurConversionImpossible{wrap(err)}
	}
	bilan := xmlBilans.BuildBilan()
	return bilan, nil
}

// NewBilanRoutineWithBytesChan fournit une goroutine traitant les sources XML envoyées.
// La routine s'interrompt lorsque bytesChannel est fermé.
func NewBilanRoutineWithBytesChan(bytesChannel chan []byte, chanSize int) chan Bilan {
	var bilanChan = make(chan Bilan, chanSize)
	go bilanRoutine(bytesChannel, bilanChan)
	return bilanChan
}

func bilanRoutine(bytesChannel chan []byte, bilanChannel chan Bilan) {
	for bilanBytes := range bytesChannel {
		bilan, err := NewBilanWithBytes(bilanBytes)
		if err != nil {
			bilan.RapportConversion = []error{err}
		}
		bilanChannel <- bilan
	}
	close(bilanChannel)
}
