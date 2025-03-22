package samples

import (
	"fmt"
	"os"
)


func ServeSample(filename string) (*os.File, error) {
	
	fn:= fmt.Sprintf("./samples/%s.wql", filename)
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}

	return file, nil
}
