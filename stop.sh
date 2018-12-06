#!/bin/bash
#
# Stop script for the Asset Custody usecase. There are 6 nodes and each node is stopped in this script.
#

sudo docker-compose -f docker-compose-regulator.yml down
sleep 10
sudo docker-compose -f docker-compose-depository.yml down
sleep 10
sudo docker-compose -f docker-compose-bank.yml down
sleep 10
sudo docker-compose -f docker-compose-exchange.yml down
sleep 10
sudo docker-compose -f docker-compose-custodian.yml down
sleep 10
sudo docker-compose -f docker-compose-investor.yml down

sudo docker ps
