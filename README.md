# Git Repository Manager (gitman)

Having trouble updating multiple git repositories? Then try using `gitman`! This tool was created specifically for calling git commands for multiple repos. For example, let's say we have a collection of repositories that need to be updated:

```bash
├── coolprojectfromgithub
├── gdpm
└── myrepo1
```
With `gitman`, set a reference to these repositories by using `gitman repo add` command, then provide a group name and collection of paths:

```bash
gitman repo add example ./coolprojectfromgithub ./gdpm ./myrepo1
```
Next, you can add command aliases to call using the `gitman command add` command:
```bash
gitman command add update "git pull"
gitman command add update "./update-script.sh"  # will overwrite previous
```
Specify a repo group and run the command or call exec to run any command:
```bash
gitman run update
gitman exec "git pull"
```
This will run git pull for all the repositories in the "example" repo group. In a future, this will be capable of running in parallel using go-routines and specifying a `--jobs` optional parameter.