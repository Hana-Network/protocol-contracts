// SPDX-License-Identifier: MIT
pragma solidity 0.8.7;

import "@openzeppelin/contracts/access/Ownable2Step.sol";
import "../tools/HanaInteractor.sol";

contract HanaInteractorMock is Ownable2Step, HanaInteractor, HanaReceiver {
    constructor(address hanaConnectorAddress) HanaInteractor(hanaConnectorAddress) {}

    function onHanaMessage(
        HanaInterfaces.HanaMessage calldata hanaMessage
    ) external override isValidMessageCall(hanaMessage) {}

    function onHanaRevert(
        HanaInterfaces.HanaRevert calldata hanaRevert
    ) external override isValidRevertCall(hanaRevert) {}
}
