package main

import (
	//"encoding/json"
	//"fmt"

	"fmt"
	"io"
	"log"
	"net/http"
	//"time"
)

type Division struct {
	DivisionId             int
	Date                   string
	PublicationUpdate      string
	Number                 int
	IsDeferred             bool
	EvelType               string
	EvelCountry            string
	Title                  string
	AyeCount               int
	NoCount                int
	DoubleMajorityAyeCount int
	DoubleMajorityNoCount  int
	FriendlyDescription    string
	FriendlyTitle          string
	RemoteVotingStart      string
	RemoteVotingEnd        string
	AyeTellers             []RecordedMember
	NoTellers              []RecordedMember
	Ayes                   []RecordedMember
	Noes                   []RecordedMember
	NoVoteRecorded         []RecordedMember
}

type RecordedMember struct {
	MemberID          int
	Name              string
	Party             string
	SubParty          string
	PartyColour       string
	PartyAbbreviation string
	MemberFrom        string
	ListAs            string
	ProxyName         string
}

func main() {

	req, err := http.NewRequest("GET", "https://commonsvotes-api.parliament.uk/data/division/2.json", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(b))

}
