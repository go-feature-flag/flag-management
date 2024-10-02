CREATE TYPE variationType AS ENUM ('string','boolean','integer','double','json');
CREATE TABLE IF NOT EXISTS feature_flags
(
    id                UUID          NOT NULL PRIMARY KEY,
    name              TEXT          NOT NULL UNIQUE CHECK (name <> ''),
    description       TEXT,
    variations        JSONB         NOT NULL,
    type              variationType NOT NULL,
    bucketing_key     TEXT,
    metadata          JSONB,
    track_events      BOOLEAN DEFAULT TRUE,
    disable           BOOLEAN DEFAULT FALSE,
    version           TEXT,
    created_date      TIMESTAMP     NOT NULL,
    last_updated_date TIMESTAMP     NOT NULL,
    last_modified_by  TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS rules
(
    id                                     UUID                  NOT NULL PRIMARY KEY,
    feature_flag_id                        UUID REFERENCES feature_flags (id) NOT NULL,
    name                                   TEXT,
    query                                  TEXT,
    variation_result                       TEXT,
    percentages                            JSONB,
    disable                                BOOLEAN DEFAULT FALSE NOT NULL,
    progressive_rollout_initial_variation  TEXT,
    progressive_rollout_end_variation      TEXT,
    progressive_rollout_initial_percentage FLOAT,
    progressive_rollout_end_percentage     FLOAT,
    progressive_rollout_start_date         TIMESTAMP,
    progressive_rollout_end_date           TIMESTAMP,
    is_default                             BOOLEAN DEFAULT FALSE NOT NULL,
    order_index                            INTEGER NOT NULL,
    CONSTRAINT rule_return_something CHECK (percentages IS NOT NULL
        OR variation_result IS NOT NULL
        OR (progressive_rollout_initial_variation IS NOT NULL
            AND progressive_rollout_end_variation IS NOT NULL
            AND progressive_rollout_start_date IS NOT NULL
            AND progressive_rollout_end_date IS NOT NULL
                                                )
        )
);

CREATE INDEX idx_feature_flags_name ON feature_flags (name);
CREATE INDEX idx_rules_feature_flag_id ON rules (feature_flag_id);
