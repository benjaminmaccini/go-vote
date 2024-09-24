CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE election (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL
);
CREATE TABLE candidate (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    election_id TEXT,
    FOREIGN KEY (election_id) REFERENCES election(election_id)
);
CREATE TABLE voter (
    id TEXT PRIMARY KEY,
    zip TEXT
);
CREATE TABLE vote (
    id TEXT PRIMARY KEY,
    candidate_id TEXT NOT NULL,
    rank INTEGER,
    timestamp DATETIME NOT NULL,
    voter_id TEXT NOT NULL,
    FOREIGN KEY (candidate_id) REFERENCES candidate(candidate_id)
);
CREATE TABLE election_result (
    id TEXT PRIMARY KEY,
    election_id TEXT,
    candidate_id TEXT,
    total_votes INTEGER,
    vote_percentage REAL,
    rank INTEGER,
    FOREIGN KEY (election_id) REFERENCES election(election_id),
    FOREIGN KEY (candidate_id) REFERENCES candidate(candidate_id)
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20240920232332');
