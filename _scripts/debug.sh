#!/bin/bash

# This script is just a way to easily test the ssh server, it needs to be in a git repository
# who has a remote named origin pointing to the ssh server

git switch -c main

echo "$RANDOM" > README.md

git add .

git commit -m "update"

ssh-keygen -R "[127.0.0.1]:2222"

git push -uf origin main
