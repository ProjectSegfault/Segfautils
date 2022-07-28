# Segfautilities
Web utilities for Project Segfault

## What does it do?
For now it powers our contact form. In the future we will expand our APIs so you can do more cool things.

## Setup

### Docker: 
```
docker run -d --restart=always -p 6893:6893 --name segfautilities projectsegfault/segfautilities:latest -e HCAPTCHA_SITE_KEY='YOURSITEKEY' -e HCAPTCHA_SECRET_KEY='YOURSECRETKEY' -e SEGFAUTILITIES_WEBHOOK_URL='YOURWEBHOOKURL'
```

We recommend using Docker as it provides better security (we suck in security, so that's why) and we are constantly updating Segfautilities. Docker makes it easy to update the program.

If you're using Portainer, you should know how to add Segfautilities.

### Manual (recommended for development)
```
git clone https://github.com/ProjectSegfault/segfautilities
cd segfautilities/
# You need to add the environment HCAPTCHA_SITE_KEY, HCAPTCHA_SECRET_KEY, SEGFAUTILITIES_WEBHOOK_URL and SEGFAUTILITIES_PORT.
go run main.go # Run this when you've done above!
```
#### NixOS
``` 
git clone https://github.com/ProjectSegfault/segfautilities
cd segfautilities/ 
nix-shell # Avoid installing Go and setting up the web port, by just running a nix shell
# You still need the environment variables HCAPTCHA_SITE_KEY, HCAPTCHA_SECRET_KEY and SEGFAUTILITIES_WEBHOOK_URL though!
go run main.go # I wonder if this is good practice or not. If this isn't good practice, make a GitHub issue please.
```
