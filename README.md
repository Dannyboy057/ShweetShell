ShweetShell
ShweetShell is a simple shell written in Go that supports basic shell commands and built-ins.

Basic shell functionalities such as executing commands and built-ins.
Support for built-in commands including cd, env, exit, ls, mkdir, and more.
Customizable with the ability to add new shell commands and built-ins.
Installation
To install ShweetShell, follow these steps:

Clone the repository:

git clone https://github.com/Dannyboy/ShweetShell.git
Accces the Shweet Shell directory:
Build the project command: 'go build'
Usage
After building the project, you can run the shell by executing the command:
'go run main.go'

Supported Commands
The following built-in commands are currently supported:

cd: Change directory.
env: List environment variables.
exit: Exit the shell.
ls: List files and directories. [optional argument 'f' or 'd' for only showing files or directories]
mkdir: Create a new directory. [argument: name of directory] ex:'mkdir test'
pwd: prints current working directory
clear: clars the console
rm: takes file or deirctory name as parameter
help: Display list of available commands.
Testing
To run the tests for this project, execute the following command:
'go test'

This will run all the test cases defined in the project.
