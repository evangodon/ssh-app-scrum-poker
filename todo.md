
# TODO

- [ ] Authorization for viewing votes / reset room
    -- [ ] Only enable view votes when everybody has voted
- [ ] Fix log viewport
- [ ] Mouse support, see https://github.com/maaslalani/gambit/blob/main/game/game.go
- [ ] Add keyboard shortcuts at bottom

*Short ones*
- [ ] Sort breakdown of votes and remove counts of zero, use https://github.com/Evertras/bubble-table
- [ ] Fix the selection of an story point 
- [ ] Clean up vote type, int/string, create enum




# Done
- [X] Create user model
  -- [X] Add user on new connection
  -- [X] Remove user on closed connection
- [X] Figure out how to have each user see their own ui
    -- Create bubble tea model on server start
    -- use program.Send(msg) to update ui from outside the tea program
- [X] Use slice for room.users instead of map
- [X] how to expose the local ssh server
   -- `ngrok tcp <ssh-port>` works
- [X] Figure out how to reconnect when connection drops on server restart
- [-] Use bubbles/viewport for logs section *Didn't work*
