package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
)

type File struct {
	Data   []byte
	Header Header
}

type Header struct {
	Version uint32

	PlayTime int64
	GameTime float32
	Scenario int32
	Money    uint32
	FloorID  [16]byte
	AreaID   int32

	PlayerName  [9][32]byte
	PlayerLevel [9]int32
	PlayerParam [9]int32

	Clear               uint8
	MainSceneState      int32
	MiniMapState        int32
	CameraLockAreaIndex int32
	FloorPassiveIndex   int32
	ItemSort            int32

	MachineryState       int32
	MachineryEnergy      int32
	MachineryEnergyMax   int32
	MachineryBoostTime   float32
	MachineryBoostEnable uint8
	MachineryScanTime    int32
	MachineryScanEnable  uint8

	PlayerPosition [3]float32
	PlayerRotY     float32

	TotalExperience int32
	TotalMoney      int32

	BattleCount                int32
	HumanBattleActionCount     int32
	MachineryBattleActionCount int32
	MachineryKillCount         int32
	SetsunaKillCount           int32
	MaxHitCount                int32
	MaxSimultaneousAttackCount int32
	MaxDamage                  int32
	WeakAttackCount            int32
	CriticalCount              int32
	AvoidCount                 int32
	SpritniteCreateCount       int32
	SublimationCount           int32
	SetsunaSystemCount         int32
	CounterCount               int32
	FishingCount               int32
	ShiningPointCount          int32
	ArtifactCreateCount        int32
	MachineryMenuCount         int32
	FoodCount                  int32
	ScanKillCount              int32

	ShopBuyMoney  int32
	ShopSellMoney int32

	FieldMoveDistance      float32
	BattleMoveDistance     float32
	RecoveryHpMoveDistance float32
	RecoveryMpMoveDistance float32
	DebuffSlipMoveDistance float32
	GroundDamageDistance   float32
	GroundDamageTime       float32
}

var flags struct {
	key    string
	iv     string
	maxout bool
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("lost-sphear: ")

	flag.StringVar(&flags.key, "key", "G3giuyVjh8F3KZTM", "specify key")
	flag.StringVar(&flags.iv, "iv", "v6LVsqPyB4jm6bgw", "specify iv")
	flag.BoolVar(&flags.maxout, "maxout", false, "maxout")
	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 2 {
		usage()
	}

	f, err := decode(flag.Arg(0), []byte(flags.key), []byte(flags.iv))
	check(err)

	err = os.WriteFile(flag.Arg(1), f.Data, 0644)
	check(err)
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: [options] file")
	flag.PrintDefaults()
	os.Exit(2)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func decode(name string, key, iv []byte) (*File, error) {
	const (
		sig = 20170721
		off = 0x510
	)

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	if len(data) < off {
		return nil, fmt.Errorf("file too small")
	}

	f := &File{
		Data: data,
	}

	if len(key) > 0 {
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}

		dec := make([]byte, len(data)-off)
		if len(dec)%block.BlockSize() != 0 {
			return nil, fmt.Errorf("encrypted file size not aligned")
		}

		cbc := cipher.NewCBCDecrypter(block, iv)
		cbc.CryptBlocks(dec, data[off:])
		copy(f.Data[off:], dec)
	}

	r := bytes.NewReader(f.Data[off:])
	err = binary.Read(r, binary.LittleEndian, &f.Header)
	if err != nil {
		return nil, err
	}

	if f.Header.Version != sig {
		return nil, fmt.Errorf("invalid signature")
	}

	return f, nil
}
