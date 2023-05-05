CREATE TABLE chat (
  chat_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  question TEXT,
  answer TEXT
);

CREATE TABLE history (
  history_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  content TEXT,
  fromUser TINYINT(1),
  created_at TIMESTAMP DEFAULT NOW()
);



INSERT INTO chat (question, answer) VALUES
('What is your name?', 'My name is ChatTMBOK.'),
('How old are you?', 'I am a computer program, so I do not have an age.'),
('What can you do?', 'I can answer your questions and have a conversation with you.'),
('What is the meaning of life?', 'That is a philosophical question that people have been pondering for centuries.'),
('Do you have any hobbies?', 'As an AI chatbot, I do not have any hobbies, but I am always ready to help you.'),
('What is your favorite color?', 'As a chatbot, I do not have a favorite color. What is your favorite color?'),
('How can I contact you?', 'You can reach me here or send an email to support@chatbot.com.'),
('What do you think of artificial intelligence?', 'As an AI chatbot, I am a big fan of artificial intelligence and its potential to change the world.'),
('What is the best programming language?', 'That depends on what you want to use it for. There is no single "best" programming language.'),
('What is your favorite food?', 'As an AI chatbot, I do not have a favorite food. But what is your favorite food?');