// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "../interfaces/HanaInterfaces.sol";

contract HanaReceiverMock is HanaReceiver {
    event MockOnHanaMessage(address destinationAddress);

    event MockOnHanaRevert(address hanaTxSenderAddress);

    function onHanaMessage(HanaInterfaces.HanaMessage calldata hanaMessage) external override {
        emit MockOnHanaMessage(hanaMessage.destinationAddress);
    }

    function onHanaRevert(HanaInterfaces.HanaRevert calldata hanaRevert) external override {
        emit MockOnHanaRevert(hanaRevert.hanaTxSenderAddress);
    }
}
