# Installation

- Install ganache-cli `npm install ganache-cli@latest --global`
- Install Docker
- Install Go 1.6. Make sure GOPATH, GOROOT are set and go module is turned on.

# Errors you might face.
- A few depdencies of this project are private repo. You need to config git to use `git` instead of `https` when downloading repo.
Run the following command to replace `https` by `git`.

```
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

To confirm, do `more ~/.gitconfig` and make sure you see the following:

```
[url "git@github.com:"]
	insteadOf = https://github.com/
```

- In case the above does not work.
  Clone the individual repos and add replace statements in the go.mod for each library that failed.
  For Example
  
```
replace github.com/sisu-network/deyes => ../deyes
```

# Generate Mock structs

- Install mockgen to generate mock structs from interfaces:

```bash
go install github.com/golang/mock/mockgen@v1.6.0
```

# From the sisu root folder

## Run Mysql
- Run Mysql - `make dev-mysql-up`
In addition, you need to install mysql locally with the following config:

```
host: localhost
port: 3306
username: root
password: password
```

You can run `make dev-mysql-up` (make sure you installed Docker) to initialize databases with default settings

## Create env file

```
cp .env.dev .env
```

## Run the sisu app.

Install all modules.
```
go mod tidy
```

Generate config file and genesis for local sisu app:

```
./scripts/gen_localnet.sh
```

To run the app on localhost:

```
./scripts/run_local.sh
```

You can now deploy ETH transaction on ganache1 at port 7545.

---
### For Running locally without TSS. (Keep it enabled though.)
---

Disable the TSS component in the file `~/.sisu/main/config/sisu.toml` by changing the default tss settings to false.

```
[tss]
enable = false
```

---
### Running with TSS (single node) one command line
---

You need to enable tss in `~/.sisu/main/config/sisu.toml` in order to run full Sisu node.

```
[tss]
enable = true
```

You will need at least 5 different terminal tabs to run a full Sisu nodes: 2 tabs for ganache-cli, 1 for dheart, 1 for deyes, 1 for sisu.

## Run ganache-cli
---

Download ganache-cli (make sure you have version 6.x) and runs the following commands on 2 different terminals:

```
ganache-cli --accounts 10 --blockTime 3 --port 7545 --defaultBalanceEther 100000 --networkId 189985 --chainId 189985 --mnemonic "draft attract behave allow rib raise puzzle frost neck curtain gentle bless letter parrot hold century diet budget paper fetch hat vanish wonder maximum"
```

```
ganache-cli --accounts 10 --blockTime 3 --port 8545 --defaultBalanceEther 100000 --networkId 189986 --chainId 189986 --mnemonic "draft attract behave allow rib raise puzzle frost neck curtain gentle bless letter parrot hold century diet budget paper fetch hat vanish wonder maximum"
```

These commands create a simulated blockchain on port 7545 and 8545.

## Run dheart and deyes
---

Runs dheart and deyes using their instruction in the repos.

## Run Sisu
---

```
./scripts/run_local.sh
```

## Interact with Sisu
---

Sisu has a number of commands for you to interact with the network. At the root folder, type `./sisu` to see a list of command options. Most of them are default commands from Cosmos SDK. We are only interested in `./sisu dev` command for local dev.

After the start of Sisu, wait until the sisu log has the following message similar to the following:

```
adding watcher address 0x1D156a3e1356b58733305e670D61018001997f6E
```

Please note that `0x1D156a3e1356b58733305e670D61018001997f6E` could be replaced by another address. It's the address of the signing key that Sisu network controls. We need to fund that address with some ETH so that Sisu can have balance to deploy contracts and invoke function calls.

Alternatively, you can connect to mysql and see the address of the generate key in the `keygen` table. Mysql Workbench is good GUI tool to view and edit myqsl data.

```
./sisu dev fund-account ganache1 7545 ganache2 8545 10
```

This commands will fund Sisu's signing key account with 10 ethereum. The last number in the command is the amount you want to fund.

Waits for 10 seconds for the transaction to finalized and for Sisu to deploy its gateway contract.

Now you can start transferring ERC20 tokens out of a blockchain. You need a deployed ERC20 contract on one of the ganache chain. You can deploy using separate service or use Sisu command line:

```
./sisu dev deploy-erc20 7545
```

Waits for a few seconds and you will see the address the newly deploy ERC20 contract. You can transfer token to second dev chain using this command format:

```
./sisu dev transfer-out [ContractType] [FromChain] [TokenAddress] [ToChain] [RecipientAddress]
```

For example

```
./sisu dev transfer-out erc20 ganache1 7545 0xf0D676183dD5ae6b370adDdbE770235F23546f9d 8545 0xE8382821BD8a0F9380D88e2c5c33bc89Df17E466
```

Waits a few seconds for the transaction to complete. Afterward, you can query the asset balance on the destination chain:

```
./sisu dev query [ContractType] [chain] [port] [AssetId] [AccountAddress]

./sisu dev query erc20 ganache1 7545 ganache1__0x3DeaCe7E9C8b6ee632bb71663315d6330914f915 0xE8382821BD8a0F9380D88e2c5c33bc89Df17E466
```

## Running with TSS (multi nodes) using docker

This section describes how to run sisu (multiple nodes) using docker. This simulates a real network operation using various services in docker container. Each node 3 instances: sisu, dheart, deyes services in the container in addition to 2 ganaches (to simulate at least 2 chains) and a mysql instance.

Because of high number of services in the docker container (3 x n + 3 where n is the number of validators in the network), it's recommended that you run TSS with single node (in the following section) for faster local development speed and do extensive unit tests. You can use docker for testing full multi nodes with TSS enabled.

## Build containers

You need to build sisu, dheart, deyes, ganache images in their corresponding repos. Assume that you save dheart, deyes in the same directory level with sisu.

```
cd ../dheart
docker build -t dheart .
cd ../deyes
docker build -t deyes .
docker build -f Dockerfile-ganache -t ganache-cli .
cd ../sisu
docker build -t sisu .
```

## Generate data
Next, generate docker-compose file and data for all nodes in the network.
```
go build -o ./sisu cmd/sisud/main.go
./sisu local-docker --v 2 --output-dir ./output
```

## Run docker-compose
Go the generate folder and start all services using docker-compose
```
cd output
docker-compose up -d
```

You can check logs for each services using `docker-compose logs`. For example:
```
docker-compose logs -f sisu0
```

The ganache cli instances expose 2 ports to the host machine: 7545 and 7546. You can interact with the blockchain through this port.

To fund accounts generated by Sisu network, use `fund-account command`
```
./sisu dev fund-account ganache1 7545 ganache2 7546 10
```

To create an ERC20 contract:
```
./sisu dev deploy-erc20 7545
```
This will deploy an ERC20 contract and its address is shown in the console.

To transfer token cross chain:
```
./sisu dev transfer-out erc20 ganache1 7545 [ERC20 contract address] ganache2 [Recipient address]
```

To close all services:
```
docker-compose down
```

NOTE: if you get into bad data state or you want to reset the blockchain, simply run the local-docker command again
```
./sisu local-docker --v 2 --output-dir ./output
```

