//SPDX-License-Identifier: Unlicense
pragma solidity ^0.8.0;

// This is a gateway contract for ERC20 and ERC721 on ethereum.
contract TokenGateway {
    event TransferIn(string assetId, address recipient, uint256 amount);
    event TransferInAssetOfThisChain(
        address assetAddr,
        address recipient,
        uint256 amount,
        bool success
    );
    event TransferWithin(string assetId, address recipient, uint256 amount);
    event TransferOut(
        string assetId,
        address from,
        string toChain,
        string recipient,
        uint256 amount
    );
    event TransferOutFromContract(
        address assetAddr,
        string toChain,
        string recipient,
        uint256 amount,
        bool success
    );

    address _owner;
    string _thisChain;

    // Map from: token -> set of tokens owner
    // Each tokens owner set is a map between owner id (identified by addr and chain) and owned amount.
    mapping(string => mapping(address => uint256)) _store;

    // List of chain that we allow user to send asset to.
    mapping(string => bool) _allowedChain;

    mapping(address => uint256) _assetBalance;

    constructor(string memory chain) {
        _owner = msg.sender;
        _thisChain = chain;
    }

    // Transfers tokens from other chains into this chain. This function can be only called by the
    // owner of this contract.
    //
    // Use this function if you want to transfer in assets NOT originated from this chain. Use
    // transferInAssetOfThisChain to transfer asset from this chain.
    function transferIn(
        string memory assetId,
        address recipient,
        uint256 amount
    ) public {
        require(msg.sender == _owner);

        _store[assetId][recipient] += amount;

        emit TransferIn(assetId, recipient, amount);
    }

    // Transfers a token from another chain to this chain. This token was created within this chain
    // and hence we do not need a chain id in the function call to identify the token.
    function transferInAssetOfThisChain(
        address assetAddr,
        address recipient,
        uint256 amount
    ) public {
        require(msg.sender == _owner);
        require(_assetBalance[assetAddr] >= amount);

        (bool success, ) = assetAddr.call(
            abi.encodeWithSignature(
                "transferFrom(address,address,uint256)",
                address(this),
                recipient,
                amount
            )
        );

        if (success) {
            _assetBalance[assetAddr] -= amount;
        }

        emit TransferInAssetOfThisChain(assetAddr, recipient, amount, success);
    }

    // Adds a new chain that we can support.
    function addAllowedChain(string memory newChain) public {
        require(msg.sender == _owner);

        _allowedChain[newChain] = true;
    }

    // ---- /

    // Transfer a token within this network.
    function transferWithin(
        string memory assetId,
        address recipient,
        uint256 amount
    ) public {
        require(amount > 0);
        require(_store[assetId][msg.sender] >= amount);

        _store[assetId][msg.sender] -= amount;
        _store[assetId][recipient] += amount;

        emit TransferWithin(assetId, recipient, amount);
    }

    // Transfer a token to outside network.
    function transferOut(
        string memory assetId,
        string memory toChain,
        string memory recipient,
        uint256 amount
    ) public {
        require(amount > 0);
        require(_store[assetId][msg.sender] >= amount);
        require(_allowedChain[toChain]);

        _store[assetId][msg.sender] -= amount;

        emit TransferOut(assetId, msg.sender, toChain, recipient, amount);
    }

    // Transfers a token from one of the existing ERC20/ERC721 contract address to outside network.
    // The caller has to make sure that it has approve enough tokens for this contract to withdraw
    // from the token contract.
    function transferOutFromContract(
        address assetAddr,
        string memory toChain,
        string memory recipient,
        uint256 amount
    ) public {
        require(amount > 0);
        require(_allowedChain[toChain]);

        (bool success, ) = assetAddr.call(
            abi.encodeWithSignature(
                "transferFrom(address,address,uint256)",
                msg.sender,
                address(this),
                amount
            )
        );

        if (success) {
            _assetBalance[assetAddr] += amount;
        }

        emit TransferOutFromContract(
            assetAddr,
            toChain,
            recipient,
            amount,
            success
        );
    }

    function changeOwner(address newOwner) public {
        require(msg.sender == _owner);

        _owner = newOwner;
    }

    function getOwner() public view returns (address) {
        return _owner;
    }

    function getBalance(string memory assetId, address account)
        public
        view
        returns (uint256)
    {
        return _store[assetId][account];
    }

    function isChainAllowed(string memory chain) public view returns (bool) {
        return _allowedChain[chain];
    }
}