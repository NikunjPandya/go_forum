# go_forum
Question answer repo in GoLang


# SQL Declarations

```
CREATE TABLE user (id SERIAL PRIMARY KEY, name VARCHAR NOT NULL, created_at timestamp DEFAULT NOW(), updated_at timestamp DEFAULT NOW(), deleted_at timestamp)
CREATE TABLE question (id SERIAL PRIMARY KEY, question VARCHAR NOT NULL, user_id integer not null references user(id), created_at timestamp DEFAULT NOW(), updated_at timestamp DEFAULT NOW(), deleted_at timestamp)
CREATE TABLE answers(id SERIAL PRIMARY KEY, answer TEXT, user_id int REFERENCES users(id) NOT NULL, question_id int REFERENCES questions(id) NOT NULL, created_at timestamp DEFAULT NOW(), updated_at timestamp DEFAULT NOW(), deleted_at timestamp)
```
# APIS
Can be found in main.go file.
