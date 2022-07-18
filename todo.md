
# TODO
- [X] Create user model
  -- [X] Add user on new connection
  -- [X] Remove user on new connection

- [X] Figure out how to have each user see their own ui
    -- Create bubble tea model on server start
    -- use program.Send(msg) to update ui from outside the tea program

- [X] Use slice for room.users instead of map

- [ ] how to expose the local ssh server
   -- `ngrok tcp <ssh-port>` works


- [X] Figure out how to reconnect when connection drops on server restart

- [ ] Mouse support, see https://github.com/maaslalani/gambit/blob/main/game/game.go
- [ ] Use bubbles/viewport for logs section


*Short ones*
- [ ] Sort breakdown of votes and remove counts of zero, use https://github.com/Evertras/bubble-table
- [ ] Fix the selection of an story point 
- [ ] Clean up vote type, int/string, create enum
