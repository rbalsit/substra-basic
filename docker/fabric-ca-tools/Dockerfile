FROM hyperledger/fabric-tools:1.4.9

# Install curl and netcat
RUN apt-get update && \
  apt-get install -y curl netcat apt-transport-https vim

# Install fabric-ca-client
RUN curl -SL https://github.com/hyperledger/fabric-ca/releases/download/v1.4.8/hyperledger-fabric-ca-linux-amd64-1.4.8.tar.gz | tar xz --strip-components=1 && \
  mv ./fabric-ca-client /bin

# Install kubectl
RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.16.6/bin/linux/amd64/kubectl && \
  chmod +x ./kubectl && \
  mv ./kubectl /bin

# Install grpcurl for convenience
RUN wget https://github.com/fullstorydev/grpcurl/releases/download/v1.3.0/grpcurl_1.3.0_linux_x86_64.tar.gz && \
  tar xvzf grpcurl_1.3.0_linux_x86_64.tar.gz && \
  mv grpcurl /bin

# Create a non-root user and Group
# RUN groupadd -g 2000 hlfpgroup
# RUN useradd -u 2000 -g 2000 hlfpuser
USER root
RUN mkdir /.fabric-ca-client
RUN mkdir /data
# Change the ownership of the Fabric config files to the non-root user
# RUN chown -R 2000:2000  /
RUN chown -R 1000:1000  /etc/hyperledger
RUN chown -R 1000:1000  /var/hyperledger
RUN chown -R 1000:1000  /tmp/

RUN chown -R 1000:1000  /data/
RUN chown -R 1000:1000  /.fabric-ca-client

# WORKDIR /etc/hyperledger
# Run Fabric as the non-root user
USER 1000
