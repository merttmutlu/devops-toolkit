FROM jenkins/jenkins

# Switch to root user to install additional dependencies
USER root

# Install necessary packages for Python
RUN apt-get update -y && \
    apt-get install -y build-essential zlib1g-dev libncurses5-dev libgdbm-dev \
                       libnss3-dev libssl-dev libreadline-dev libffi-dev \
                       libsqlite3-dev wget libbz2-dev python3-distutils python3-apt

# Install Python3
RUN apt-get install -y python3

# Install Ansible using pip
RUN curl https://bootstrap.pypa.io/get-pip.py -o get-pip.py && \
    python3 get-pip.py && \
    python3 -m pip install ansible

# Switch back to the Jenkins user for security reasons
USER jenkins
