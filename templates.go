package main

const MAIN_TEMPLATE = `
Welcome to SyncEnv!!!
Simple Env Variables Management tool

Supported Actions

Intialise: --init 
This action adds the current directory to SycnEnv store
Usage: SyncEnv --init

Addtion: --add 
This action adds a new variable to the SyncEnv
Usage: SyncEnv --add foo=bar --add baz=gaz

Updation: --update
This action updates an existing variable
Usage: SyncEnv --update foo=gaz --update baz=bar

Peek: --peek
This action lets you have a glance at the stored variables
Usage: SyncEnv --peek

Unpacking: --unpack
This action unpacks the stored variables to a loadable format
Usage: SyncEnv --unpack

`

const LOAD_TEMPLATE = `Loading: load
This action loads the latest unpacked variables to the current shell session
This has to be ran as shell's eval.
If you want to load from local .env file, add --from-file
Additionally --no-debug flag can be paused to prevent message outputs on load action
Usage: SyncEnv load, SyncEnv load --from-file <path>, SyncEnv load --no-debug
`

const SYNCENV_HOOK = `syncenv_hook() {
  local pdir="$PWD"
  builtin cd "$@" || return
  if [[ "$pdir" != "$PWD" ]]
  then
      eval "$(SyncEnv load --no-debug)"
  fi
}

cd() {
  syncenv_hook "$@"
}
`
