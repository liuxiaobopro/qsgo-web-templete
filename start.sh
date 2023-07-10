#!/bin/bash

project_name="qsgo-web-templete"

# Set log directory path
log_directory="/www/wwwroot/$project_name/log"

# Check if log directory exists, if not, create it
if [ ! -d "$log_directory" ]; then
    mkdir "$log_directory"
fi

# Get current datetime
current_date=$(date +%Y%m%d%H%M%S)

# Construct filename
filename="$log_directory/my.log"

# Check if file exists, if not, create it
if [ ! -f "$filename" ]; then
    touch "$filename"
else
    # If the file exists, make a copy with the current date as the name
    cp "$filename" "$log_directory/$current_date.log"
fi

# Run the command and redirect the output to the file
/www/wwwroot/"$project_name"/"$project_name" -p 10011 -c prod > "$filename"
