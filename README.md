# blockchain_go
A simplified blockchain implementation in Golang (Jeiwan ver.)

### Try it Out

```bash
❯ blockchain_go createwallet
Your new address: 1DEBUe5gFAKUA8AxbbSsf4ETXbCga8uqiT

❯ blockchain_go createwallet
Your new address: 1M7dipSJMmEJjUwjDvaG6p49SR4oHUrC1m

❯ blockchain_go createblockchain -address 1DEBUe5gFAKUA8AxbbSsf4ETXbCga8uqiT
00000210dd0b996e9e139aa62090046ea231a594498f680561bce4663a223508

Done!

❯ blockchain_go getbalance -address 1DEBUe5gFAKUA8AxbbSsf4ETXbCga8uqiT
Balance of '1DEBUe5gFAKUA8AxbbSsf4ETXbCga8uqiT': 10

❯ blockchain_go send -from 1DEBUe5gFAKUA8AxbbSsf4ETXbCga8uqiT -to 1M7dipSJMmEJjUwjDvaG6p49SR4oHUrC1m -amount 3
000008812f538631d683460f8e3ed9f7d7fbfa3e6e02f154cd8e1a31001bf52e

Success!

❯ blockchain_go getbalance -address 1DEBUe5gFAKUA8AxbbSsf4ETXbCga8uqiT
Balance of '1DEBUe5gFAKUA8AxbbSsf4ETXbCga8uqiT': 7

❯ blockchain_go getbalance -address 1M7dipSJMmEJjUwjDvaG6p49SR4oHUrC1m
Balance of '1M7dipSJMmEJjUwjDvaG6p49SR4oHUrC1m': 3
```