CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

INSERT INTO users (id, username, email, password, created_at, updated_at, first_name, last_name, bio, website)
VALUES
    ('bd80f4c4-59ab-4ade-82e6-5600f2b39c0a', 'user1', 'user1@example.com', 'password1', NOW(), NULL, 'John', 'Doe', 'Bio for user 1', 'https://example.com/user1'),
    ('21c3ebd2-8a2b-466c-955f-58ad0d539ba9', 'user2', 'user2@example.com', 'password2', NOW(), NULL, 'Jane', 'Smith', 'Bio for user 2', 'https://example.com/user2'),
    ('fa147c55-f451-4290-a8a7-21266cf9df4f', 'user3', 'user3@example.com', 'password3', NOW(), NULL, 'Alice', 'Johnson', 'Bio for user 3', 'https://example.com/user3'),
    ('424a6582-ffa2-447a-a04b-b2040ebda0e3', 'user4', 'user4@example.com', 'password4', NOW(), NULL, 'Bob', 'Williams', 'Bio for user 4', 'https://example.com/user4'),
    ('2074860c-0348-4c4f-ba4a-ac6982550664', 'user5', 'user5@example.com', 'password5', NOW(), NULL, 'Emily', 'Brown', 'Bio for user 5', 'https://example.com/user5'),
    ('fc1b0622-214c-4613-954e-44ec04f0ed5c', 'user6', 'user6@example.com', 'password6', NOW(), NULL, 'David', 'Jones', 'Bio for user 6', 'https://example.com/user6'),
    ('4d96224e-5dfd-4be3-b26d-ab98f6b5cbca', 'user7', 'user7@example.com', 'password7', NOW(), NULL, 'Sarah', 'Miller', 'Bio for user 7', 'https://example.com/user7'),
    ('c701306b-8949-49de-a647-2602cf5a1c07', 'user8', 'user8@example.com', 'password8', NOW(), NULL, 'Michael', 'Davis', 'Bio for user 8', 'https://example.com/user8'),
    ('ef463075-fb91-4de1-bad0-7181a9ac8fb5', 'user9', 'user9@example.com', 'password9', NOW(), NULL, 'Rachel', 'Taylor', 'Bio for user 9', 'https://example.com/user9'),
    ('c0be9ae0-8980-4808-b185-62333825ea26', 'user10', 'user10@example.com', 'password10', NOW(), NULL, 'Kevin', 'Clark', 'Bio for user 10', 'https://example.com/user10');

INSERT INTO posts (id, user_id, content, title, likes, dislikes, views, category, created_at, update_at, deleted_at)
VALUES 
    ('b3fd920c-ef7c-4ec9-8a75-e04e64306269', 'bd80f4c4-59ab-4ade-82e6-5600f2b39c0a', 'Content 1', 'Post 1', 5, 0, 100, 'Category 1', NOW(), NULL, NULL),
    ('337451bd-6178-43e8-86dc-8e3a440bbc92', '21c3ebd2-8a2b-466c-955f-58ad0d539ba9', 'Content 2', 'Post 2', 10, 2, 200, 'Category 2', NOW(), NULL, NULL),
    ('7798b238-79ca-42b2-bd19-b5fe5017388a', 'fa147c55-f451-4290-a8a7-21266cf9df4f', 'Content 3', 'Post 3', 15, 4, 150, 'Category 3', NOW(), NULL, NULL),
    ('c3fa6df7-b2f2-4414-879b-4939552a2875', '424a6582-ffa2-447a-a04b-b2040ebda0e3', 'Content 4', 'Post 4', 20, 3, 180, 'Category 1', NOW(), NULL, NULL),
    ('772a1a45-70f7-40e9-8f44-bb9a03986808', '2074860c-0348-4c4f-ba4a-ac6982550664', 'Content 5', 'Post 5', 8, 5, 220, 'Category 2', NOW(), NULL, NULL),
    ('59c0d999-e42b-4a9f-afe7-8be8cab926a7', 'fc1b0622-214c-4613-954e-44ec04f0ed5c', 'Content 6', 'Post 6', 25, 2, 130, 'Category 3', NOW(), NULL, NULL),
    ('4a1a337f-d74d-4cf1-b40b-a7085883020f', '4d96224e-5dfd-4be3-b26d-ab98f6b5cbca', 'Content 7', 'Post 7', 18, 6, 190, 'Category 1', NOW(), NULL, NULL),
    ('e504c892-804f-4996-ba90-6b28b6447b5a', 'ef463075-fb91-4de1-bad0-7181a9ac8fb5', 'Content 8', 'Post 8', 22, 1, 210, 'Category 2', NOW(), NULL, NULL),
    ('06676bd9-1dac-4a66-bc68-404ebcdb7d35', 'c701306b-8949-49de-a647-2602cf5a1c07', 'Content 9', 'Post 9', 30, 0, 240, 'Category 3', NOW(), NULL, NULL),
    ('c0be9ae0-8980-4808-b185-62333825ea26', 'bd80f4c4-59ab-4ade-82e6-5600f2b39c0a', 'Content 10', 'Post 10', 17, 3, 170, 'Category 1', NOW(), NULL, NULL);

