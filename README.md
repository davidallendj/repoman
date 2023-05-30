# Git Repository Manager (gitman)

Gitman is a tool created specifically for updating multiple git repositories with a single command. For example, let's say we have a collection of git repositories that need to be updated:

```bash
├── mycoolprojectfromgithub
├── gdpm
└── repo1
```
With `gitman`, we can set a reference to these repositories by using `gitman repo add` command, then provide a name and path:

```bash
gitman repo add coolprojectfromgithub path/to/mycoolprojectfromgithub
gitman repo add gdpm path/to/gdpm
gitman repo add repo1 path/to/repo1
```
Next, you can add command aliases to call using the `gitman command add` command:
```bash
gitman command add update "git pull"
gitman command add update "./update-script.sh"  # will overwrite previous add
```
Finally, run a command by calling an alias define with the `gitman repo add` command. Repositories can be specified with the `--repos` flag. Otherwise, the command will run for all repositories added.
```bash
gitman run update							# update all repos
gitman run --repos gdpm,repo1 update update	# only update gdpm and repo1
```
You can already run arbitrary commands using `exec` instead.
```bash
gitman exec "git pull" --repos coolprojectfromgithub
```
This will run `git pull` only for the repositories specified. In a future, this will be capable of running in parallel using goroutines and specifying a `--jobs` optional parameter.

When `gitman` is ran the first time, a `config.yaml` file will be created in the `$HOME/.config/gitman` directory. This file can be modified directly to add more commands and repositories.
