#!/bin/bash

pushd sql/schema
goose postgres postgres://postgres:postgres@localhost:5432/chirpy up
popd

sqlc generate