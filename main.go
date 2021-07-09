package main

type Division struct {
	DivisionId int
	Date string
	PublicationUpdate string
	Number int
	IsDeferred bool
	EvelType string
	EvelCountry string
	Title string
	AyeCount int
	NoCount int
	DoubleMajorityAyeCount int
	DoubleMajorityNoCount int
	FriendlyDescription string
	FriendlyTitle string
	RemoteVotingStart string
	RemoteVotingEnd string

}

type AyeTellers struct {
	MemberID int
	Name string
	Party string
	SubParty string
	PartyColour string
	PartyAbbreviation string
	MemberFrom string
	ListAs string
	ProxyName string
}

type NoTellers struct {
	MemberID int
	Name string
	Party string
	SubParty string
	PartyColour string
	PartyAbbreviation string
	MemberFrom string
	ListAs string
	ProxyName string
}

type Ayes struct {
	MemberID int
	Name string
	Party string
	SubParty string
	PartyColour string
	PartyAbbreviation string
	MemberFrom string
	ListAs string
	ProxyName string
}

type Noes struct {
	MemberID int
	Name string
	Party string
	SubParty string
	PartyColour string
	PartyAbbreviation string
	MemberFrom string
	ListAs string
	ProxyName string
}


func main() {



}