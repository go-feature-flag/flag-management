CREATE TABLE IF NOT EXISTS feature_flags_history
(
--     attributes of the history table
    history_timestamp  TIMESTAMPTZ NOT NULL DEFAULT now(),
    history_deleted_by VARCHAR,
    history_op         VARCHAR     NOT NULL,
--     attributes of the feature_flags table
    id                UUID          NOT NULL,
    name              TEXT          NOT NULL,
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


CREATE TABLE IF NOT EXISTS rules_history
(
    --     attributes of the history table
    history_timestamp  TIMESTAMPTZ NOT NULL DEFAULT now(),
    history_deleted_by VARCHAR,
    history_op         VARCHAR     NOT NULL,
    --     attributes of the rules table
    id                                     UUID                  NOT NULL,
    feature_flag_id                        UUID  NOT NULL,
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
    order_index                            INTEGER NOT NULL
);