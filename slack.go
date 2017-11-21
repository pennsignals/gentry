package gentry

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

type Color string

const (
	ColorDanger  Color = "danger"
	ColorGood          = "good"
	ColorWarning       = "warning"
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

type Fields []Field
