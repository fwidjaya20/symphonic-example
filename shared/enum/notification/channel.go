package notification

import "encoding/json"

type Channel int

const (
	FirebaseCloudMessaging Channel = iota
)

func ToChannelEnum(s string) Channel {
	var t Channel
	switch s {
	case "firebase_cloud_messaging":
		t = FirebaseCloudMessaging
	default:
		t = FirebaseCloudMessaging
	}
	return t
}

func (c *Channel) ToString() string {
	switch *c {
	case FirebaseCloudMessaging:
		return "firebase_cloud_messaging"
	default:
		return ""
	}
}

func (c *Channel) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.ToString())
}

func (c *Channel) UnmarshalJSON(b []byte) error {
	var str string

	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}

	*c = ToChannelEnum(str)
	return nil
}
