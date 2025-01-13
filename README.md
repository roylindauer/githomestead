# Git Homestead

> Your personal land for cultivating repositories.


**Requirements**

* Docker
* Ansible


**A git repository server for your personal use**

I use GitHub for nearly all of my personal projects. 
I was thinking that  while it's fine mostly, 
it would be really nice to be able to keep 
my repos in my homelab on a VM I control. 

This project is how I am going to do that. 

This is built in Go, runs in a docker compose env, and is deployed to 
my homelab using Ansible. 

## Features

* REST API for git repositories
* Git WebUI to view git repositories
* Git running in daemon mode to serve git repositories

**Future Features**

* Fleshed out REST API
* CLI for REST API to have a local tool to manage repos
* Maybe a nicer webui (hello htmx?) using that same REST API

## Setup

1. Create `config.yml` from `example.config.yml`.
2. Create `inventory.ini` from `example.inventory.ini`.
3. Edit config.yml and inventory.ini to suit your environment
4. Run `ansible-playbook main.yml` to deploy.<br>
    That playbook will:
    * Install and configure Docker on the remote host
    * Copy the entire project to the remote host
    * Run docker compose build and then start the services

You can then access the services at:

* http://yourhost:8080/repos (REST API)
* http://yourhost:8081 (WebUI)
* git://yourhost:9418 (Git Server)
