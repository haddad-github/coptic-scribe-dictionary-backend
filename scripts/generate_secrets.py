import json
import os

print("=== Create secrets.json ===")
db_name = input("Database name (default: coptic_dictionary_db): ") or "coptic_dictionary_db"
db_user = input("Database user (default: postgres): ") or "postgres"
db_password = input("Database password: ")
db_host = input("Database host (default: localhost): ") or "localhost"
db_port = input("Database port (default: 5432): ") or "5432"

secrets = {
    "DB_NAME": db_name,
    "DB_USER": db_user,
    "DB_PASSWORD": db_password,
    "DB_HOST": db_host,
    "DB_PORT": db_port
}

env = f"""DB_HOST=host.docker.internal
DB_USER={db_user}
DB_PASSWORD={db_password}
DB_NAME={db_name}
DB_PORT={db_port}
DB_SSLMODE=disable
"""

# Save secrets.json
with open("secrets.json", "w") as f:
    json.dump(secrets, f, indent=4)
    print("secrets.json created.")

# Save .env
os.makedirs("api", exist_ok=True)
with open("api/.env", "w") as f:
    f.write(env)
    print("api/.env created.")