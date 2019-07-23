# tgmsm - Telegram Session Manager
> ### Simple(and useless) CLI-tool for switching between multiple Telegram sessions

## Install
```
git clone https://github.com/mordorsun/tgmsm.git
cd tgmsm 
go install
```

## Usage 
```
tgmsm command_name [arg]
``` 

## Commands
```
list - list of saved sessions
save session_name - save active telegram session
rm session_name - remove saved session
purge - remove ALL saved sessions
back - back to previos session
switch session_name - switch to selected session
```

## TODO
 - test coverage
 - Windows & mac support
 - Compressing saving sessions
