#!/bin/bash

rm -rf _scripts/output/client*

rm -rf _scripts/output/server*

mkdir -p _scripts/output/client

mkdir -p _scripts/output/server

ssh-keygen -t ed25519 -C "userA@user.fr" -f _scripts/output/client/id_ed25519 -q -N ""

ssh-keygen -t ed25519 -C "git@paastech.fr" -f _scripts/output/server/id_ed25519 -q -N ""
