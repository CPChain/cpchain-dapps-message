# CPChain DApp - Message

Message 是 CPChain DApp 生态中的基础消息服务，提供加密通信功能，其基于 Identity 身份 DApp。

用户可通过 Identity Dapp 获取通信对象的公钥，然后使用其公钥对信息进行加密，加密后通过 Message 进行消息投递（消息通过事件发出）。通信对象收到消息后，使用 Identity 的私钥进行消息解密，从而实现加密通信。

## 合约方法

+ `sendMessage(to: address, message: string)`: 发送消息
+ `count()`: 发送消息总数
+ `sentCount(addr: address)`: 获取 addr 的发送消息数
+ `receivedCount(addr: address)`: 获取 addr 的接收消息数，可通过在客户端缓存此数字，以次判断是否需要向 explorer 拉取新消息

发送消息时，合约将触发事件，explorer 负责同步此事件，事件有以下信息：

+ `from`: 发送人
+ `to`: 接收人
+ `sentID`: 发送者的消息ID（自增）
+ `recvID`: 接收者的消息ID（自增）
+ `message`: 消息内容

消息内容请使用 JSON 进行序列化：

```json

{
    "message": "message"
}

```
