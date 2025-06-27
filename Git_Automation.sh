#!/bin/bash



# This script is used just to push code to 
# github.com/anicse37/Library_Management
# and does not serve any purpose in this software.
echo "Hey!!!"

git status 

read  -p "Enter yout commit: " msg

git add .

git status

git commit -m "$msg"

git push
