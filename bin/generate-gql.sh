#!/bin/bash
set -o errexit

pushd "cmd/graphql/gqlgen" > /dev/null

gqlgen generate
