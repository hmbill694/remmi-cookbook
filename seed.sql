-- Clear existing data
TRUNCATE chat_messages,
chats,
recipe_chunks,
recipes,
users RESTART IDENTITY CASCADE;

-- Insert Users
INSERT INTO
    users (username, email, password_hash)
VALUES
    (
        'john_doe',
        'john@example.com',
        'hashed_password_1'
    ),
    (
        'jane_smith',
        'jane@example.com',
        'hashed_password_2'
    ),
    (
        'chef_master',
        'chef@example.com',
        'hashed_password_3'
    );

-- Insert Recipes
INSERT INTO
    recipes (user_id, title, description)
VALUES
    (
        1,
        'Spaghetti Bolognese',
        'A classic Italian pasta dish.'
    ),
    (
        2,
        'Chocolate Cake',
        'Rich and moist chocolate cake.'
    ),
    (
        3,
        'Vegetable Stir Fry',
        'Quick and healthy stir fry with fresh vegetables.'
    );

-- Insert Recipe Chunks
INSERT INTO
    recipe_chunks (recipe_id, chunk_order, content)
VALUES
    (1, 1, 'Step 1: Heat olive oil in a large pan.'),
    (
        1,
        2,
        'Step 2: Add ground beef and cook until browned.'
    ),
    (
        1,
        3,
        'Step 3: Add tomatoes and simmer for 30 minutes.'
    ),
    (2, 1, 'Step 1: Preheat oven to 350°F (175°C).'),
    (
        2,
        2,
        'Step 2: Mix flour, cocoa powder, and sugar in a bowl.'
    ),
    (
        2,
        3,
        'Step 3: Bake for 30 minutes until a toothpick comes out clean.'
    ),
    (3, 1, 'Step 1: Heat a wok and add oil.'),
    (
        3,
        2,
        'Step 2: Stir-fry garlic and ginger until fragrant.'
    ),
    (
        3,
        3,
        'Step 3: Add vegetables and stir-fry for 5 minutes.'
    );

-- Insert Chats
INSERT INTO
    chats (recipe_id, user_id, started_at)
VALUES
    (1, 1, NOW ()),
    (2, 2, NOW ()),
    (3, 3, NOW ());

-- Insert Chat Messages
INSERT INTO
    chat_messages (chat_id, sender, message, sent_at)
VALUES
    (1, 'user', 'How do I make the sauce?', NOW ()),
    (
        1,
        'llm',
        'Start by heating olive oil and adding ground beef.',
        NOW ()
    ),
    (
        2,
        'user',
        'Can I use butter instead of oil?',
        NOW ()
    ),
    (
        2,
        'llm',
        'Yes, butter works well for this recipe.',
        NOW ()
    ),
    (
        3,
        'user',
        'What vegetables are best for this stir fry?',
        NOW ()
    ),
    (
        3,
        'llm',
        'Use broccoli, bell peppers, and carrots for the best flavor.',
        NOW ()
    );
