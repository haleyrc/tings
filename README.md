## Deploying a new version

1. Push a new image:
   ```bash
   make release
   ```
2. Connect to the NAS:
   ```bash
   ssh <IP>
   ```
3. Pull the new image:
   ```bash
   sudo docker pull ghcr.io/haleyrc/tings:latest
   ```
4. Restart the service:
   ```bash
   sudo docker-compose up -d server
   ```
