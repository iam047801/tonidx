package contract

import (
	"context"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/allisson/go-env"
	"github.com/rs/zerolog/log"

	"github.com/iam047801/tonidx/internal/core"
	"github.com/iam047801/tonidx/internal/core/db"
)

func InsertInterface() {
	var err error

	contract := new(core.ContractInterface)

	f := flag.NewFlagSet(os.Args[1], flag.ExitOnError)
	name := f.String("name", "", "Unique contract name (example: getgems_nft_sale)")
	address := f.String("address", "", "Contract address")
	code := f.String("code", "", "Contract code encoded to hex")
	getMethods := f.String("getmethods", "", "Contract get methods separated with commas")
	_ = f.Parse(os.Args[2:])

	if *name == "" {
		log.Fatal().Msg("contract name must be set")
	}
	if *address == "" && *code == "" && *getMethods == "" {
		log.Fatal().Msg("contract address or code or get methods must be set")
	}

	contract.Name = core.ContractType(*name)
	if *address != "" {
		contract.Address = *address
	}
	if *code != "" {
		contract.Code, err = hex.DecodeString(*code)
		if err != nil {
			log.Fatal().Err(err).Msg("cannot parse contract code")
		}
	}
	if *getMethods != "" {
		contract.GetMethods = strings.Split(*getMethods, ",")
	}

	conn, err := db.Connect(context.Background(), env.GetString("DB_URL", ""))
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to a database")
	}
	_, err = conn.NewInsert().Model(contract).Exec(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("cannot insert contract interface")
	}

	log.Info().
		Str("name", string(contract.Name)).
		Str("address", contract.Address).
		Str("get_methods", fmt.Sprintf("%+v", contract.GetMethods)).
		Str("code", hex.EncodeToString(contract.Code)).
		Msg("inserted new contract interface")
}
