// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestChannelBookmarkIsValid(t *testing.T) {
	testCases := []struct {
		Description     string
		Bookmark        *ChannelBookmark
		ExpectedIsValid bool
	}{
		{
			"nil bookmark",
			&ChannelBookmark{},
			false,
		},
		{
			"bookmark without create at timestamp",
			&ChannelBookmark{
				Id:          NewId(),
				OwnerId:     NewId(),
				ChannelId:   "",
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    0,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark without update at timestamp",
			&ChannelBookmark{
				Id:          NewId(),
				OwnerId:     NewId(),
				ChannelId:   "",
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    0,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with missing channel id",
			&ChannelBookmark{
				Id:          NewId(),
				OwnerId:     NewId(),
				ChannelId:   "",
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with invalid channel id",
			&ChannelBookmark{
				Id:          NewId(),
				OwnerId:     NewId(),
				ChannelId:   "invalid",
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with missing owner id",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     "",
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with invalid user id",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     "invalid",
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with missing displayname",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      "",
				DisplayName: "",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with missing type",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        "",
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark with invalid type",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        "invalid",
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark of type link with missing link url",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark of type link with invalid link url",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "invalid",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark of type link with valid link url",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "https://mattermost.com",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			true,
		},
		{
			"bookmark of type link with empty image url",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "https://mattermost.com",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			true,
		},
		{
			"bookmark of type link with invalid image url",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "https://mattermost.com",
				ImageUrl:    "invalid",
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark of type link with invalid image url",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "https://mattermost.com",
				ImageUrl:    "https://mattermost.com/some-image-without-extension", // we don't care if the URL is an actual image as the client should handle the error
				Emoji:       "",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			true,
		},
		{
			"bookmark of type file with missing file id",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      "",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkFile,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark of type file with invalid file id",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      "invalid",
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkFile,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			false,
		},
		{
			"bookmark of type file with valid file id",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      NewId(),
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkFile,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
			},
			true,
		},
		{
			"bookmark of type file with invalid original id",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      NewId(),
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkFile,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
				OriginalId:  "invalid",
			},
			false,
		},
		{
			"bookmark of type file with invalid parent id",
			&ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				FileId:      NewId(),
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "",
				ImageUrl:    "",
				Emoji:       "",
				Type:        ChannelBookmarkFile,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    4,
				ParentId:    "invalid",
			},
			false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Description, func(t *testing.T) {
			if testCase.ExpectedIsValid {
				require.Nil(t, testCase.Bookmark.IsValid())
			} else {
				require.NotNil(t, testCase.Bookmark.IsValid())
			}
		})
	}
}

func TestChannelBookmarkPreSave(t *testing.T) {
	bookmark := &ChannelBookmark{
		Id:          NewId(),
		ChannelId:   NewId(),
		OwnerId:     NewId(),
		DisplayName: "display name",
		SortOrder:   0,
		LinkUrl:     "https://mattermost.com",
		Type:        ChannelBookmarkLink,
		DeleteAt:    0,
	}

	originalBookmark := &ChannelBookmark{
		Id:          bookmark.Id,
		ChannelId:   bookmark.ChannelId,
		OwnerId:     bookmark.OwnerId,
		DisplayName: bookmark.DisplayName,
		SortOrder:   bookmark.SortOrder,
		LinkUrl:     bookmark.LinkUrl,
		Type:        bookmark.Type,
		DeleteAt:    bookmark.DeleteAt,
	}

	bookmark.PreSave()
	assert.NotEqual(t, 0, bookmark.CreateAt)
	assert.NotEqual(t, 0, bookmark.UpdateAt)

	originalBookmark.CreateAt = bookmark.CreateAt
	originalBookmark.UpdateAt = bookmark.UpdateAt
	assert.Equal(t, originalBookmark, bookmark)
}

func TestChannelBookmarkPreUpdate(t *testing.T) {
	bookmark := &ChannelBookmark{
		Id:          NewId(),
		ChannelId:   NewId(),
		OwnerId:     NewId(),
		DisplayName: "display name",
		SortOrder:   0,
		LinkUrl:     "https://mattermost.com",
		Type:        ChannelBookmarkLink,
		CreateAt:    2,
		DeleteAt:    0,
	}

	originalBookmark := &ChannelBookmark{
		Id:          bookmark.Id,
		ChannelId:   bookmark.ChannelId,
		OwnerId:     bookmark.OwnerId,
		DisplayName: bookmark.DisplayName,
		SortOrder:   bookmark.SortOrder,
		LinkUrl:     bookmark.LinkUrl,
		Type:        bookmark.Type,
		DeleteAt:    bookmark.DeleteAt,
	}

	bookmark.PreSave()
	assert.NotEqual(t, 0, bookmark.UpdateAt)

	originalBookmark.CreateAt = bookmark.CreateAt
	originalBookmark.UpdateAt = bookmark.UpdateAt
	assert.Equal(t, originalBookmark, bookmark)

	bookmark.PreUpdate()
	assert.Greater(t, bookmark.UpdateAt, originalBookmark.UpdateAt)
}

func TestChannelBookmarkEtag(t *testing.T) {
	t.Run("same etags", func(t *testing.T) {
		bookmark := &ChannelBookmark{
			Id:          NewId(),
			ChannelId:   NewId(),
			OwnerId:     NewId(),
			DisplayName: "display name",
			SortOrder:   0,
			LinkUrl:     "https://mattermost.com",
			Type:        ChannelBookmarkLink,
			CreateAt:    2,
			DeleteAt:    0,
		}
		bookamrk2 := bookmark

		assert.Equal(t, bookmark.Etag(), bookamrk2.Etag())
	})
	t.Run("different etags", func(t *testing.T) {
		t.Run("different Id", func(t *testing.T) {
			bookmark1 := &ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "https://mattermost.com",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    2,
				DeleteAt:    0,
			}
			bookmark2 := &ChannelBookmark{
				Id:          NewId(),
				ChannelId:   bookmark1.ChannelId,
				OwnerId:     bookmark1.OwnerId,
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "https://mattermost.com",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    2,
				DeleteAt:    0,
			}

			assert.NotEqual(t, bookmark1.Etag(), bookmark2.Etag())
		})
		t.Run("different update at", func(t *testing.T) {
			bookmark1 := &ChannelBookmark{
				Id:          NewId(),
				ChannelId:   NewId(),
				OwnerId:     NewId(),
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "https://mattermost.com",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    2,
				DeleteAt:    0,
			}
			bookmark2 := &ChannelBookmark{
				Id:          bookmark1.Id,
				ChannelId:   bookmark1.ChannelId,
				OwnerId:     bookmark1.OwnerId,
				DisplayName: "display name",
				SortOrder:   0,
				LinkUrl:     "https://mattermost.com",
				Type:        ChannelBookmarkLink,
				CreateAt:    2,
				UpdateAt:    3,
				DeleteAt:    0,
			}

			assert.NotEqual(t, bookmark1.Etag(), bookmark2.Etag())
		})
	})
}
