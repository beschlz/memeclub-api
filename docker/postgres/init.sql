SELECT 'CREATE DATABASE memeclub'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'memeclub')\gexec