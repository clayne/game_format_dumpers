// TODO: checksum
package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	maxout = flag.Bool("maxout", false, "max out")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("lunar-story-2-savemod: ")

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
	fmt.Fprintln(os.Stderr, "usage: [options] cardxx_SAVE_DATA")
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
	Lvl     uint8
	Unk1    [2]uint8
	Exp     uint32
	NextExp uint32
	HP      uint16
	MaxHP   uint16
	MP      uint16
	MaxMP   uint16
	Atk     uint16
	Atks    uint16
	Def     uint16
	Agi     uint16
	Spd     uint16
	Wis     uint16
	MgcDef  uint16
	Range   uint16
	Luck    uint16
}

func decode(name string) (*File, error) {
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
		fmt.Printf("Attacks: %v\n", ch.Atks)
		fmt.Printf("Defense: %v\n", ch.Def)
		fmt.Printf("Agility: %v\n", ch.Agi)
		fmt.Printf("Speed: %v\n", ch.Spd)
		fmt.Printf("Wisdom: %v\n", ch.Wis)
		fmt.Printf("Magic Defense: %v\n", ch.MgcDef)
		fmt.Printf("Range: %v\n", ch.Range)
		fmt.Printf("Luck: %v\n", ch.Luck)
		fmt.Printf("\n")
	}
}

func (f *File) MaxOut() {
	fmt.Println("MAX OUT!")
	for i, ch := range f.Chars {
		ch.HP = 9999
		ch.MaxHP = 9999
		ch.MP = 9999
		ch.MaxMP = 9999
		ch.Atk = 9999
		ch.Atks = 9999
		ch.Def = 9999
		ch.Agi = 9999
		ch.Spd = 9999
		ch.Wis = 9999
		ch.MgcDef = 9999
		ch.Range = 9999
		ch.Luck = 9999
		copy(f.Data[STATS[i]:], makebytes(ch))
	}
}

func makebytes(v any) []byte {
	w := new(bytes.Buffer)
	binary.Write(w, binary.LittleEndian, v)
	return w.Bytes()
}

const (
	GOLD = 0x0
)

var STATS = []int{
	0x311, // hiro
}
