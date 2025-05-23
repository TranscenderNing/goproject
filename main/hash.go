package main

import (
	"demo/balance"
	"fmt"
	"hash/crc32"
	"math/rand"
)

func init() {
	balance.RegisterBalancer("hash", &HashBalance{})
}

type HashBalance struct {

}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {
	var defalutKey string = fmt.Sprintf("%d", rand.Intn(1000))
	lens := len(insts)
	if lens <= 0 {
		err = fmt.Errorf("no baackend instance")
		return 
	}
	crcTable := crc32.MakeTable(crc32.IEEE)
	hashval := crc32.Checksum([]byte(defalutKey), crcTable)
	index := int(hashval) % (lens)
	inst = insts[index]
	return 
}