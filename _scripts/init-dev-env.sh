#!/bin/bash

rm -rf _scripts/output/client*

mkdir -p _scripts/output/client

ssh-keygen -t ed25519 -C "userA@user.fr" -f _scripts/output/client/id_ed25519 -q -N ""
