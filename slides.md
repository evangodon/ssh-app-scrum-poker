
---

# Topic chosen: Mini multiplayer video game
Workshop on building video game with multiplayer functionality.
It can be 2d or 3d or even ASCII based game.

---

# But I actually built this... **A multiplayer SSH App**

- An app where uses can connect to it via the ssh protocol.
- There's no need to install anything.

---

# Easy to build with libraries from *charm.sh*

Libraries I used:
- `github.com/charmbracelet/wish` *The server library*
- `github.com/charmbracelet/bubbletea` *The TUI framework*
- `github.com/charmbracelet/lipgloss` *The style and layout library for terminals*

### See more at https://charm.sh/ or by running `ssh git.charm.sh`
---

# Examples of other ssh apps:

- `ssh clidle.duckdns.org -p 3000` 
- `ssh git.charm.sh`

---

# The app I built: Poker Scrum
- Estimate stories by having each member of the team anonymously vote the complexity
- All votes are revealed once everybody voted

---

# Demo time

**Command to connect:**
`ssh 172.105.6.132 -p 8080` 


- Make sure you terminal window is big enough
- Need to have a public key in `~/.ssh/`
- Command to generate a ssh key: `ssh-keygen -t ed25519 -C "your_email@example.com"` 

---

# The process of building it

---

## Verify this was even possible 
- No official support or docs for multi-users apps 
- Some challenges with syncing UI
- Ensure every user has their own UI 

---

## Working with BubbleTea, Creating the data model
- The "TEA" architecture: **Model, Update, View**
- Objects: **room, user**

---

## Build the UI
- The ui is one long string
- All dimensions are hard-coded
- No flex-box for layout 
- Scrollable logs

## Deploy the server somewhere for demo
- *Railway*
- *Fly.io*
- *Ngrok*
- *Linode*

---

## Q/A

