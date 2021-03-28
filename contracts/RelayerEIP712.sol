pragma solidity 0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract RelayerEIP712 is Ownable {
    using SafeMath for uint256;

    string public constant name = "RelayerEIP712";


    event Executed(uint256 nonce, address signer);


    bytes32 public constant DOMAIN_TYPEHASH = keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");
    bytes32 public constant TRANSFER_FROM_TYPEHASH = keccak256("TransferFrom(address receiver,uint256 amount,uint256 fee,address erc20,uint256 nonce)");

    mapping(address => uint) public nonces;
    mapping(address => bool) public whitelisted;

    constructor(address[] memory tokens){
        for (uint i = 0; i < tokens.length; i++) {
            whitelisted[tokens[i]] = true;
        }
    }

    function transferFromBySig(address receiver, address erc20, uint256 amount, uint256 fee, uint256 nonce, bytes32 r, bytes32 s, uint8 v) public {
        require(whitelisted[erc20], "RelayerEIP712: erc20 is not whitelisted");

        address signer = address(0);
        {
            bytes32 domainSeparator = keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes(name)), getChainId(), address(this)));
            bytes32 structHash = keccak256(abi.encode(TRANSFER_FROM_TYPEHASH, receiver, amount, fee, erc20, nonce));
            bytes32 digest = keccak256(abi.encodePacked("\x19\x01", domainSeparator, structHash));

            signer = ecrecover(digest, v, r, s);
            require(signer != address(0), "RelayerEIP712: bad signature");
            require(nonce == nonces[signer], "RelayerEIP712: problem with signer nonce");
        }

        nonces[signer] += 1;
        IERC20 ierc20 = IERC20(erc20);

        uint256 sum = amount.add(fee);
        require(ierc20.balanceOf(signer) >= sum, "RelayerEIP712: signer balance is too low");
        require(ierc20.allowance(signer, address(this)) >= amount, "RelayerEIP712: allowance is too low");

        ierc20.transferFrom(signer, receiver, amount);
        ierc20.transferFrom(signer, msg.sender, fee);

        emit Executed(nonce, signer);

    }

    function list(address erc20) external onlyOwner {
        require(!whitelisted[erc20], "RelayerEIP712: erc20 is already whitelisted");
        whitelisted[erc20] = true;
    }

    function unlist(address erc20) external onlyOwner {
        require(whitelisted[erc20], "RelayerEIP712: erc20 is not whitelisted");
        delete whitelisted[erc20];
    }

    function getChainId() internal view returns (uint) {
        uint256 chainId;
        assembly {chainId := chainid()}
        return chainId;
    }
}