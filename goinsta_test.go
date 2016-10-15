package goinsta

import (
	"encoding/json"
	"os"
	"testing"
)

var (
	username string
	password string
	insta    *Instagram
	skip     bool
)

func TestHandlesNonExistingItems(t *testing.T) {
	username = os.Getenv("INSTA_USERNAME")
	password = os.Getenv("INSTA_PASSWORD")
	if len(username)*len(password) == 0 {
		skip = true
		t.Fatal("Username or Password is empty")
	}
}

func TestDeviceID(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	insta = New(username, password)
	t.Log(insta.Informations.DeviceID)
}

func TestLogin(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	err := insta.Login()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log("Logged in user", insta.LoggedInUser.FullName, insta.Informations.UsernameId)
}

func TestUserFollowings(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	bytes, err := insta.UserFollowings(insta.Informations.UsernameId, "")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(string(bytes)[:15])
}

func TestUserFollowers(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	bytes, err := insta.UserFollowers(insta.Informations.UsernameId, "")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(string(bytes)[:15])
}

func TestSelfUserFeed(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	bytes, err := insta.UserFeed()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(string(bytes))
}

func TestMediaLikers(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}
	bytes, err := insta.UserFeed()
	if err != nil {
		t.Fatal(err)
		return
	}

	type Item struct {
		Id string `json:"id"`
	}

	var Result struct {
		Status string `json:"status"`
		Items  []Item `json:"items"`
	}

	err = json.Unmarshal(bytes, &Result)
	if err != nil {
		t.Fatal(err)
		return
	}

	if len(Result.Items) > 0 {
		bytes, err = insta.MediaLikers(Result.Items[0].Id)
		if err != nil {
			t.Fatal(err)
			return
		}
		t.Log(string(bytes)[:30])
	} else {
		t.Skip("Empty feed")
	}
}

func TestFollow(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	bytes, err := insta.Follow("1572292791") // ahmdrz (creator) instagram usernameID
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(string(bytes))
}

func TestLike(t *testing.T) {
	if skip {
		t.Skip("Empty username or password , Skipping ...")
	}

	bytes, err := insta.Like("1336846574982263293") // one of ahmdrz images
	if err != nil {
		t.Fatal(err)
		return
	}

	t.Log(string(bytes))
}
