# blockchain_go
A simplified blockchain implementation in Golang (Jeiwan ver.)

### Try it Out

```bash
❯ alias blockchain_go="go run ."

❯ blockchain_go
Usage:
  addblock -data BLOCK_DATA - add a block to the blockchain
  printchain - print all the blocks of the blockchain

exit status 1

❯ blockchain_go printchain
Prev. hash: 
Data: Genesis Block
Hash: 00000f21a0214565f0d194483e1e38410e170a93f2d0a2a9f43317323724fed9
PoW: true

❯ blockchain_go addblock -data "Send 1 BTC to Ivan"
Mining the block containing "Send 1 BTC to Ivan"
0000049f6a9d32b039fb39a011a71643ddb38997ff65aba79a1338ba6d633176

Success!

❯ blockchain_go printchain
Prev. hash: 00000f21a0214565f0d194483e1e38410e170a93f2d0a2a9f43317323724fed9
Data: Send 1 BTC to Ivan
Hash: 0000049f6a9d32b039fb39a011a71643ddb38997ff65aba79a1338ba6d633176
PoW: true

Prev. hash: 
Data: Genesis Block
Hash: 00000f21a0214565f0d194483e1e38410e170a93f2d0a2a9f43317323724fed9
PoW: true
```