package main

// Config represents the configuration for a role.
type Config struct {
	CertPath     string `json:"certPath"`
	KeyDirectory string `json:"keyPath"`
	TLSCertPath  string `json:"tlsCertPath"`
	PeerEndpoint string `json:"peerEndpoint"`
	GatewayPeer  string `json:"gatewayPeer"`
	MSPID        string `json:"mspID"`
}

// Create a Profile map
var profile = map[string]Config{

	"charityorg": {
		CertPath:     "../Charity-network/organizations/peerOrganizations/charityorg.charity.com/users/User1@charityorg.charity.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Charity-network/organizations/peerOrganizations/charityorg.charity.com/users/User1@charityorg.charity.com/msp/keystore/",
		TLSCertPath:  "../Charity-network/organizations/peerOrganizations/charityorg.charity.com/peers/peer0.charityorg.charity.com/tls/ca.crt",
		PeerEndpoint: "localhost:7051",
		GatewayPeer:  "peer0.charityorg.charity.com",
		MSPID:        "CharityorgMSP",
	},

	"donar": {
		CertPath:     "../Charity-network/organizations/peerOrganizations/donar.charity.com/users/User1@donar.charity.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Charity-network/organizations/peerOrganizations/donar.charity.com/users/User1@donar.charity.com/msp/keystore/",
		TLSCertPath:  "../Charity-network/organizations/peerOrganizations/donar.charity.com/peers/peer0.donar.charity.com/tls/ca.crt",
		PeerEndpoint: "localhost:9051",
		GatewayPeer:  "peer0.donar.charity.com",
		MSPID:        "DonarMSP",
	},

	"auditor": {
		CertPath:     "../Charity-network/organizations/peerOrganizations/auditor.charity.com/users/User1@auditor.charity.com/msp/signcerts/cert.pem",
		KeyDirectory: "../Charity-network/organizations/peerOrganizations/auditor.charity.com/users/User1@auditor.charity.com/msp/keystore/",
		TLSCertPath:  "../Charity-network/organizations/peerOrganizations/auditor.charity.com/peers/peer0.auditor.charity.com/tls/ca.crt",
		PeerEndpoint: "localhost:11051",
		GatewayPeer:  "peer0.auditor.charity.com",
		MSPID:        "AuditorMSP",
	},
}