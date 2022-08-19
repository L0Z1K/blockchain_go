# blockchain_go
A simplified blockchain implementation in Golang (Jeiwan ver.)

### Try it Out

```bash
❯ alias blockchain_go="go run ."

❯ blockchain_go
Usage:
  createblockchain -address ADDRESS - Create a blockchain and send genesis block reward to ADDRESS
  getbalance -address ADDRESS - Get balance of ADDRESS
  send -from FROM -to TO -amount AMOUNT - Send AMOUNT of coins from FROM to TO
exit status 1

❯ blockchain_go send -from Ivan -to Baek -amount 1
00000d995a2583a2bed48360431081d9aa5d639d13092fc35495ab2ba25140e5

Success!

❯ blockchain_go getbalance -address Baek
Balance of 'Baek': 6
```