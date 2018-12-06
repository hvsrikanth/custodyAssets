#!/bin/bash
#
# Start script for the Asset Custody usecase. There are 6 nodes and each node is stopped / started in this script.
#
# Exit on first error, print all commands.
set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

SLEEPY=10
# STOP all the containers
sudo docker-compose -f docker-compose-regulator.yml down
sleep $SLEEPY
sudo docker-compose -f docker-compose-depository.yml down
sleep $SLEEPY
sudo docker-compose -f docker-compose-bank.yml down
sleep $SLEEPY
sudo docker-compose -f docker-compose-exchange.yml down
sleep 10
sudo docker-compose -f docker-compose-custodian.yml down
sleep $SLEEPY
sudo docker-compose -f docker-compose-investor.yml down

# START all the containers
sudo docker-compose -f docker-compose-investor.yml up -d
sleep $SLEEPY
sudo docker-compose -f docker-compose-custodian.yml up -d
sleep $SLEEPY
sudo docker-compose -f docker-compose-exchange.yml up -d
sleep $SLEEPY
sudo docker-compose -f docker-compose-bank.yml up -d
sleep $SLEEPY
sudo docker-compose -f docker-compose-depository.yml up -d
sleep $SLEEPY
sudo docker-compose -f docker-compose-regulator.yml up -d
sleep $SLEEPY

# Create the channel
sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer channel create -o orderer.example.com:7050 -c tradingchannel -f /etc/hyperledger/configtx/channel.tx
sleep $SLEEPY

# Join peer0.investor.example.com to the channel.
sudo docker exec -e "CORE_PEER_LOCALMSPID=InvestorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/investor.example.com/users/Admin@investor.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.investor.example.com:7051" cli peer channel join -b tradingchannel.block
sleep $SLEEPY

# Join peer0.custodian.example.com to the channel.
sudo docker exec -e "CORE_PEER_LOCALMSPID=CustodianMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/custodian.example.com/users/Admin@custodian.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.custodian.example.com:7051" cli peer channel join -b tradingchannel.block
sleep $SLEEPY

# Join peer0.exchange.example.com to the channel.
sudo docker exec -e "CORE_PEER_LOCALMSPID=ExchangeMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/exchange.example.com/users/Admin@exchange.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.exchange.example.com:7051" cli peer channel join -b tradingchannel.block
sleep $SLEEPY

# Join peer0.bank.example.com to the channel.
sudo docker exec -e "CORE_PEER_LOCALMSPID=BankMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/bank.example.com/users/Admin@bank.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.bank.example.com:7051" cli peer channel join -b tradingchannel.block
sleep $SLEEPY

# Join peer0.depository.example.com to the channel.
sudo docker exec -e "CORE_PEER_LOCALMSPID=DepositoryMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/depository.example.com/users/Admin@depository.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.depository.example.com:7051" cli peer channel join -b tradingchannel.block
sleep $SLEEPY

# Join peer0.regulator.example.com to the channel.
sudo docker exec -e "CORE_PEER_LOCALMSPID=RegulatorMSP" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/regulator.example.com/users/Admin@regulator.example.com/msp" -e "CORE_PEER_ADDRESS=peer0.regulator.example.com:7051" cli peer channel join -b tradingchannel.block
sleep $SLEEPY
