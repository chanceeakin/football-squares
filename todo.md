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

- database
  - THINK WE'RE DONE!

* error handler

  - figure out a way to deal with inputs missing their proper response bodies
  - validate input in response body
  - https://husobee.github.io/golang/validation/2016/01/08/input-validation.html

* session storage

  - redis

* endpoints
  - login/logout
    - JWT?
  - users
    - password
  - games
    - SET game admin(s)
    - put
    - join a game
    - leave a game
    - open a game
    - close a game
  - bets
    - getBets
    - postaBet
* testing strategy for golang
  - do some research

- frontend
  - reimagine it in yew. because why not
