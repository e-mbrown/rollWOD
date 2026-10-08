package samples

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

func ServeSample(filename string) (*os.File, error) {
	_, fn, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("no caller info; %s", fn)
	}
	fp := fmt.Sprintf("%s/%s.wql", path.Dir(fn), filename)
	file, err := os.Open(fp)
	if err != nil {
		return nil, err
	}

	return file, nil
}
