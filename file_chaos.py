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

# Function which returns a list of files
def list_files():
    # Creates a list of files from a private directory
    files = [os.path.join(TARGET_DIR, f) for f in os.listdir(TARGET_DIR) if os.path.isfile(os.path.join(TARGET_DIR, f))]
    # Returns the files
    return files

# Function to create file
def create_file():
    # Creates the random files path
    path = os.path.join(TARGET_DIR, random_filename())
    # Opens the new file at the path
    with open(path, "w") as f:
        # Writes content from the random content function to the file
        f.write(random_content())
    # Logs a file has been created
    print(f"[CREATE] {path}")

# Function to modify file
def modify_file(path):
    # Opens the file path in the parameter
    with open(path, "a") as f:
        # Writes adds extra content to files
        f.write("\n" + random_content())
    # Outputs the path that has been modified
    print(f"[MODIFY] {path}")

# Function to overright file
def overwrite_file(path):
    # Opens the new file at the path
    with open(path, "w") as f:
        # Writes content from the random content function to the file
        f.write(random_content())
    # Outputs the path that has been modified
    print(f"[OVERWRITE] {path}")

# Function to delete file
def delete_file(path):
    # Removes the file at the path
    os.remove(path)
    # Outputs the delete file function
    print(f"[DELETE] {path}")

# Main function
def main():
    # Runs while true
    while True:
        # Lists the files in the directory
        files = list_files()
        # Chooses a random action
        action = random.choice(["create", "modify", "overwrite", "delete"])

        # Checks if the action is create and the length of files is less than the max amount
        if action == "create" and len(files) < MAX_FILES:
            # Creates a random file
            create_file()

        # Checks if the action is in modify and overright and if there are files
        elif action in ("modify", "overwrite") and files:
            # Chooses a random file
            path = random.choice(files)
            # Checsk if the action is modify
            if action == "modify":
                # Calls the modify file function with the path
                modify_file(path)
            else:
                # Calls the overwrite file path
                overwrite_file(path)

        # Checks for the delete action and if there are files
        elif action == "delete" and files:
            # Deletes a random file
            delete_file(random.choice(files))
        # Sleeps for a random amount of time
        time.sleep(random.uniform(SLEEP_MIN, SLEEP_MAX))

# Starts the program
if __name__ == "__main__":
    print("🔥 File chaos generator started")
    main()