INSERT INTO comments (id, post_id, user_id, content, likes, dislikes, created_at, updated_at, deleted_at)
VALUES
    (uuid_generate_v4(), 'b3fd920c-ef7c-4ec9-8a75-e04e64306269', 'bd80f4c4-59ab-4ade-82e6-5600f2b39c0a', 'Comment 1 on Post 1', 2, 0, NOW(), NULL, NULL),
    (uuid_generate_v4(), 'b3fd920c-ef7c-4ec9-8a75-e04e64306269', '21c3ebd2-8a2b-466c-955f-58ad0d539ba9', 'Comment 2 on Post 1', 3, 1, NOW(), NULL, NULL),
    (uuid_generate_v4(), '337451bd-6178-43e8-86dc-8e3a440bbc92', '21c3ebd2-8a2b-466c-955f-58ad0d539ba9', 'Comment 1 on Post 2', 5, 2, NOW(), NULL, NULL),
    (uuid_generate_v4(), '7798b238-79ca-42b2-bd19-b5fe5017388a', '424a6582-ffa2-447a-a04b-b2040ebda0e3', 'Comment 1 on Post 3', 4, 0, NOW(), NULL, NULL),
    (uuid_generate_v4(), 'c3fa6df7-b2f2-4414-879b-4939552a2875', 'fa147c55-f451-4290-a8a7-21266cf9df4f', 'Comment 1 on Post 4', 3, 1, NOW(), NULL, NULL),
    (uuid_generate_v4(), '772a1a45-70f7-40e9-8f44-bb9a03986808', 'b40e9230-0f0e-408d-851f-cdb2778d1c26', 'Comment 1 on Post 5', 6, 3, NOW(), NULL, NULL),
    (uuid_generate_v4(), '59c0d999-e42b-4a9f-afe7-8be8cab926a7', 'fa147c55-f451-4290-a8a7-21266cf9df4f', 'Comment 2 on Post 6', 2, 0, NOW(), NULL, NULL),
    (uuid_generate_v4(), '4a1a337f-d74d-4cf1-b40b-a7085883020f', '4d96224e-5dfd-4be3-b26d-ab98f6b5cbca', 'Comment 1 on Post 7', 5, 1, NOW(), NULL, NULL),
    (uuid_generate_v4(), 'e504c892-804f-4996-ba90-6b28b6447b5a', '424a6582-ffa2-447a-a04b-b2040ebda0e3', 'Comment 2 on Post 8', 3, 0, NOW(), NULL, NULL),
    (uuid_generate_v4(), '06676bd9-1dac-4a66-bc68-404ebcdb7d35', 'bd80f4c4-59ab-4ade-82e6-5600f2b39c0a', 'Comment 2 on Post 9', 4, 2, NOW(), NULL, NULL);
