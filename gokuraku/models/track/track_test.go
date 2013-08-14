package track

import (
	"github.com/ToQoz/Gokuraku/gokuraku/test_helpers"
	"strconv"
	"testing"
	"time"
)

func TestValidateId(t *testing.T) {
	var err error
	now := strconv.FormatInt(time.Now().Unix(), 10)

	track := Track{
		Id:           "",
		Genre:        "Genre",
		Title:        "Title",
		Description:  "Description",
		Url:          "Url",
		ImageUrl:     "ImageUrl",
		UserImageUrl: "UserImageUrl",
		UserUrl:      "UserUrl",
		UserName:     "UserName",
		Type:         "soundcloud",
		CreatedAt:    now,
	}

	err = track.Validate()
	if err == nil {
		t.Error("Track#Validate() should return error if Track#Title is empty")
		return
	}

	test_helpers.AssertEqual(t, "This item doesn't have ID", err.Error())
}

func TestValidateTitle(t *testing.T) {
	var err error
	now := strconv.FormatInt(time.Now().Unix(), 10)

	track := Track{
		Id:           "123",
		Genre:        "Genre",
		Title:        "",
		Description:  "Description",
		Url:          "Url",
		ImageUrl:     "ImageUrl",
		UserImageUrl: "UserImageUrl",
		UserUrl:      "UserUrl",
		UserName:     "UserName",
		Type:         "soundcloud",
		CreatedAt:    now,
	}

	err = track.Validate()
	if err == nil {
		t.Error("Track#Validate() should return error if Track#Title is empty")
		return
	}

	test_helpers.AssertEqual(t, "This item doesn't have Title", err.Error())
}

func TestValidateUrl(t *testing.T) {
	var err error
	now := strconv.FormatInt(time.Now().Unix(), 10)

	track := Track{
		Id:           "123",
		Genre:        "Genre",
		Title:        "Title",
		Description:  "Description",
		Url:          "",
		ImageUrl:     "ImageUrl",
		UserImageUrl: "UserImageUrl",
		UserUrl:      "UserUrl",
		UserName:     "UserName",
		Type:         "soundcloud",
		CreatedAt:    now,
	}

	err = track.Validate()
	if err == nil {
		t.Error("Track#Validate() should return error if Track#Url is empty")
		return
	}

	test_helpers.AssertEqual(t, "This item doesn't have Url", err.Error())
}

func TestValidateUserName(t *testing.T) {
	var err error
	now := strconv.FormatInt(time.Now().Unix(), 10)

	track := Track{
		Id:           "123",
		Genre:        "Genre",
		Title:        "Title",
		Description:  "Description",
		Url:          "Url",
		ImageUrl:     "ImageUrl",
		UserImageUrl: "UserImageUrl",
		UserUrl:      "UserUrl",
		UserName:     "",
		Type:         "soundcloud",
		CreatedAt:    now,
	}

	err = track.Validate()
	if err == nil {
		t.Error("Track#Validate() should return error if Track#UserName is empty")
		return
	}

	test_helpers.AssertEqual(t, "This item doesn't have UserName", err.Error())
}

func TestValidateUserUrl(t *testing.T) {
	var err error
	now := strconv.FormatInt(time.Now().Unix(), 10)

	track := Track{
		Id:           "123",
		Genre:        "Genre",
		Title:        "Title",
		Description:  "Description",
		Url:          "Url",
		ImageUrl:     "ImageUrl",
		UserImageUrl: "UserImageUrl",
		UserUrl:      "",
		UserName:     "UserName",
		Type:         "soundcloud",
		CreatedAt:    now,
	}

	err = track.Validate()
	if err == nil {
		t.Error("Track#Validate() should return error if Track#UserUrl is empty")
		return
	}

	test_helpers.AssertEqual(t, "This item doesn't have UserUrl", err.Error())
}
