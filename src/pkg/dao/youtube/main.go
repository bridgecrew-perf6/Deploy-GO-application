package youtube

import "time"

type AccessRefreshToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
}

type AccessAfterRefreshToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type Broadcast struct {
	Id             string         `json:"id"`
	Snippet        Snippet        `json:"snippet"`
	Status         Status         `json:"status"`
	ContentDetails ContentDetails `json:"contentDetails"`
}

type Snippet struct {
	ChannelId          string      `json:"channelId"`
	ScheduledStartTime time.Time   `json:"scheduledStartTime"`
	Title              string      `json:"title"`
	Description        string      `json:"description"`
	Thumbnails         interface{} `json:"thumbnails"`
}

type Status struct {
	PrivacyStatus   string `json:"privacyStatus"`
	RecordingStatus string `json:"recordingStatus"`
	MadeForKids     bool   `json:"madeForKids"`
	LifeCycleStatus string `json:"lifeCycleStatus"`
}

type ContentDetails struct {
	MonitorStream               MonitorStream `json:"monitorStream"`
	LatencyPreference           string        `json:"latencyPreference"`
	EnableAutoStart             bool          `json:"enableAutoStart"`
	EnableAutoStop              bool          `json:"enableAutoStop"`
	BoundStreamId               string        `json:"boundStreamId"`
	BoundStreamLastUpdateTimeMs string        `json:"boundStreamLastUpdateTimeMs"`
}

type MonitorStream struct {
	EmbedHtml string `json:"embedHtml"`
}

// *=*=* STREAM KEY *=*=*

type Stream struct {
	Id      string        `json:"id"`
	Snippet StreamSnippet `json:"snippet"`
	Status  StreamStatus  `json:"status"`
	Cdn     StreamCdn     `json:"cdn"`
}

type StreamSnippet struct {
	ChannelId   string    `json:"channelId"`
	PublishedAt time.Time `json:"publishedAt"`
	Title       string    `json:"title"`
}

type StreamStatus struct {
	StreamStatus string `json:"streamStatus"`
	HealthStatus struct {
		Status string `json:"status"`
	} `json:"healthStatus"`
}

type StreamCdn struct {
	IngestionType string        `json:"ingestionType"`
	IngestionInfo IngestionInfo `json:"ingestionInfo"`
	Resolution    string        `json:"resolution"`
	FrameRate     string        `json:"frameRate"`
}

type IngestionInfo struct {
	StreamName       string `json:"streamName"`
	IngestionAddress string `json:"ingestionAddress"`
}

// *=*=* TRANSITION KEY *=*=*

type Transition struct {
	Id     string       `json:"id"`
	Status StreamStatus `json:"status"`
}
