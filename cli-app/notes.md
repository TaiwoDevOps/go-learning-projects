# Designing a CLI application

- A neatly design CLI application must be concise and must be clear on what to do, and this application must be small.
- A CLI app must have an action to perform - the command eg. cd
  --> It must be clear on whats its to do - eg, cd means "change directory", ls means "list current directory"
- it can take args[s] eg, cp [file] [newfile] "copy from file to newfile
  --> the order matters, the instruction must be carried out sequentially
  --> spaces matter, it is what defines the sequence at which the instruction must be followed
- It can have flags, meaning optional parameters
  --> A flag modifies behavior eg, rm -rf podfile.lock
