#!/bin/bash

if [ ! -e domainfinder/domainfinder ]; then

  echo building... domainfinder
  cd ./domainfinder/
  go build -o domainfinder

  echo building... synonyms
  cd ../synonyms
  go build -o ../domainfinder/lib/synonyms

  echo building... available
  cd ../available
  go build -o ../domainfinder/lib/available

  echo building... sprinkle
  cd ../sprinkle
  go build -o ../domainfinder/lib/sprinkle

  echo building... coolify
  cd ../coolify
  go build -o ../domainfinder/lib/coolify

  echo building... domainify
  cd ../domainify
  go build -o ../domainfinder/lib/domainify

  cd ../

  echo

fi

echo Please enter the word.
echo

# execute binary file
domainfinder/domainfinder
