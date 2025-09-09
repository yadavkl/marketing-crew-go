package metadata

type Content struct {
	ContentType    string   `json:"content_type" jsonschema_description:"The type of content to be created (e.g., blog post, social media post, video)"`
	Topic          string   `json:"topic" jsonschema_description:"The topic of the content"`
	TargetAudience string   `json:"target_audience" jsonschema_description:"The target audience for the content"`
	Tags           []string `json:"tags" jsonschema_description:"Tags to be used for the content"`
	Text           string   `json:"content" jsonschema_description:"The content itself"`
}
