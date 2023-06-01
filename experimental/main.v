module main

import flag
import cli
import os

struct Config {
mut:
	repositories map[string]string
	commands     map[string][]string
	path         string
}

__global (
	config       Config
	repositories []string
	show_window  bool
)

fn main() {
	config = load_config('./config.json')
	mut app := cli.Command{
		name: 'gitman'
		description: 'Manage a collection of git repositories efficiently'
		execute: fn (cmd cli.Command) ! {
			cmd.execute_help()
		}
		commands: [
			cli.Command{
				name: 'exec'
				description: 'Execute a command or script'
				execute: fn (cmd cli.Command) ! {
					// Execute any arbitrary command on a group
					if cmd.args.len == 0 {
						return
					}

					mut repos := config.repositories.keys()
					if repositories.len > 0 {
						repos = repositories.clone() 
					} 

					// Execute single command directly on provided groups
					for name in repos {
						path := config.repositories[name] 
						if path == "" {
							continue
						}

						os.chdir(path) or { eprintln('could not chdir: ${err}') }
						result := os.execute(cmd.args[1..].join(' '))
						if result.exit_code != 0 {
							eprintln('could not execute process: ${result.exit_code}')
						}
						println("${name}: ${result.output}")
					}
				}
			},
			cli.Command{
				name: 'run'
				description: 'Run a command alias'
				execute: fn (cmd cli.Command) ! {
					if cmd.args.len == 0 {
						eprintln('could not run command alias')
						return
					}

					// Run a command alias on a collection of repos
					mut repos := config.repositories.keys()
					if repositories.len > 0 {
						repos = repositories.clone()
					} 
					for command in cmd.args {
						for name in repos {

							path := config.repositories[name]
							if path == "" {
								continue
							}

							run := config.commands[command]
							os.chdir(path) or {
								eprintln('could not chdir: ${err}')
								return
							}
							result := os.execute(run.join(' '))
							if result.exit_code != 0 {
								eprintln('could not execute process: ${result.exit_code}')
							}
							println("${name}: ${result.output}")
						}
					}
				}
			},
			cli.Command{
				name: 'repo'
				description: 'Manage repositories'
				execute: fn (cmd cli.Command) ! {
					if cmd.args.len == 0 {
						for k, v in config.repositories {
							print('${k}: ${v}\n')
						}
						return
					}

					sub := cmd.args[0]
					name := cmd.args[1]
					if cmd.args.len >= 2 {
						if sub == 'remove' {
							for i in 1..cmd.args.len {
								config.repositories.delete(cmd.args[i])
								println('${cmd.args[i]}')
							}
							return
						}
					}

					if cmd.args.len < 3 {
						eprintln('could not add repository')
						return
					}

					path := cmd.args[2]
					if sub == 'add' {
						config.repositories[name] = path
					}

					println('${cmd.args[1]}: ${cmd.args[2]}')

					// Add a new repo and update config
					// mut repos map[string]string =
				}
			},
			cli.Command{
				name: 'command'
				description: 'Manage command aliases'
				execute: fn (cmd cli.Command) ! {
					if cmd.args.len == 0 {
						for k, v in config.commands {
							print('${k}: ${v}\n')
						}
						return
					}

					sub := cmd.args[0]
					alias := cmd.args[1]
					if cmd.args.len == 2 {
						if sub == 'remove' {
							config.commands.delete(alias)
							println('${alias}')
							return
						}
					}

					if cmd.args.len < 3 {
						eprintln('could not add command')
						return
					}

					args := cmd.args[2..]
					if sub == 'add' {
						config.commands[alias] = args
					}

					println('${cmd.args[1]}: ${cmd.args[2]}')
				}
			},
		]
	}
	mut fp := flag.new_flag_parser(os.args)
	fp.application('gitman')
	fp.version('v0.0.1')
	fp.description('An experimental vlang version of `gitman`')
	repositories = fp.string_multi("repos", "r".u8(), "Set repositories to use")
	fp.skip_executable()

	config.path = fp.string('config-path', 0, '.', '')

	_ := fp.finalize() or {
		eprintln(err)
		println(fp.usage())
		return
	}

	app.setup()
	app.parse(os.args)

	save_config('./config.json', config)
}


