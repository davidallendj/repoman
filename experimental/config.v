
import json
import os

fn load_config(path string) Config {
	contents := os.read_file(path) or {
		eprintln('could not open file: ${err}')
		return Config{}
	}
	config := json.decode(Config, contents) or {
		eprintln('could not decode JSON: ${err}')
		return Config{}
	}
	return config
}

fn save_config(path string, config Config) {
	s := json.encode(config)

	mut file := os.create(path) or {
		eprintln('could not open file: ${err}')
		return
	}
	os.write_file(path, s) or {
		eprintln('could not write to file: ${err}')
		return
	}
	file.close()
}