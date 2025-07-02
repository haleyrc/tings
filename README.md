## Deploying a new version

1. Push a new image:
   ```bash
   make release
   ```
2. Connect to the NAS:
   ```bash
   ssh <IP>
   ```
3. Navigate to the project directory:
   ```bash
   cd /volume1/docker/tings
   ```
4. Pull the new image:
   ```bash
   sudo docker pull ghcr.io/haleyrc/tings:latest
   ```
5. Restart the service:
   ```bash
   sudo docker-compose up -d server
   ```
