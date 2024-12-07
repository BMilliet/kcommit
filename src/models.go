package src

type CommitType struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Example     string `json:"example"`
}

type UserCustomRc struct {
	CommitTypes []CommitType `json:"commitTypes"`
}
