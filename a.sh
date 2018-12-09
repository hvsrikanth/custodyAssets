#!/bin/bash
#
#

if [ `sudo docker ps | grep "regulator" | wc -l` != 0 ]; then
    sudo docker-compose -f docker-compose-regulator.yml down
fi

if [ `sudo docker ps | grep "abc" | wc -l` == 0 ]; then
echo "shit works"
fi
