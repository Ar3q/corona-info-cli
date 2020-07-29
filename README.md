# corona-info-cli

## Install

For using corona-info-cli from terminal in any directory:

```bash
go install
# run program
corona-info-cli
```

For using only from directory where source code lives:

```bash
go build
# run program
./corona-info-cli
```

## Usage

Getting table with:
 - all countries: `corona-info-cli`
 - one specified country: `corona-info-cli -c poland` or `corona-info-cli -c pl`
 - first x countries: `corona-info-cli -t 5`


