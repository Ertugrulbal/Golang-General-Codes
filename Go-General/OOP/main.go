package main

import (
	"fmt"
	"time"

	"yemeksepeti.com/ecommerce/dev-team/models"
	"yemeksepeti.com/ecommerce/dev-team/utils"
)

// Interface, Composition, Method Overriding

func main() {
	// Gamze, Batuhan, Ayberk, Busra
	gamze := models.Developer{
		Employee: models.Employee{
			FirstName: "Gamze",
			LastName:  "Karadayı",
			Dob:       time.Date(2000, time.February, 15, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Senior Go Developer",
			Location:  "İzmir",
		},
		Skills: []string{"Go", "Docker", "Kubernetes", "C#"},
	}
	ellerhavada_batuhan := models.Developer{
		Employee: models.Employee{
			FirstName: "Batuhan",
			LastName:  "D.",
			Dob:       time.Date(1999, time.February, 18, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Junior Go Developer",
			Location:  "İzmir",
		},
		Skills: []string{"Go", "MySQL", "PHP", "ANV"},
	}
	ayberk := models.Developer{
		Employee: models.Employee{
			FirstName: "Ay",
			LastName:  "Bi",
			Dob:       time.Date(1998, time.July, 14, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Go Rapper",
			Location:  "İstanbul",
		},
		Skills: []string{"Go", "Python", "PostgreSQL", "Rust"},
	}
	mslogic := models.Manager{
		Employee: models.Employee{
			FirstName: "Bayan",
			LastName:  "Mantik",
			Dob:       time.Date(1987, time.March, 13, 0, 0, 0, 0, time.UTC),
			JobTitle:  "Team Manager",
			Location:  "Yozgat",
		},
		Projects:  []string{"Face Detection", "e-Commerce"},
		Locations: []string{"San Francisco", "Yozgat"},
	}

	team := models.Team{
		"Go",
		"Golang Engineering Team",
		[]models.TeamMember{gamze, ayberk, ellerhavada_batuhan, mslogic},
	}
	team.PrintTeamDetails()

	xx := 90
	fmt.Println(utils.Itoa(xx))
}
