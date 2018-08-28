package gov

import (
	"fmt"
	"os"

	"github.com/irisnet/irishub/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/irisnet/irishub/modules/gov"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/irisnet/irishub/app"
	"encoding/json"
)

const (
	flagProposalID        = "proposal-id"
	flagTitle             = "title"
	flagDescription       = "description"
	flagProposalType      = "type"
	flagDeposit           = "deposit"
	flagVoter             = "voter"
	flagOption            = "option"
	flagDepositer         = "depositer"
	flagStatus            = "status"
	flagLatestProposalIDs = "latest"

	flagParams       = "params"
)

// GetCmdSubmitProposal implements submitting a proposal transaction command.
func GetCmdSubmitProposal(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proposal",
		Short: "Submit a proposal along with an initial deposit",
		RunE: func(cmd *cobra.Command, args []string) error {
			title := viper.GetString(flagTitle)
			description := viper.GetString(flagDescription)
			strProposalType := viper.GetString(flagProposalType)
			initialDeposit := viper.GetString(flagDeposit)
			paramsStr := viper.GetString(flagParams)

			ctx := app.NewContext().WithCodeC(cdc)
			ctx = ctx.WithCLIContext(ctx.WithLogger(os.Stdout).
				WithAccountDecoder(authcmd.GetAccountDecoder(cdc)))

			fromAddr, err := ctx.GetFromAddress()
			if err != nil {
				return err
			}

			amount, err := ctx.ParseCoins(initialDeposit)
			if err != nil {
				return err
			}

			proposalType, err := gov.ProposalTypeFromString(strProposalType)
			if err != nil {
				return err
			}

			var params gov.Params
			if proposalType == gov.ProposalTypeParameterChange {
				if err := json.Unmarshal([]byte(paramsStr),&params);err != nil{
					fmt.Println(err.Error())
					return nil
				}
			}

			msg := gov.NewMsgSubmitProposal(title, description, proposalType, fromAddr, amount,params)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			// Build and sign the transaction, then broadcast to Tendermint
			// proposalID must be returned, and it is a part of response.
			ctx.PrintResponse = true
			return utils.SendTx(ctx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagTitle, "", "title of proposal")
	cmd.Flags().String(flagDescription, "", "description of proposal")
	cmd.Flags().String(flagProposalType, "", "proposalType of proposal,eg:Text/ParameterChange/SoftwareUpgrade")
	cmd.Flags().String(flagDeposit, "", "deposit of proposal")
	cmd.Flags().String(flagParams, "", "parameter of proposal,eg. [{key:key,value:value,op:update}]")

	return cmd
}

// GetCmdDeposit implements depositing tokens for an active proposal.
func GetCmdDeposit(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deposit",
		Short: "deposit tokens for activing proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := app.NewContext().WithCodeC(cdc)
			ctx = ctx.WithCLIContext(ctx.WithLogger(os.Stdout).
				WithAccountDecoder(authcmd.GetAccountDecoder(cdc)))

			depositerAddr, err := ctx.GetFromAddress()
			if err != nil {
				return err
			}

			proposalID := viper.GetInt64(flagProposalID)

			amount, err := sdk.ParseCoins(viper.GetString(flagDeposit))
			if err != nil {
				return err
			}

			msg := gov.NewMsgDeposit(depositerAddr, proposalID, amount)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			// Build and sign the transaction, then broadcast to a Tendermint
			// node.
			return utils.SendTx(ctx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagProposalID, "", "proposalID of proposal depositing on")
	cmd.Flags().String(flagDeposit, "", "amount of deposit")

	return cmd
}

// GetCmdVote implements creating a new vote command.
func GetCmdVote(cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote",
		Short: "vote for an active proposal, options: Yes/No/NoWithVeto/Abstain",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := app.NewContext().WithCodeC(cdc)
			ctx = ctx.WithCLIContext(ctx.WithLogger(os.Stdout).
				WithAccountDecoder(authcmd.GetAccountDecoder(cdc)))

			voterAddr, err := ctx.GetFromAddress()
			if err != nil {
				return err
			}

			proposalID := viper.GetInt64(flagProposalID)
			option := viper.GetString(flagOption)

			byteVoteOption, err := gov.VoteOptionFromString(option)
			if err != nil {
				return err
			}

			msg := gov.NewMsgVote(voterAddr, proposalID, byteVoteOption)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			fmt.Printf("Vote[Voter:%s,ProposalID:%d,Option:%s]",
				voterAddr.String(), msg.ProposalID, msg.Option.String(),
			)

			// Build and sign the transaction, then broadcast to a Tendermint
			// node.
			return utils.SendTx(ctx, []sdk.Msg{msg})
		},
	}

	cmd.Flags().String(flagProposalID, "", "proposalID of proposal voting on")
	cmd.Flags().String(flagOption, "", "vote option {Yes, No, NoWithVeto, Abstain}")

	return cmd
}

