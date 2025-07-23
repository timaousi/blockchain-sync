// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Counter {
    uint public count;

    // 添加一个 string 类型的 message 参数
    event CountChanged(uint newCount, address indexed changer, string message);

    constructor() {
        count = 0;
    }

    function increment() public {
        count++;
        // 触发事件，传入中文消息
        emit CountChanged(count, msg.sender, unicode"计数已增加 😄");
    }

    function decrement() public {
        if (count > 0) {
            count--;
            // 触发事件，传入中文消息
            emit CountChanged(count, msg.sender, unicode"计数已减少 😥");
        }
    }

    function getCount() public view returns (uint) {
        return count;
    }
}