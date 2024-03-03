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

cric player --help // search term, stats of player

cric ranking  // icc ranking of teams and players (Done)
--help
--batting default
--bowling
--all-rounders
--teams

Project Scope
1. Focused on men's international cricket
2. Subcommands work

1. Ranking - Men's and team ranking across batting, bowling and all rounders
--help
--batting default
--bowling
--all-rounders
--teams

2. Player - with search term gives list of players and id, and next enter id to explore stats of that player
flags
--help
--all
Player information show format
Jos Buttler - England - WK-Batsman
Batting Summary
        M   Runs    HS  100 200 50  
Test    
ODI
T20
IPL

Bowling Summary
        M   Wkts    Best(BBM) Econ  5W  10W
Test
ODI
T20
IPL

Career Information


List Down Tasks and Record Time spent on each task

1.