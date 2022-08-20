package main

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Output is indivisible.
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

type TXOutputs struct {
	Outputs []TXOutput
}

func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
	out.PubKeyHash = pubKeyHash
}

func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}

func (outs TXOutputs) Serialize() []byte {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(outs)
	if err != nil {
		log.Panic(err)
	}

	return buf.Bytes()
}

func DeserializeOutputs(data []byte) TXOutputs {
	var outputs TXOutputs

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}

	return outputs
}
