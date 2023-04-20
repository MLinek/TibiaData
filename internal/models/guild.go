package models

type Guild struct {
	Name             string `json:"name"`
	World            string `json:"world"`
	Logo             string `json:"logo"`
	Description      string `json:"description"`
	Active           bool   `json:"active"`
	Founded          string `json:"founded"`
	OpenApplications bool   `json:"open_applications"`
	Homepage         string `json:"homepage"`
	InWar            bool   `json:"in_war"`
	DisbandDate      string `json:"disband_date"`
	DisbandCondition string `json:"disband_condition"`
	PlayersOnline    int    `json:"players_online"`
	PlayersOffline   int    `json:"players_offline"`
	MembersTotal     int    `json:"members_total"`
	MembersInvited   int    `json:"members_invited"`
	Guildhalls       []struct {
		Name      string `json:"name"`
		World     string `json:"world"`
		PaidUntil string `json:"paid_until"`
	} `json:"guildhalls"`

	Members []Member `json:"members"`
}
