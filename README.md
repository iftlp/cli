# CLI

CLI app to test different components of the Shizima app

**Prerequisites:**

- _go 1.21+_

## How to use

> [!WARNING]  
> Current states isn't advanced at all
> Things may change often in a quick time span

### Installation

1. clone the repo  
```sh
git clone https://gitlab.isima.fr/shizima/clients/cli.git
```

2. run the following
```sh
go mod tidy
go mod download 
```

### Run it

1. To use the program add the **.mp3** musics you want to decode

2. Call it like so `test_sound_1.mp3` 

3. Run it
```sh
go run .
```