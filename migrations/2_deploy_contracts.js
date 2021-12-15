// Deploy Message
var Message = artifacts.require("./Message.sol");

module.exports = function(deployer) {
        deployer.deploy(Message); //"参数在第二个变量携带"
};
