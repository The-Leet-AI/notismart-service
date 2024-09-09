CREATE TABLE if not exists users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL UNIQUE,
    password_hash TEXT NOT NULL,
    phone_number TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    email_verified BOOLEAN DEFAULT FALSE,
    verification_token UUID DEFAULT gen_random_uuid(),
    phone_verified BOOLEAN DEFAULT FALSE,
    phone_verification_code TEXT
);

CREATE TABLE if not exists notifications (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    send_at TIMESTAMP WITH TIME ZONE,
    status TEXT NOT NULL CHECK (status IN ('Pending', 'Sent', 'Failed'))
);

CREATE TABLE if not exists user_preferences (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id),
    preferred_time TIME,
    preferred_method TEXT NOT NULL CHECK (preferred_method IN ('Email', 'SMS', 'Push'))
);
