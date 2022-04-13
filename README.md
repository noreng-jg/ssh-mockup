# ssh-mockup 

Concept proof ssh server possible connection types

## Requirements

- Docker
- docker-compose

## Setup

Before starting the ssh server, run:

```
./bin/keygen
```

Run the project 

```
docker-compose build && docker-compose up
```

## Ssh Pty checks 

### SCP

```
./bin/scp
I am a scp
```

### Pty not iteractive 

```
./bin/sshcmd
I am cmdpty
```

### Pty iteractive 

```
./bin/sshi
I am a itpty
Connection to localhost closed
```
