# API Gateway with NGINX
POC diatas menggunakan endpoint dari project 1 GogoManager dengan port 8081. Sila ganti port sesuai dengan environment masing-masing

# Informasi
1. Asumsi aplikasi GogoManager berjalan menggunakan docker container, karena konfigurasinya adalah `host.docker.internal:8081`

## How to test
1. Jalankan GogoManager diatas docker
2. Jalankan APIGateway dengan `docker-compose up -d`
3. Buka swagger untuk GogoManager dengan akses `http://localhost:8080/swagger/index.html`
4. Coba hit api dengan postman dengan port `8080`

# Konfigurasi
```
...
    upstream [ur_awesome_service_1] {
        server host.docker.internal:[ur_port_1];
    }

    upstream [ur_awesome_service_n] {
        server host.docker.internal:[ur_port_n];
    }

server {
        listen 80;
        # The server block defines settings for a virtual server.
        # This server listens on port 80.

        server_name localhost;
        # The server_name directive specifies the domain name for this server.

        location / {
            proxy_pass http://[ur_awesome_service_1];
            # For requests to /users, forward the request to the user_service upstream.
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
            # Set various headers to forward to the upstream server.
        }

        location / {
            proxy_pass http://[ur_awesome_service_n];
            # For requests to /users, forward the request to the user_service upstream.
            proxy_set_header Host $host;
            proxy_set_header X-Real-IP $remote_addr;
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header X-Forwarded-Proto $scheme;
            # Set various headers to forward to the upstream server.
        }

        # location / {
        #     return 404;
        #     # Return a 404 Not Found status for any requests not matching the above locations.
        # }
    }
```
