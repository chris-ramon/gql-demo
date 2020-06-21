#!/bin/bash
set -o errexit

MYSQL_USER=${MYSQL_USER} MYSQL_PASSWORD=${MYSQL_PASSWORD} MYSQL_DB=${MYSQL_DB} go run server.go
