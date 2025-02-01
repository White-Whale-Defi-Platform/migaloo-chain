package e2e

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	wasmvmtypes "github.com/CosmWasm/wasmvm/v2/types"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmWasm/wasmd/tests/ibctesting"
	"github.com/CosmWasm/wasmd/x/wasm/keeper/testdata"
	"github.com/CosmWasm/wasmd/x/wasm/types"
)

// InstantiateReflectContract store and instantiate a reflect contract instance
func InstantiateReflectContract(t *testing.T, chain *ibctesting.TestChain) sdk.AccAddress {
	t.Helper()
	codeID := chain.StoreCodeFile("testdata/reflect_2_0.wasm").CodeID
	contractAddr := chain.InstantiateContract(codeID, []byte(`{}`))
	require.NotEmpty(t, contractAddr)
	return contractAddr
}

// MustExecViaReflectContract submit execute message to send payload to reflect contract
func MustExecViaReflectContract(t *testing.T, chain *ibctesting.TestChain, contractAddr sdk.AccAddress, msgs ...wasmvmtypes.CosmosMsg) *abci.ExecTxResult {
	t.Helper()
	rsp, err := ExecViaReflectContract(t, chain, contractAddr, msgs)
	require.NoError(t, err)
	return rsp
}

type sdkMessageType interface {
	codec.ProtoMarshaler
	sdk.Msg
}

func MustExecViaStargateReflectContract[T sdkMessageType](t *testing.T, chain *ibctesting.TestChain, contractAddr sdk.AccAddress, msgs ...T) *abci.ExecTxResult {
	t.Helper()
	// convert messages to stargate variant
	vmMsgs := make([]string, len(msgs))
	for i, m := range msgs {
		bz, err := chain.Codec.Marshal(m)
		require.NoError(t, err)
		// json is built manually because the wasmvm CosmosMsg does not have the `Stargate` variant anymore
		vmMsgs[i] = fmt.Sprintf("{\"stargate\":{\"type_url\":\"%s\",\"value\":\"%s\"}}", sdk.MsgTypeURL(m), base64.StdEncoding.EncodeToString(bz))
	}
	// build the complete reflect message
	reflectSendBz := []byte(fmt.Sprintf("{\"reflect_msg\":{\"msgs\":%s}}", vmMsgs))

	execMsg := &types.MsgExecuteContract{
		Sender:   chain.SenderAccount.GetAddress().String(),
		Contract: contractAddr.String(),
		Msg:      reflectSendBz,
	}
	rsp, err := chain.SendMsgs(execMsg)
	require.NoError(t, err)
	return rsp
}

// ExecViaReflectContract submit execute message to send payload to reflect contract
func ExecViaReflectContract(t *testing.T, chain *ibctesting.TestChain, contractAddr sdk.AccAddress, msgs []wasmvmtypes.CosmosMsg) (*abci.ExecTxResult, error) {
	t.Helper()
	require.NotEmpty(t, msgs)
	reflectSend := testdata.ReflectHandleMsg{
		Reflect: &testdata.ReflectPayload{Msgs: msgs},
	}
	reflectSendBz, err := json.Marshal(reflectSend)
	require.NoError(t, err)
	execMsg := &types.MsgExecuteContract{
		Sender:   chain.SenderAccount.GetAddress().String(),
		Contract: contractAddr.String(),
		Msg:      reflectSendBz,
	}
	return chain.SendMsgs(execMsg)
}
