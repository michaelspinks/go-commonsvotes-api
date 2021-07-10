package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Division struct this contains elements relating to the division of House of Commons
// details about the division
// slices of RecordedMembers for Aye, NoTellers, the Ayes, the Nos, and NoVotesRecorded
type Division struct {
	DivisionId             int    `json:"DivisionId"`
	Date                   string `json:"Date"`
	PublicationUpdate      string `json:"PublicationUpdate"`
	Number                 int    `json:"Number"`
	IsDeferred             bool   `json:"IsDeferred"`
	EVELType               string `json:"EVELType"`
	EVELCountry            string `json:"EVELCountry"`
	Title                  string `json:"Title"`
	AyeCount               int    `json:"AyeCount"`
	NoCount                int    `json:"NoCount"`
	DoubleMajorityAyeCount int    `json:"DoubleMajorityAyeCount"`
	DoubleMajorityNoCount  int    `json:"DoubleMajorityNoCount"`
	FriendlyDescription    string `json:"FriendlyDescription"`
	FriendlyTitle          string `json:"FriendlyTitle"`
	RemoteVotingStart      string `json:"RemoteVotingStart"`
	RemoteVotingEnd        string `json:"RemoteVotingEnd"`
	AyeTellers             []RecordedMember
	NoTellers              []RecordedMember
	Ayes                   []RecordedMember
	Noes                   []RecordedMember
	NoVoteRecorded         []RecordedMember
}

type RecordedMember struct {
	MemberID          int    `json:"MemberID"`
	Name              string `json:"Name"`
	Party             string `json:"Party"`
	SubParty          string `json:"SubParty"`
	PartyColour       string `json:"PartyColour"`
	PartyAbbreviation string `json:"PartyAbbreviation"`
	MemberFrom        string `json:"MemberFrom"`
	ListAs            string `json:"ListAs"`
	ProxyName         string `json:"ProxyName"`
}

func main() {

	baseURL := "https://commonsvotes-api.parliament.uk/data/division/2.json"

	// define our http request
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// set request header to accept application/json as per the Model on the commons site
	req.Header.Set("Accept", "application/json")

	// send our HTTP GET request and receive the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// defer the closing of our response so we can parse it later on
	defer resp.Body.Close()

	// read the response body into a byte array b
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// initialize our Division struct
	var division Division

	// unmarshal the byte array with our response body
	// into division
	jsonErr := json.Unmarshal(b, &division)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// print out number of AyeTellers
	// loop through the AyeTellers printing out their relevant information
	fmt.Println("======================================================")
	fmt.Println("Ayetellers: " + strconv.Itoa(len(division.AyeTellers)))
	fmt.Println("-----------------------------------------------------")
	for i := 0; i < len(division.AyeTellers); i++ {
		fmt.Println("MemberID: " + strconv.Itoa(division.AyeTellers[i].MemberID))
		fmt.Println("Name: " + division.AyeTellers[i].Name)
		fmt.Println("Party: " + division.AyeTellers[i].Party)
		fmt.Println("SubParty: " + division.AyeTellers[i].SubParty)
		fmt.Println("PartyColour: " + division.AyeTellers[i].PartyColour)
		fmt.Println("PartyAbbreviation: " + division.AyeTellers[i].PartyAbbreviation)
		fmt.Println("MemberFrom: " + division.AyeTellers[i].MemberFrom)
		fmt.Println("ListAs: " + division.AyeTellers[i].ListAs)
		fmt.Println("ProxyName: " + division.AyeTellers[i].ProxyName)
		fmt.Println("-----------------------------------------------------")
	}
	fmt.Println("======================================================")

		// print out number of AyeTellers
	// loop through the NoTellers printing out their relevant information
	fmt.Println("======================================================")
	fmt.Println("Notellers: " + strconv.Itoa(len(division.NoTellers)))
	fmt.Println("-----------------------------------------------------")
	for i := 0; i < len(division.AyeTellers); i++ {
		fmt.Println("MemberID: " + strconv.Itoa(division.NoTellers[i].MemberID))
		fmt.Println("Name: " + division.NoTellers[i].Name)
		fmt.Println("Party: " + division.NoTellers[i].Party)
		fmt.Println("SubParty: " + division.NoTellers[i].SubParty)
		fmt.Println("PartyColour: " + division.NoTellers[i].PartyColour)
		fmt.Println("PartyAbbreviation: " + division.NoTellers[i].PartyAbbreviation)
		fmt.Println("MemberFrom: " + division.NoTellers[i].MemberFrom)
		fmt.Println("ListAs: " + division.NoTellers[i].ListAs)
		fmt.Println("ProxyName: " + division.NoTellers[i].ProxyName)
		fmt.Println("-----------------------------------------------------")
	}
	fmt.Println("======================================================")

	fmt.Println("======================================================")
	fmt.Println("Ayes: " + strconv.Itoa(len(division.Ayes)))
	fmt.Println("-----------------------------------------------------")
	for i := 0; i < len(division.Ayes); i++ {
		fmt.Println("MemberID: " + strconv.Itoa(division.Ayes[i].MemberID))
		fmt.Println("Name: " + division.Ayes[i].Name)
		fmt.Println("Party: " + division.Ayes[i].Party)
		fmt.Println("SubParty: " + division.Ayes[i].SubParty)
		fmt.Println("PartyColour: " + division.Ayes[i].PartyColour)
		fmt.Println("PartyAbbreviation: " + division.Ayes[i].PartyAbbreviation)
		fmt.Println("MemberFrom: " + division.Ayes[i].MemberFrom)
		fmt.Println("ListAs: " + division.Ayes[i].ListAs)
		fmt.Println("ProxyName: " + division.Ayes[i].ProxyName)
		fmt.Println("-----------------------------------------------------")
	}
	fmt.Println("======================================================")
}
