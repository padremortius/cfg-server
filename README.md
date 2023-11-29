# cfg-server

## About project

cfg-server - web-server publishing configuration from git-server. Now support work only through ssh channel

## Build project

For build project start in root directory
`go build .`

## Params to start

`./cfg-server --help` - short help for params
`./cfg-server --repoUrl=git@github.com:padremortius/configs-example.git --repoBranch=main --searchPath=dev`

For current time ssh key and password for ssh-key (maybe empty) must be save in cfg-server.json.
For example:
`cat cfg-server.json`

    ```{
        "git.pKey": "-----BEGIN OPENSSH PRIVATE KEY-----\n......w=\n-----END OPENSSH PRIVATE KEY-----",
        "git.password": "testPassword"
    }```

## How used

`curl http://localhost:8080/dev/downdetector.yaml` - get configuration without using searchPath
`curl http://localhost:8080/downdetector.json` - get configuration with using searchPath

## Tasks for future

- Add support params (${param}) in configuration files
- Add using external ssh key and set path to it in env variable or startup flag
- Add support configuration in json files
