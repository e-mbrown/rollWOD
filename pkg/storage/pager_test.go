package storage_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/e-mbrown/rollWOD/pkg/storage"
)

/*
	Expand testInit
		- add case just header
		- header w\ initialized btree read slots and root page id
*/

func TestNewPager(t *testing.T) {
	tests := []struct {
		fileExist bool
		mb string
	}{
		{false, "pro"},
		{true, "boop"},
	}

	for _, tt := range tests{
		p := storage.NewPager(tt.mb)

		_, mb, _, _ := p.PrintInfo()

		if  tt.mb != mb {
			t.Fatalf("Magic byte mismatch got: %s, expected: %s", mb, tt.mb)
		}
		
		
		if !tt.fileExist {
			err := os.Remove(tt.mb)
			if err != nil {
				t.Fatalf("Failed to remove file. Err: %s", err.Error())
			}
		} else {
			// check if page info matches header info
		}
	}
}

func TestInit(t *testing.T) {
	pager := storage.NewPager("boop")
	fmt.Println(pager.PrintInfo())

	err := pager.InitHeader()
	if err != nil {
		t.Fatalf("Got an error %s", err.Error())
	}

	// Update to read header back
}


