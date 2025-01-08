FROM bitnami/git:latest 

CMD ["git", "daemon", "--reuseaddr", "--export-all", "--enable=receive-pack", "--verbose", "--base-path=/srv/git/", "/srv/git/"]
