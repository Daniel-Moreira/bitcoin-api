FROM golang:latest

WORKDIR /go/src/bitcoin-api

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

COPY . .
RUN make build

# Install node and npm
RUN apt-get update && \
    apt-get install -y nodejs && \
    apt-get install -y npm && \
    npm install npm@latest -g

# Install packages
RUN npm i serverless -g
RUN npm i

# Install dynamodb local
# RUN sls dynamodb install

# Install OpenJDK-8
# RUN apt-get update && \
#     apt-get install -y openjdk-8-jdk && \
#     apt-get install -y ant && \
#     apt-get clean;

# Fix certificate issues
# RUN apt-get update && \
#     apt-get install ca-certificates-java && \
#     apt-get clean && \
#     update-ca-certificates -f;

# Setup JAVA_HOME -- useful for docker commandline
# ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64/
# RUN export JAVA_HOME

# Setup AWS credentials
ENV AWS_ACCESS_KEY_ID teste
ENV AWS_SECRET_ACCESS_KEY teste

EXPOSE 3000

CMD ["sls", "offline", "start"]
