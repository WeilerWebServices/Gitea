curl "https://api.github.com/orgs/go-gitea/repos?per_page=1000" | grep -o 'git@[^"]*' | xargs -L1 git clone

git clone https://github.com/go-gitea/go-sdk.git
git clone https://github.com/go-gitea/gitea.git
git clone https://github.com/go-gitea/lgtm.git
git clone https://github.com/go-gitea/git.git
git clone https://github.com/go-gitea/website.git
git clone https://github.com/go-gitea/docs.git
git clone https://github.com/go-gitea/blog.git
git clone https://github.com/go-gitea/infrastructure.git
git clone https://github.com/go-gitea/redirects.git
git clone https://github.com/go-gitea/theme.git
git clone https://github.com/go-gitea/homebrew-gitea.git
git clone https://github.com/go-gitea/lgtm-cli.git
git clone https://github.com/go-gitea/lgtm-go.git
git clone https://github.com/go-gitea/lgtm-docs.git
git clone https://github.com/go-gitea/bolt.git
git clone https://github.com/go-gitea/chardet.git
git clone https://github.com/go-gitea/changelog.git
git clone https://github.com/go-gitea/yaml.git
git clone https://github.com/go-gitea/test-openldap.git
git clone https://github.com/go-gitea/gitea-darkmode.git
git clone https://github.com/go-gitea/u2f-api.git
git clone https://github.com/go-gitea/debian-packaging.git
git clone https://github.com/go-gitea/tea.git
git clone https://github.com/go-gitea/test_repo.git
