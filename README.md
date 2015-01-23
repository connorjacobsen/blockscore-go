# blockscore
Go client library for the BlockScore API.

## Installation

```bash
go get github.com/connorjacobsen/blockscore-go
```

## Usage

Import the package:

```go
import (
    "github.com/connorjacobsen/blockscore-go"
)
```

Set your BlockScore API key in an environment variable:

```bash
export BLOCKSCORE_API_KEY="your key"
```

OR you can also set your BlockScore API key in a file:

```go
blockscore.SetKey("YOUR_API_KEY")
```

*Note: Make sure to use your test key when running the test suite, otherwise the test suite will make live API requests.*

### People

Create a new person:

```go
params := blockscore.PeopleParams{
    NameFirst:          "John",
    NameMiddle:         "P",
    NameLast:           "Denver",
    DocumentType:       "ssn",
    DocumentValue:      "0000",
    BirthDay:           7,
    BirthMonth:         6,
    BirthYear:          1980,
    AddressStreet1:     "1234 Main Street",
    AddressStreet2:     "APT 12",
    AddressCity:        "Palo Alto",
    AddressSubdivision: "California",
    AddressPostalCode:  "94025",
    AddressCountryCode: "US",
    PhoneNumber:        "123-456-78910",
    IPAddress:          "127.0.0.1",
    Note:               "Hello, world",
}

person, err := blockscore.People.Create(&params)
```

Get a single existing People:

```go
person, err := blockscore.People.Retrieve("DESIRED_PERSON_ID")
```

Get a list of People:

```go
people, err := blockscore.People.List()
```

Get a list of X people, offset by Y people:

```go
people, err := blockscore.People.ListN(X, Y)
```

### Question Sets

Create a QuestionSet:

```go
// Get the Person we want a QuestionSet for.
person, err := blockscore.People.Retrieve("DESIRED_PERSON_ID")

questionSet, err := blockscore.QuestionSets.Create(person.Id)
```

Score a QuestionSet:

```go
params := ScoreParams{
    Answers: []ScoreAnswer{
        ScoreAnswer{QuestionId: 1, AnswerId: 1},
        ScoreAnswer{QuestionId: 2, AnswerId: 1},
        ScoreAnswer{QuestionId: 3, AnswerId: 1},
        ScoreAnswer{QuestionId: 4, AnswerId: 1},
        ScoreAnswer{QuestionId: 5, AnswerId: 1},
    },
}

questionSet, err := blockscore.QuestionSets.Score("QUESTION_SET_ID", &params)

if err != nil {
    // Error handling for your application.
}

// Print the QuestionSet's score.
fmt.Printf("Score: %s\n", questionSet.Score())
```

Retrieve a QuestionSet:

```go
questionSet, err := blockscore.QuestionSets.Retrieve("QUESTION_SET_ID")
```

Get a list of QuestionSets:

```go
questionSets, err := blockscore.QuestionSets.List()
```

Get a list of X QuestionSets, offset by Y QuestionSets:

```go
questionSets, err := blockscore.QuestionSets.ListN(X, Y)
```

### Companies

Create a new company:

```go
params := blockscore.CompanyParams{
    EntityName:               "BlockScore",
    TaxId:                    "123410000",
    IncorporationState:       "DE",
    IncorporationCountryCode: "US",
    IncorporationType:        "corporation",
    IncorporationDay:         23,
    IncorporationMonth:       8,
    IncorporationYear:        1993,
    Dbas:                     "BitRemit",
    RegistrationNumber:       "123123123",
    Email:                    "test@example.com",
    Url:                      "https://blockscore.com",
    PhoneNumber:              "6505555555",
    IPAddress:                "67.160.8.182",
    Note:                     "Much wow",
    AddressStreet1:           "12345 Main Street",
    AddressStreet2:           "#12",
    AddressCity:              "Palo Alto",
    AddressSubdivision:       "CA",
    AddressPostalCode:        "94025",
    AddressCountryCode:       "US",
}

company, err := blockscore.Companies.Create(&params)
```

Get a single existing Company:

```go
company, err := blockscore.Companies.Retrieve("DESIRED_COMPANY_ID")
```

Get a list of Companies:

```go
companies, err := blockscore.Companies.List()
```

Get a list of X companies, offset by Y companies:

```go
companies, err := blockscore.Companies.ListN(X, Y)
```

### Candidates

Create a Candidate:

```go
params := CandidateParams{
    NameFirst:          "John",
    NameLast:           "Brendekamp",
    Note:               "12341234",
    Ssn:                "0001",
    DateOfBirth:        "1940-08-11",
    AddressStreet1:     "1 Infinite Loop",
    AddressCity:        "Harare",
    AddressCountryCode: "ZW",
}

candidate, err := blockscore.Candidates.Create(&params)
```

Retrieve a Candidate:

```go
candidate, err := blockscore.Candidates.Retrieve("DESIRED_CANDIDATE_ID")
```

Update a Candidate:

```go
params := CandidateParams{
    Passport: "123456789",
}

candidate, err := blockscore.Candidate.Update("DESIRED_CANDIDATE_ID", &params)
```

Delete a Candidate:

```go
candidate, err := blockscore.Candidate.Delete("DESIRED_CANDIDATE_ID")
```

Get a list of Candidates:

```go
candidates, err := blockscore.Candidates.List()
```

Get a list of X Candidates, offset by Y Candidates:

```go
candidates, err := blockscore.Candidates.ListN(X, Y)
```

View the Revision History of a Candidate:

```go
candidateHistory, err := blockscore.Candidates.History("CANDIDATE_ID")
```

View past Watchlist Hits of a Candidate:

```go
candidateHits, err := blockscore.Candidates.Hits("CANDIDATE_ID")
```

### Watchlists

Search a Watchlist:

```go
watchlistParams := WatchlistParams{
    CandidateId: "DESIRED_CANDIDATE_ID",
    MatchType:   "person" // search for individuals only
}

search, err := blockscore.Watchlists.Search(&watchlistParams)
```
