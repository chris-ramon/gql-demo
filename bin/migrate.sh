#!/bin/bash
set -o errexit

mysql gql_demo_dev -u ${MYSQL_USER} -p${MYSQL_PASSWORD} < schema.sql
