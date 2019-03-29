database stuff
users
update for login and password hash? later.
users and games have a many:many association.
1:many associaton with games
1:1 with bets
game
also have a 1:many associaton with games
1:many with bets
messages
bets
reference users

scripting

- seed script for DB
- up/down for migrations
- some sort of CI/CD stuff as well
- dockerize this shit.
  - Is there a way to do hot building in go with a dockerized container? Why? Why the fuck not?
- I really need to learn GCP. and I really don't want to learn GCP. Ugh.
