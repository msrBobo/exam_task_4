INSERT INTO comments (id, post_id, user_id, content, likes, dislikes, created_at, updated_at, deleted_at)
VALUES
  (uuid_generate_v4(), '547115a7-cc1f-4201-bce4-a0dce9a97618', 'bb23d098-9cbe-462f-897a-9a3a95d4f132', 'content1', 5, 0, NOW(), NULL, NULL),
  (uuid_generate_v4(), '2399a1f6-d5d6-4eb1-8059-8d702227262a', '7466b7a2-b62e-4cc7-b902-d5c7b0f22ce3', 'content2', 4, 0, NOW(), NULL, NULL),
  (uuid_generate_v4(), 'a0abb2be-50cf-40a8-ab9b-3430bf05a384', '7466b7a2-b62e-4cc7-b902-d5c7b0f22ce3', 'content3', 55, 0, NOW(), NULL, NULL),
  (uuid_generate_v4(), 'a0bf6462-59f3-4dba-a705-5483540c6a3a', '7466b7a2-b62e-4cc7-b902-d5c7b0f22ce3', 'content4', 32, 7, NOW(), NULL, NULL),
  (uuid_generate_v4(), 'd17b8ae5-2d1a-45db-bdd2-73cbdfe9577b', 'f1a1801d-68a3-43dd-ba08-3171fb3f2ef1', 'content5', 76, 6, NOW(), NULL, NULL),
  (uuid_generate_v4(), '9005a294-74c2-4264-80af-1352dcf0962f', 'f1a1801d-68a3-43dd-ba08-3171fb3f2ef1', 'content6', 44, 0, NOW(), NULL, NULL),
  (uuid_generate_v4(), '531c119a-d0cf-4e86-993f-d2714053f3aa', 'f1a1801d-68a3-43dd-ba08-3171fb3f2ef1', 'content7', 65, 5, NOW(), NULL, NULL),
  (uuid_generate_v4(), 'fa257936-d1e8-4cc9-83c4-5dfce38c077a', 'f1a1801d-68a3-43dd-ba08-3171fb3f2ef1', 'content8', 77, 0, NOW(), NULL, NULL),
  (uuid_generate_v4(), '2cf4353f-7ddc-45a3-b673-8d24bffd540a', 'dafdf864-3e4f-4f9f-8fba-e17821c2f7ab', 'content9', 73, 3, NOW(), NULL, NULL),
  (uuid_generate_v4(), '56da3637-e2e4-49b6-811a-8e8a52bf300d', 'bb23d098-9cbe-462f-897a-9a3a95d4f132', 'content10', 33, 0, NOW(),NULL, NULL);
