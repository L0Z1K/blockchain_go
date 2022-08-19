package main

type TXInput struct {
	Txid      []byte
	Vout      int // index of output
	ScriptSig string
}
