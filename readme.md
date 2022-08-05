# Segfautils
Web utilities for Project Segfault

## What does it do?
For now it powers our contact form. In the future we will expand our APIs so you can do more cool things.

## Setup

### Docker:
```
docker run -d --restart=always -p 6893:6893 --name segfautils --env-file ./docker.env -v "$(pwd)"/data:/segfautils/data projectsegfault/segfautils:latest
```
docker.env should be the environment file located in this repository, customized to your settings. The env file is self-documenting so I don't need to go in any detail here.


We recommend using Docker as it provides better security (we suck in security, so that's why) and we are constantly updating Segfautils. Docker makes it easy to update the program.

If you're using Portainer, you should know how to add Segfautils.

### Manual (recommended for development)
```
git clone https://github.com/ProjectSegfault/segfautils
cd segfautils/
# You need to add the environment HCAPTCHA_SITE_KEY, HCAPTCHA_SECRET_KEY, SEGFAUTILS_WEBHOOK_URL and SEGFAUTILS_PORT.
go run main.go # Run this when you've done above, and you're planning on developing, if not, do below
go build . -o segfautils
./segfautils
```
#### NixOS
```
git clone https://github.com/ProjectSegfault/segfautils
cd segfautils/
nix-shell # Avoid installing Go and setting up the web port, by just running a nix shell
# You still need the environment variables HCAPTCHA_SITE_KEY, HCAPTCHA_SECRET_KEY and SEGFAUTILS_WEBHOOK_URL though!
go run main.go # If you're developing
go build . -o segfautils && ./segfautils # If you're intending to use Segfautils for production.
```
