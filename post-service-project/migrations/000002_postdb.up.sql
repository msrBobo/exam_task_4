INSERT INTO posts (id, user_id, content, title, likes, dislikes, views, category, created_at, update_at, deleted_at)
VALUES 
  (uuid_generate_v4(), 'bb23d098-9cbe-462f-897a-9a3a95d4f132', 'content1', 'title1', 10, 5, 100, 'category1', NOW(), NULL, NULL),
  (uuid_generate_v4(), '7466b7a2-b62e-4cc7-b902-d5c7b0f22ce3', 'content2', 'title2', 15, 3, 200, 'category2', NOW(), NULL, NULL),
  (uuid_generate_v4(), '7466b7a2-b62e-4cc7-b902-d5c7b0f22ce3', 'content3', 'title3', 20, 2, 150, 'category3', NOW(), NULL, NULL),
  (uuid_generate_v4(), '7466b7a2-b62e-4cc7-b902-d5c7b0f22ce3', 'content4', 'title4', 8, 7, 180, 'category1', NOW(),  NULL, NULL),
  (uuid_generate_v4(), 'f1a1801d-68a3-43dd-ba08-3171fb3f2ef1', 'content5', 'title5', 25, 1, 220, 'category2', NOW(), NULL, NULL),
  (uuid_generate_v4(), 'f1a1801d-68a3-43dd-ba08-3171fb3f2ef1', 'content6', 'title6', 12, 4, 130, 'category3', NOW(), NULL, NULL),
  (uuid_generate_v4(), 'f1a1801d-68a3-43dd-ba08-3171fb3f2ef1', 'content7', 'title7', 18, 6, 190, 'category1', NOW(), NULL, NULL),
  (uuid_generate_v4(), 'f1a1801d-68a3-43dd-ba08-3171fb3f2ef1', 'content8', 'title8', 22, 2, 210, 'category2', NOW(), NULL, NULL),
  (uuid_generate_v4(), 'dafdf864-3e4f-4f9f-8fba-e17821c2f7ab', 'content9', 'title9', 30, 0, 240, 'category3', NOW(), NULL, NULL),
  (uuid_generate_v4(), 'bb23d098-9cbe-462f-897a-9a3a95d4f132', 'content10', 'title10', 17, 3, 170, 'category1', NOW(),NULL, NULL);

