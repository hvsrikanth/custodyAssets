#!/bin/bash
#
# Start script for the Asset Custody usecase. There are 6 nodes and each node is stopped / started in this script.
#
# Exit on first error, print all commands.
set -ev

# Test harness - onboard_investor
sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C tradingchannel -n custodycc -c '{"function":"onboard_investor","Args":["hvsrikanth","Srikanth","Harathi","AACPH1111G","ok","11223344","090999909"]}'

# Test harness - init_bank
sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer chaincode invoke -o orderer.example.com:7050 -C tradingchannel -n custodycc -c '{"function":"init_bank","Args":[]}'

