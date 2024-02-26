package campaign

import "strings"

type CampaignFormatter struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	Name          string `json:"name"`
	ShortDesc     string `json:"short_description"`
	Slug          string `json:"slug"`
	ImageURL      string `json:"image_url"`
	GoalAmount    int    `json:"goal_amount"`
	CurrentAmount int    `json:"current_amount"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	CampaignFormatter := CampaignFormatter{}
	CampaignFormatter.ID = campaign.ID
	CampaignFormatter.UserID = campaign.UserID
	CampaignFormatter.Name = campaign.Name
	CampaignFormatter.ShortDesc = campaign.ShortDesc
	CampaignFormatter.Slug = campaign.Slug
	CampaignFormatter.GoalAmount = campaign.GoalAmount
	CampaignFormatter.CurrentAmount = campaign.CurrentAmount
	CampaignFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		CampaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return CampaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		CampaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, CampaignFormatter)
	}
	return campaignsFormatter
}

type CampaignDetailFormatter struct {
	ID            int                      `json:"id"`
	Name          string                   `json:"name"`
	ShortDesc     string                   `json:"short_description"`
	Description   string                   `json:"description"`
	ImageURL      string                   `json:"image_url"`
	GoalAmount    int                      `json:"goal_amount"`
	CurrentAmount int                      `json:"current_amount"`
	UserID        int                      `json:"user_id"`
	Slug          string                   `json:"slug"`
	Perks         []string                 `json:"perks"`
	User          CampaignUserFormatter    `json:"user"`
	Images        []CampaignImageFormatter `json:"image"`
}

type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDesc = campaign.ShortDesc
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.ImageURL = ""

	if len(campaign.CampaignImages) > 0 {
		var image_url string
		for _, image := range campaign.CampaignImages {
			if image.IsPrimary == 1 {
				image_url = image.FileName
			}
		}
		campaignDetailFormatter.ImageURL = image_url
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetailFormatter.Perks = perks

	user := campaign.User
	CampaignUserFormatter := CampaignUserFormatter{}
	CampaignUserFormatter.Name = user.Name
	CampaignUserFormatter.ImageURL = user.AvatarFileName

	campaignDetailFormatter.User = CampaignUserFormatter

	images := []CampaignImageFormatter{}

	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary

		images = append(images, campaignImageFormatter)
	}

	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}
