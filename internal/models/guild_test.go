package models

import "testing"

func TestGuild_OnlinePlayers(t *testing.T) {
	guild := Guild{
		Name: "Example guild",
		Members: []Member{
			{
				Name:   "Offline1",
				Status: "offline",
			},
			{
				Name:   "Online1",
				Status: "online",
			},
			{
				Name:   "Online2",
				Status: "online",
			},
			{
				Name:   "Offline2",
				Status: "offline",
			},
		},
	}
	online, _ := guild.OnlinePlayers()
	if len(online) != 2 {
		t.Fatalf(`len(OnlinePlayers()) = %d, want 2`, len(online))
	}

	if online[0].Name != "Online1" {
		t.Fatalf(`online[0].Name = %s, want "Online1""`, online[0].Name)
	}

	if online[1].Name != "Online2" {
		t.Fatalf(`online[1].Name = %s, want "Online2""`, online[1].Name)
	}
}
