# Git Ssh Server

## Description

This is a simple ssh server that allows you to use git over ssh. It allows `git-receive-pack` command only.

## Requirements

- [git](https://git-scm.com/) (to test the server)
- [go](https://golang.org/) (to build the server)
- [docker](https://www.docker.com/) (to run the server in a container)
- [compose](https://docs.docker.com/compose/) (to run a sample environment)

## Controlling access

The server does not use an `authorized_keys` file like [openssh-server](https://www.openssh.com/) would. Instead, it queries a [postgresql](https://www.postgresql.org/) database to check if the user's public key is present in the database.

It then checks if the user is allowed to push to the repository he is trying to push to.

We extract that information from the `git-receive-pack` command the user is trying to execute. Here's an example:

```bash
git-receive-pack '/reponame'
```

The server will then check if the user is allowed to push to the repository `reponame` by checking if the user's public key is linked to the repository in the database.

Once this is validated, the server will execute the `git-receive-pack` command.

## Configuration

The server is configured using environment variables.

| Variable | Description |
| --- | --- |
| GIT_REPOSITORIES_FULL_BASE_PATH | The path to the directory containing all your repositories. |
| GIT_POSTGRESQL_USERNAME | The username to use to connect to the postgresql database. |
| GIT_POSTGRESQL_PASSWORD | The password to use to connect to the postgresql database. |
| GIT_POSTGRESQL_DATABASE_NAME | The name of the postgresql database. |
| GIT_POSTGRESQL_PORT | The port to use to connect to the postgresql database. |
| GIT_POSTGRESQL_HOST | The host to use to connect to the postgresql database. |
| GIT_HOST_SIGNER_PATH | The path to the host key. |

## Docker

### Configuration

The server requires a few volumes to be mounted in the container. The following volumes are required:

- A volume containing the host key
- A volume containing the git repositories
- A volume to the docker socket to be able to build the repositories via pack and docker

You can use the provided [compose file](compose.yml) to run the server in a docker container.

It bootstraps a postgresql database, this application and a debian container with git installed and a sample nodejs project.

```bash
# Prepare the environment
# This script creates a pair of ssh keys
./_scripts/init-dev-env.sh

# Launch the containers
docker compose up -d

# Populate the database
cd /path/to/projects

git clone git@github.com:paastech-cloud/api.git

cd api

# Create a .env file based on the .env.example file where DATABASE_URL is set to postgresql://paastech:paastech@postgres:5432/paastech
# Also add the public key you generated earlier to the .env file based on the .env.example file
npm install
# Create the database schema
npx prisma db push
# Seed the database
npx prisma db seed

# Back into the root directory of this project
# Exec into the client container
docker compose exec client bash

# You should now be in a git repository with a sample nodejs project
# You can now push to the server using this command
debug
```
