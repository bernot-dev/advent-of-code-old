#!/bin/bash

# Figure out where we are
BASEDIR=$(git rev-parse --show-toplevel)
PWD=$(pwd)
RELATIVE=`realpath --relative-to="$BASEDIR" $PWD`

# Validate current directory (check year 2015-2029 🤞, day 1-25)
if [[ $RELATIVE =~ ^20(1[5-9]|2[0-9])/day/([1-9]|1[0-9]|2[0-5])$ ]]; then
	echo "Current Directory: $RELATIVE"
else
	echo "Not a valid AOC Directory: $RELATIVE";
	exit 1
fi

if [ -z "$(ls -A $PWD)" ]; then
	echo "OK: Directory is empty"
else
	echo "ABORT: Directory is not empty"
	exit 1
fi

# Write input to input.txt
curl --fail -b $BASEDIR/.session https://adventofcode.com/$RELATIVE/input -o input.txt
if [[ $? -ne 0 ]]; then
	echo "Unable to retrieve puzzle input. Is this year/day valid and available?"
	exit 1
fi

# Copy template
cp $BASEDIR/template/* $PWD

# Initialize go module
go mod init "github.com/bernot-dev/advent-of-code/$RELATIVE" &&
go get github.com/bernot-dev/advent-of-code/helpers@latest
