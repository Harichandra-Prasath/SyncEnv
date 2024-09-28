# SyncEnv

Simple tool to Track and load the environment variables specified to directory   

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [Future Work](#future-work)

## Installation

1. Install the binary  

```bash
go install github.com/Harichandra-Prasath/SyncEnv
```

2. Run the binary  

```bash 
SyncEnv
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

**Unpacking**  
In this action, SyncEnv reads the stored variables and unpacks them to a loadable file  
```bash
SyncEnv --unpack
```

**Loading**  
This is the action where SyncEnv loads the latest unpacked variables. As child process cannot write to it's Parent, you need to run this action with bash eval.   
```bash
eval `SyncEnv --load`  
```

**Peek**
You can see all the variables in the store   
```bash
SyncEnv --peek
```

**Quick Notes**   
1. Currently only one action is supported per run  
2. You can skip unpacking if there are no changes made in the store and you just wants to load the previous state  

## Future-Work

1. Loading from local .env   
2. Multiple Actions per Run   