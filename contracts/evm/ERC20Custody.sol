// SPDX-License-Identifier: MIT
// v1.0, 2023-01-10
pragma solidity 0.8.7;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

/// @title ERC20Custody.
/// @notice ERC20Custody for depositing ERC20 assets into HanaNetwork and making operations with them.
contract ERC20Custody is ReentrancyGuard {
    using SafeERC20 for IERC20;

    error NotWhitelisted();
    error NotPaused();
    error InvalidSender();
    error InvalidTSSUpdater();
    error ZeroAddress();
    error IsPaused();
    error HanaMaxFeeExceeded();
    error ZeroFee();

    /// @notice If custody operations are paused.
    bool public paused;
    /// @notice TSSAddress is the TSS address collectively possessed by Hana blockchain validators.
    address public TSSAddress;
    /// @notice Threshold Signature Scheme (TSS) [GG20] is a multi-sig ECDSA/EdDSA protocol.
    address public TSSAddressUpdater;
    /// @notice Current hana fee for depositing funds into HanaNetwork.
    uint256 public hanaFee;
    /// @notice Maximum hana fee for transaction.
    uint256 public immutable hanaMaxFee;
    /// @notice Hana ERC20 token .
    IERC20 public immutable hana;
    /// @notice Mapping of whitelisted token => true/false.
    mapping(IERC20 => bool) public whitelisted;

    event Paused(address sender);
    event Unpaused(address sender);
    event Whitelisted(IERC20 indexed asset);
    event Unwhitelisted(IERC20 indexed asset);
    event Deposited(bytes recipient, IERC20 indexed asset, uint256 amount, bytes message);
    event Withdrawn(address indexed recipient, IERC20 indexed asset, uint256 amount);
    event RenouncedTSSUpdater(address TSSAddressUpdater_);
    event UpdatedTSSAddress(address TSSAddress_);
    event UpdatedHanaFee(uint256 hanaFee_);

    /**
     * @dev Only TSS address allowed modifier.
     */
    modifier onlyTSS() {
        if (msg.sender != TSSAddress) {
            revert InvalidSender();
        }
        _;
    }

    /**
     * @dev Only TSS address updater allowed modifier.
     */
    modifier onlyTSSUpdater() {
        if (msg.sender != TSSAddressUpdater) {
            revert InvalidTSSUpdater();
        }
        _;
    }

    constructor(address TSSAddress_, address TSSAddressUpdater_, uint256 hanaFee_, uint256 hanaMaxFee_, IERC20 hana_) {
        TSSAddress = TSSAddress_;
        TSSAddressUpdater = TSSAddressUpdater_;
        hanaFee = hanaFee_;
        hana = hana_;
        hanaMaxFee = hanaMaxFee_;
    }

    /**
     * @dev Update the TSSAddress in case of Hana blockchain validator nodes churn.
     * @param TSSAddress_, new tss address.
     */
    function updateTSSAddress(address TSSAddress_) external onlyTSSUpdater {
        if (TSSAddress_ == address(0)) {
            revert ZeroAddress();
        }
        TSSAddress = TSSAddress_;
        emit UpdatedTSSAddress(TSSAddress_);
    }

    /**
     * @dev Update hana fee
     * @param hanaFee_, new hana fee
     */
    function updateHanaFee(uint256 hanaFee_) external onlyTSS {
        if (hanaFee_ == 0) {
            revert ZeroFee();
        }
        if (hanaFee_ > hanaMaxFee) {
            revert HanaMaxFeeExceeded();
        }
        hanaFee = hanaFee_;
        emit UpdatedHanaFee(hanaFee_);
    }

    /**
     * @dev Change the ownership of TSSAddressUpdater to the Hana blockchain TSS nodes.
     * Effectively, only Hana blockchain validators collectively can update TSSAddress afterwards.
     */
    function renounceTSSAddressUpdater() external onlyTSSUpdater {
        if (TSSAddress == address(0)) {
            revert ZeroAddress();
        }
        TSSAddressUpdater = TSSAddress;
        emit RenouncedTSSUpdater(msg.sender);
    }

    /**
     * @dev Pause custody operations.
     */
    function pause() external onlyTSS {
        if (paused) {
            revert IsPaused();
        }
        if (TSSAddress == address(0)) {
            revert ZeroAddress();
        }
        paused = true;
        emit Paused(msg.sender);
    }

    /**
     * @dev Unpause custody operations.
     */
    function unpause() external onlyTSS {
        if (!paused) {
            revert NotPaused();
        }
        paused = false;
        emit Unpaused(msg.sender);
    }

    /**
     * @dev Whitelist asset.
     * @param asset, ERC20 asset.
     */
    function whitelist(IERC20 asset) external onlyTSS {
        whitelisted[asset] = true;
        emit Whitelisted(asset);
    }

    /**
     * @dev Unwhitelist asset.
     * @param asset, ERC20 asset.
     */
    function unwhitelist(IERC20 asset) external onlyTSS {
        whitelisted[asset] = false;
        emit Unwhitelisted(asset);
    }

    /**
     * @dev Deposit asset amount to recipient with message that encodes additional hananetwork evm call or message.
     * @param recipient, recipient address.
     * @param asset, ERC20 asset.
     * @param amount, asset amount.
     * @param message, bytes message or encoded zetechain call.
     */
    function deposit(
        bytes calldata recipient,
        IERC20 asset,
        uint256 amount,
        bytes calldata message
    ) external nonReentrant {
        if (paused) {
            revert IsPaused();
        }
        if (!whitelisted[asset]) {
            revert NotWhitelisted();
        }
        if (hanaFee != 0 && address(hana) != address(0)) {
            hana.safeTransferFrom(msg.sender, TSSAddress, hanaFee);
        }
        uint256 oldBalance = asset.balanceOf(address(this));
        asset.safeTransferFrom(msg.sender, address(this), amount);
        // In case if there is a fee on a token transfer, we might not receive a full exepected amount
        // and we need to correctly process that, o we subtract an old balance from a new balance, which should be an actual received amount.
        emit Deposited(recipient, asset, asset.balanceOf(address(this)) - oldBalance, message);
    }

    /**
     * @dev Withdraw asset amount to recipient by custody TSS owner.
     * @param recipient, recipient address.
     * @param asset, ERC20 asset.
     * @param amount, asset amount.
     */
    function withdraw(address recipient, IERC20 asset, uint256 amount) external nonReentrant onlyTSS {
        if (!whitelisted[asset]) {
            revert NotWhitelisted();
        }
        IERC20(asset).safeTransfer(recipient, amount);
        emit Withdrawn(recipient, asset, amount);
    }
}
