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

Porting: --port
This actions ports the SyncEnv variables to requested file
Usage: SyncEnv --port <path>, SyncEnv --port

`

const LOAD_TEMPLATE = `Loading: load
This action loads the latest variables to the current shell session
This has to be ran as shell's eval.
If you want to load from local .env file, add --from-file
Additionally --no-debug flag can be paused to prevent message outputs on load action
Usage: SyncEnv load, SyncEnv load --from-file <path>, SyncEnv load --no-debug
`

// Check syncenv_hook, if not found, add it to the PROMPT_COMMAND
const BASH_HOOK = `
syncenv_hook(){
local prev_exit=$?;
eval "$(SyncEnv load --no-debug)";
return $prev_exit;
};
if [[ ";${PROMPT_COMMAND[*]:-};" != *";syncenv_hook;"* ]]; then
  PROMPT_COMMAND="syncenv_hook${PROMPT_COMMAND:+;$PROMPT_COMMAND}"
fi
`

// append the syncenv_hook to zsh chpwd_functions
const ZSH_HOOK = `
syncenv_hook(){
eval "$(SyncEnv load --no-debug)"
}
chpwd_functions=("syncenv_hook" "${chpwd_functions[@]}")
`
