# redacted-environment-variables
A small replacement for `env` that redacts sensitive values

## Introduction

`redacted-environment-variables` replaces using `env` to show your
environment. It prints each environment variable for your current
process but replaces the actual values for specified variable names
with a redacted string (`********` by default) to avoid exposing
secrets or tokens.

It was written as a safer way to show my current environment when
pair programming or streaming.

## Usage

I host a lot of my code on GitHub so I often have a GitHub token defined
as an environment variable. This token allows you to modify my account
so it shouldn't be exposed to other people. Such as when typing `env`
while streaming to Twitch. By using `redacted-environment-variables` I
can reduce the chances of this happening.

First we'll set a fake `GITHUB_TOKEN` so we can confirm how things are working.

    export GITHUB_TOKEN=myrealtoken

We run `env` to ensure we can see the token

    env | grep GITHUB_TOKEN
    # oh no, I am undone!
    GITHUB_TOKEN=myrealtoken

If instead we run `redacted-environment-variables` (which I alias to `renv`)

    redacted-environment-variables | grep GIT
    # My modesty has been preserved.
    GITHUB_TOKEN=********

The tokens value is protected but you can see that it is set.

## Aliasing

`redacted-environment-variables` is a long name to remember so in my
shell I have an `renv` alias to make the command easier to use.

    alias renv=redacted-environment-variables

I can't directly alias it to `env` as a number of scripts use it to
locate their interpreter

    /usr/bin/env python

## Configuration

By default `redacted-environment-variables` will redact the following environment
variables:

    AWS_ACCESS_KEY_ID
    AWS_SECRET_ACCESS_KEY
    GITHUB_TOKEN
    GITHUB_AUTH_TOKEN
    _TOKEN

These are done as substring matches so `_TOKEN` will match any variable
including that phrase, such as `TWILLIO_TOKEN` and `CD_USER_TOKEN`

You can replace these values by specifying your own
[TOML](https://en.wikipedia.org/wiki/TOML) configuration file.


    cat .redacted-environment.toml
    [config]
    redacted = "YOUCANTSEEME"
    
    [redact]
    names = [
      "AWS_ACCESS_KEY_ID",
      "AWS_SECRET_ACCESS_KEY",
      "GITHUB_TOKEN",
      "GITHUB_AUTH_TOKEN",
      "MY_LOCAL_VALUE",
      "_TOKEN_"
    ]

`redacted-environment-variables` will check for a `.redacted-environment.toml`
file in the current directory, the users home directory, and the `/etc`
directory and will stop and use the first of those it finds.

## License

 * GPLv2

## Author

 * [Dean Wilson](https://www.unixdaemon.net)
