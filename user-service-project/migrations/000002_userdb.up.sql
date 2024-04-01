INSERT INTO users (id, username, email, password, created_at, updated_at, first_name, last_name, bio, website)
VALUES 
(uuid_generate_v4(), 'user1', 'user1@example.com', 'password1', NOW(), NULL, 'John', 'Doe', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.', 'https://example.com/user1'),
(uuid_generate_v4(), 'user2', 'user2@example.com', 'password2', NOW(), NULL, 'Jane', 'Smith', 'Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.', 'https://example.com/user2'),
(uuid_generate_v4(), 'user3', 'user3@example.com', 'password3', NOW(), NULL, 'Alice', 'Johnson', 'Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.', 'https://example.com/user3'),
(uuid_generate_v4(), 'user4', 'user4@example.com', 'password4', NOW(), NULL, 'Bob', 'Williams', 'Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.', 'https://example.com/user4'),
(uuid_generate_v4(), 'user5', 'user5@example.com', 'password5', NOW(), NULL, 'Emily', 'Brown', 'Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'https://example.com/user5'),
(uuid_generate_v4(), 'user6', 'user6@example.com', 'password6', NOW(), NULL, 'David', 'Jones', 'Fugiat quo voluptas nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.', 'https://example.com/user6'),
(uuid_generate_v4(), 'user7', 'user7@example.com', 'password7', NOW(), NULL, 'Sarah', 'Miller', 'Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium.', 'https://example.com/user7'),
(uuid_generate_v4(), 'user8', 'user8@example.com', 'password8', NOW(), NULL, 'Michael', 'Davis', 'Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit.', 'https://example.com/user8'),
(uuid_generate_v4(), 'user9', 'user9@example.com', 'password9', NOW(), NULL, 'Rachel', 'Taylor', 'Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit.', 'https://example.com/user9'),
(uuid_generate_v4(), 'user10', 'user10@example.com','password10', NOW(),NULL, 'Kevin', 'Clark', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.', 'https://example.com/user10');
