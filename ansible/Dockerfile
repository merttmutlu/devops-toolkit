FROM python:3.9.18-slim

RUN apt-get update \
	&& apt-get install -y openssh-client \
	&& rm -rf /var/lib/apt/lists/*

RUN python3 -m pip install --user ansible

ENV PATH=/root/.local/bin:$PATH

RUN mkdir -p \
	/ansible/inventories \
	/ansible/roles \
	/ansible/keys \
	/ansible/playbooks

WORKDIR /ansible

COPY roles/ /ansible/roles/
COPY playbooks /ansible/playbooks/

CMD [ "ansible", "--version" ]
