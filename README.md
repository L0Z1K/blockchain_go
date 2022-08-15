# blockchain_go
A simplified blockchain implementation in Golang (Jeiwan ver.)

### Try it Out

```bash
❯ go run .
Mining the block containing "Genesis Block"
00000292c9433165ad07a6243b491f405aecb4588f82faf109b5759aa1f1f793

Mining the block containing "Send 1 BTC to Ivan"
0000071eb55bd044f87c6bdd932c0694d2179611dce0a0fb1f7fb3d4457f102c

Mining the block containing "Send 2 more BTC to Ivan"
000007dec8e227bbf9e1459e3ab325d306106e31da8f285b7c8082fb6797985c

Prev. hash: 
Data: Genesis Block
Hash: 00000292c9433165ad07a6243b491f405aecb4588f82faf109b5759aa1f1f793

Prev. hash: 00000292c9433165ad07a6243b491f405aecb4588f82faf109b5759aa1f1f793
Data: Send 1 BTC to Ivan
Hash: 0000071eb55bd044f87c6bdd932c0694d2179611dce0a0fb1f7fb3d4457f102c

Prev. hash: 0000071eb55bd044f87c6bdd932c0694d2179611dce0a0fb1f7fb3d4457f102c
Data: Send 2 more BTC to Ivan
Hash: 000007dec8e227bbf9e1459e3ab325d306106e31da8f285b7c8082fb6797985c
```