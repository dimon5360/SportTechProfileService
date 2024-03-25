package storage

import (
	"app/main/internal/models"
	"fmt"
	"log"
)

func (s *ProfileUsersService) GetProfileWithFormat(sql string) (*models.Profile, error) {

	rows, err := s.db.Query(sql)
	if err != nil {
		log.Print(err)
		return &models.Profile{}, err
	}

	var profile models.Profile

	for rows.Next() {
		err := rows.Scan(&profile.Id, &profile.Username, &profile.Firstname,
			&profile.Lastname, &profile.UserId, &profile.Created_at, &profile.Updated_at)
		if err != nil {
			log.Print(err)
			return &models.Profile{}, err
		}
	}

	if profile.Id == 0 {
		return &models.Profile{}, fmt.Errorf("%s", "Profile not found")
	}

	if err := rows.Err(); err != nil {
		log.Print(err)
		return &models.Profile{}, err
	}
	return &profile, nil
}

func (s *ProfileUsersService) GetProfileByIdFromDatabase(uuid uint64) (*models.Profile, error) {
	return s.GetProfileWithFormat(fmt.Sprintf("select * from profiles where user_id = %d;", uuid))
}

func (s *ProfileUsersService) AddProfileToDatabase(username string, firstname string, lastname string, userId uint64) (*models.Profile, error) {

	format := "insert into profiles (username, firstname, lastname, user_id) values ('%s', '%s', '%s', '%d') ON conflict DO NOTHING RETURNING id;"

	var profile models.Profile
	rows := s.db.QueryRow(fmt.Sprintf(format, username, firstname, lastname, userId))
	if err := rows.Err(); err != nil {
		log.Print("Failed inserting new profile")
		return &profile, err
	}

	err := rows.Scan(&profile.Id)
	if err != nil {
		log.Println("Failed scaning new profile")
		return &profile, err
	}

	if profile.Id == 0 {
		return &models.Profile{}, fmt.Errorf("%s", "Profile wasn't created")
	}

	return &profile, nil
}
