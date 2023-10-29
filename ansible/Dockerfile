FROM python:3.9.18-slim

RUN apt-get update \
	&& apt-get install -y openssh-client \
	&& rm -rf /var/lib/apt/lists/*

RUN python3 -m pip install --user ansible

ENV PATH=/root/.local/bin:$PATH

RUN mkdir -p \
	/ansible/inventory \
	/ansible/playbooks/files \
	/ansible/roles/common \
	/ansible/roles/tasks \
	/ansible/roles/handlers \
	/ansible/roles/files \
	/ansible/roles/templates

WORKDIR /ansible

CMD [ "ansible", "--version" ]
