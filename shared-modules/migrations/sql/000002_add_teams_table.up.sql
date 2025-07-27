CREATE TABLE IF NOT EXISTS teams (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS team_members (
    team_id UUID NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(userid) ON DELETE CASCADE,
    PRIMARY KEY (team_id, user_id)
);

CREATE TABLE IF NOT EXISTS team_managers (
    team_id UUID NOT NULL REFERENCES teams(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(userid) ON DELETE CASCADE,
    is_main_manager BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (team_id, user_id)
);