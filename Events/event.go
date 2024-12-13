package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

func blockEventListener(organization string, channelName string) {

	orgProfile := profile[organization]
	mspID := orgProfile.MSPID
	certPath := orgProfile.CertPath
	keyPath := orgProfile.KeyDirectory
	tlsCertPath := orgProfile.TLSCertPath
	gatewayPeer := orgProfile.GatewayPeer
	peerEndpoint := orgProfile.PeerEndpoint

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection(tlsCertPath, gatewayPeer, peerEndpoint)
	defer clientConnection.Close()

	id := newIdentity(certPath, mspID)
	sign := newSign(keyPath)

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	network := gw.GetNetwork(channelName)

	// Context used for event listening
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("\n*** Start Block event listening")

	events, err := network.BlockEvents(ctx, client.WithStartBlock(2))
	if err != nil {
		panic(fmt.Errorf("failed to start Block event listening: %w", err))
	}

	for event := range events {
		fmt.Printf("Received block number %d\n", event.GetHeader().GetNumber())
	}

}

func chaincodeEventListener(organization string, channelName string, chaincodeName string) {

	orgProfile := profile[organization]
	mspID := orgProfile.MSPID
	certPath := orgProfile.CertPath
	keyPath := orgProfile.KeyDirectory
	tlsCertPath := orgProfile.TLSCertPath
	gatewayPeer := orgProfile.GatewayPeer
	peerEndpoint := orgProfile.PeerEndpoint

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection(tlsCertPath, gatewayPeer, peerEndpoint)
	defer clientConnection.Close()

	id := newIdentity(certPath, mspID)
	sign := newSign(keyPath)

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	network := gw.GetNetwork(channelName)

	// Context used for event listening
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fmt.Println("\n*** Start Chaincode event listening")

	events, err := network.ChaincodeEvents(ctx, chaincodeName)
	if err != nil {
		panic(fmt.Errorf("failed to start Chaincode event listening: %w", err))
	}

	for event := range events {
		fmt.Printf("Received Chaincode Event: %s\n Data: %s\n \n", event.EventName, event.Payload)
	}

}

func pvtBlockEventListener(organization string, channelName string) {

	orgProfile := profile[organization]
	mspID := orgProfile.MSPID
	certPath := orgProfile.CertPath
	keyPath := orgProfile.KeyDirectory
	tlsCertPath := orgProfile.TLSCertPath
	gatewayPeer := orgProfile.GatewayPeer
	peerEndpoint := orgProfile.PeerEndpoint

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	clientConnection := newGrpcConnection(tlsCertPath, gatewayPeer, peerEndpoint)
	defer clientConnection.Close()

	id := newIdentity(certPath, mspID)
	sign := newSign(keyPath)

	// Create a Gateway connection for a specific client identity
	gw, err := client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}
	defer gw.Close()

	network := gw.GetNetwork(channelName)

	// Context used for event listening
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("\n*** Start Block event listening")

	events, err := network.BlockAndPrivateDataEvents(ctx, client.WithStartBlock(1))
	if err != nil {
		panic(fmt.Errorf("failed to start Block event listening: %w", err))
	}
	// .GetHeader().GetNumber()

	for event := range events {
		//fmt.Println("Private block event",event)
		if event.GetPrivateDataMap() != nil {
			//fmt.Println("private data event is:",event)
			fmt.Printf("Received block: %d with pvt data \n", event.GetBlock().GetHeader().GetNumber())
		}
	}

}