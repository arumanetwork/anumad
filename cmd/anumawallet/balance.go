package main

import (
	"context"
	"fmt"

	"github.com/AnumaNetwork/anumad-testnet/cmd/anumawallet/daemon/client"
	"github.com/AnumaNetwork/anumad-testnet/cmd/anumawallet/daemon/pb"
	"github.com/AnumaNetwork/anumad-testnet/cmd/anumawallet/utils"
)

func balance(conf *balanceConfig) error {
	daemonClient, tearDown, err := client.Connect(conf.DaemonAddress)
	if err != nil {
		return err
	}
	defer tearDown()

	ctx, cancel := context.WithTimeout(context.Background(), daemonTimeout)
	defer cancel()
	response, err := daemonClient.GetBalance(ctx, &pb.GetBalanceRequest{})
	if err != nil {
		return err
	}

	pendingSuffix := ""
	if response.Pending > 0 {
		pendingSuffix = " (pending)"
	}
	if conf.Verbose {
		pendingSuffix = ""
		println("Address                                                                       Available             Pending")
		println("-----------------------------------------------------------------------------------------------------------")
		for _, addressBalance := range response.AddressBalances {
			fmt.Printf("%s %s %s\n", addressBalance.Address, utils.FormatSec(addressBalance.Available), utils.FormatSec(addressBalance.Pending))
		}
		println("-----------------------------------------------------------------------------------------------------------")
		print("                                                 ")
	}
	fmt.Printf("Total balance, ANUM %s %s%s\n", utils.FormatSec(response.Available), utils.FormatSec(response.Pending), pendingSuffix)

	return nil
}
