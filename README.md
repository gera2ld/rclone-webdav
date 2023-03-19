# rclone-webdav

Create a WebDAV server with `rclone serve webdav --auth-proxy`.

See [the documentation](https://rclone.org/commands/rclone_serve_webdav/) for more details.

## Usage

Create `./data/auth_data.json`:

```json
{
  "user1:pwd1": {
    "type": "local"
  },
  "user2:pwd2": {
    "type": "webdav",
    "_root": "/Documents/my-app",
    "url": "https://cloud.example.com/remote.php/dav/files/gerald/",
    "vendor": "nextcloud",
    "user": "gerald",
    "pass": "app-password"
  }
}
```

Create a directory for DAV files:

```bash
$ mkdir data/webdav
```

For `local` backends, the files will be stored in `data/webdav/$username`.

Create `docker-compose.yml`:

```yaml
version: '3'

services:
  rclone:
    image: rclone-webdav
    build:
      context: https://github.com/gera2ld/rclone-webdav.git#main
    restart: unless-stopped
    volumes:
      - ./data/auth_data.json:/data/auth_data.json:ro
      - ./data/webdav:/data/webdav
    ports:
      - 8080:8080
```

Then start the service by `docker-compose up -d`.
