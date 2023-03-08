import json

# Set this to your home directory.
INPUT_GENESIS_PATH = "/???/.migalood/config/genesis.json"
OUTPUT_GENESIS_PATH = "/???/.migalood/config/pre-genesis.json"

# Convenience variables
DENOM = "uwhale"
DECIMALS = "000000"
MILLION = 1000000
SECONDS_PER_DAY = 24 * 60 * 60

NUM_GENESIS_VALIDATORS = 5
INITIAL_GENESIS_ALLOCATION = 10
INITIAL_AMOUNT = NUM_GENESIS_VALIDATORS * INITIAL_GENESIS_ALLOCATION

# Do not edit. Used to check the total balances in the genesis.
THREE_MONTH_AMOUNT = 0
TWELVE_MONTH_AMOUNT = 0
TWENTY_FOUR_MONTH_AMOUNT = 0
THIRTY_SIX_MONTH_AMOUNT = 0
COMMUNITY_POOL_AMOUNT = 25 * MILLION
MULTI_SIG_AMOUNT = 380.3848 * MILLION - INITIAL_AMOUNT


# see: https://www.epochconverter.com for more infos on the unix time stamp
GENESIS_TIME_UNIX = 1676300400
THREE_MONTH_UNIX = 1683990000
TWELVE_MONTH_UNIX = 1707836400
TWENTY_FOUR_MONTH_UNIX = 1739458800
THIRTY_SIX_MONTH_UNIX = 1770994800

