stage 'Commit'

node {
  def workspace = pwd()
  env.GOPATH = "${workspace}"
  sh 'go get github.com/superboum/atuin/...'
  sh 'go test -v github.com/superboum/atuin/...'
}
