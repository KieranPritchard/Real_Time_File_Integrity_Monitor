import os
import random
import string
import time

TARGET_DIR = "testdata" # Stores the target directory
MAX_FILES = 10 # Stores the maximum number of files
SLEEP_MIN = 0.5 # Stores the minimum time to sleep
SLEEP_MAX = 2.0 # Stores the longest time to sleep

# Creates the target directory
os.makedirs(TARGET_DIR, exist_ok=True)

# Function creates a random file name
def random_filename():
    # Creates the name by joining the random choice of a string of ascii characters
    name = "".join(random.choices(string.ascii_lowercase, k=6))
    # Returns the file name
    return f"{name}.txt"

# Function to generate random content
def random_content():
    # Decides the random length of the file
    length = random.randint(10, 100)
    # Creates the content by joining the random choice of a string of ascii character
    return "".join(random.choices(string.ascii_letters + string.digits, k=length))