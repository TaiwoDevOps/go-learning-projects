# Platform Reminder app

- This is an application that reminds you of your TODO directly from the computer.

-- Use cases:

- this is useful for reminding yourself of what do directly from your machine(laptop)
- Useful fo those who work heavily on their laptop

## App Features:

# simple alert

- set simple reminders and have an alert dialog

# strict alerts (This are timer countdown based and can eb extensible)

- set strict reminders such as "shut down laptop" to help user take breaks
- set reminders to lock user away from screen

## Todo:

- Algo for this project:-
  -- get device time stamp
  time.Now() or use when package to make it more English readable

  -- using system dialog api
  check for platform

  - os := runtime.GOOS
    if OS is "darwin" Mac, use the sound option to compliment the beeep Notifier
    else, just show log

## Project dependencies

- beeep from gen2brain - this is to use platform dialog API
- when from olebedev - this is to work with a more robust time func
