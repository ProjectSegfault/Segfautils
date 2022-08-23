# Segfautils
Web utilities for Project Segfault

## What does it do?
For now it powers our contact form. In the future we will expand our APIs so you can do more cool things.

## Setup

### Docker:
```
docker run -d --restart=always -p 6893:6893 --name segfautils -v "$(pwd)"/data:/segfautils/data projectsegfault/segfautils:latest
```
You need to copy the ``config.example.toml`` to ```config.toml`` and customize the values in the config file. YAML works as well, if you'd like to use that instead.

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