# Use the "python:slim-bookworm" base image for building this Docker image.
FROM python:slim-bookworm

# Install Ansible using pip3 with the --user flag to install it for the current user only.
RUN python3 -m pip install --user ansible

# Add the directory "/root/.local/bin" to the PATH environment variable for the current user
# by appending it to the "~/.bashrc" file. This ensures Ansible executable is in the PATH.
RUN echo "export PATH="$PATH:/root/.local/bin" >> /root/.bashrc"

# Create directories within the container for organizing Ansible playbooks, sources, and credentials.
RUN mkdir -p \
    /opt/ansible/playbooks \
    /opt/ansible/sources \
    /opt/ansible/credentials
