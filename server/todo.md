

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
