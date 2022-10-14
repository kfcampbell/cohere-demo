# cohere-demo

This repo is for integrating GitHub issues with cohere.ai's APIs to make open source maintainers' daily tasks easier.

## usage

ensure labels in a repository are applied to at least 2 issues for the label to be elligible as a classification target.

## developing

this project uses a `.env` file to configure environment variables.

see [example.env](./example.env) for requirements.

running locally can be done via:

```sh
env $(cat .env) go run main.go classify
env $(cat .env) dlv debug main.go -- classify
```
