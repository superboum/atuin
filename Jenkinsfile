stage 'Commit'

node {
  def workspace = pwd()
  env.GOPATH = "${workspace}"
  checkout scm
  sh 'mkdir -p src/github.com/superboum && mv atuin src/github.com/superboum'
  sh 'go test -v github.com/superboum/atuin/...'
}
