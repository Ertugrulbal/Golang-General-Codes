package models

import "fmt"

// Interface

type TeamMember interface {
	PrintName()
	PrintDetails()
}

// Struct -> Team

type Team struct {
	Name, Description string
	TeamMembers       []TeamMember
}

func (t Team) PrintTeamDetails() {
	fmt.Printf("Team: %s - %s\n", t.Name, t.Description)
	fmt.Println("Details of the team members: ")
	for _, v := range t.TeamMembers {
		v.PrintName()
		v.PrintDetails()
	}
}
