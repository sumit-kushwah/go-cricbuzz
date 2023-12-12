## Cricbuzz Things to Implement
- [] current and future series
- [] 5 matches by team, last, live, and upcoming

## Go Depedencies
Colly
Cobra
Viper

## Installation and build exectuable file
git clone <url>
go build -o ./cric

## move exectuble file into /usr/local/bin/

## Try out commands
cric -h


`commands`

cric --help
cric match
--help
--team
--live default
--recent
--count 5(default)
--upcoming
--international(default)
--league
--domestic
--women

cric scorecard <match-id> 
--help
--first
--second
First Innings(Team 1 Name)
Batting
Bowling
Second Innings(Team 2 Name)
Batting
Bowling

cric series
--help
--count 5
--team

cric team <name> // general information about team
--help

cric player --help // search term, stats of player

cric ranking  // icc ranking of teams and players
--help
--men
--batting
--bowling
--all-rounders
--teams
--women