package lmjrfll

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/datastore"
	"time"
)

// Expo is the day the lego expo will be held, where the teams present their project
type Expo struct {
	Year             int
	Name             string
	ExpoDate         time.Time
	ExpoEndTime      time.Time
	IsActive         bool
	IsCurrent        bool
	NumTeams         int
	NumRegistered    int
	RegistrationCost int // in dollars
	LocationName     string
}

// Team is the team name for a team of people
type Team struct {
	Name                 string
	CreateDate           time.Time
	ExpoKey              string
	JrFLLTeamNumber      string
	IsLookingForMemebers bool
	IsPaid               bool
	LeaderEmail          string
	NumMembers           int
	IsActive             bool
	IsWaitList           bool
	InviteCode           string
}

type TeamMember struct {
	TeamName  string
	TeamKey   string
	FirstName string
	LastName  string
	Email     string
}

// Returns the current Expo, or an error that there is no current Expo
func GetCurrentExpo(c context.Context) (*Expo, error) {
	var expo Expo
	query := datastore.NewQuery("Expo").Filter("IsCurrent =", true)
	results := query.Run(c)
	_, err := results.Next(&expo)
	if err != nil {
		return nil, err
	}
	return &expo, nil
}