# List of all accounts that get continously vested tokens.
# Address: Address of the account.
# Amount: Amount of vested tokens of the Address.
# Vesting Duration: Duration of the vesting period in unix format starting from genesis.
VESTING_ACCOUNTS = [
    # Pre-Seed Investors
    {
        "address": "migaloo1spge4nh9r5l9lepst79rvksdz0qjgtate3d4mu",
        "amount": 2 * MILLION,
        "duration": THREE_MONTH_UNIX,
    },
    {
        "address": "migaloo1e98m6za0tc446jjdzt7ux2muajf0uj4k4p538s",
        "amount": 2 * MILLION,
        "duration": THREE_MONTH_UNIX,
    },
    {
        "address": "migaloo1lm0ewsgtysmntfzqyvf89n43g33zkdytlmsshj",
        "amount": 2 * MILLION,
        "duration": THREE_MONTH_UNIX,
    },
    {
        "address": "migaloo1vnpc90shne5ncs2ddr7g4rrttqjudz4cyp4tu4",
        "amount": 2 * MILLION,
        "duration": THREE_MONTH_UNIX,
    },
    {
        "address": "migaloo1sksw9c946ann4w5a2d7kgphrm06frtsqqgp2sv",
        "amount": 1.2 * MILLION,
        "duration": THREE_MONTH_UNIX,
    },
    {
        "address": "migaloo1e82da9n6jz4t42eh0wn5hrt6hdmf7jyqvq0np0",
        "amount": 3 * MILLION,
        "duration": THREE_MONTH_UNIX,
    },
    {
        "address": "migaloo19chphxqsduplg5qfsjae2ync94c3cgh9h049fj",
        "amount": 2 * MILLION,
        "duration": THREE_MONTH_UNIX,
    },
    # Seed Investors
    {
        "address": "migaloo1wuu8l4srf3k7phewhc5sqjvvqtax0ch8mntfc3",
        "amount": 3 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1nat09n7vfkgrv3p78vyan203umugmrkxzr3e3s",
        "amount": 1 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1ycaqrfzyfue40vey36q02l2zrvqzw6sfryur5l",
        "amount": 3 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1nfk64t5cygztmgf7ycq0yjkh72udgpqftmrsjp",
        "amount": 1 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1la4nlstynp2wl0r236du0evjpqh8h9fqxs8xpa",
        "amount": 3 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1z62c730l6u8v2q8kadenjtp8tf45m3yjwmsdz7",
        "amount": 1.5 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo19lrmt9pl06hz6d96pqvcs0jl86wdldhj5dd8fp",
        "amount": 1 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1cgdjruw7602qerr72wpzl4jyg46d0rq802v93l",
        "amount": 0.5 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1z46x9kfmp64qqk7rxun6ll9peqapmswpjqyx6q",
        "amount": 1 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1l5py295agfhqfmnh76n428s6f7tkhaypzjdras",
        "amount": 0.5 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1w8yjx8ku3xj4zv3nxrycv5536qq5kvy7nsn7y0",
        "amount": 1 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo156fmcjwlzgj8gt9h8zh43phmc4p9xmjc06jw0u",
        "amount": 3 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo10ep3rpmjjhumzeglgxfd9zs5cmt5awxh0zreuh",
        "amount": 1 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1zkdvn6j7k6dqtknhxqyysdpne5mft4r97p0k4g",
        "amount": 0.5 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1cev3vxz9t0vcef68s4ghgm6yj9lxum9h7535yh",
        "amount": 3 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1dfx3smd93xhz5j04sncdkgnrqnawkt9svt99e0",
        "amount": 2 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo14uwnqvrjkux5fah6m2r7z24ztzdw9ns4n58wyq",
        "amount": 0.5 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1mnpjf7ay5n08nw0hrzx0mpg324nh6pumxwqrv8",
        "amount": 1.5 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1lwzh84q7ffa7xrnj8x3k6nlxcuz8k5c32xcl3v",
        "amount": 0.5 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo12743c6vye8nnluqvxdylc9flsj4fjhzmth6l4l",
        "amount": 1 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1ng379qqcdpw8nhc8fa9s006dukt0h6p6uyhr6k",
        "amount": 1 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1nyjpegzqq47jvwc94g8lf5t5mhng976lmszz9r",
        "amount": 1.5 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1lns0rwwzpahe0hx7ctlz66nmlwhm6dzrp5dqk2",
        "amount": 1.5 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1cnykulmrjx928ylhtms5cq66nwac32a67lpsq2",
        "amount": 3 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1nty7kgguz5fhxz0fta5f593wygnzdkr6ut2k2z",
        "amount": 1 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    {
        "address": "migaloo1nnumc3e3nykgqtr8nr0gk09js2k4qtt4n8a93q",
        "amount": 5 * MILLION,
        "duration": TWELVE_MONTH_UNIX,
    },
    # Team
    {
        "address": "migaloo1v8n2jpmq8a97xtv9j28c557q3ztqn7rkrxtykg",
        "amount": 24 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    {
        "address": "migaloo16u8ljxqr2sjfpxur5nmkpdu0j86vxh2udhunsu",
        "amount": 15 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    {
        "address": "migaloo13w00xz4t7gftnur9q6w8ychlpxc4pw02hzdx5z",
        "amount": 14 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    {
        "address": "migaloo18a9m9stu3dyvewwcq9qmp85euxqcvln5mefync",
        "amount": 12.5 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    {
        "address": "migaloo1eenffmetkkyzf5rejtn0jq5l2p7nz35klxlgff",
        "amount": 8 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    {
        "address": "migaloo1jr6m4eu65kgf6flncd7nrhd4gcaexexr0swxm9",
        "amount": 8 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    {
        "address": "migaloo1typsgj0nvketwusvaakyerjp9exr2fuxtmda9v",
        "amount": 4.25 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    {
        "address": "migaloo1pen8w6rshxqrt9t5fsd8zhc0ux54mrlw4al4mw",
        "amount": 4 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    {
        "address": "migaloo10layywwtg6kcjr0nunvnsr9g78n6veae85zx8j",
        "amount": 2 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    {
        "address": "migaloo16k4493edgkpy4apthg6m3m5wp0s36fvmkkampn",
        "amount": 4 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    {
        "address": "migaloo1ug5h3wn9qlmn00h7dvws8h68vdeh8c0paazv8z",
        "amount": 4 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    {
        "address": "migaloo1938p0dxg9s2wdfcy8j747gfuyf40wqzkdcmuxt",
        "amount": 2 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    {
        "address": "migaloo1qp8flhtcev6ag8av3mwr3eessuq8sxjrukh7tw",
        "amount": 1 * MILLION,
        "duration": TWENTY_FOUR_MONTH_UNIX,
    },
    # Genesis Validators
    {
        "address": "migaloo1alga5e8vr6ccr9yrg0kgxevpt5xgmgrvqgujs6",
        "amount": 22.1652 * MILLION,
        "duration": THIRTY_SIX_MONTH_UNIX,
    },
    {
        "address": "migaloo1fc4kjfau480nr503yl0r8ml7vvn07d2r7ztjky",
        "amount": 3 * MILLION,
        "duration": THIRTY_SIX_MONTH_UNIX,
    },
    {
        "address": "migaloo17gcjmzpz2sfjj9waa6f2e8pr7s70v3hcudtsqj",
        "amount": 3 * MILLION,
        "duration": THIRTY_SIX_MONTH_UNIX,
    },
    {
        "address": "migaloo14kqr9fjjzl24gwfndf05wkelncm76ynkk04zjk",
        "amount": 4 * MILLION,
        "duration": THIRTY_SIX_MONTH_UNIX,
    },
    {
        "address": "migaloo1zz9ppl2wy4ruzzv8mmnx6cente9ygvcx2r3qap",
        "amount": 3 * MILLION,
        "duration": THIRTY_SIX_MONTH_UNIX,
    },
]


def create_vesting_genesis_entry(address, amount, end_time):
    return {
        "@type": "/cosmos.vesting.v1beta1.ContinuousVestingAccount",
        "base_vesting_account": {
            "base_account": {
                "address": address,
                "pub_key": None,
                "account_number": "0",
                "sequence": "0",
            },
            "original_vesting": [{"denom": DENOM, "amount": amount}],
            "delegated_free": [],
            "delegated_vesting": [],
            "end_time": end_time,
        },
        "start_time": "%d" % GENESIS_TIME_UNIX,
    }


def create_account_genesis_entry(address, amount):
    return {"address": address, "coins": [{"denom": DENOM, "amount": amount}]}


if __name__ == "__main__":
    # Load genesis
    with open(INPUT_GENESIS_PATH, "r") as FILE:
        genesis = json.load(FILE)

    # Modify genesis parameters
    genesis["genesis_time"] = "2023-02-13T15:00:00.000000Z"
    genesis["chain_id"] = "migaloo-1"
    genesis["app_state"]["auth"]["params"]["max_memo_characters"] = "512"
    genesis["app_state"]["crisis"]["constant_fee"]["denom"] = DENOM
    genesis["app_state"]["distribution"]["params"][
        "community_tax"
    ] = "0.100000000000000000"
    genesis["app_state"]["gov"]["deposit_params"]["min_deposit"] = [
        {"denom": DENOM, "amount": "25000" + DECIMALS}
    ]
    genesis["app_state"]["gov"]["deposit_params"]["max_deposit_period"] = "%ds" % (
        14 * SECONDS_PER_DAY
    )
    genesis["app_state"]["gov"]["voting_params"]["voting_period"] = "%ds" % (
        7 * SECONDS_PER_DAY
    )
    genesis["app_state"]["mint"]["minter"]["inflation"] = "0.040000000000000000"
    genesis["app_state"]["mint"]["params"]["mint_denom"] = DENOM
    genesis["app_state"]["mint"]["params"]["inflation_max"] = "0.040000000000000000"
    genesis["app_state"]["mint"]["params"]["inflation_min"] = "0.040000000000000000"
    genesis["app_state"]["mint"]["params"]["goal_bonded"] = "0.750000000000000000"
    genesis["app_state"]["slashing"]["params"]["signed_blocks_window"] = "10000"
    genesis["app_state"]["slashing"]["params"][
        "min_signed_per_window"
    ] = "0.100000000000000000"
    genesis["app_state"]["slashing"]["params"][
        "slash_fraction_downtime"
    ] = "0.001000000000000000"
    genesis["app_state"]["staking"]["params"]["max_validators"] = 50
    genesis["app_state"]["staking"]["params"]["bond_denom"] = DENOM
    genesis["app_state"]["staking"]["params"][
        "min_commission_rate"
    ] = "0.050000000000000000"

    # Add community pool
    genesis["app_state"]["distribution"]["fee_pool"]["community_pool"].append(
        {"denom": DENOM, "amount": "%d%s" % (COMMUNITY_POOL_AMOUNT, DECIMALS)}
    )
    genesis["app_state"]["bank"]["balances"].append(
        create_account_genesis_entry(
            address="migaloo1jv65s3grqf6v6jl3dp4t6c9t9rk99cd82tdxu3",
            amount="%d%s" % (COMMUNITY_POOL_AMOUNT, DECIMALS),
        )
    )

    # Add multi-sig
    genesis["app_state"]["bank"]["balances"].append(
        create_account_genesis_entry(
            address="migaloo1zfpqclh862kcdr8czul2k2lu2gvwk5gfg26fpx",
            amount="%d%s" % (MULTI_SIG_AMOUNT, DECIMALS),
        )
    )

    # Add vesting accounts
    for account in VESTING_ACCOUNTS:
        genesis["app_state"]["bank"]["balances"].append(
            create_account_genesis_entry(
                address=account["address"],
                amount="%d%s" % (account["amount"], DECIMALS),
            )
        )
        genesis["app_state"]["auth"]["accounts"].append(
            create_vesting_genesis_entry(
                address=account["address"],
                amount="%d%s" % (account["amount"], DECIMALS),
                end_time="%d" % account["duration"],
            )
        )

    # Check how many tokens are vested to each category of initial holders.
    for account in VESTING_ACCOUNTS:
        if account["duration"] == THREE_MONTH_UNIX:
            THREE_MONTH_AMOUNT += account["amount"]
        elif account["duration"] == TWELVE_MONTH_UNIX:
            TWELVE_MONTH_AMOUNT += account["amount"]
        elif account["duration"] == TWENTY_FOUR_MONTH_UNIX:
            TWENTY_FOUR_MONTH_AMOUNT += account["amount"]
        elif account["duration"] == THIRTY_SIX_MONTH_UNIX:
            THIRTY_SIX_MONTH_AMOUNT += account["amount"]
        else:
            ...

    # Print results with some minor adjustments
    # migaloo14kqr9fjjzl24gwfndf05wkelncm76ynkk04zjk: Substract 1M from 36 and add to 12
    # migaloo18a9m9stu3dyvewwcq9qmp85euxqcvln5mefync: Substract 0.5M from 24 and add to 12
    # migaloo1typsgj0nvketwusvaakyerjp9exr2fuxtmda9v: Substract 0.25M from 24 and add to 12
    print("Initial", INITIAL_AMOUNT / MILLION)
    print("Three Month", THREE_MONTH_AMOUNT / MILLION)
    print("Twelve Month", TWELVE_MONTH_AMOUNT / MILLION + 0.25 + 0.5)
    print("Twenty Four Month", TWENTY_FOUR_MONTH_AMOUNT / MILLION - 0.25 - 0.5)
    print("Thirty Six Month", (THIRTY_SIX_MONTH_AMOUNT) / MILLION - 1)
    print("Community Pool", COMMUNITY_POOL_AMOUNT / MILLION)
    print("Multi Sig", MULTI_SIG_AMOUNT / MILLION)
    TOTAL = (
        THREE_MONTH_AMOUNT
        + TWELVE_MONTH_AMOUNT
        + TWENTY_FOUR_MONTH_AMOUNT
        + THIRTY_SIX_MONTH_AMOUNT
        + COMMUNITY_POOL_AMOUNT
        + MULTI_SIG_AMOUNT
    )
    print("Total", (TOTAL + INITIAL_AMOUNT) / MILLION)

    # Update total supply
    genesis["app_state"]["bank"]["supply"].append(
        {
            "amount": "%d%s" % (TOTAL, DECIMALS),
            "denom": DENOM,
        }
    )
    # Store genesis
    with open(OUTPUT_GENESIS_PATH, "w") as FILE:
        json.dump(genesis, FILE, indent=2)
