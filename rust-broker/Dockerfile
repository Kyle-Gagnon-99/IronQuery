# Start from the official Rust image
FROM rust:latest

# Install cargo-watch
RUN cargo install cargo-watch

# Set the working directory
WORKDIR /usr/src/rust-broker

# Copy the source code into the container
COPY . .

# Finally start the container
CMD ["cargo", "watch", "-x", "run"]