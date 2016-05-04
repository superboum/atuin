stage 'Dev'

node {
  def workspace = pwd()
  env.GOPATH = "${workspace}"
  sh 'go get github.com/superboum/atuin/...'
  sh 'go test github.com/superboum/atuin/...'
}
