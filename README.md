# CLI

CLI app to test different components of the Shizima app

**Prerequisites:**

- _go 1.21+_

### How to use

#### Installation

after you installed a [version of go](https://go.dev/doc/install)

1. Clone the repo  

```sh
git clone https://gitlab.isima.fr/shizima/clients/cli.git
```

2. Run the following

```sh
go mod tidy
go mod download 
```

3. Set config

You can change the app config depending on your terminal theme. \
Open the already existing `shizima.config` file or create it on the project's root folder.

It should looks like so :

```json
{
    "theme": "light"
}
```

The possible values or `dark` or `light`

#### Run it

After your set up is finished you can start the app directly from root folder

```sh
go run .
```

You will then have to choose the `.wav` file you want to try to recognize. \
Further explanation on how to use the app are directly displayed inside the app.
