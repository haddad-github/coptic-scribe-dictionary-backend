import json
import os
import pandas as pd
import psycopg2
from psycopg2 import sql
from psycopg2.extensions import ISOLATION_LEVEL_AUTOCOMMIT

#Load database credentials from secrets.json
secrets_file = os.path.join(os.path.dirname(__file__), "../secrets.json")

try:
    with open(secrets_file, "r") as f:
        secrets = json.load(f)

    DB_NAME = secrets["DB_NAME"]
    DB_USER = secrets["DB_USER"]
    DB_PASSWORD = secrets["DB_PASSWORD"]
    DB_HOST = secrets["DB_HOST"]
    DB_PORT = secrets["DB_PORT"]
except Exception as e:
    print("Error loading secrets.json:", e)
    exit(1)

#Path of the Excel file
file_path = "dictionary_sources_digital/coptic_dict_2024.xlsx"

#Create the PostgreSQL Database (if it doesn't exist)
try:
    conn = psycopg2.connect(
        dbname="postgres", #temporarily connect to the default db, in case the coptic dictionary db doesn't exist yet
        user=DB_USER,
        password=DB_PASSWORD,
        host=DB_HOST,
        port=DB_PORT
    )
    conn.set_isolation_level(ISOLATION_LEVEL_AUTOCOMMIT)  #allow creating databases
    cursor = conn.cursor()
    cursor.execute(f"SELECT 1 FROM pg_database WHERE datname = '{DB_NAME}';")
    exists = cursor.fetchone()

    #if doesn't exist, create database
    if not exists:
        cursor.execute(f"CREATE DATABASE {DB_NAME};")
        print(f"Database '{DB_NAME}' created successfully.")
    else:
        print(f"Database '{DB_NAME}' already exists.")

    cursor.close()
    conn.close()
except Exception as e:
    print("Error creating database:", e)

#Connect to the newly created database
try:
    conn = psycopg2.connect(
        dbname=DB_NAME,
        user=DB_USER,
        password=DB_PASSWORD,
        host=DB_HOST,
        port=DB_PORT
    )
    cursor = conn.cursor()
    print(f"Connected to database '{DB_NAME}'.")

    #Create table, if it doesn't exist
    create_table_query = """
    CREATE TABLE IF NOT EXISTS coptic_dictionary (
        id SERIAL PRIMARY KEY,
        coptic_word TEXT,
        arabic_translation TEXT,
        english_translation TEXT,
        greek_script_coptic_word TEXT,
        grammatical_modification TEXT,
        word_category TEXT,
        word_gender TEXT,
        word_category_2 TEXT,
        word_gender_2 TEXT,
        greek_word TEXT,
        coptic_word_alt TEXT
    );
    """
    cursor.execute(create_table_query)
    conn.commit()
    print(f"Table {DB_NAME} created successfully.")

    #Load Excel file
    df = pd.read_excel(file_path, header=None)

    #Columns in order
    df.columns = [
        "coptic_word",
        "arabic_translation",
        "english_translation",
        "greek_script_coptic_word",
        "grammatical_modification",
        "word_category",
        "word_gender",
        None,  #ignore
        "word_category_2",
        "word_gender_2",
        "greek_word",
        None,  #ignore
        None,  #ignore
        "coptic_word_alt",
        None  #extra ignored column to match 15 columns
    ]

    #Drop ignored columns
    df = df.drop(columns=[None])

    #Remove empty rows
    df = df.dropna(how="all")

    #Remove rows where "coptic_word" is missing
    df = df.dropna(subset=["coptic_word"])

    #Remove duplicate Coptic words, keeping the first occurrence
    df = df.drop_duplicates(subset=["coptic_word"], keep="first")

    #Insert data into PostgreSQL
    insert_query = sql.SQL("""
        INSERT INTO coptic_dictionary (
            coptic_word, arabic_translation, english_translation, 
            greek_script_coptic_word, grammatical_modification, word_category, 
            word_gender, word_category_2, word_gender_2, greek_word, coptic_word_alt
        ) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s);
    """)

    #Batch insert rows
    records = df.to_records(index=False)
    cursor.executemany(insert_query, records)
    conn.commit()
    print(f"Inserted {len(records)} records into PostgreSQL.")

    #Add index for each entry
    cursor.execute("CREATE INDEX IF NOT EXISTS idx_coptic_word ON coptic_dictionary (coptic_word);")
    conn.commit()
    print("Indexing completed.")

except Exception as e:
    print("Error:", e)
finally:
    cursor.close()
    conn.close()
    print("PostgreSQL connection closed.")
