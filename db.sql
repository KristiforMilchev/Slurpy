CREATE TABLE IF NOT EXISTS deployments (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    contract TEXT NOT NULL,
    name TEXT NOT NULL,
    created_at TEXT DEFAULT (datetime('now', 'localtime')),
    group_name TEXT
);

CREATE TABLE IF NOT EXISTS deployment_parameters (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    parameter TEXT,
    deploymentId INTEGER,
    FOREIGN KEY (deploymentId) REFERENCES deployments(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS wallets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    wallet_key TEXT,
    network TEXT,
    FOREIGN KEY (network) REFERENCES networks(network_id) ON DELETE CASCADE

);

CREATE TABLE IF NOT EXISTS networks (
    network_name TEXT PRIMARY KEY,
    rpc TEXT,
    network_id INTEGER
)
