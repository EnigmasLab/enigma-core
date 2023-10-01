---
sidebar_position: 3
title: Join Testnet
---

# Join the Network

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
wget https://github.com/EnigmasLab/testnet/releases/download/v1.10.2-1678712142/linux-amd64.zip
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
# the Enigma Chain has a chain-id of "enigma-888"
enigmad init $MONIKER --chain-id enigma-888
```

Running this command will create `enigmad` default configuration files at `~/.enigmad`.

## Prepare configuration to join Testnet

You should now update the default configuration with the Testnet's genesis file and application config file, as well as configure your persistent peers with a seed node.  

```bash
git clone https://github.com/EnigmasLab/testnet.git

# copy genesis file to config directory
aws s3 cp s3://enigma-snapshots/testnet/genesis.json . --no-sign-request
mv genesis.json ~/.enigmad/config/

# copy config file to config directory
cp testnet/corfu/70001/app.toml  ~/.enigmad/config/app.toml
cp testnet/corfu/70001/config.toml ~/.enigmad/config/config.toml
```

You can also run verify the checksum of the genesis checksum - a4abe4e1f5511d4c2f821c1c05ecb44b493eec185c0eec13b1dcd03d36e1a779
```bash
sha256sum ~/.enigmad/config/genesis.json
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

```bash
sudo systemctl stop enigmad
aws s3 sync --acl public-read --no-sign-request --delete s3://enigma-snapshots/testnet/enigmad/data $HOME/.enigmad/data
aws s3 sync --acl public-read --no-sign-request --delete s3://enigma-snapshots/testnet/enigmad/wasm $HOME/.enigmad/wasm
sudo systemctl start enigmad
```


### Support

For any further questions, you can always connect with the Enigma Team via Discord, Telegram, and email.

[Discord](https://discord.gg/enigma)
[Telegram](https://t.me/joinenigma)
[E-mail](mailto:contact@injectivelabs.org)
