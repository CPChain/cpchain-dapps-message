const Message = artifacts.require("Message");

contract("Message", (accounts) => {
    it("Greet", async () => {
        const instance = await Message.deployed()
        await instance.sendMessage(accounts[1], "hello")
    })
})
