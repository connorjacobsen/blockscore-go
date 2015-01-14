# blockscore
Go client library for the BlockScore API.

Note:
* Some of the APIs may not be idiomatic Go, this is in order to keep those APIs as similar to existing BlockScore client libraries as possible.
* Some things will be grammatically incorrect (i.e., "get a People" rather than "get a Person"), this is in order to remain consistent with BlockScore API endpoint names, for which things make sense.

## Usage

Set your BlockScore API key in an environment variable:

```bash
export BLOCKSCORE_API_KEY="your key"
```

*Note: Make sure to use your test key when running the test suite, otherwise the test suite will make live API requests.*

### Client Creation

```go
client = blockscore.Client.new("YOUR_BLOCKSCORE_API_KEY")
```

### People

Create a new person:

```go
params := blockscore.PeopleParams{
    // Fill this in later
}

person, err := client.CreatePeople(&params)
```

Get a single existing People:

```go
person, err = client.RetrievePeople("DESIRED_PERSON_ID")
```

Get a list of People:

```go
```