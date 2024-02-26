package campaign

import "time"

type Campaign struct {
	ID             int
	UserID         int
	Name           string
	ShortDesc      string
	Description    string
	Perks          string
	BackerCount    int
	GoalAmount     int
	CurrentAmount  int
	Slug           string
	CreatedAt      time.Time
	UpdateddAt     time.Time
	CampaignImages []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdateddAt time.Time
}
