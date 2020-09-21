# redacted-environment-variables
A small replacement for `env` that redacts sensitive values

## Introduction

A replacement for using `env` to show your environment that protects your
important values.

## Usage

I host a lot of my code on GitHub so I often have a GitHub token defined
in my shell as an environment variable. This token allows you to modify
my account so it shouldn't be exposed to other people. Such as when
typing `env` while streaming to Twitch.

First we'll set a fake GITHUB_TOKEN so we can confirm how things are working.

    export GITHUB_TOKEN=myrealtoken

And then we run `env` to ensure we can see it

    env | grep GITHUB_TOKEN
    # oh no, I am undone!
    GITHUB_TOKEN=myrealtoken

If instead we run (#TODO: replace with released binary)

    go run main.go | grep GIT
    # My modesty has been preserved.
    GITHUB_TOKEN=XXXXXXXX

## License

 * GPLv2

## Author

 * [Dean Wilson](https://www.unixdaemon.net)
