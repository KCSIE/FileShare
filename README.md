# FileShare
FileShare, a web app based on gin &amp; gorm for file sharing.

# Features
- Support multiple files
- QR code
- Clean up scheduler

# Deployment

1. Use [xgo](https://github.com/techknowlogick/xgo) to cross-compile. For example, I want to cross complie on windows for linux. So under the root folder of your project:  

   ```go
   xgo -targets=linux/amd64 -ldflags="-w -s" .
   ```

   Then you will get the executable file (I named it fileshare).

2. Upload `the executable file`, `config.yaml`, `the public folder` to somewhere. I moved them t o `/home/fileshare/` of my server.

3. Configure the IP(or domain name) and port in `config.yaml`.

4. Run the executable file. 

   ```shell
   setsid ./FileName
   ```

   In my case, `setsid ./fileshare`.

5. Access the application through the IP and port you just configured.

# Inspired By / Acknowlegement

- https://airportal.cn/
- https://github.com/Superioz/aqua
- https://github.com/schollz/share
- https://github.com/goupfile/server
- https://github.com/heyLu/share

