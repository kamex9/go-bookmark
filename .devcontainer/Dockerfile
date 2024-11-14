FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

# よく使用するGoのツールのインストール（オプション）
RUN su vscode -c "go install golang.org/x/tools/gopls@latest" && \
    su vscode -c "go install github.com/go-delve/delve/cmd/dlv@latest" && \
    su vscode -c "go install github.com/spf13/cobra-cli@latest"

# 作業ディレクトリの作成と権限設定
WORKDIR "/workspaces/go-bookmark"
RUN chown vscode:vscode /workspaces/go-bookmark
