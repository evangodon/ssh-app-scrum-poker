
# TODO

- [X] Make selected card more apparent


- [ ] Mouse support, see https://github.com/maaslalani/gambit/blob/main/game/game.go *Didn't work*
- [ ] Clean up vote type, int/string, create enum
- [ ] Fix the selection of an story point 


# Done
- [X] Create user model
  -- [X] Add user on new connection
  -- [X] Remove user on closed connection
- [X] Figure out how to have each user see their own ui
    -- Create bubble tea model on server start
    -- use program.Send(msg) to update ui from outside the tea program
- [X] Use slice for room.users instead of map
- [X] how to expose the local ssh server
   -- `ngrok tcp <ssh-port>` and then `ssh 2.tcp.ngrok.io -p <port>`
- [X] Figure out how to reconnect when connection drops on server restart
- [-] Use bubbles/viewport for logs section *Didn't work*
- [X] Fix countdown
- [X] Add keyboard shortcuts at bottom
- [X] Fix: Logs disapears on every vote
- [X] Sort breakdown of votes and remove counts of zero, use https://github.com/Evertras/bubble-table
- [X] Authorization for viewing votes / reset room
    -- [X] Add isHost field to user
    -- [X] Only enable view votes when everybody has voted
- [X] Manage height of every section
- [ ] Add background color to each user block *Nevermind*
- [X] handle ties in votes table
- [X] user circle for story points in vote table
- [X] trim username if too long 
- [X] Update logs section, add divider line
