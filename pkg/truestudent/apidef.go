package truestudent

type userInfo struct {
	LoginURL              string `json:"loginURL"`
	MaxSessionTime        int64  `json:"maxSessionTime"`
	RemainingSessionTime  int    `json:"remainingSessionTime"`
	ElapsedSessionTime    int    `json:"elapsedSessionTime"`
	MaxOutputOctets       string `json:"maxOutputOctets"`
	RemainingOutputOctets string `json:"remainingOutputOctets"`
	UsedOutputOctets      string `json:"usedOutputOctets"`
	MaxInputOctets        string `json:"maxInputOctets"`
	RemainingInputOctets  string `json:"remainingInputOctets"`
	UsedInputOctets       string `json:"usedInputOctets"`
	IsLoggedIn            string `json:"isLoggedIn"`
	Nasid                 string `json:"nasid"`
	Vlan                  string `json:"vlan"`
	CustMac               string `json:"CustMac"`
	CustIP                string `json:"CustIP"`
}

type accountInfo struct {
	Connected bool   `json:"connected"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}
