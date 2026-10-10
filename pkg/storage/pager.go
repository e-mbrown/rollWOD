package storage

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/e-mbrown/rollWOD/pkg/btree"
)

type Pager struct {
	mb string 	// magicbyte
	pgS uint32 	// page size
	pgCt uint32
	RootID *btree.Node
}

type Header struct {
	slots [16]byte
	mb [4]byte
	pgS [4]byte
	pgCt [4]byte
}

/* 
TODO: Add....node id? or somethin lol and finish
Do I need to close cause docs dont have close...but I feel a memory leak in my future
also this looks like booty clean up
*/
func MakeExtractHeader(p *Pager, file *os.File) (Header, error) {
	var slots [16]byte
	var mb, pgS, pgCt[4]byte
	var header Header
	var newHeader bool

	if file == nil {
		newHeader = true
	} else {
		fileInfo, err := file.Stat()
		if err != nil {
			return header, err
		}

		if fileInfo.Size() == 0 {
			newHeader = true
		}
	}

	
	if newHeader {
		copy(mb[:], p.mb)
		binary.LittleEndian.PutUint32(pgS[:], 4096)
		binary.LittleEndian.PutUint32(pgCt[:], 1)
		header =  Header{
			slots : slots,
			mb: mb,
			pgS: pgS,
			pgCt: pgCt,
		}
	} else {
		// check mb or error
		//extract
		//deterministic 16, 4,4,4 for hardcode
	}


	return header, nil
}

func NewPager(mb string) *Pager {
	p :=  &Pager{
		mb: mb,
	}

	// change eventually to direct to take actual dir
	p.InitHeader(mb)


	return p
}

func (p *Pager) InitHeader(dir string) error {
	var header Header
	var root *os.File
	var err error

	root, err = os.Open(p.mb)
	if err == nil {
		header, err = MakeExtractHeader(p, root)
		if err != nil {
			return err
		}
			
	}else {
		header,err = MakeExtractHeader(p, nil)
		if err != nil {
			return err
		}

		root, err = os.Create(p.mb)
		if err != nil {
			return err
		}

		root.Write(header.slots[:])
		root.Write(header.mb[:])
		root.Write(header.pgS[:])
		root.Write(header.pgCt[:])
	}

	p.pgS = binary.LittleEndian.Uint32(header.pgS[:])
	p.pgCt = binary.LittleEndian.Uint32(header.pgS[:])
	

	return nil
}

func (p *Pager) Read() error {
	root, err := os.Open(p.mb)
	if err != nil {
		return err
	}

	
	var output [4]byte
	offset := int64(16)
	for i:= 0; i < 3; i++ {
		n, err := root.ReadAt(output[:], offset)
		if err != nil {
			return err
		}
		fmt.Println(n, output)

		offset += 4
	}
	return nil
}

func (p *Pager) PrintInfo() (string, string, int, int) {
	return fmt.Sprintf(
		"Pager mb %s, Pager page Size %d and count %d",
		p.mb,p.pgS, p.pgCt,
	), p.mb, int(p.pgS), int(p.pgCt)
}
