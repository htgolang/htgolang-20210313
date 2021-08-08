package forms

type AgentdHeartbeatForm struct {
	UUID       string `json:"uuid"`
	ClientTime int64  `json:"now"`
}

type AgentdRegisterForm struct {
	UUID     string `json:"uuid"`
	Addr     string `json:"addr"`
	Hostname string `json:"hostname"`
	Now      int64  `json:"now"`
}

type AgentdConfigForm struct {
	UUID    string `form:"uuid"`
	Version int64  `form:"version"`
}
