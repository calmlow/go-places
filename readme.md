# Go-Places

Show your places (folders and files) and go there.

It was used to be called reposelector, now just "Places", probably legacy left

## Purpose

There are many "terminal tools" out there that does this, automatically keep track of most visited "places" and can
generate some kind of drop down for you to easily navigate there. Probably a plugin in zsh...

But the files I visit are many times the same and I want to be in control of how this is handled. Also to implement
this in Go and TView for educational purposes, along with the setup in Bash (bind) & Zsh (bindkey)

## SETUP

Not part of installation.. but this is how I initialized the project.

```bash
# installed version of go: 1.23.4
go mod init github.com/calmlow/go-places
go get github.com/rivo/tview
go get gopkg.in/yaml.v3
```

## build and install

```bash
go build -o go-places main.go
# or
make build
```

## Bash and Zsh Setup

To make the tool useful, I have the following bind setup. Also to get some icons in the tool I have
my terminal setup with [nerd fonts](https://www.nerdfonts.com/cheat-sheet).


```bash
# bash
ANSI_ORANGE='\033[0;33m'
ANSI_PURPLE='\033[0;35m'
NO_C='\x1b[0m'
N_HOME='\ue780'


show_repo_selector_tool() {
  ~/bin/go-places # or other path to the executable
  STATUS_OF_LAST_CMD=$?
  if [ $STATUS_OF_LAST_CMD -eq 0 ]; then
    TMP_DIR1=$(cat /tmp/selected-repo.txt.tmp)
    TMP_DIR2=$(echo $TMP_DIR1 | cut -d'=' -f 2)
    echo -e "Going to ${ANSI_ORANGE}${N_HOME} ${ANSI_PURPLE}${TMP_DIR2} ${NO_C}"
    if [[ -d "$TMP_DIR2" ]]; then
      cd $TMP_DIR2
    else
      # assume text file..open with vscode
      code $TMP_DIR2
    fi
  fi
}

# Shift+ArrowLeft & Alt+R
bind -x '"\C-[r": "show_repo_selector_tool"'
bind -x '"\e[1;2D": show_repo_selector_tool' # Shift + ArrowLeft
# if your terminal supports some kind of select mode - maybe you don't want to bind the shift+left key.

# if you want to run it as a command instead of a bind shortcut, you need to add:
alias go-places="$DIR_TO_BIN/go-places && source /tmp/selected-repo.txt.tmp && cd \$TMP_GO_SELECTED_REPO"
```

```zsh
# zsh
show_repo_selector_tool() {
  ~/bin/go-places # or other path to the executable
  STATUS=$?
  if [ $STATUS -eq 0 ]; then
    TMP_DIR1=$(cat /tmp/selected-repo.txt.tmp)
    TMP_DIR2=$(echo $TMP_DIR1 | cut -d'=' -f 2)
    echo -e "Going to $ANSI_ORANGE\ue780 $ANSI_PURPLE$TMP_DIR2 $NO_C"
    if [[ -d "$TMP_DIR2" ]]; then
      cd $TMP_DIR2
    else
      code $TMP_DIR2
    fi
  fi
  zle reset-prompt
}

zle -N show_repo_selector_tool

# Shift+ArrowLeft & Alt+R
bindkey '\e[1;2D' show_repo_selector_tool
bindkey '^[r' show_repo_selector_tool
```

## Dependencies

A section about the dependencies used

### TVIEW

https://github.com/rivo/tview/wiki

To install tview run:

```console
go get github.com/rivo/tview
```

## Local Config File

Create a config file in your home directory:

```bash
mkdir -p ~/.config/go-places && touch ~/.config/go-places/go-places-config.yaml
```

The structure of the config file is:

```yaml
places:
  - name: my-important-repo1
    path: ~/my-repos/my-important-repo1
  - name: my-important-repo1
    path: ~/my-repos/my-important-repo1
  - name: my-important-repo2
    path: ~/my-repos/my-important-repo2
```

```bash
# populate above example into the config file you just created
sed -n '123,129p' ~/r/go-places/readme.md > ~/.config/go-places/go-places-config.yaml
```

See the .json schema in ./asssets 

## Testing

```bash
go test -v ./...
# or
go test -v ./internal/config
go test -v ./internal/config -run=Test_should_read_local_config_file
# or
make test
```

Todo: 
Use https://github.com/stretchr/testify/
