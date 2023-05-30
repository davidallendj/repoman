package defaults

func MakeDefaultCommands() map[string]string {
	return map[string]string{
		"update": "git pull",
	}
}
