# Server Backup CLI

## Overview
Server Backup CLI is a command-line application designed to manage backups of server projects. 

## Features
- Manage applications (CRUD)
- Manage the app's backup frequency
- Backup full and custom paths
- upload the backup locally and externally
- list backups and manage it
- CLI API
- Rest API
- Providers to save the (ssh,digitalocean, aws s3, google drive)


### Notes
- sort apps by space
- list apps with their spaces
- check the available free space on disk before taking the backup
- each app has paths and each path has its own backup frequency
- each app's database has its own frequency too
- each database can be full backup  or custom tables or exclude tables
- database dump for each db engine and the execution command can be locally or using docker run
- auto configure cronjob
- notify discord for failure
- notify discord with daily backup summary

## todo 

i was working on the perform backup ,test if the get stats is working and save it to database and then perform the backup process