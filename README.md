# blockchain_go
A simplified blockchain implementation in Golang (Jeiwan ver.)

### Try it Out

```bash
❯ go run .
Mining the block containing "Genesis Block"
00000b24d04b24281f0d8616dc1e992a7ba7558bec998a171e8e511282e7b70c

Mining the block containing "Send 1 BTC to Ivan"
00000b4781887f105726db8b2ec8bc439829b0d4aa92fd0cc2543fcd90ee3c97

Mining the block containing "Send 2 more BTC to Ivan"
000005bbd5072ef46bfc89a50824b1e9390e10f55754935ba314b87169fdb3d2

Prev. hash: 
Data: Genesis Block
Hash: 00000b24d04b24281f0d8616dc1e992a7ba7558bec998a171e8e511282e7b70c
PoW: true

Prev. hash: 00000b24d04b24281f0d8616dc1e992a7ba7558bec998a171e8e511282e7b70c
Data: Send 1 BTC to Ivan
Hash: 00000b4781887f105726db8b2ec8bc439829b0d4aa92fd0cc2543fcd90ee3c97
PoW: true

Prev. hash: 00000b4781887f105726db8b2ec8bc439829b0d4aa92fd0cc2543fcd90ee3c97
Data: Send 2 more BTC to Ivan
Hash: 000005bbd5072ef46bfc89a50824b1e9390e10f55754935ba314b87169fdb3d2
PoW: true
```