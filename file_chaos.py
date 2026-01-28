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

