package webhook

type Auth struct {
	Addr     string
	Port     string
	User     string
	Pass     string
	ClientID string
}

type Publish struct {
	ClientID  string
	Usernme   string
	Addr      string
	Port      string
	Msg       string
	Topic     string
	TimeStamp int32
}

type Ping struct {
	ClientID  string
	Usernme   string
	Addr      string
	Port      string
	TimeStamp int32
}