// GetCmdQueryProposal implements the query proposal command.
func GetCmdQueryProposal(storeName string, cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-proposal",
		Short: "query proposal details",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := app.NewContext().WithCodeC(cdc)
			proposalID := viper.GetInt64(flagProposalID)

			res, err := ctx.QueryStore(gov.KeyProposal(proposalID), storeName)
			if len(res) == 0 || err != nil {
				return errors.Errorf("proposalID [%d] is not existed", proposalID)
			}

			var proposal gov.Proposal
			cdc.MustUnmarshalBinary(res, &proposal)

			output, err := wire.MarshalJSONIndent(cdc, proposal)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	cmd.Flags().String(flagProposalID, "", "proposalID of proposal being queried")

	return cmd
}

// nolint: gocyclo
// GetCmdQueryProposals implements a query proposals command.
func GetCmdQueryProposals(storeName string, cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-proposals",
		Short: "query proposals with optional filters",
		RunE: func(cmd *cobra.Command, args []string) error {
			bechDepositerAddr := viper.GetString(flagDepositer)
			bechVoterAddr := viper.GetString(flagVoter)
			strProposalStatus := viper.GetString(flagStatus)
			latestProposalsIDs := viper.GetInt64(flagLatestProposalIDs)

			var err error
			var voterAddr sdk.AccAddress
			var depositerAddr sdk.AccAddress
			var proposalStatus gov.ProposalStatus

			if len(bechDepositerAddr) != 0 {
				depositerAddr, err = sdk.AccAddressFromBech32(bechDepositerAddr)
				if err != nil {
					return err
				}
			}

			if len(bechVoterAddr) != 0 {
				voterAddr, err = sdk.AccAddressFromBech32(bechVoterAddr)
				if err != nil {
					return err
				}
			}

			if len(strProposalStatus) != 0 {
				proposalStatus, err = gov.ProposalStatusFromString(strProposalStatus)
				if err != nil {
					return err
				}
			}

			ctx := app.NewContext().WithCodeC(cdc)

			res, err := ctx.QueryStore(gov.KeyNextProposalID, storeName)
			if err != nil {
				return err
			}
			var maxProposalID int64
			cdc.MustUnmarshalBinary(res, &maxProposalID)

			matchingProposals := []gov.Proposal{}

			if latestProposalsIDs == 0 {
				latestProposalsIDs = maxProposalID
			}

			for proposalID := maxProposalID - latestProposalsIDs; proposalID < maxProposalID; proposalID++ {
				if voterAddr != nil {
					res, err = ctx.QueryStore(gov.KeyVote(proposalID, voterAddr), storeName)
					if err != nil || len(res) == 0 {
						continue
					}
				}

				if depositerAddr != nil {
					res, err = ctx.QueryStore(gov.KeyDeposit(proposalID, depositerAddr), storeName)
					if err != nil || len(res) == 0 {
						continue
					}
				}

				res, err = ctx.QueryStore(gov.KeyProposal(proposalID), storeName)
				if err != nil || len(res) == 0 {
					continue
				}

				var proposal gov.Proposal
				cdc.MustUnmarshalBinary(res, &proposal)

				if len(strProposalStatus) != 0 {
					if proposal.GetStatus() != proposalStatus {
						continue
					}
				}

				matchingProposals = append(matchingProposals, proposal)
			}

			if len(matchingProposals) == 0 {
				fmt.Println("No matching proposals found")
				return nil
			}

			for _, proposal := range matchingProposals {
				fmt.Printf("  %d - %s\n", proposal.GetProposalID(), proposal.GetTitle())
			}

			return nil
		},
	}

	cmd.Flags().String(flagLatestProposalIDs, "", "(optional) limit to latest [number] proposals. Defaults to all proposals")
	cmd.Flags().String(flagDepositer, "", "(optional) filter by proposals deposited on by depositer")
	cmd.Flags().String(flagVoter, "", "(optional) filter by proposals voted on by voted")
	cmd.Flags().String(flagStatus, "", "(optional) filter proposals by proposal status")

	return cmd
}

