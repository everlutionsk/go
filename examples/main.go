package main

import (
	"fmt"
	"github.com/everlutionsk/go/qzila/sdk/citadel"
	"log"
	"os"
	"time"
)

const (
	CitadelBaseUrl   = "CITADEL_BASE_URL"
	CitadelApiKey    = "CITADEL_API_KEY"
	CitadelSharedKey = "CITADEL_SHARED_KEY"
)

var (
	client = citadel.NewClient(&citadel.ClientConfig{
		BaseUrl:      os.Getenv(CitadelBaseUrl),
		ApiKey:       os.Getenv(CitadelApiKey),
		PreSharedKey: os.Getenv(CitadelSharedKey),
	})
)

func main() {
	// Invite User
	createdUser, err := client.InviteUser(&citadel.InviteUserRequest{
		Username:         "USERNAME",
		EmailAddress:     "MUST BE AN EMAIL ADDRESS",
		AllowedAuthFlows: []string{citadel.AuthFlowEmailCode, citadel.AuthFlowPassword},
		// Where to redirect after successful login
		RedirectUri:         "https://REDIRECT_URL/",
		ExpirationInSeconds: int(time.Minute * 15),
		Language:            "en",
	})

	if err != nil {
		log.Printf("Failed to create user: %v\n", err)
		return
	}
	log.Printf("Created user with id %s\n", createdUser.UserId)

	// if you wish you can set user metadata whilst onboarding him
	setUserMetadataResponse, err := client.SetUserMetadata(&citadel.SetUserMetadataRequest{
		UserId: createdUser.UserId,
		Metadata: []citadel.MetadataItem{
			{
				Key:   "profile:email",
				Value: "MUST BE AN EMAIL ADDRESS",
			},
			{
				Key:   "profile:fullName",
				Value: "Your Name",
			},
		},
	})
	if err != nil {
		log.Printf("Failed to set user metadata: %v\n", err)
		return
	}
	fmt.Printf("Metadata set %v\n", setUserMetadataResponse)
}
