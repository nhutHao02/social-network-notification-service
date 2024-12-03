package constants

type ActionTweet string

const (
	Love     ActionTweet = "Love"
	Bookmark ActionTweet = "Bookmark"
	Repost   ActionTweet = "Repost"
	Comment  ActionTweet = "Comment"
)

func (a ActionTweet) IsValid() bool {
	switch a {
	case Love, Bookmark, Repost, Comment:
		return true
	default:
		return false
	}
}

func (a ActionTweet) Message() string {
	switch a {
	case Love:
		return "Loved your tweet"
	case Bookmark:
		return "Bookmarked your tweet"
	case Repost:
		return "Reposted your tweet"
	case Comment:
		return "Commented your tweet"
	default:
		return ""
	}
}
