package swarm

func generateCompose() string {
	result := "version: '3.8'\n"
	result += "network\n:
			\b\btest\n"
	result += "services:\n"
	
	return result
}
