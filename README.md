# kinniku-manager

Reference: https://musclewiki.com/

## Terminologies

- Training Exercise
- Training Set
- Training Menu

## How to use

After cloning this repository to your local environment, run following command.

```shell
make run-server
```

### Read all training exercises

```shell
curl localhost:8080/training_exercise/read_all
```

### Register new training exercise

```shell
curl -X POST -H "Content-Type: application/json" -d @sample.json localhost:8080/training_exercise/save
```

### Update existing training exercise

TBD

### Delete existing training exercise

TBD
