# CPChain DApp - Message

Message DApp is the fundamental DApp of the CPChain DApp Ecosystem, you can use this DApp to send and receive encrypted messae，This DApp base on [Identity DApp](https://github.com/CPChain/cpchain-dapps-identity).

The sender can get the public key of the receiver from [Identity]((https://github.com/CPChain/cpchain-dapps-identity)), then use the public key encrypts the message, send the message via contract (Contract will emit an event). When the receiver gets the message, use their own private key to decrypted the message.

## Setup

```bash

npm install

# test
truffle test


```

Contract Address： 0x856c36486163dB6f9aEbeD1407a3c6C51FD7566E

## Methods

+ `sendMessage(to: address, message: string)`: Send the message to an address
+ `count()`: The count of all messages
+ `sentCount(addr: address)`: get the count of messages sent by the address
+ `receivedCount(addr: address)`: get the count of messages received by the address

When there a new message be sent, the DApp when emitting an event, information as below:

+ `from`: sender
+ `to`: receiver
+ `sentID`: message ID of the sender(auto increment)
+ `recvID`: message ID of the receiver(auto increment)
+ `message`: content of the message

Message content need to be encoded in JSON, required fields as below:

```json

{
    "message": "message",
    "version": "1.0"
}

```

+ *version is for the parser on the client*

### Version 2.0

Now, the latest CPC-Wallet(v1.4.3) support you send `password` to others.

```json

{
    "message": "message",
    "type": "password",
    "version": "2.0"
}


```

## Test with Cli

```bash

# send message
build/main message send --to d7b93566d41b6dc3858b8dace06a63ac8f272c81 --msg "HelloWorld" --keystore ./dapps-admin/keystore/ --endpoint http://52.220.174.168:8501 --contractaddr 0x856c36486163dB6f9aEbeD1407a3c6C51FD7566E

build/main message show-configs --keystore ./dapps-admin/keystore/ --endpoint http://52.220.174.168:8501 --contractaddr 0x856c36486163dB6f9aEbeD1407a3c6C51FD7566E

```
