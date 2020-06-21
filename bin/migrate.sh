#!/bin/bash
set -o errexit

mysql ${MYSQL_DB} -u ${MYSQL_USER} -p${MYSQL_PASSWORD} < schema.sql
