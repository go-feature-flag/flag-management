INSERT INTO feature_flags
    (id, name, description, variations, type, bucketing_key, metadata, track_events, disable, version, created_date, last_updated_date, last_modified_by)
VALUES (
    '69aa10ec-ec3e-4139-8cdf-6902a5746e2d',
    'my-feature-flag',
    'This is a feature flag',
    '{"variationA": "valueA", "variationB": "valueB"}',
    'string',
    NULL,
    '{"key": "value"}',
    TRUE,
    FALSE,
    '1.0.0',
    '2020-01-01 00:00:00' AT TIME ZONE 'UTC',
    '2020-01-01 00:00:00' AT TIME ZONE 'UTC',
    'admin'
);