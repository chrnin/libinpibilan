# libinpibilan: Déchiffrer les bilans INPI en go pur

## Usage
### Lire un fichier
```go
package main

import (
	"fmt"
	"github.com/chrnin/libinpibilan"
)

func main() {
	source, err := os.Open("path/bilan.xml")
    if err != nil {
        panic(err)
    }
	if bilan, err := libinpibilan.NewBilanWithReader(source); err != nil {
        fmt.Printf("lecture du fichier impossible: %+v\n", err)
    } else {
        fmt.Println(bilan)
    }
}
```

### Lire un tableau de bytes
```go
package main

import (
	"fmt"
	"io/ioutil"
	"github.com/chrnin/libinpibilan"
)

func main() {
	source, err := ioutil.ReadFile("test/test.xml")
	if err != nil {
        panic(err)
    }
	if bilan, err := libinpibilan.NewBilanWithBytes(source); err != nil {
        fmt.Printf("lecture du fichier impossible: %+v\n", err)
    } else {
        fmt.Println(bilan)
    }
}
```

### Convertir les bilans dans une goroutine
```go
package main

import (
	"fmt"
	"io/ioutil"
	"github.com/chrnin/libinpibilan"
)

func main() {
	source, err := ioutil.ReadFile("test/test.xml")
    if err != nil {
        panic(err)
    }
	
	bytesChannel := make(chan []byte)
	bilanChannel := libinpibilan.NewBilanRoutineWithBytesChan(bytesChannel, 1)
	
	bytesChannel <- source
	bilan := <- bilanChannel

	fmt.Println(bilan)
}
```

