---
sidebar_position: 4
title: Join Mainnet
---

# Join Enigma Mainnet

## Hardware Specification
Validators should expect to provision one or more data center locations with redundant power, networking, firewalls, HSMs and servers.

We initially recommend this minimum hardware specifications and they might rise as network usage increases.

```
4+ vCPU x64 2.0+ GHz
32+ GB RAM
1TB+ SSD
```

## Install enigmad and peggo

```bash
wget https://github.com/EnigmasLab/enigma-chain-releases/releases/download/v1.9.0-1673640888/linux-amd64.zip
unzip linux-amd64.zip
sudo mv peggo /usr/bin
sudo mv enigmad /usr/bin
sudo mv libwasmvm.x86_64.so /usr/lib 
```

## Initialize a new Enigma Chain node

Before actually running the Enigma Chain node, we need to initialize the chain, and most importantly its genesis file.

```
# The argument <moniker> is the custom username of your node, it should be human-readable.
export MONIKER=<moniker>
# the Enigma Chain has a chain-id of "enigma-1"
enigmad init $MONIKER --chain-id enigma-1
```

Running this command will create `enigmad` default configuration files at `~/.enigmad`.

## Prepare configuration to join Mainnet

You should now update the default configuration with the Mainnet's genesis file and application config file, as well as configure your persistent peers with a seed node.  

```bash
git clone https://github.com/EnigmasLab/mainnet-config

# copy genesis file to config directory
cp mainnet-config/10001/genesis.json ~/.enigmad/config/genesis.json

# copy config file to config directory
cp mainnet-config/10001/app.toml  ~/.enigmad/config/app.toml
```

You can also run verify the checksum of the genesis checksum - 573b89727e42b41d43156cd6605c0c8ad4a1ce16d9aad1e1604b02864015d528
```bash
sha256sum ~/.enigmad/config/genesis.json
```

Then open update the persistent_peers field present in ~/.enigmad/config/config.toml with the contents of mainnet-config/10001/seeds.txt and update the `timeout_commit` to `300ms`.
```bash
cat mainnet-config/10001/seeds.txt
nano ~/.enigmad/config/config.toml
```

## Configure systemd service for enigmad

Edit the config at `/etc/systemd/system/enigmad.service`:
```bash
[Unit]
  Description=enigmad

[Service]
  WorkingDirectory=/usr/bin
  ExecStart=/bin/bash -c '/usr/bin/enigmad --log-level=error start'
  Type=simple
  Restart=always
  RestartSec=5
  User=root

[Install]
  WantedBy=multi-user.target
```

Starting and restarting the systemd service
```bash
sudo systemctl daemon-reload
sudo systemctl restart enigmad
sudo systemctl status enigmad

# enable start on system boot
sudo systemctl enable enigmad

# To check Logs
journalctl -u enigmad -f
```

## Sync with the network

### Option 1. State-Sync

You can use state-sync to join the network by following the below instructions.

```bash
#!/bin/bash
sudo systemctl stop enigmad
sudo enigmad tendermint unsafe-reset-all --home ~/.enigmad
CUR_HEIGHT=$(curl -sS https://tm.enigma.network/block | jq .result.block.header.height | tr -d '"')
SNAPSHOT_INTERVAL=1000
RPC_SERVERS="23d0eea9bb42316ff5ea2f8b4cd8475ef3f35209\@65.109.36.70:11750,38c18461209694e1f667ff2c8636ba827cc01c86\@176.9.143.252:11750,4f9025feca44211eddc26cd983372114947b2e85\@176.9.140.49:11750,c98bb1b889ddb58b46e4ad3726c1382d37cd5609\@65.109.51.80:11750,f9ae40fb4a37b63bea573cc0509b4a63baa1a37a\@15.235.144.80:11750,7f3473ddab10322b63789acb4ac58647929111ba\@15.235.13.116:11750"
TRUST_HEIGHT=$(( CUR_HEIGHT - SNAPSHOT_INTERVAL ))
TRUSTED_HASH=$(curl -sS https://tm.enigma.network/block?height=$TRUST_HEIGHT | jq .result.block_id.hash)
perl -i -pe 's|enable = false|enable = true|g' ~/.enigmad/config/config.toml
perl -i -pe 's|rpc_servers = ".*?"|rpc_servers = "'$RPC_SERVERS'"|g' ~/.enigmad/config/config.toml
perl -i -pe 's/^trust_height = \d+/trust_height = '$TRUST_HEIGHT'/' ~/.enigmad/config/config.toml
perl -i -pe 's/^trust_hash = ".*?"/trust_hash = '$TRUSTED_HASH'/' ~/.enigmad/config/config.toml
sudo systemctl start enigmad
```

### Option 2. Snapshots

You can find pruned snapshots on:

1. [ChainLayer](https://quicksync.io/networks/enigma.html).
2. [Polkachu](https://polkachu.com/tendermint_snapshots/enigma).
3. [HighStakes](https://tools.highstakes.ch/files/enigma.tar.gz).
4. [AutoStake](http://snapshots.autostake.net/enigma-1/).

**Archival** (>7TB)
```bash
aws s3 sync --acl public-read --delete --no-sign-request s3://enigma-snapshots/mainnet/enigmad/daily/data $HOME/.enigmad/data
aws s3 sync --acl public-read --delete --no-sign-request s3://enigma-snapshots/testnet/enigmad/daily/wasm $HOME/.enigmad/wasm
```


### Support

For any further questions, you can always connect with the Enigma Team via Discord, Telegram, and email.

[Discord](https://discord.gg/enigma)

[Telegram](https://t.me/joinenigma)

[E-mail](mailto: contact@injectivelabs.org)
