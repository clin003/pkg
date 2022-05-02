#!/bin/bash

git add .
git commit -m "debug"
#git remote rename origin gitee
git push -u gitee main

#git remote add github git@github.com:clin003/pkg.git
#git branch -M main
git push -u github main