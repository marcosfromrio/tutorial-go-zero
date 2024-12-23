# Go CLI Tutorial

Tutorial from JetBrains: https://www.jetbrains.com/guide/go/tutorials/cli-apps-go-cobra/
The app is called `Zero` and implements basic math operations.

## Using the binary file to run CLI commands

Cloning the repo:

1. `git clone git@github.com:marcosfromrio/tutorial-go-zero.git`
2. `cd tutorial-go-zero`

Example using binary in Linux:

- `./Zero add 5 2 // Addition of 5 and 2 = 7.000000.`
- `./Zero sub 10 2 // Subtraction of 2 from 10 = 8.000000.`
- `./Zero divide 10 5 // Division of 10 by 5 = 2.000000.`
- `./Zero multiply 3 3 // Multiplication of 3 and 3 = 9.000000.`

If you want to generate a new binary file, run the command: `go build` inside the repo dir.