// Command to Get a Proposal Information
// GetCmdQueryVote implements the query proposal vote command.
func GetCmdQueryVote(storeName string, cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-vote",
		Short: "query vote",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := app.NewContext().WithCodeC(cdc)
			proposalID := viper.GetInt64(flagProposalID)

			voterAddr, err := sdk.AccAddressFromBech32(viper.GetString(flagVoter))
			if err != nil {
				return err
			}

			res, err := ctx.QueryStore(gov.KeyVote(proposalID, voterAddr), storeName)
			if len(res) == 0 || err != nil {
				return errors.Errorf("proposalID [%d] does not exist", proposalID)
			}

			var vote gov.Vote
			cdc.MustUnmarshalBinary(res, &vote)

			output, err := wire.MarshalJSONIndent(cdc, vote)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	cmd.Flags().String(flagProposalID, "", "proposalID of proposal voting on")
	cmd.Flags().String(flagVoter, "", "bech32 voter address")

	return cmd
}

// GetCmdQueryVotes implements the command to query for proposal votes.
func GetCmdQueryVotes(storeName string, cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-votes",
		Short: "query votes on a proposal",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := app.NewContext().WithCodeC(cdc)
			proposalID := viper.GetInt64(flagProposalID)

			res, err := ctx.QueryStore(gov.KeyProposal(proposalID), storeName)
			if len(res) == 0 || err != nil {
				return errors.Errorf("proposalID [%d] does not exist", proposalID)
			}

			var proposal gov.Proposal
			cdc.MustUnmarshalBinary(res, &proposal)

			if proposal.GetStatus() != gov.StatusVotingPeriod {
				fmt.Println("Proposal not in voting period.")
				return nil
			}

			res2, err := ctx.QuerySubspace(gov.KeyVotesSubspace(proposalID), storeName)
			if err != nil {
				return err
			}

			var votes []gov.Vote
			for i := 0; i < len(res2); i++ {
				var vote gov.Vote
				cdc.MustUnmarshalBinary(res2[i].Value, &vote)
				votes = append(votes, vote)
			}

			output, err := wire.MarshalJSONIndent(cdc, votes)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	cmd.Flags().String(flagProposalID, "", "proposalID of which proposal's votes are being queried")

	return cmd
}

func GetCmdQueryConfig(storeName string, cdc *wire.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-params",
		Short: "query parameter proposal's config",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := app.NewContext().WithCodeC(cdc)
			res , err  := ctx.QuerySubspace([]byte(gov.Prefix),storeName)

			var kvs []KvPair
			for _,kv := range res {
				var v string
				cdc.UnmarshalBinary(kv.Value, &v)
				kv := KvPair{
					K: string(kv.Key),
					V: v,
				}
				kvs = append(kvs, kv)
			}
			output, err := wire.MarshalJSONIndent(cdc, kvs)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	return cmd
}

type KvPair struct {
	K string `json:"key"`
	V string `json:"value"`
}
