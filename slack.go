package main

type Attachment struct {
	AuthorIcon string `json:"author_icon"`
	AuthorLink string `json:"author_link"`
	AuthorName string `json:"author_name"`
	Color      Color  `json:"color"`
	Fallback   string `json:"fallback"`
	Fields     Fields `json:"fields"`
	Footer     string `json:"footer"`
	FooterIcon string `json:"footer_icon"`
	ImageURL   string `json:"image_url"`
	Pretext    string `json:"pretext"`
	Text       string `json:"text"`
	ThumbURL   string `json:"thumb_url"`
	Title      string `json:"title"`
	TitleLink  string `json:"title_link"`
	TS         int    `json:"ts"`
}

type Attachments []*Attachment

func (a *Attachments) Add(attachments ...*Attachment) {
	*a = append(*a, attachments...)
}

type Color string

const (
	ColorDanger  Color = "danger"
	ColorGood    Color = "good"
	ColorWarning Color = "warning"
)

// String returns the literal text of the color.
func (c Color) String() string {
	return string(c)
}

type Field struct {
	Short bool   `json:"short"`
	Title string `json:"title"`
	Value string `json:"value"`
}

type Fields []*Field

type PostMessage struct {
	AsUser         bool        `json:"as_user"`
	Attachments    Attachments `json:"attachments"`
	Channel        string      `json:"channel"`
	IconEmoji      string      `json:"icon_emoji"`
	IconURL        string      `json:"icon_url"`
	LinkNames      bool        `json:"link_names"`
	Parse          string      `json:"parse"`
	ReplyBroadcast bool        `json:"reply_broadcast"`
	Text           string      `json:"text"`
	ThreadTS       float32     `json:"thread_ts"`
	Token          string      `json:"token"`
	UnfurlLinks    bool        `json:"unfurl_links"`
	UnfurlMedia    bool        `json:"unfurl_media"`
	Username       string      `json:"username"`
}
