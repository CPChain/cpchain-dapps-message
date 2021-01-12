
pragma solidity ^0.4.24;

import "./lib/safeMath.sol";
import "./lib/set.sol";

contract Message {
    using Set for Set.Data;
    using SafeMath for uint256;

    address owner; // owner has permissions to modify parameters
    bool public enabled = true; // if upgrade contract, then the old contract should be disabled

    uint256 _count = 0; // users _count

    mapping (address => uint256) sent; // sent message count per address
    mapping (address => uint256) received; // received message count per address

    event NewMessage(address from, address to, uint256 sentID, uint256 recvID, string message);

    modifier onlyOwner() {require(msg.sender == owner);_;}
    modifier onlyEnabled() {require(enabled);_;}

    constructor () public {
        owner = msg.sender;
    }

    function sendMessage(address to, string message) public payable onlyEnabled {
        _count += 1;
        sent[msg.sender] += 1;
        received[to] += 1;
        emit NewMessage(msg.sender, to, sent[msg.sender], received[to], message);
    }

    function sentCount(address addr) public view returns (uint256) {
        return sent[addr];
    }

    function receivedCount(address addr) public view returns (uint256) {
        return received[addr];
    }

    function count() public view returns (uint256) {
        return _count;
    }

    // owner can enable and disable rnode contract
    function enableContract() public onlyOwner {
        enabled = true;
    }

    function disableContract() public onlyOwner {
        enabled = false;
    }
}
