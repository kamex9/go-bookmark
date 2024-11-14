FROM mcr.microsoft.com/devcontainers/ruby:3.3-bullseye

RUN apt update && \
    apt-get install --no-install-recommends -y \
    git \
    curl \
    wget \
    procps \
    net-tools \
    htop \
    strace \
    lsof && \
    unzip && \
    rm -rf /var/lib/apt/lists/*

# AWS CLI v2のインストール
RUN curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && \
    ./aws/install && \
    rm -rf awscliv2.zip aws

# Solargraphのインストール
RUN su vscode -c "gem install solargraph"
RUN su vscode -c "/usr/local/rvm/bin/rvm fix-permissions"

# Terraform最新版のインストール
RUN wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg && \
   echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list && \
   sudo apt update && sudo apt install terraform && \
   terraform -install-autocomplete 

# ecspressoのインストール
RUN curl -L https://github.com/kayac/ecspresso/releases/download/v2.4.1/ecspresso_2.4.1_linux_amd64.tar.gz -o ecspresso.tar.gz \
   && tar -zxvf ecspresso.tar.gz -C /usr/local/bin \
   && chmod +x /usr/local/bin/ecspresso \
   && rm ecspresso.tar.gz

# Golangのインストール
RUN wget -O go.tar.gz https://go.dev/dl/go1.23.3.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go.tar.gz && \
    rm go.tar.gz

# 環境変数の設定
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/home/vscode/go
ENV PATH=$PATH:$GOPATH/bin

# GOPATHディレクトリの作成と権限設定
RUN mkdir -p $GOPATH && \
    chown -R vscode:vscode $GOPATH

# よく使用するGoのツールのインストール（オプション）
RUN su vscode -c "go install golang.org/x/tools/gopls@latest" && \
    su vscode -c "go install github.com/go-delve/delve/cmd/dlv@latest"

# 作業ディレクトリの作成と権限設定
WORKDIR "/workspaces/infra-tools"
RUN chown vscode:vscode /workspaces/infra-tools
RUN bundle config set --local path vendor/bundle

# AWS認証情報ファイルのコピー
# COPY ./aws-config /home/vscode/.aws/config
# COPY ./aws-credentials /home/vscode/.aws/credentials

# ~/.aws ディレクトリの設定
RUN mkdir -p /home/vscode/.aws
RUN chown -R vscode:vscode /home/vscode/.aws