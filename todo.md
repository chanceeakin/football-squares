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

## CI/CD

- up/down for migrations
  - when you actually have migrations
- some sort of CI/CD stuff as well
  - get your app engine set up, maybe?
- I really need to learn GCP. and I really don't want to learn GCP. Ugh.

## Code

- endpoints
  - messages
    - getMessagesByGame
  - bets
    - getBets
    - postaBet
- testing strategy for golang
  - do some research
