package environments

func AzurePublicCanary() Environment {
	// Canary is Azure Public with a different Microsoft Graph endpoint
	env := AzurePublic()
	env.Name = "Canary"
	env.MicrosoftGraph = MicrosoftGraphAPI("https://canary.graph.microsoft.com")
	return env
}
