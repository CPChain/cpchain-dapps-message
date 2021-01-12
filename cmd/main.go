package main

import (
	"fmt"
	"message"
	"os"
	"path/filepath"
	"sort"

	"bitbucket.org/cpchain/chain/accounts/abi/bind"
	"bitbucket.org/cpchain/chain/api/cpclient"
	"bitbucket.org/cpchain/chain/commons/log"
	"bitbucket.org/cpchain/chain/tools/contract-admin/flags"
	"bitbucket.org/cpchain/chain/tools/contract-admin/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
)

var sendFlag = []cli.Flag{
	cli.StringFlag{
		Name:  "to",
		Usage: "to address",
	},
	cli.StringFlag{
		Name:  "msg",
		Usage: "message",
	},
}

var (
	// MessageCommand message contract
	MessageCommand = cli.Command{
		Name:  "message",
		Usage: "Manage Message Contract",
		Description: `
		Manage Message Contract
		`,
		Flags: flags.GeneralFlags,
		Subcommands: []cli.Command{
			{
				Name:        "deploy",
				Usage:       "message deploy",
				Action:      deploy,
				Flags:       flags.GeneralFlags,
				ArgsUsage:   "string",
				Description: `deploy contract`,
			},
			{
				Name:        "send",
				Usage:       "send --to <address> --msg <message>",
				Action:      send,
				Flags:       append(flags.GeneralFlags, sendFlag...),
				Description: `register message`,
			},
			{
				Name:        "show-configs",
				Usage:       "message show-configs",
				Action:      showConfigs,
				Flags:       flags.GeneralFlags,
				Description: `show configs`,
			},
			{
				Name:        "disable",
				Usage:       "disbale",
				Action:      disable,
				Flags:       flags.GeneralFlags,
				Description: `disable contract`,
			},
			{
				Name:        "enable",
				Usage:       "enable",
				Action:      enable,
				Flags:       flags.GeneralFlags,
				Description: `enable contract`,
			},
		},
	}
)

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Authors = []cli.Author{
		{
			Name:  "The cpchain authors",
			Email: "info@cpchain.io",
		},
	}
	app.Copyright = "LGPL"
	app.Usage = "Executable for the cpchain official contract admin"

	app.Action = cli.ShowAppHelp

	app.Commands = []cli.Command{
		MessageCommand,
	}

	// maintain order
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}

func main() {
	if err := newApp().Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func deploy(ctx *cli.Context) error {
	endpoint, err := flags.GetEndpoint(ctx)
	if err != nil {
		log.Fatal("endpoint error", "err", err)
	}
	client, err := utils.PrepareCpclient(endpoint)
	if err != nil {
		log.Fatal("prepare client error", "err", err)
	}
	_ = client
	keystoreFile, err := flags.GetKeystorePath(ctx)
	if err != nil {
		log.Fatal("get keystore path error", "err", err)
	}
	password := utils.GetPassword()
	_, key := utils.GetAddressAndKey(keystoreFile, password)
	auth := bind.NewKeyedTransactor(key.PrivateKey)
	_ = auth

	addr, tx, _, err := message.DeployMessage(auth, client)
	if err != nil {
		log.Fatal("deploy message contract error", "err", err)
	}
	log.Info("message contract address", "addr", addr.Hex())
	return utils.WaitMined(client, tx)
}

func showConfigs(ctx *cli.Context) error {
	instance, _, _, err := createContractInstanceAndTransactor(ctx, false)
	if err != nil {
		return err
	}
	size, err := instance.Count(nil)
	if err != nil {
		log.Error("Maybe your message contract address is wrong, please check it.")
		return err
	}
	log.Info("message size", "value", size)

	enabled, _ := instance.Enabled(nil)
	log.Info("contract enabled", "value", enabled)

	return nil
}

func send(ctx *cli.Context) error {
	instance, opts, client, err := createContractInstanceAndTransactor(ctx, true)
	if err != nil {
		return err
	}
	to := ctx.String("to")
	addr := common.HexToAddress(to)
	msg := ctx.String("msg")
	tx, err := instance.SendMessage(opts, addr, msg)
	if err != nil {
		return err
	}
	return utils.WaitMined(client, tx)
}

func disable(ctx *cli.Context) error {
	instance, opts, client, err := createContractInstanceAndTransactor(ctx, true)
	if err != nil {
		return err
	}
	tx, err := instance.DisableContract(opts)
	if err != nil {
		return err
	}
	return utils.WaitMined(client, tx)
}

func enable(ctx *cli.Context) error {
	instance, opts, client, err := createContractInstanceAndTransactor(ctx, true)
	if err != nil {
		return err
	}
	tx, err := instance.EnableContract(opts)
	if err != nil {
		return err
	}
	return utils.WaitMined(client, tx)
}

func createContractInstanceAndTransactor(ctx *cli.Context, withTransactor bool) (contract *message.Message, opts *bind.TransactOpts, client *cpclient.Client, err error) {
	contractAddr, client, key, err := utils.PrepareAll(ctx, withTransactor)
	if err != nil {
		return &message.Message{}, &bind.TransactOpts{}, &cpclient.Client{}, err
	}
	if withTransactor {
		opts = bind.NewKeyedTransactor(key.PrivateKey)
	}

	contract, err = message.NewMessage(contractAddr, client)
	if err != nil {
		log.Info("Failed to create new contract instance", "err", err)
		return &message.Message{}, &bind.TransactOpts{}, &cpclient.Client{}, err
	}

	return contract, opts, client, nil
}
