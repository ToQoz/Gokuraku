package soundcloud

import (
	"encoding/json"
	"github.com/ToQoz/Gokuraku/gokuraku/test_helpers"
	"testing"
)

var itemJson = []byte(`

{
  "attachments_uri": "http://api.soundcloud.com/tracks/93919708/attachments",
  "comment_count": 126,
  "favoritings_count": 1160,
  "download_count": 0,
  "playback_count": 42190,
  "stream_url": "http://api.soundcloud.com/tracks/93919708/stream",
  "waveform_url": "http://w1.sndcdn.com/mMv80QIobRnu_m.png",
  "artwork_url": "http://i1.sndcdn.com/artworks-000048980129-sa1mue-large.jpg?5ffe3cd",
  "permalink_url": "http://soundcloud.com/knxwledge/ignorntsht-twrk",
  "user": {
    "avatar_url": "http://i1.sndcdn.com/avatars-000043150214-63rk2k-large.jpg?5ffe3cd",
    "permalink_url": "http://soundcloud.com/knxwledge",
    "uri": "http://api.soundcloud.com/users/2158",
    "username": "Knx.",
    "permalink": "knxwledge",
    "kind": "user",
    "id": 2158
  },
  "uri": "http://api.soundcloud.com/tracks/93919708",
  "label_id": null,
  "purchase_url": "http://gloof.bandcamp.com/track/5-ignorntsht-twrk",
  "downloadable": false,
  "embeddable_by": "all",
  "streamable": true,
  "permalink": "ignorntsht-twrk",
  "tag_list": "knxwledge loops twrklfe",
  "sharing": "public",
  "kind": "track",
  "id": 93919708,
  "created_at": "2013/05/26 03:59:27 +0000",
  "user_id": 2158,
  "duration": 240112,
  "commentable": true,
  "state": "finished",
  "original_content_size": 11392752,
  "purchase_title": null,
  "genre": "Hop.Hip",
  "title": "ignorntsht[TWRK]",
  "description": "http://gloof.bandcamp.com/album/wraptaypes-prt-5",
  "label_name": "",
  "release": "",
  "track_type": "",
  "key_signature": "",
  "isrc": "",
  "video_url": null,
  "bpm": null,
  "release_year": null,
  "release_month": null,
  "release_day": null,
  "original_format": "mp3",
  "license": "all-rights-reserved"
}
`)

func TestMappingJsonToItem(t *testing.T) {
	var item Item
	err := json.Unmarshal(itemJson, &item)

	if err != nil {
		panic(err)
	}

	test_helpers.AssertEqual(t, 93919708, item.Id)
	test_helpers.AssertEqual(t, "Hop.Hip", item.Genre)
	test_helpers.AssertEqual(t, "ignorntsht[TWRK]", item.Title)
	test_helpers.AssertEqual(t, "http://soundcloud.com/knxwledge/ignorntsht-twrk", item.Url)
	test_helpers.AssertEqual(t, "http://gloof.bandcamp.com/album/wraptaypes-prt-5", item.Description)
	test_helpers.AssertEqual(t, "http://i1.sndcdn.com/artworks-000048980129-sa1mue-large.jpg?5ffe3cd", item.ImageUrl)
	test_helpers.AssertEqual(t, "http://i1.sndcdn.com/avatars-000043150214-63rk2k-large.jpg?5ffe3cd", item.User.ImageUrl)
	test_helpers.AssertEqual(t, "http://soundcloud.com/knxwledge", item.User.Url)
}
