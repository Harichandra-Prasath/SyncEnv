package main

const MAIN_TEMPLATE = `Welcome to SyncEnv!!!
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
const LOAD_TEMPLATE = "\nLoading: --load\nThis action loads the unpacked varaibles to parent process.\nIt has to be run as bash eval\nUsage: eval `SyncEnv --load`\n"

const LOAD_FROM_FILE_TEMPLATE = "\nLoad from File: --load-from-file\nThis action loads the variables defined in the local file.\nIt has to be run as bash eval\nUsage: eval`SyncEnv --load-from-file <path>`"
