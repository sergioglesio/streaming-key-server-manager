[![Streaming Key Server][repo_logo_img]][repo_url]

# Streaming Key Server

[![Go version][go_version_img]][go_dev_url]
[![Go report][go_report_img]][go_report_url]
[![Code coverage][go_code_coverage_img]][repo_url]
[![Wiki][repo_wiki_img]][repo_wiki_url]
[![License][repo_license_img]][repo_license_url]

## Overview

Streaming Key Server is a backend application developed in Go for managing streaming passkey authentication, integrating PostgreSQL for database management, and Docker for deployment automation.

## Features

- **Nginx**: Utilizes the RTMP module for stream ingestion and HLS configuration.
- **PostgreSQL**: Stores streaming keys and manages authentication.
- **Docker**: Ensures the environment is containerized for easy deployment and management.
- **Authentication**: Implements a Go-based server for validating streaming keys.
- **HLS**: Configures Nginx to create HLS playlists and serve .ts segments via HTTP.

## Getting Started

### Prerequisites

Ensure you have the following installed:

- [Go](https://golang.org/dl/) (version `1.22` or higher)
- [Docker](https://www.docker.com/get-started)
- [docker-compose](https://docs.docker.com/compose/install/)

### Installation

Clone the repository:

```bash
git clone https://github.com/sergioglesio/streaming-key-server.git
cd streaming-key-server 
```

## Build and run the Docker containers
```bash
docker-compose up --build
```

## Usage
```bash
# Start a new stream: Use vMIX or any RTMP streaming software to stream:
rtmp://localhost:1935/live/stream_key

# Fetch HLS Playlist: Access the HLS playlist via the HTTP server: 
http://localhost:8081/live/stream_key.m3u8
```

![cgapp_create][cgapp_create_gif]

## Technical Stack

- **Golang**: Backend server for key authentication.
- **PostgreSQL**: Database for storing and managing streaming keys.
- **Nginx**: Server for handling RTMP streams and generating HLS playlists.
- **Docker**: Containerization for deployment and environment consistency.

## Documentation

For detailed documentation and usage instructions, visit the [project wiki][repo_wiki_url].

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.


## ⚠️ License

This project is free and open-source software licensed under the [Apache 2.0 License][repo_license_url]. Official logo was created by [Sérgio Glésio][author] and distributed under [Creative Commons][repo_cc_url] license (CC BY-SA 4.0 International).


[repo_wiki_url]: https://github.com/sergioglesio/streaming-key-server-manager
[repo_license_url]: https://github.com/sergioglesio/streaming-key-server-manager
[repo_cc_url]: https://github.com/sergioglesio/streaming-key-server-manager
[author]: https://github.com/sergioglesio
[cgapp_create_gif]: https://github.com/sergioglesio/streaming-key-server-manager/image/image.png
