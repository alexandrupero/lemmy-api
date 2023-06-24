//  Source: lemmy/crates/db_schema/src/source/post.rs
// Code generated by go.elara.ws/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.

package types

type Post struct {
	ID                int              `json:"id" url:"id,omitempty"`
	Name              string           `json:"name" url:"name,omitempty"`
	URL               Optional[string] `json:"url" url:"url,omitempty"`
	Body              Optional[string] `json:"body" url:"body,omitempty"`
	CreatorID         int              `json:"creator_id" url:"creator_id,omitempty"`
	CommunityID       int              `json:"community_id" url:"community_id,omitempty"`
	Removed           bool             `json:"removed" url:"removed,omitempty"`
	Locked            bool             `json:"locked" url:"locked,omitempty"`
	Published         LemmyTime        `json:"published" url:"published,omitempty"`
	Updated           LemmyTime        `json:"updated" url:"updated,omitempty"`
	Deleted           bool             `json:"deleted" url:"deleted,omitempty"`
	NSFW              bool             `json:"nsfw" url:"nsfw,omitempty"`
	EmbedTitle        Optional[string] `json:"embed_title" url:"embed_title,omitempty"`
	EmbedDescription  Optional[string] `json:"embed_description" url:"embed_description,omitempty"`
	ThumbnailURL      Optional[string] `json:"thumbnail_url" url:"thumbnail_url,omitempty"`
	ApID              string           `json:"ap_id" url:"ap_id,omitempty"`
	Local             bool             `json:"local" url:"local,omitempty"`
	EmbedVideoURL     Optional[string] `json:"embed_video_url" url:"embed_video_url,omitempty"`
	LanguageID        int              `json:"language_id" url:"language_id,omitempty"`
	FeaturedCommunity bool             `json:"featured_community" url:"featured_community,omitempty"`
	FeaturedLocal     bool             `json:"featured_local" url:"featured_local,omitempty"`
}
type PostInsertForm struct {
	Name              string           `json:"name" url:"name,omitempty"`
	CreatorID         int              `json:"creator_id" url:"creator_id,omitempty"`
	CommunityID       int              `json:"community_id" url:"community_id,omitempty"`
	NSFW              Optional[bool]   `json:"nsfw" url:"nsfw,omitempty"`
	URL               Optional[string] `json:"url" url:"url,omitempty"`
	Body              Optional[string] `json:"body" url:"body,omitempty"`
	Removed           Optional[bool]   `json:"removed" url:"removed,omitempty"`
	Locked            Optional[bool]   `json:"locked" url:"locked,omitempty"`
	Updated           LemmyTime        `json:"updated" url:"updated,omitempty"`
	Published         LemmyTime        `json:"published" url:"published,omitempty"`
	Deleted           Optional[bool]   `json:"deleted" url:"deleted,omitempty"`
	EmbedTitle        Optional[string] `json:"embed_title" url:"embed_title,omitempty"`
	EmbedDescription  Optional[string] `json:"embed_description" url:"embed_description,omitempty"`
	EmbedVideoURL     Optional[string] `json:"embed_video_url" url:"embed_video_url,omitempty"`
	ThumbnailURL      Optional[string] `json:"thumbnail_url" url:"thumbnail_url,omitempty"`
	ApID              Optional[string] `json:"ap_id" url:"ap_id,omitempty"`
	Local             Optional[bool]   `json:"local" url:"local,omitempty"`
	LanguageID        Optional[int]    `json:"language_id" url:"language_id,omitempty"`
	FeaturedCommunity Optional[bool]   `json:"featured_community" url:"featured_community,omitempty"`
	FeaturedLocal     Optional[bool]   `json:"featured_local" url:"featured_local,omitempty"`
}
type PostUpdateForm struct {
	Name              Optional[string]           `json:"name" url:"name,omitempty"`
	NSFW              Optional[bool]             `json:"nsfw" url:"nsfw,omitempty"`
	URL               Optional[Optional[string]] `json:"url" url:"url,omitempty"`
	Body              Optional[Optional[string]] `json:"body" url:"body,omitempty"`
	Removed           Optional[bool]             `json:"removed" url:"removed,omitempty"`
	Locked            Optional[bool]             `json:"locked" url:"locked,omitempty"`
	Published         LemmyTime                  `json:"published" url:"published,omitempty"`
	Updated           LemmyTime                  `json:"updated" url:"updated,omitempty"`
	Deleted           Optional[bool]             `json:"deleted" url:"deleted,omitempty"`
	EmbedTitle        Optional[Optional[string]] `json:"embed_title" url:"embed_title,omitempty"`
	EmbedDescription  Optional[Optional[string]] `json:"embed_description" url:"embed_description,omitempty"`
	EmbedVideoURL     Optional[Optional[string]] `json:"embed_video_url" url:"embed_video_url,omitempty"`
	ThumbnailURL      Optional[Optional[string]] `json:"thumbnail_url" url:"thumbnail_url,omitempty"`
	ApID              Optional[string]           `json:"ap_id" url:"ap_id,omitempty"`
	Local             Optional[bool]             `json:"local" url:"local,omitempty"`
	LanguageID        Optional[int]              `json:"language_id" url:"language_id,omitempty"`
	FeaturedCommunity Optional[bool]             `json:"featured_community" url:"featured_community,omitempty"`
	FeaturedLocal     Optional[bool]             `json:"featured_local" url:"featured_local,omitempty"`
}
type PostLike struct {
	ID        int32     `json:"id" url:"id,omitempty"`
	PostID    int       `json:"post_id" url:"post_id,omitempty"`
	PersonID  int       `json:"person_id" url:"person_id,omitempty"`
	Score     int16     `json:"score" url:"score,omitempty"`
	Published LemmyTime `json:"published" url:"published,omitempty"`
}
type PostLikeForm struct {
	PostID   int   `json:"post_id" url:"post_id,omitempty"`
	PersonID int   `json:"person_id" url:"person_id,omitempty"`
	Score    int16 `json:"score" url:"score,omitempty"`
}
type PostSaved struct {
	ID        int32     `json:"id" url:"id,omitempty"`
	PostID    int       `json:"post_id" url:"post_id,omitempty"`
	PersonID  int       `json:"person_id" url:"person_id,omitempty"`
	Published LemmyTime `json:"published" url:"published,omitempty"`
}
type PostSavedForm struct {
	PostID   int `json:"post_id" url:"post_id,omitempty"`
	PersonID int `json:"person_id" url:"person_id,omitempty"`
}
type PostRead struct {
	ID        int32     `json:"id" url:"id,omitempty"`
	PostID    int       `json:"post_id" url:"post_id,omitempty"`
	PersonID  int       `json:"person_id" url:"person_id,omitempty"`
	Published LemmyTime `json:"published" url:"published,omitempty"`
}
type PostReadForm struct {
	PostID   int `json:"post_id" url:"post_id,omitempty"`
	PersonID int `json:"person_id" url:"person_id,omitempty"`
}
