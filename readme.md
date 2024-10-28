# SyncEnv

Simple Open Source tool to Track and load the environment variables specified to directory   

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [Hooking](#hooking)
4. [Future Work](#future-work)
5. [Contributions](#contributions)

## Installation

1. Install the binary  

```bash
go install github.com/Harichandra-Prasath/SyncEnv
```

## Usage

**Initialise**  
Start by running the below to add the current directory to SyncEnv     
```bash
SyncEnv --init
```
  
**Additon**   
You can add new variables  
```bash
SyncEnv --add foo=bar --add baz=gaz  
``` 
**Updation**  
You can update existing variables   
```bash
SyncEnv --update foo=gaz --update baz=bar
```

**Loading**  
This is the action where SyncEnv loads the attached variables. As child process cannot write to it's Parent, you need to run this action with bash eval.   
```bash
eval `SyncEnv load`  
```   
   
If you want to load from a local file, add --from-file flag.   

```bash
eval `SyncEnv load --from-file <path-to-file>`
```
   
Additionally, --no-debug flag can be passed to prevent message outputs.  

```bash
eval `SyncEnv load --no-debug`
```
  
**Peek**   
You can see all the variables in the store   
```bash
SyncEnv --peek
```

**Port**   
You can Port your SyncEnv variables to a standalone file.  
```bash
SyncEnv --port <path-to-file>
```


**Help**  
To see the help menu  
```bash
SyncEnv --help
```
  

## Hooking

For auto-loading of variables when entering the directory, you can create a hook in your shell profile.  

Currently SyncEnv supports zsh and bash.     

For Zsh, add the below in your .zshrc  
```bash
eval "$(SyncEnv hook --shell zsh)"
```

For Bash, add the below in your .bashrc  
```bash
eval "$(SyncEnv hook --shell bash)"
```
     
  
## Future-Work
  
1. Port to Windows  
2. Additional Shells like fish,etc..  
3. Reducing Write actions if possible  

## Contributions  

This is a small project open to everyone.Contributions and improvements are always welcome.  
