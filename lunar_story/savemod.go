// TODO: checksum
package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"unsafe"
)

var (
	maxout = flag.Bool("maxout", false, "max out")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("lunar-story-savemod: ")

	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 1 {
		usage()
	}

	f, err := decode(flag.Arg(0))
	check(err)

	f.Dump()

	if *maxout {
		f.MaxOut()
		err := os.WriteFile(flag.Arg(0), f.Data, 0644)
		check(err)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: [options] cardxx_LUNAR000")
	flag.PrintDefaults()
	os.Exit(2)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type File struct {
	Data  []byte
	Chars []Char
}

type Char struct {
	HP     uint16
	MaxHP  uint16
	MP     uint16
	MaxMP  uint16
	Atk    uint16
	Def    uint16
	Wis    uint16
	MgcEn  uint16
	Agi    uint16
	NumAtk uint8
	Range  uint8
	Luck   uint8
	Unk    [42]byte
}

func decode(name string) (*File, error) {
	fmt.Println(unsafe.Sizeof(Char{}))
	const sig = "EX"

	b, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	if len(b) != 0x2000 {
		return nil, fmt.Errorf("%v: save file too small", name)
	}

	if bytes.Compare(b[:len(sig)], []byte(sig)) != 0 {
		return nil, fmt.Errorf("%v: invalid header signature", name)
	}

	f := &File{
		Data: b,
	}

	var ch Char
	for _, o := range STATS {
		r := bytes.NewReader(b[o:])
		binary.Read(r, binary.LittleEndian, &ch)
		f.Chars = append(f.Chars, ch)
	}

	return f, nil
}

func (f *File) Dump() {
	for i, ch := range f.Chars {
		fmt.Printf("Char #%v\n", i+1)
		fmt.Printf("HP: %v/%v\n", ch.HP, ch.MaxHP)
		fmt.Printf("MP: %v/%v\n", ch.MP, ch.MaxMP)
		fmt.Printf("Attack: %v\n", ch.Atk)
		fmt.Printf("Defense: %v\n", ch.Def)
		fmt.Printf("Agility: %v\n", ch.Agi)
		fmt.Printf("Wisdom: %v\n", ch.Wis)
		fmt.Printf("Magic Enhance: %v\n", ch.MgcEn)
		fmt.Printf("Num Attack: %v\n", ch.NumAtk)
		fmt.Printf("Range: %v\n", ch.Range)
		fmt.Printf("Luck: %v\n", ch.Luck)
		fmt.Printf("\n")
	}
	fmt.Printf("Gold: %v\n", binary.LittleEndian.Uint32(f.Data[MONEY:]))
}

func (f *File) MaxOut() {
	fmt.Println("MAX OUT!")
	for i, ch := range f.Chars {
		copy(f.Data[STATS[i]:], makebytes(ch))
	}
	binary.LittleEndian.PutUint32(f.Data[MONEY:], 10000)
}

func makebytes(v any) []byte {
	w := new(bytes.Buffer)
	binary.Write(w, binary.LittleEndian, v)
	return w.Bytes()
}

const (
	MONEY = 0x714
)

var STATS = []int{
	0x318, // alex
	0x498, // luna
	0x4d8, // ramus
}
