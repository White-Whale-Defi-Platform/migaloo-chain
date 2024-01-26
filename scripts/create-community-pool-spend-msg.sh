BINARY=${1:-migalood}
CONTINUE=${CONTINUE:-"false"}
HOME_DIR=$(pwd)/mytestnet
SCRIPTS_FOLDER=$(pwd)/scripts
ENV=${ENV:-""}

CHAIN_ID="localmigaloo"
KEYRING="test"
KEY="test0"
KEY1="test1"
KEY2="test2"
DENOM=${2:-uwhale}

COMMUNITY_POOL_AMOUNT=1000000000

## Test0 fund community pool
$BINARY tx distribution fund-community-pool $COMMUNITY_POOL_AMOUNT$DENOM  --from $KEY --keyring-backend $KEYRING --chain-id $CHAIN_ID --home $HOME_DIR -y

## Show receipient balance before proposal.
test1_addr=$($BINARY keys show $KEY1 -a --keyring-backend $KEYRING --home $HOME_DIR)
recipient=$($BINARY keys show $KEY2 -a --keyring-backend $KEYRING --home $HOME_DIR)

PRE_AMOUNT=$($BINARY query bank balances $recipient --chain-id $CHAIN_ID --home $HOME_DIR -o json | jq -r ".balances[0].amount")
echo "Recipient: $recipient"
echo "Pre receipient amount: $PRE_AMOUNT"


AMOUNT_REQUEST=$COMMUNITY_POOL_AMOUNT
proposal_file=$SCRIPTS_FOLDER/proposal.json

# chihuahua10d07y265gmmuvt4z0w9aw880jnsr700jeh7th3
authority=$($BINARY query auth module-account gov -o json | jq -r '.account.base_account.address')

cat << EOF > $proposal_file
{
  "title": "Community Spend for Coinhall Integration",
  "metadata": "ipfs://commonwealth.im/chihuahua/discussion/14608-improve-meme-trading-on-chihuahua-with-coinhall",
  "summary": "This proposal is to request funds for ...",
  "messages": [
    {
      "@type": "/cosmos.distribution.v1beta1.MsgCommunityPoolSpend",
      "recipient": "chihuahua1rqya48t6u5vtj55uumrx3puwucq6tpppz6tj6w",
      "amount": [
        {
          "denom": "uhuahua",
          "amount": "275000000000000000"
        }
      ]
    }
  ]
}
EOF


echo "Proposal file: $proposal_file"
sleep 3
$BINARY tx gov submit-proposal $proposal_file --from test1  --keyring-backend test --chain-id $CHAIN_ID --home $HOME_DIR -y


## Validator vote yes 
sleep 3
$BINARY tx gov vote 1 yes --from test0  --keyring-backend test --chain-id $CHAIN_ID --home $HOME_DIR -y

## Check recipient balance after proposal.
sleep 20 
POST_AMOUNT=$($BINARY query bank balances $recipient --chain-id $CHAIN_ID --home $HOME_DIR -o json | jq -r ".balances[0].amount")
## assert post amount 
EXPECTED_POST_AMOUNT=$((PRE_AMOUNT + COMMUNITY_POOL_AMOUNT))
if [ "$POST_AMOUNT" -eq "$EXPECTED_POST_AMOUNT" ]; then
  echo "Post recipient amount is as expected: $POST_AMOUNT"
else
  echo "Error: Post recipient amount $POST_AMOUNT does not match expected amount $EXPECTED_POST_AMOUNT"
  exit 1
fi
