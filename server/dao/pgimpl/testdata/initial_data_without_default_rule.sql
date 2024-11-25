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

INSERT INTO rules
(id, feature_flag_id, name, query, variation_result, percentages, disable, progressive_rollout_initial_variation, progressive_rollout_end_variation, progressive_rollout_initial_percentage, progressive_rollout_end_percentage, progressive_rollout_start_date, progressive_rollout_end_date, is_default, order_index)
VALUES (
   '546939a9-6df8-4a0b-b9cf-1d69ff300eb5',
   '69aa10ec-ec3e-4139-8cdf-6902a5746e2d',
   'rule 1',
   'targetingKey eq "valueA"',
   NULL,
   '{"variationA": 10, "variationB": 90}',
   FALSE,
   NULL,
   NULL,
   NULL,
   NULL,
   NULL,
   NULL,
   FALSE,
   1
);

INSERT INTO rules
(id, feature_flag_id, name, query, variation_result, percentages, disable, progressive_rollout_initial_variation, progressive_rollout_end_variation, progressive_rollout_initial_percentage, progressive_rollout_end_percentage, progressive_rollout_start_date, progressive_rollout_end_date, is_default, order_index)
VALUES (
   '9f82fe80-b4b6-426a-869a-e4436de66d0d',
   '69aa10ec-ec3e-4139-8cdf-6902a5746e2d',
   'rule 2',
   'targetingKey eq "1234"',
   NULL,
   NULL,
   TRUE,
   'variationA',
   'variationB',
   0,
   100,
   '2023-01-01 00:00:00' AT TIME ZONE 'UTC',
   '2024-01-01 00:00:00' AT TIME ZONE 'UTC',
   FALSE,
   0
);