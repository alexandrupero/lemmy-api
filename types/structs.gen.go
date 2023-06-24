//  Source: lemmy/crates/db_views_actor/src/structs.rs
// Code generated by go.elara.ws/go-lemmy/cmd/gen (struct generator). DO NOT EDIT.

package types

type CommunityBlockView struct {
	Person    Person    `json:"person" url:"person,omitempty"`
	Community Community `json:"community" url:"community,omitempty"`
}
type CommunityFollowerView struct {
	Community Community `json:"community" url:"community,omitempty"`
	Follower  Person    `json:"follower" url:"follower,omitempty"`
}
type CommunityModeratorView struct {
	Community Community `json:"community" url:"community,omitempty"`
	Moderator Person    `json:"moderator" url:"moderator,omitempty"`
}
type CommunityPersonBanView struct {
	Community Community `json:"community" url:"community,omitempty"`
	Person    Person    `json:"person" url:"person,omitempty"`
}
type CommunityView struct {
	Community  Community           `json:"community" url:"community,omitempty"`
	Subscribed SubscribedType      `json:"subscribed" url:"subscribed,omitempty"`
	Blocked    bool                `json:"blocked" url:"blocked,omitempty"`
	Counts     CommunityAggregates `json:"counts" url:"counts,omitempty"`
}
type PersonBlockView struct {
	Person Person `json:"person" url:"person,omitempty"`
	Target Person `json:"target" url:"target,omitempty"`
}
type PersonMentionView struct {
	PersonMention              PersonMention     `json:"person_mention" url:"person_mention,omitempty"`
	Comment                    Comment           `json:"comment" url:"comment,omitempty"`
	Creator                    Person            `json:"creator" url:"creator,omitempty"`
	Post                       Post              `json:"post" url:"post,omitempty"`
	Community                  Community         `json:"community" url:"community,omitempty"`
	Recipient                  Person            `json:"recipient" url:"recipient,omitempty"`
	Counts                     CommentAggregates `json:"counts" url:"counts,omitempty"`
	CreatorBannedFromCommunity bool              `json:"creator_banned_from_community" url:"creator_banned_from_community,omitempty"`
	Subscribed                 SubscribedType    `json:"subscribed" url:"subscribed,omitempty"`
	Saved                      bool              `json:"saved" url:"saved,omitempty"`
	CreatorBlocked             bool              `json:"creator_blocked" url:"creator_blocked,omitempty"`
	MyVote                     Optional[int16]   `json:"my_vote" url:"my_vote,omitempty"`
}
type CommentReplyView struct {
	CommentReply               CommentReply      `json:"comment_reply" url:"comment_reply,omitempty"`
	Comment                    Comment           `json:"comment" url:"comment,omitempty"`
	Creator                    Person            `json:"creator" url:"creator,omitempty"`
	Post                       Post              `json:"post" url:"post,omitempty"`
	Community                  Community         `json:"community" url:"community,omitempty"`
	Recipient                  Person            `json:"recipient" url:"recipient,omitempty"`
	Counts                     CommentAggregates `json:"counts" url:"counts,omitempty"`
	CreatorBannedFromCommunity bool              `json:"creator_banned_from_community" url:"creator_banned_from_community,omitempty"`
	Subscribed                 SubscribedType    `json:"subscribed" url:"subscribed,omitempty"`
	Saved                      bool              `json:"saved" url:"saved,omitempty"`
	CreatorBlocked             bool              `json:"creator_blocked" url:"creator_blocked,omitempty"`
	MyVote                     Optional[int16]   `json:"my_vote" url:"my_vote,omitempty"`
}
type PersonView struct {
	Person Person           `json:"person" url:"person,omitempty"`
	Counts PersonAggregates `json:"counts" url:"counts,omitempty"`
}
