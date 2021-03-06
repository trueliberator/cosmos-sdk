package tx

import (
	"testing"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

var (
	_, _, addr1 = testdata.KeyTestPubAddr()
	_, _, addr2 = testdata.KeyTestPubAddr()

	coins = sdk.Coins{sdk.NewInt64Coin("foocoin", 10)}
	gas   = uint64(10000)
	msg   = testdata.NewTestMsg(addr1, addr2)
	memo  = "foo"
)

func buildTx(t *testing.T, bldr *builder) {
	bldr.SetFeeAmount(coins)
	bldr.SetGasLimit(gas)
	bldr.SetMemo(memo)
	require.NoError(t, bldr.SetMsgs(msg))
}

func TestLegacyAminoJSONHandler_GetSignBytes(t *testing.T) {
	bldr := newBuilder(std.DefaultPublicKeyCodec{})
	buildTx(t, bldr)
	tx := bldr.GetTx()

	var (
		chainId        = "test-chain"
		accNum  uint64 = 7
		seqNum  uint64 = 7
	)

	handler := signModeLegacyAminoJSONHandler{}
	signingData := signing.SignerData{
		ChainID:         chainId,
		AccountNumber:   accNum,
		AccountSequence: seqNum,
	}
	signBz, err := handler.GetSignBytes(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, signingData, tx)
	require.NoError(t, err)

	expectedSignBz := types.StdSignBytes(chainId, accNum, seqNum, types.StdFee{
		Amount: coins,
		Gas:    gas,
	}, []sdk.Msg{msg}, memo)

	require.Equal(t, expectedSignBz, signBz)

	// expect error with wrong sign mode
	_, err = handler.GetSignBytes(signingtypes.SignMode_SIGN_MODE_DIRECT, signingData, tx)
	require.Error(t, err)

	// expect error with timeout height
	bldr = newBuilder(std.DefaultPublicKeyCodec{})
	buildTx(t, bldr)
	bldr.tx.Body.TimeoutHeight = 10
	tx = bldr.GetTx()
	signBz, err = handler.GetSignBytes(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, signingData, tx)
	require.Error(t, err)

	// expect error with extension options
	bldr = newBuilder(std.DefaultPublicKeyCodec{})
	buildTx(t, bldr)
	any, err := cdctypes.NewAnyWithValue(testdata.NewTestMsg())
	require.NoError(t, err)
	bldr.tx.Body.ExtensionOptions = []*cdctypes.Any{any}
	tx = bldr.GetTx()
	signBz, err = handler.GetSignBytes(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, signingData, tx)
	require.Error(t, err)

	// expect error with non-critical extension options
	bldr = newBuilder(std.DefaultPublicKeyCodec{})
	buildTx(t, bldr)
	bldr.tx.Body.NonCriticalExtensionOptions = []*cdctypes.Any{any}
	tx = bldr.GetTx()
	signBz, err = handler.GetSignBytes(signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, signingData, tx)
	require.Error(t, err)
}

func TestLegacyAminoJSONHandler_DefaultMode(t *testing.T) {
	handler := signModeLegacyAminoJSONHandler{}
	require.Equal(t, signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON, handler.DefaultMode())
}

func TestLegacyAminoJSONHandler_Modes(t *testing.T) {
	handler := signModeLegacyAminoJSONHandler{}
	require.Equal(t, []signingtypes.SignMode{signingtypes.SignMode_SIGN_MODE_LEGACY_AMINO_JSON}, handler.Modes())
}
