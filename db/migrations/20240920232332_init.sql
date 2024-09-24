-- migrate:up
-- Create the database tables

CREATE TABLE IF NOT EXISTS election (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS candidate (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    election_id TEXT,
    FOREIGN KEY (election_id) REFERENCES election(election_id)
);

CREATE TABLE IF NOT EXISTS voter (
    id TEXT PRIMARY KEY,
    zip TEXT
);

CREATE TABLE IF NOT EXISTS vote (
    id TEXT PRIMARY KEY,
    candidate_id TEXT NOT NULL,
    rank INTEGER,
    timestamp DATETIME NOT NULL,
    voter_id TEXT NOT NULL,
    FOREIGN KEY (candidate_id) REFERENCES candidate(candidate_id)
);

CREATE TABLE IF NOT EXISTS election_result (
    id TEXT PRIMARY KEY,
    election_id TEXT,
    candidate_id TEXT,
    total_votes INTEGER,
    vote_percentage REAL,
    rank INTEGER,
    FOREIGN KEY (election_id) REFERENCES election(election_id),
    FOREIGN KEY (candidate_id) REFERENCES candidate(candidate_id)
);

-- migrate:down

DROP TABLE IF EXISTS election_result;
DROP TABLE IF EXISTS vote;
DROP TABLE IF EXISTS voter;
DROP TABLE IF EXISTS candidate;
DROP TABLE IF EXISTS election;
DROP TABLE IF EXISTS voting_system;
