pragma solidity 0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";


contract VerifierEIP712 {
    uint256 public constant version = 1;

    bytes32 public constant DOMAIN_TYPEHASH = keccak256("EIP712Domain(string name,uint256 chainId,address verifyingContract)");
    bytes32 public constant DELEGATION_TYPEHASH = keccak256("TransferFrom(address receiver,uint256 amount,address erc20,uint256 nonce)");

    mapping(address => uint) public nonces;
    mapping(address => bool) public whitelisted;

    function transferFromBySig(address receiver, address erc20, uint nonce, bytes32 r, bytes32 s, uint8 v) public {
        require(whitelisted(erc20), "VerifierEIP712: erc20 is not whitelisted");
        bytes32 domainSeparator = keccak256(abi.encode(DOMAIN_TYPEHASH, keccak256(bytes(name)), getChainId(), address(this)));
        bytes32 structHash = keccak256(abi.encode(DELEGATION_TYPEHASH, delegatee, nonce));
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", domainSeparator, structHash));

        address signer = ecrecover(digest, v, r, s);

        require(signer != address(0), "VerifierEIP712: bad signature");
        require(nonce == nonces[signer]++, "VerifierEIP712: problem with signer nonce");

        IERC20(erc20).transferFrom(signer, receiver);
    }



    function getChainId() internal pure returns (uint) {
        uint256 chainId;
        assembly {chainId := chainid()}
        return chainId;
    }
}