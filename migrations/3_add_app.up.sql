INSERT INTO apps (id, name, secret)
VALUES (1, 'text', 'test-secret')
ON CONFLICT DO NOTHING;