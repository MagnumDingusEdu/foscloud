-- Create the root element beforehand
INSERT INTO nodes
(parent_id, name, depth, lineage, is_dir)
VALUES (null, 'root', 1, '/1/', true)
RETURNING *;