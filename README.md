# sorta

**sorta** is a simple CLI tool to sort files based on MIME type

## Usage

```

  A simple CLI tool to sort files based on MIME type

  USAGE


    sorta [command] [directory] [--flags]


  COMMANDS

    completion [command]  Generate the autocompletion script for the specified shell
    help [command]        Help about any command

  FLAGS

    -d --dry              Dry run; don't actually move any files
    -h --help             Help for sorta
    -v --version          Version for sorta
```

## Installation

### with Nix

Simply add the repo to your flake inputs...

```nix
inputs = {
  sorta.url = "github:indium114/sorta";
};
```

...and pass it into your `environment.systemPackages`...

```nix
environment.systemPackages = [
  inputs.sorta.packages.${pkgs.stdenv.hostPlatform.system}.sorta
];
```

### with Go

Simply run the following command:

```shell
go install github.com/indium114/sorta@latest
```
