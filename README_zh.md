# CPChain DApp - Message

Message 是 CPChain DApp 生态中的基础消息服务，提供加密通信功能，其基于 Identity 身份 DApp。

用户可通过 Identity Dapp 获取通信对象的公钥，然后使用其公钥对信息进行加密，加密后通过 Message 进行消息投递（消息通过事件发出）。通信对象收到消息后，使用 Identity 的私钥进行消息解密，从而实现加密通信。
