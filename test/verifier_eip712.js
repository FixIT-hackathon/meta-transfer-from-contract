const RelayerEIP712 = artifacts.require("RelayerEIP712");
const Web3 = require('web3');
const EIP712 = require('./helps/eip.js');
const Token = artifacts.require("TestToken");
const truffleAssert = require('truffle-assertions');

const web3 = new Web3('http://localhost:8545');

const createTransferFromBySigMessage = (contract, erc20, receiver, amount, chainId = 1, nonce = 0) => {

    const primaryType = 'TransferFrom';
    const domain = {name: 'RelayerEIP712', chainId: chainId, verifyingContract: contract};
    const message = {receiver, amount, erc20, nonce};

    return JSON.stringify({types, primaryType, domain, message});
};

contract("RelayerEIP712", function (accounts) {
    let verifier
    let token

    const SENDER = "0x98E003e1A08317DFac1C72172Cad9FFbB0a232aB",
        TX_SENDER = "0x5176b00E912494b18772B543Da98CeF28539af0B",
        RECEIVER = "0xd1F98B4dF6E832d3e591E2fD2dD64D962c105d71",
        SENDER_PRIV_KEY = "0xd9c1f4f836f257cb9b2a4c99a09374255d7badd4254c6c9364b949e5f4a0ae7d"


    it("deploying", async () => {
        token = await Token.new([SENDER], [web3.utils.toWei("100000", "ether")])
        verifier = await RelayerEIP712.new([token.address])
    })

    it("should deploy token", async () => {
        let addr = token.address

        assert.notEqual(addr, 0x0)
        assert.notEqual(addr, 0)
        assert.notEqual(addr, undefined)
    });

    it("should deploy verifier", async () => {
        let addr = verifier.address

        assert.notEqual(addr, 0x0)
        assert.notEqual(addr, 0)
        assert.notEqual(addr, undefined)
    });

    it("should transfer", async () => {
        await token.approve(verifier.address, 1010, {
            from: SENDER
        })

        const name = 'RelayerEIP712';
        const chainId = 1;


        const domain = {name: name, chainId: chainId, verifyingContract: verifier.address};

        let types = {
            TransferFrom: [
                {name: 'receiver', type: 'address'},
                {name: 'amount', type: 'uint256'},
                {name: 'fee', type: 'uint256'},
                {name: 'erc20', type: 'address'},
                {name: 'nonce', type: 'uint256'}
            ]
        }

        let amount = 1000
        let nonce = 0
        let fee = 10
        let message = {receiver: RECEIVER, amount: amount, fee: fee, erc20: token.address, nonce: nonce};

        const {v, r, s} = EIP712.sign(domain, 'TransferFrom', message, types, SENDER_PRIV_KEY);


        let result = await verifier.transferFromBySig(RECEIVER, token.address, amount, fee, nonce, r, s, v, {
                from: TX_SENDER
            }
        )

        assert.strictEqual(result.logs[0].args['signer'].toString(), SENDER.toString())
        assert.strictEqual(result.logs[0].args['nonce'].toString(), String(nonce))
        assert.strictEqual((await token.balanceOf(TX_SENDER)).toString(), String(fee))
        assert.strictEqual((await token.balanceOf(RECEIVER)).toString(), String(amount))
        assert.strictEqual((await verifier.nonces.call(RECEIVER)).toString(), String(nonce))
    })
})