# Installation

## Install latest binaries

Download suitable binary from [realizes page](https://github.com/tariel-x/anzer/releases).
For example for linux:

```bash
sudo curl -L "https://github.com/tariel-x/anzer/releases/download/latest/linux_amd64_anzer" -o /usr/local/bin/anzer
sudo chmod +x /usr/local/bin/anzer
anzer help
```

For OS X:

```bash
sudo curl -L "https://github.com/tariel-x/anzer/releases/download/latest/darwin_amd64_anzer" -o /usr/local/bin/anzer
sudo chmod +x /usr/local/bin/anzer
anzer help
```

## Install from sources

```bash
git clone https://github.com/tariel-x/anzer.git
cd anzer
go install
anzer help
```