package main

import (
	"errors"
	"fmt"
)

type Profile struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

type IProfileService interface {
	Validate(c *Profile) bool
	CreateProfile(c *Profile) error
	GetProfile(name string) (*Profile, error)
}

type profileService struct {
	defaultAge int
	storage    map[string]*Profile
}

func NewProfileService(defaultAge int) IProfileService {
	return &profileService{
		defaultAge: defaultAge,
		storage:    make(map[string]*Profile),
	}
}

func main() {
	service := NewProfileService(18)
	profile := &Profile{Firstname: "Thamrong", Age: 20}
	if service.Validate(profile) {
		err := service.CreateProfile(profile)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	profileData, err := service.GetProfile("Thamrong")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Retrieved Profile:", profileData)
	}

}

func (p *profileService) Validate(c *Profile) bool {
	return c.Age >= p.defaultAge
}

func (p *profileService) CreateProfile(c *Profile) error {
	if _, exists := p.storage[c.Firstname]; exists {
		return errors.New("profile already exists")
	}
	p.storage[c.Firstname] = c
	fmt.Println("Profile created:", c.Firstname)
	return nil
}

func (p *profileService) GetProfile(name string) (*Profile, error) {
	profile, exists := p.storage[name]
	if !exists {
		return nil, errors.New("profile not found")
	}
	return profile, nil
}
