# aisleriot
[![Go Report Card](https://goreportcard.com/badge/github.com/philhanna/aisleriot)][idGoReportCard]
[![PkgGoDev](https://pkg.go.dev/badge/github.com/philhanna/aisleriot)][idPkgGoDev]

Displays statistics for Linux Aisleriot card games

## Usage
```
Usage: arstats [OPTION]...

Shows statistics for Aisleriot games played by the current user.

Options:
  -g, --game=GAMENAME	Name of game for which statistics are desired
                        (Default is most recently played game)
  -l, --list            List the names of all games played

Output includes:
  - Game name
  - Number of wins
  - Number of losses
  - Total games played
  - Best time
  - Average time
  - Worst time
  - Winning percentage
  - Number of wins to next higher percent
  - Number of losses to next lower percent
```
## Installation
```bash
cd /tmp
git clone git@github.com:philhanna/aisleriot.git
cd aisleriot
go install cmd/arstats.go
```

[idGoReportCard]: https://goreportcard.com/report/github.com/philhanna/aisleriot
[idPkgGoDev]: https://pkg.go.dev/github.com/philhanna/aisleriot
