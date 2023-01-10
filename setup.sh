#!/bin/bash

cd build/auth-database
docker-compose down
docker-compose up -d
sleep 3