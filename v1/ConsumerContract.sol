pragma solidity ^0.8.12;

interface SanctionsList {
    function isSanctioned(address addr) external view returns (bool);
}

contract ConsumerContract {
    address constant SANCTIONS_CONTRACT = 0x40C57923924B5c5c5455c48D93317139ADDaC8fb;

    function transfer(address to, uint256 amount) external {
        SanctionsList sanctionsList = SanctionsList(SANCTIONS_CONTRACT);
        bool isToSanctioned = sanctionsList.isSanctioned(to);
        require(!isToSanctioned, "Transfer to sanctioned address");
    }
}