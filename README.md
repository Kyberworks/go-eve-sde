# go-eve-sde

This repository contains Go models generated from the EVE Online Static Data Export (SDE). The models are automatically updated nightly using a GitHub Action that downloads the latest SDE from Fuzzwork and generates Go structs using `sql2struct`.

## Usage

Add this module to your Go project:

```bash
go get github.com/Kyberworks/go-eve-sde@v1.23456  # Replace with desired build version
```

Then import and use the models:

```go
import "github.com/Kyberworks/go-eve-sde/models"

func main() {
    // Example usage
    var item models.InvType
    // ... populate from JSON or database
}
```

## Data Source

The SDE data is sourced from the official CCP EVE Online Static Data Export in JSON Lines format available at https://developers.eveonline.com/static-data/eve-online-static-data-latest-jsonl.zip.

## Automation

The repository is updated nightly at 2 AM UTC via a GitHub Action that:

1. Checks the latest SDE build number from https://developers.eveonline.com/static-data/tranquility/latest.jsonl
2. If the build number has changed, downloads the latest official SDE JSONL zip from CCP
3. Extracts the JSONL data files
4. Generates Go structs using `quicktype` from sample JSON objects
5. Commits any changes with the new build number

## Releases

Each update creates a new GitHub release tagged with the SDE build number (e.g., `v12345`). You can reference specific versions of the models by using the appropriate release tag in your Go module dependencies.
