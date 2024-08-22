# Design

This is a basic set of design decisions so I don't forget.

- Core library of voting protocols that persist the results to a SQLite database, wrapped by either a
CLI or web service.

```
erDiagram
    ELECTION {
        int election_id PK
        string name
        date start_date
        date end_date
        int voting_system_id FK
    }
    VOTING_SYSTEM {
        int voting_system_id PK
        string name
        string description
    }
    CANDIDATE {
        int candidate_id PK
        string name
        string party
        int election_id FK
    }
    VOTER {
        int voter_id PK
        string name
        string address
        date date_of_birth
    }
    BALLOT {
        int ballot_id PK
        int election_id FK
        int voter_id FK
        datetime timestamp
    }
    VOTE {
        int vote_id PK
        int ballot_id FK
        int candidate_id FK
        int rank
        boolean approved
    }
    ELECTION_RESULT {
        int result_id PK
        int election_id FK
        int candidate_id FK
        int total_votes
        float vote_percentage
        int rank
    }

    ELECTION ||--o{ CANDIDATE : "has"
    ELECTION ||--|| VOTING_SYSTEM : "uses"
    ELECTION ||--o{ BALLOT : "contains"
    VOTER ||--o{ BALLOT : "casts"
    BALLOT ||--o{ VOTE : "includes"
    VOTE }o--|| CANDIDATE : "for"
    ELECTION ||--o{ ELECTION_RESULT : "produces"
    CANDIDATE ||--o{ ELECTION_RESULT : "receives"
```

# Implementation

- If possible, don't import new packages. Of course, if it's an extra convienent package use it.
- Err on the side of global utilities that other parts of the library can use
- Configuration is file-driven in the home directory. Like `neovim`.
- The core data model should be protocol agnostic and bashed into whatever data structure makes the most
sense for calculating results.
