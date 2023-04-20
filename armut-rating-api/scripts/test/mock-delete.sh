#!/bin/bash

echo "STARTING TO DELETE MOCKS"

find ../../internal -type f -iname \*_mock.go -delete

echo "END"