#!/bin/bash
set -o errexit

pushd "internal" > /dev/null

sqlboiler mysql
